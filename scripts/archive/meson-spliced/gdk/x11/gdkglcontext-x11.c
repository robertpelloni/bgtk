/* GDK - The GIMP Drawing Kit
 *
 * gdkglcontext-x11.c: X11 specific OpenGL wrappers
 *
 * Copyright © 2014  Emmanuele Bassi
 *
 * This library is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Library General Public
 * License as published by the Free Software Foundation; either
 * version 2 of the License, or (at your option) any later version.
 *
 * This library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * Library General Public License for more details.
 *
 * You should have received a copy of the GNU Library General Public
 * License along with this library. If not, see <http://www.gnu.org/licenses/>.
 */

#include "config.h"

#include "gdkglcontext-x11.h"
#include "gdkdisplay-x11.h"
#include "gdkprivate-x11.h"
#include "gdkscreen-x11.h"

#include "gdkx11display.h"
#include "gdkx11glcontext.h"
#include "gdkx11screen.h"
#include "gdkx11property.h"
#include <X11/Xatom.h>

#include <glib/gi18n-lib.h>

#include <cairo-xlib.h>

#include <epoxy/glx.h>

G_GNUC_BEGIN_IGNORE_DEPRECATIONS

G_DEFINE_ABSTRACT_TYPE (GdkX11GLContext, gdk_x11_gl_context, GDK_TYPE_GL_CONTEXT)

static void
gdk_x11_gl_context_empty_frame (GdkDrawContext *draw_context)
{
  DrawableInfo *data = data_;
  Display *dpy;

  gdk_x11_display_error_trap_push (data->display);

  dpy = gdk_x11_display_get_xdisplay (data->display);

  if (data->glx_drawable)
    glXDestroyWindow (dpy, data->glx_drawable);

  if (data->dummy_glx)
    glXDestroyWindow (dpy, data->dummy_glx);

  if (data->dummy_xwin)
    XDestroyWindow (dpy, data->dummy_xwin);

  gdk_x11_display_error_trap_pop_ignored (data->display);

  g_slice_free (DrawableInfo, data);
}

static DrawableInfo *
get_glx_drawable_info (GdkWindow *window)
{
  return g_object_get_data (G_OBJECT (window), "-gdk-x11-window-glx-info");
}

static void
set_glx_drawable_info (GdkWindow    *window,
                       DrawableInfo *info)
{
  g_object_set_data_full (G_OBJECT (window), "-gdk-x11-window-glx-info",
                          info,
                          drawable_info_free);
}

static void
maybe_wait_for_vblank (GdkDisplay  *display,
                       GLXDrawable  drawable)
{
  GdkX11Display *display_x11 = GDK_X11_DISPLAY (display);
  Display *dpy = gdk_x11_display_get_xdisplay (display);

  if (display_x11->has_glx_sync_control)
    {
      gint64 ust, msc, sbc;

      glXGetSyncValuesOML (dpy, drawable, &ust, &msc, &sbc);
      glXWaitForMscOML (dpy, drawable,
                        0, 2, (msc + 1) % 2,
                        &ust, &msc, &sbc);
    }
  else if (display_x11->has_glx_video_sync)
    {
      guint32 current_count;

      glXGetVideoSyncSGI (&current_count);
      glXWaitVideoSyncSGI (2, (current_count + 1) % 2, &current_count);
    }
}

void
gdk_x11_window_invalidate_for_new_frame (GdkWindow      *window,
                                         cairo_region_t *update_area)
{
  cairo_rectangle_int_t window_rect;
  GdkDisplay *display = gdk_window_get_display (window);
  GdkX11Display *display_x11 = GDK_X11_DISPLAY (display);
  Display *dpy = gdk_x11_display_get_xdisplay (display);
  GdkX11GLContext *context_x11;
  unsigned int buffer_age;
  gboolean invalidate_all;

  /* Minimal update is ok if we're not drawing with gl */
  if (window->gl_paint_context == NULL)
    return;

  context_x11 = GDK_X11_GL_CONTEXT (window->gl_paint_context);

  buffer_age = 0;

  context_x11->do_blit_swap = FALSE;

  if (display_x11->has_glx_buffer_age)
    {
      gdk_gl_context_make_current (window->gl_paint_context);
      glXQueryDrawable(dpy, context_x11->drawable,
		       GLX_BACK_BUFFER_AGE_EXT, &buffer_age);
    }


  invalidate_all = FALSE;
  if (buffer_age >= 4)
    {
      cairo_rectangle_int_t whole_window = { 0, 0, gdk_window_get_width (window), gdk_window_get_height (window) };

      if (gdk_gl_context_has_framebuffer_blit (window->gl_paint_context) &&
          cairo_region_contains_rectangle (update_area, &whole_window) != CAIRO_REGION_OVERLAP_IN)
        {
          context_x11->do_blit_swap = TRUE;
        }
      else
        invalidate_all = TRUE;
    }
  else
    {
      if (buffer_age == 0)
        {
          invalidate_all = TRUE;
        }
      if (buffer_age >= 2)
        {
          if (window->old_updated_area[0])
            cairo_region_union (update_area, window->old_updated_area[0]);
          else
            invalidate_all = TRUE;
        }
      if (buffer_age >= 3)
        {
          if (window->old_updated_area[1])
            cairo_region_union (update_area, window->old_updated_area[1]);
          else
            invalidate_all = TRUE;
        }
    }

  if (invalidate_all)
    {
      window_rect.x = 0;
      window_rect.y = 0;
      window_rect.width = gdk_window_get_width (window);
      window_rect.height = gdk_window_get_height (window);

      /* If nothing else is known, repaint everything so that the back
         buffer is fully up-to-date for the swapbuffer */
      cairo_region_union_rectangle (update_area, &window_rect);
    }

}

static void
gdk_gl_blit_region (GdkWindow *window, cairo_region_t *region)
{
  int n_rects, i;
  int scale = gdk_window_get_scale_factor (window);
  int wh = gdk_window_get_height (window);
  cairo_rectangle_int_t rect;

  n_rects = cairo_region_num_rectangles (region);
  for (i = 0; i < n_rects; i++)
    {
      cairo_region_get_rectangle (region, i, &rect);
      glScissor (rect.x * scale, (wh - rect.y - rect.height) * scale, rect.width * scale, rect.height * scale);
      glBlitFramebuffer (rect.x * scale, (wh - rect.y - rect.height) * scale, (rect.x + rect.width) * scale, (wh - rect.y) * scale,
                         rect.x * scale, (wh - rect.y - rect.height) * scale, (rect.x + rect.width) * scale, (wh - rect.y) * scale,
                         GL_COLOR_BUFFER_BIT, GL_NEAREST);
    }
}

static void
gdk_x11_gl_context_end_frame (GdkGLContext *context,
                              cairo_region_t *painted,
                              cairo_region_t *damage)
{
  GdkX11GLContext *context_x11 = GDK_X11_GL_CONTEXT (context);
  GdkWindow *window = gdk_gl_context_get_window (context);
  GdkDisplay *display = gdk_gl_context_get_display (context);
  Display *dpy = gdk_x11_display_get_xdisplay (display);
  GdkX11Display *display_x11 = GDK_X11_DISPLAY (display);
  DrawableInfo *info;
  GLXDrawable drawable;

  gdk_gl_context_make_current (context);

  info = get_glx_drawable_info (window);

  drawable = context_x11->drawable;

  GDK_NOTE (OPENGL,
            g_message ("Flushing GLX buffers for drawable %lu (window: %lu), frame sync: %s",
                       (unsigned long) drawable,
                       (unsigned long) gdk_x11_window_get_xid (window),
                       context_x11->do_frame_sync ? "yes" : "no"));

  /* if we are going to wait for the vertical refresh manually
   * we need to flush pending redraws, and we also need to wait
   * for that to finish, otherwise we are going to tear.
   *
   * obviously, this condition should not be hit if we have
   * GLX_SGI_swap_control, and we ask the driver to do the right
   * thing.
   */
  if (context_x11->do_frame_sync)
    {
      guint32 end_frame_counter = 0;
      gboolean has_counter = display_x11->has_glx_video_sync;
      gboolean can_wait = display_x11->has_glx_video_sync || display_x11->has_glx_sync_control;

      if (display_x11->has_glx_video_sync)
        glXGetVideoSyncSGI (&end_frame_counter);

      if (context_x11->do_frame_sync && !display_x11->has_glx_swap_interval)
        {
          glFinish ();

          if (has_counter && can_wait)
            {
              guint32 last_counter = info != NULL ? info->last_frame_counter : 0;

              if (last_counter == end_frame_counter)
                maybe_wait_for_vblank (display, drawable);
            }
          else if (can_wait)
            maybe_wait_for_vblank (display, drawable);
        }
    }

  if (context_x11->do_blit_swap)
    {
      glDrawBuffer(GL_FRONT);
      glReadBuffer(GL_BACK);
      gdk_gl_blit_region (window, painted);
      glDrawBuffer(GL_BACK);
      glFlush();

      if (gdk_gl_context_has_frame_terminator (context))
        glFrameTerminatorGREMEDY ();
    }
  else
    glXSwapBuffers (dpy, drawable);

  if (context_x11->do_frame_sync && info != NULL && display_x11->has_glx_video_sync)
    glXGetVideoSyncSGI (&info->last_frame_counter);
}

typedef struct {
  Display *display;
  GLXDrawable drawable;
  gboolean y_inverted;
} GdkGLXPixmap;

static void
glx_pixmap_destroy (void *data)
{
  GdkGLXPixmap *glx_pixmap = data;

  glXDestroyPixmap (glx_pixmap->display, glx_pixmap->drawable);

  g_slice_free (GdkGLXPixmap, glx_pixmap);
}

static GdkGLXPixmap *
glx_pixmap_get (cairo_surface_t *surface, guint texture_target)
{
  Display *display = cairo_xlib_surface_get_display (surface);
  Screen *screen = cairo_xlib_surface_get_screen (surface);
  Visual *visual = cairo_xlib_surface_get_visual (surface);
  GdkGLXPixmap *glx_pixmap;
  GLXFBConfig *fbconfigs, config;
  int nfbconfigs;
  XVisualInfo *visinfo;
  VisualID visualid;
  int i, value;
  gboolean y_inverted;
  gboolean with_alpha;
  guint target = 0;
  guint format = 0;
  int pixmap_attributes[] = {
    GLX_TEXTURE_TARGET_EXT, 0,
    GLX_TEXTURE_FORMAT_EXT, 0,
    None
  };

  if (visual == NULL)
    return NULL;

  with_alpha = cairo_surface_get_content (surface) == CAIRO_CONTENT_COLOR_ALPHA;

  y_inverted = FALSE;
  fbconfigs = glXGetFBConfigs (display, XScreenNumberOfScreen (screen), &nfbconfigs);
  for (i = 0; i < nfbconfigs; i++)
    {
      visinfo = glXGetVisualFromFBConfig (display, fbconfigs[i]);
      if (!visinfo)
        continue;

      visualid = visinfo->visualid;
      XFree (visinfo);

      if (visualid != XVisualIDFromVisual (visual))
        continue;

      glXGetFBConfigAttrib (display, fbconfigs[i], GLX_DRAWABLE_TYPE, &value);
      if (!(value & GLX_PIXMAP_BIT))
        continue;

      glXGetFBConfigAttrib (display, fbconfigs[i],
                            GLX_BIND_TO_TEXTURE_TARGETS_EXT,
                            &value);
      if (texture_target == GL_TEXTURE_2D)
        {
          if (value & GLX_TEXTURE_2D_BIT_EXT)
            target = GLX_TEXTURE_2D_EXT;
          else
            continue;
        }
      else if (texture_target == GL_TEXTURE_RECTANGLE_ARB)
        {
          if (value & GLX_TEXTURE_RECTANGLE_BIT_EXT)
            target = GLX_TEXTURE_RECTANGLE_EXT;
          else
            continue;
        }
      else
        continue;

      if (!with_alpha)
        {
          glXGetFBConfigAttrib (display, fbconfigs[i],
                                GLX_BIND_TO_TEXTURE_RGB_EXT,
                                &value);
          if (!value)
            continue;

          format = GLX_TEXTURE_FORMAT_RGB_EXT;
        }
      else
        {
          glXGetFBConfigAttrib (display, fbconfigs[i],
                                GLX_BIND_TO_TEXTURE_RGBA_EXT,
                                &value);
          if (!value)
            continue;

          format = GLX_TEXTURE_FORMAT_RGBA_EXT;
        }

      glXGetFBConfigAttrib (display, fbconfigs[i],
                            GLX_Y_INVERTED_EXT,
                            &value);
      if (value == TRUE)
        y_inverted = TRUE;

      config = fbconfigs[i];
      break;
    }

  XFree (fbconfigs);

  if (i == nfbconfigs)
    return NULL;

  pixmap_attributes[1] = target;
  pixmap_attributes[3] = format;

  glx_pixmap = g_slice_new0 (GdkGLXPixmap);
  glx_pixmap->y_inverted = y_inverted;
  glx_pixmap->display = display;
  glx_pixmap->drawable = glXCreatePixmap (display, config,
					  cairo_xlib_surface_get_drawable (surface),
					  pixmap_attributes);

  return glx_pixmap;
}

static gboolean
gdk_x11_gl_context_texture_from_surface (GdkGLContext *paint_context,
					 cairo_surface_t *surface,
					 cairo_region_t *region)
{
  GdkGLXPixmap *glx_pixmap;
  double device_x_offset, device_y_offset;
  cairo_rectangle_int_t rect;
  int n_rects, i;
  GdkWindow *window;
  int unscaled_window_height;
  int window_scale;
  unsigned int texture_id;
  gboolean use_texture_rectangle;
  guint target;
  double sx, sy;
  float uscale, vscale;
  GdkTexturedQuad *quads;
  GdkX11Display *display_x11;

  display_x11 = GDK_X11_DISPLAY (gdk_gl_context_get_display (paint_context));
  if (!display_x11->has_glx_texture_from_pixmap)
    return FALSE;

  if (cairo_surface_get_type (surface) != CAIRO_SURFACE_TYPE_XLIB)
    return FALSE;

  use_texture_rectangle = gdk_gl_context_use_texture_rectangle (paint_context);
  if (use_texture_rectangle)
    target = GL_TEXTURE_RECTANGLE_ARB;
  else
    target = GL_TEXTURE_2D;

  glx_pixmap = glx_pixmap_get (surface, target);
  if (glx_pixmap == NULL)
    return FALSE;

  GDK_NOTE (OPENGL, g_message ("Using GLX_EXT_texture_from_pixmap to draw surface"));

  window = gdk_gl_context_get_window (paint_context)->impl_window;
  window_scale = gdk_window_get_scale_factor (window);
  gdk_window_get_unscaled_size (window, NULL, &unscaled_window_height);

  sx = sy = 1;
  cairo_surface_get_device_scale (window->current_paint.surface, &sx, &sy);
  cairo_surface_get_device_offset (surface, &device_x_offset, &device_y_offset);

  /* Ensure all the X stuff are synced before we read it back via texture-from-pixmap */
  glXWaitX();

  glGenTextures (1, &texture_id);
  glBindTexture (target, texture_id);

  glTexParameteri (target, GL_TEXTURE_WRAP_S, GL_CLAMP_TO_EDGE);
  glTexParameteri (target, GL_TEXTURE_WRAP_T, GL_CLAMP_TO_EDGE);
  glTexParameteri (target, GL_TEXTURE_MIN_FILTER, GL_NEAREST);
  glTexParameteri (target, GL_TEXTURE_MAG_FILTER, GL_NEAREST);

  glXBindTexImageEXT (glx_pixmap->display, glx_pixmap->drawable,
		      GLX_FRONT_LEFT_EXT, NULL);

  glEnable (GL_SCISSOR_TEST);

  n_rects = cairo_region_num_rectangles (region);
  quads = g_new (GdkTexturedQuad, n_rects);

#define FLIP_Y(_y) (unscaled_window_height - (_y))

  cairo_region_get_extents (region, &rect);
  glScissor (rect.x * window_scale, FLIP_Y((rect.y + rect.height) * window_scale),
             rect.width * window_scale, rect.height * window_scale);

  for (i = 0; i < n_rects; i++)
    {
      int src_x, src_y, src_height, src_width;

      cairo_region_get_rectangle (region, i, &rect);

      src_x = rect.x * sx + device_x_offset;
      src_y = rect.y * sy + device_y_offset;
      src_width = rect.width * sx;
      src_height = rect.height * sy;

      if (use_texture_rectangle)
        {
          uscale = 1.0;
          vscale = 1.0;
        }
      else
        {
          uscale = 1.0 / cairo_xlib_surface_get_width (surface);
          vscale = 1.0 / cairo_xlib_surface_get_height (surface);
        }

      {
        GdkTexturedQuad quad = {
          rect.x * window_scale, FLIP_Y(rect.y * window_scale),
          (rect.x + rect.width) * window_scale, FLIP_Y((rect.y + rect.height) * window_scale),
          uscale * src_x, vscale * src_y,
          uscale * (src_x + src_width), vscale * (src_y + src_height),
        };

        quads[i] = quad;
      }
    }

#undef FLIP_Y

  gdk_gl_texture_quads (paint_context, target, n_rects, quads, FALSE);
  g_free (quads);

  glDisable (GL_SCISSOR_TEST);

  glXReleaseTexImageEXT (glx_pixmap->display, glx_pixmap->drawable,
  			 GLX_FRONT_LEFT_EXT);

  glDeleteTextures (1, &texture_id);

  glx_pixmap_destroy(glx_pixmap);

  return TRUE;
}

static XVisualInfo *
find_xvisinfo_for_fbconfig (GdkDisplay  *display,
                            GLXFBConfig  config)
{
  Display *dpy = gdk_x11_display_get_xdisplay (display);

  return glXGetVisualFromFBConfig (dpy, config);
}

static GLXContext
create_gl3_context (GdkDisplay   *display,
                    GLXFBConfig   config,
                    GdkGLContext *share,
                    int           profile,
                    int           flags,
                    int           major,
                    int           minor)
{
  int attrib_list[] = {
    GLX_CONTEXT_PROFILE_MASK_ARB, profile,
    GLX_CONTEXT_MAJOR_VERSION_ARB, major,
    GLX_CONTEXT_MINOR_VERSION_ARB, minor,
    GLX_CONTEXT_FLAGS_ARB, flags,
    None,
  };
  GLXContext res;

  GdkX11GLContext *share_x11 = NULL;

  if (share != NULL)
    share_x11 = GDK_X11_GL_CONTEXT (share);

  gdk_x11_display_error_trap_push (display);

  res = glXCreateContextAttribsARB (gdk_x11_display_get_xdisplay (display),
                                    config,
                                    share_x11 != NULL ? share_x11->glx_context : NULL,
                                    True,
                                    attrib_list);

  if (gdk_x11_display_error_trap_pop (display))
    return NULL;

  return res;
}

static GLXContext
create_legacy_context (GdkDisplay   *display,
                       GLXFBConfig   config,
                       GdkGLContext *share)
{
  GdkX11GLContext *share_x11 = NULL;
  GLXContext res;

  if (share != NULL)
    share_x11 = GDK_X11_GL_CONTEXT (share);

  gdk_x11_display_error_trap_push (display);

  res = glXCreateNewContext (gdk_x11_display_get_xdisplay (display),
                             config,
                             GLX_RGBA_TYPE,
                             share_x11 != NULL ? share_x11->glx_context : NULL,
                             TRUE);

  if (gdk_x11_display_error_trap_pop (display))
    return NULL;

  return res;
}

static gboolean
gdk_x11_gl_context_realize (GdkGLContext  *context,
                            GError       **error)
{
  GdkX11Display *display_x11;
  GdkDisplay *display;
  GdkX11GLContext *context_x11;
  GLXWindow drawable;
  XVisualInfo *xvisinfo;
  Display *dpy;
  DrawableInfo *info;
  GdkGLContext *share;
  GdkWindow *window;
  gboolean debug_bit, compat_bit, legacy_bit, es_bit;
  int major, minor, flags;

  window = gdk_gl_context_get_window (context);
  display = gdk_window_get_display (window);
  dpy = gdk_x11_display_get_xdisplay (display);
  context_x11 = GDK_X11_GL_CONTEXT (context);
  display_x11 = GDK_X11_DISPLAY (display);
  share = gdk_gl_context_get_shared_context (context);

  gdk_gl_context_get_required_version (context, &major, &minor);
  debug_bit = gdk_gl_context_get_debug_enabled (context);
  compat_bit = gdk_gl_context_get_forward_compatible (context);

  /* If there is no glXCreateContextAttribsARB() then we default to legacy */
  legacy_bit = !display_x11->has_glx_create_context ||
               (_gdk_gl_flags & GDK_GL_LEGACY) != 0;

  es_bit = ((_gdk_gl_flags & GDK_GL_GLES) != 0 ||
            (share != NULL && gdk_gl_context_get_use_es (share))) &&
           (display_x11->has_glx_create_context && display_x11->has_glx_create_es2_context);

  /* We cannot share legacy contexts with core profile ones, so the
   * shared context is the one that decides if we're going to create
   * a legacy context or not.
   */
  if (share != NULL && gdk_gl_context_is_legacy (share))
    legacy_bit = TRUE;

  flags = 0;
  if (debug_bit)
    flags |= GLX_CONTEXT_DEBUG_BIT_ARB;
  if (compat_bit)
    flags |= GLX_CONTEXT_FORWARD_COMPATIBLE_BIT_ARB;

  GDK_NOTE (OPENGL,
            g_message ("Creating GLX context (version:%d.%d, debug:%s, forward:%s, legacy:%s, es:%s)",
                       major, minor,
                       debug_bit ? "yes" : "no",
                       compat_bit ? "yes" : "no",
                       legacy_bit ? "yes" : "no",
                       es_bit ? "yes" : "no"));

  /* If we have access to GLX_ARB_create_context_profile then we can ask for
   * a compatibility profile; if we don't, then we have to fall back to the
   * old GLX 1.3 API.
   */
  if (legacy_bit && !GDK_X11_DISPLAY (display)->has_glx_create_context)
    {
      GDK_NOTE (OPENGL, g_message ("Creating legacy GL context on request"));
      context_x11->glx_context = create_legacy_context (display, context_x11->glx_config, share);
    }
  else
    {
      int profile;

      if (es_bit)
        profile = GLX_CONTEXT_ES2_PROFILE_BIT_EXT;
      else
        profile = legacy_bit ? GLX_CONTEXT_COMPATIBILITY_PROFILE_BIT_ARB
                             : GLX_CONTEXT_CORE_PROFILE_BIT_ARB;

      /* We need to tweak the version, otherwise we may end up requesting
       * a compatibility context with a minimum version of 3.2, which is
       * an error
       */
      if (legacy_bit)
        {
          major = 3;
          minor = 0;
        }

      GDK_NOTE (OPENGL, g_message ("Creating GL3 context"));
      context_x11->glx_context = create_gl3_context (display,
                                                     context_x11->glx_config,
                                                     share,
                                                     profile, flags, major, minor);

      /* Fall back to legacy in case the GL3 context creation failed */
      if (context_x11->glx_context == NULL)
        {
          GDK_NOTE (OPENGL, g_message ("Creating fallback legacy context"));
          context_x11->glx_context = create_legacy_context (display, context_x11->glx_config, share);
          legacy_bit = TRUE;
          es_bit = FALSE;
        }
    }

  if (context_x11->glx_context == NULL)
    {
      g_set_error_literal (error, GDK_GL_ERROR,
                           GDK_GL_ERROR_NOT_AVAILABLE,
                           _("Unable to create a GL context"));
      return FALSE;
    }

  /* Ensure that any other context is created with a legacy bit set */
  gdk_gl_context_set_is_legacy (context, legacy_bit);

  /* Ensure that any other context is created with a ES bit set */
  gdk_gl_context_set_use_es (context, es_bit);

  xvisinfo = find_xvisinfo_for_fbconfig (display, context_x11->glx_config);

  info = get_glx_drawable_info (window->impl_window);
  if (info == NULL)
    {
      XSetWindowAttributes attrs;
      unsigned long mask;

      gdk_x11_display_error_trap_push (display);

      info = g_slice_new0 (DrawableInfo);
      info->display = display;
      info->last_frame_counter = 0;

      attrs.override_redirect = True;
      attrs.colormap = XCreateColormap (dpy, DefaultRootWindow (dpy), xvisinfo->visual, AllocNone);
      attrs.border_pixel = 0;
      mask = CWOverrideRedirect | CWColormap | CWBorderPixel;
      info->dummy_xwin = XCreateWindow (dpy, DefaultRootWindow (dpy),
                                        -100, -100, 1, 1,
                                        0,
                                        xvisinfo->depth,
                                        CopyFromParent,
                                        xvisinfo->visual,
                                        mask,
                                        &attrs);
      XMapWindow(dpy, info->dummy_xwin);

      if (GDK_X11_DISPLAY (display)->glx_version >= 13)
        {
          info->glx_drawable = glXCreateWindow (dpy, context_x11->glx_config,
                                                gdk_x11_window_get_xid (window->impl_window),
                                                NULL);
          info->dummy_glx = glXCreateWindow (dpy, context_x11->glx_config, info->dummy_xwin, NULL);
        }

      if (gdk_x11_display_error_trap_pop (display))
        {
          g_set_error_literal (error, GDK_GL_ERROR,
                               GDK_GL_ERROR_NOT_AVAILABLE,
                               _("Unable to create a GL context"));

          XFree (xvisinfo);
          drawable_info_free (info);
          glXDestroyContext (dpy, context_x11->glx_context);
          context_x11->glx_context = NULL;

          return FALSE;
        }

      set_glx_drawable_info (window->impl_window, info);
    }

  XFree (xvisinfo);

  if (context_x11->is_attached)
    drawable = info->glx_drawable ? info->glx_drawable : gdk_x11_window_get_xid (window->impl_window);
  else
    drawable = info->dummy_glx ? info->dummy_glx : info->dummy_xwin;

  context_x11->is_direct = glXIsDirect (dpy, context_x11->glx_context);
  context_x11->drawable = drawable;

  GDK_NOTE (OPENGL,
            g_message ("Realized GLX context[%p], %s",
                       context_x11->glx_context,
                       context_x11->is_direct ? "direct" : "indirect"));

  return TRUE;
}

static void
gdk_x11_gl_context_dispose (GObject *gobject)
{
  GdkX11GLContext *context_x11 = GDK_X11_GL_CONTEXT (gobject);

  if (context_x11->glx_context != NULL)
    {
      GdkGLContext *context = GDK_GL_CONTEXT (gobject);
      GdkDisplay *display = gdk_gl_context_get_display (context);
      Display *dpy = gdk_x11_display_get_xdisplay (display);

      if (glXGetCurrentContext () == context_x11->glx_context)
        glXMakeContextCurrent (dpy, None, None, NULL);

      GDK_NOTE (OPENGL, g_message ("Destroying GLX context"));
      glXDestroyContext (dpy, context_x11->glx_context);
      context_x11->glx_context = NULL;
    }

  G_OBJECT_CLASS (gdk_x11_gl_context_parent_class)->dispose (gobject);
}

static void
gdk_x11_gl_context_class_init (GdkX11GLContextClass *klass)
{
  GdkDrawContextClass *draw_context_class = GDK_DRAW_CONTEXT_CLASS (klass);

  draw_context_class->empty_frame = gdk_x11_gl_context_empty_frame;
}

static void
gdk_x11_gl_context_init (GdkX11GLContext *self)
{
}

