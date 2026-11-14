/* GDK - The GIMP Drawing Kit
 * Copyright (C) 2014 Red Hat, Inc.
 *
 * This library is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 2 of the License, or (at your option) any later version.
 *
 * This library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public
 * License along with this library. If not, see <http://www.gnu.org/licenses/>.
 */

#include "config.h"

#include "gdkcairo.h"

#include "gdkglcontextprivate.h"
#include "gdksurfaceprivate.h"

#ifdef GDK_WINDOWING_WIN32
# include "win32/gdkwin32.h"
#endif

#include <epoxy/gl.h>
#include <math.h>
#include <string.h>

/* x,y,width,height describes a rectangle in the gl render buffer
   coordinate space, and its top left corner is drawn at the current
   position according to the cairo translation. */

/**
 * gdk_cairo_draw_from_gl:
 * @cr: a cairo context
 * @surface: The surface we're rendering for (not necessarily into)
 * @source: The GL ID of the source buffer
 * @source_type: The type of the @source
 * @buffer_scale: The scale-factor that the @source buffer is allocated for
 * @x: The source x position in @source to start copying from in GL coordinates
 * @y: The source y position in @source to start copying from in GL coordinates
 * @width: The width of the region to draw
 * @height: The height of the region to draw
 *
 * Draws GL content onto a cairo context.
 *
 * It takes a render buffer ID (@source_type == GL_RENDERBUFFER) or a texture
 * id (@source_type == GL_TEXTURE) and draws it onto @cr with an OVER operation,
 * respecting the current clip. The top left corner of the rectangle specified
 * by @x, @y, @width and @height will be drawn at the current (0,0) position of
 * the `cairo_t`.
 *
 * This will work for *all* `cairo_t`, as long as @surface is realized, but the
 * fallback implementation that reads back the pixels from the buffer may be
 * used in the general case. In the case of direct drawing to a surface with
 * no special effects applied to @cr it will however use a more efficient
 * approach.
 *
 * For GL_RENDERBUFFER the code will always fall back to software for buffers
 * with alpha components, so make sure you use GL_TEXTURE if using alpha.
 *
 * Calling this may change the current GL context.
 *
 * Deprecated: 4.6: The function is overly complex and produces broken output
 *   in various combinations of arguments. If you want to draw with GL textures
 *   in BOBGUI, use [ctor@Gdk.GLTexture.new]; if you want to use that texture in
 *   Cairo, use [method@Gdk.Texture.download] to download the data into a Cairo
 *   image surface.
 */
void
gdk_cairo_draw_from_gl (cairo_t    *cr,
                        GdkSurface *surface,
                        int         source,
                        int         source_type,
                        int         buffer_scale,
                        int         x,
                        int         y,
                        int         width,
                        int         height)
{
  GdkGLContext *paint_context;
  cairo_surface_t *image;
  guint framebuffer;
  int alpha_size = 0;
  gboolean es_use_bgra = FALSE;

  paint_context = gdk_surface_get_paint_gl_context (surface, NULL);
  if (paint_context == NULL)
    {
      g_warning ("gdk_cairo_draw_gl_render_buffer failed - no paint context");
      return;
    }

  es_use_bgra = gdk_gl_context_use_es_bgra (paint_context);

  gdk_gl_context_make_current (paint_context);

  glGenFramebuffers (1, &framebuffer);
  if (paint_data->tmp_framebuffer == 0)
    glGenFramebuffers (1, &paint_data->tmp_framebuffer);

  if (source_type == GL_RENDERBUFFER)
    {
      glBindRenderbuffer (GL_RENDERBUFFER, source);
      glGetRenderbufferParameteriv (GL_RENDERBUFFER, GL_RENDERBUFFER_ALPHA_SIZE,  &alpha_size);
    }
  else if (source_type == GL_TEXTURE)
    {
      glBindTexture (GL_TEXTURE_2D, source);

      if (gdk_gl_context_get_use_es (paint_context))
        alpha_size = 1;
      else
        glGetTexLevelParameteriv (GL_TEXTURE_2D, 0, GL_TEXTURE_ALPHA_SIZE,  &alpha_size);
    }
  else
    {
      g_warning ("Unsupported gl source type %d\n", source_type);
      return;
    }

  /* TODO: avoid reading back non-required data due to dest clip */
  image = cairo_surface_create_similar_image (cairo_get_target (cr),
                                              (alpha_size == 0) ? CAIRO_FORMAT_RGB24 : CAIRO_FORMAT_ARGB32,
                                              width, height);

  cairo_surface_set_device_scale (image, buffer_scale, buffer_scale);

  glBindFramebuffer (GL_FRAMEBUFFER, framebuffer);

  if (source_type == GL_RENDERBUFFER)
    {
      /* Create a framebuffer with the source renderbuffer and
         make it the current target for reads */
      glFramebufferRenderbuffer (GL_FRAMEBUFFER, GL_COLOR_ATTACHMENT0,
                                 GL_RENDERBUFFER, source);
    }
  else
    {
      glFramebufferTexture2D (GL_FRAMEBUFFER, GL_COLOR_ATTACHMENT0,
                              GL_TEXTURE_2D, source, 0);
      framebuffer = paint_data->tmp_framebuffer;
      glBindFramebuffer (GL_FRAMEBUFFER, framebuffer);
      glFramebufferRenderbuffer (GL_FRAMEBUFFER, GL_COLOR_ATTACHMENT0,
                                 GL_RENDERBUFFER, source);
      glBindFramebuffer (GL_DRAW_FRAMEBUFFER, 0);

      /* Translate to impl coords */
      cairo_region_translate (clip_region, dx, dy);

      glEnable (GL_SCISSOR_TEST);

      gdk_window_get_unscaled_size (impl_window, NULL, &unscaled_window_height);

      /* We can use glDrawBuffer on OpenGL only; on GLES 2.0 we are already
       * double buffered so we don't need it...
       */
      if (!gdk_gl_context_get_use_es (paint_context))
        glDrawBuffer (GL_BACK);
      else
        {
          int maj, min;

          gdk_gl_context_get_version (paint_context, &maj, &min);

          /* ... but on GLES 3.0 we can use the vectorized glDrawBuffers
           * call.
           */
          if ((maj * 100 + min) >= 300)
            {
              static const GLenum buffers[] = { GL_BACK };

              glDrawBuffers (G_N_ELEMENTS (buffers), buffers);
            }
        }

#define FLIP_Y(_y) (unscaled_window_height - (_y))

      for (i = 0; i < cairo_region_num_rectangles (clip_region); i++)
        {
          cairo_rectangle_int_t clip_rect, dest;

          cairo_region_get_rectangle (clip_region, i, &clip_rect);
          clip_rect.x *= window_scale;
          clip_rect.y *= window_scale;
          clip_rect.width *= window_scale;
          clip_rect.height *= window_scale;

          glScissor (clip_rect.x, FLIP_Y (clip_rect.y + clip_rect.height),
                     clip_rect.width, clip_rect.height);

          dest.x = dx * window_scale;
          dest.y = dy * window_scale;
          dest.width = width * window_scale / buffer_scale;
          dest.height = height * window_scale / buffer_scale;

          if (gdk_rectangle_intersect (&clip_rect, &dest, &dest))
            {
              int clipped_src_x = x + (dest.x - dx * window_scale);
              int clipped_src_y = y + (height - dest.height - (dest.y - dy * window_scale));
              glBlitFramebuffer (clipped_src_x, clipped_src_y,
                                 (clipped_src_x + dest.width), (clipped_src_y + dest.height),
                                 dest.x, FLIP_Y(dest.y + dest.height),
                                 dest.x + dest.width, FLIP_Y(dest.y),
                                 GL_COLOR_BUFFER_BIT, GL_NEAREST);
              if (impl_window->current_paint.flushed_region)
                {
                  cairo_rectangle_int_t flushed_rect;

                  flushed_rect.x = dest.x / window_scale;
                  flushed_rect.y = dest.y / window_scale;
                  flushed_rect.width = (dest.x + dest.width + window_scale - 1) / window_scale - flushed_rect.x;
                  flushed_rect.height = (dest.y + dest.height + window_scale - 1) / window_scale - flushed_rect.y;

                  cairo_region_union_rectangle (impl_window->current_paint.flushed_region,
                                                &flushed_rect);
                  cairo_region_subtract_rectangle (impl_window->current_paint.need_blend_region,
                                                   &flushed_rect);
                }
            }
        }

      glDisable (GL_SCISSOR_TEST);

      glBindFramebuffer (GL_FRAMEBUFFER, 0);

#undef FLIP_Y

    }
  /* For direct paint of alpha or non-alpha textures we can use texturing */
  else if ((_gdk_gl_flags & GDK_GL_SOFTWARE_DRAW_GL) == 0 &&
           source_type == GL_TEXTURE &&
           direct_window != NULL &&
           direct_window->current_paint.use_gl &&
           trivial_transform &&
           clip_region != NULL)
    {
      int unscaled_window_height;
      GLint texture_width;
      GLint texture_height;
      int i, n_rects, n_quads;
      GdkTexturedQuad *quads;
      cairo_rectangle_int_t clip_rect;

      /* Translate to impl coords */
      cairo_region_translate (clip_region, dx, dy);

      if (alpha_size != 0)
        {
          cairo_region_t *opaque_region, *blend_region;

          opaque_region = cairo_region_copy (clip_region);
          cairo_region_subtract (opaque_region, impl_window->current_paint.flushed_region);
          cairo_region_subtract (opaque_region, impl_window->current_paint.need_blend_region);

          if (!cairo_region_is_empty (opaque_region))
            gdk_gl_texture_from_surface (impl_window->current_paint.surface,
                                         opaque_region);

          blend_region = cairo_region_copy (clip_region);
          cairo_region_intersect (blend_region, impl_window->current_paint.need_blend_region);

          glEnable (GL_BLEND);
          if (!cairo_region_is_empty (blend_region))
            gdk_gl_texture_from_surface (impl_window->current_paint.surface,
                                         blend_region);

          cairo_region_destroy (opaque_region);
          cairo_region_destroy (blend_region);
        }

      glBindTexture (GL_TEXTURE_2D, source);

      if (gdk_gl_context_get_use_es (paint_context))
        {
          texture_width = width;
          texture_height = height;
        }
      else
        {
          glGetTexLevelParameteriv (GL_TEXTURE_2D, 0, GL_TEXTURE_WIDTH, &texture_width);
          glGetTexLevelParameteriv (GL_TEXTURE_2D, 0, GL_TEXTURE_HEIGHT, &texture_height);
        }

      glTexParameteri (GL_TEXTURE_2D, GL_TEXTURE_WRAP_S, GL_REPEAT);
      glTexParameteri (GL_TEXTURE_2D, GL_TEXTURE_WRAP_T, GL_REPEAT);
      glTexParameteri (GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, GL_NEAREST);
      glTexParameteri (GL_TEXTURE_2D, GL_TEXTURE_MAG_FILTER, GL_NEAREST);

      glEnable (GL_SCISSOR_TEST);

      gdk_window_get_unscaled_size (impl_window, NULL, &unscaled_window_height);

#define FLIP_Y(_y) (unscaled_window_height - (_y))

      cairo_region_get_extents (clip_region, &clip_rect);

      glScissor (clip_rect.x * window_scale, FLIP_Y ((clip_rect.y + clip_rect.height) * window_scale),
                 clip_rect.width * window_scale, clip_rect.height * window_scale);

      n_quads = 0;
      n_rects = cairo_region_num_rectangles (clip_region);
      quads = g_new (GdkTexturedQuad, n_rects);
      for (i = 0; i < n_rects; i++)
        {
          cairo_rectangle_int_t dest;

          cairo_region_get_rectangle (clip_region, i, &clip_rect);

          clip_rect.x *= window_scale;
          clip_rect.y *= window_scale;
          clip_rect.width *= window_scale;
          clip_rect.height *= window_scale;

          dest.x = dx * window_scale;
          dest.y = dy * window_scale;
          dest.width = width * window_scale / buffer_scale;
          dest.height = height * window_scale / buffer_scale;

          if (gdk_rectangle_intersect (&clip_rect, &dest, &dest))
            {
              int clipped_src_x = x + (dest.x - dx * window_scale);
              int clipped_src_y = y + (height - dest.height - (dest.y - dy * window_scale));
              GdkTexturedQuad quad = {
                dest.x, FLIP_Y(dest.y),
                dest.x + dest.width, FLIP_Y(dest.y + dest.height),
                clipped_src_x / (float)texture_width, (clipped_src_y + dest.height) / (float)texture_height,
                (clipped_src_x + dest.width) / (float)texture_width, clipped_src_y / (float)texture_height,
              };

              quads[n_quads++] = quad;

              if (impl_window->current_paint.flushed_region)
                {
                  cairo_rectangle_int_t flushed_rect;

                  flushed_rect.x = dest.x / window_scale;
                  flushed_rect.y = dest.y / window_scale;
                  flushed_rect.width = (dest.x + dest.width + window_scale - 1) / window_scale - flushed_rect.x;
                  flushed_rect.height = (dest.y + dest.height + window_scale - 1) / window_scale - flushed_rect.y;

                  cairo_region_union_rectangle (impl_window->current_paint.flushed_region,
                                                &flushed_rect);
                  cairo_region_subtract_rectangle (impl_window->current_paint.need_blend_region,
                                                   &flushed_rect);
                }
            }
        }

      if (n_quads > 0)
        gdk_gl_texture_quads (paint_context, GL_TEXTURE_2D, n_quads, quads, FALSE);

      g_free (quads);

      if (alpha_size != 0)
        glDisable (GL_BLEND);

#undef FLIP_Y

    }
  else
    {
      /* Software fallback */
      int major, minor, version;
      gboolean es_read_bgra = FALSE;

#ifdef GDK_WINDOWING_WIN32
      /* on ANGLE GLES, we need to set the glReadPixel() format as GL_BGRA instead */
      if (GDK_WIN32_IS_GL_CONTEXT(paint_context))
        es_read_bgra = TRUE;
#endif

      gdk_gl_context_get_version (paint_context, &major, &minor);
      version = major * 100 + minor;

      /* TODO: Use glTexSubImage2D() and do a row-by-row copy to replace
       * the GL_UNPACK_ROW_LENGTH support
       */
      if (gdk_gl_context_get_use_es (paint_context) &&
          !(version >= 300 || gdk_gl_context_has_unpack_subimage (paint_context)))
        goto out;

      /* TODO: avoid reading back non-required data due to dest clip */
      image = cairo_surface_create_similar_image (cairo_get_target (cr),
                                                  (alpha_size == 0) ? CAIRO_FORMAT_RGB24 : CAIRO_FORMAT_ARGB32,
                                                  width, height);

      cairo_surface_set_device_scale (image, buffer_scale, buffer_scale);

      framebuffer = paint_data->tmp_framebuffer;
      glBindFramebuffer (GL_FRAMEBUFFER, framebuffer);

      if (source_type == GL_RENDERBUFFER)
        {
          /* Create a framebuffer with the source renderbuffer and
             make it the current target for reads */
          glFramebufferRenderbuffer (GL_FRAMEBUFFER, GL_COLOR_ATTACHMENT0,
                                     GL_RENDERBUFFER, source);
        }
      else
        {
          glFramebufferTexture2D (GL_FRAMEBUFFER, GL_COLOR_ATTACHMENT0,
                                  GL_TEXTURE_2D, source, 0);
        }

      glPixelStorei (GL_PACK_ALIGNMENT, 4);
      glPixelStorei (GL_PACK_ROW_LENGTH, cairo_image_surface_get_stride (image) / 4);

      /* The implicit format conversion is going to make this path slower */
      if (!gdk_gl_context_get_use_es (paint_context))
        glReadPixels (x, y, width, height, GL_BGRA, GL_UNSIGNED_INT_8_8_8_8_REV,
                      cairo_image_surface_get_data (image));
      else
        glReadPixels (x, y, width, height, es_read_bgra ? GL_BGRA : GL_RGBA, GL_UNSIGNED_BYTE,
                      cairo_image_surface_get_data (image));

      glPixelStorei (GL_PACK_ROW_LENGTH, 0);

      glBindFramebuffer (GL_FRAMEBUFFER, 0);

      cairo_surface_mark_dirty (image);

      /* Invert due to opengl having different origin */
      cairo_scale (cr, 1, -1);
      cairo_translate (cr, 0, -height / buffer_scale);

      cairo_set_source_surface (cr, image, 0, 0);
      cairo_set_operator (cr, CAIRO_OPERATOR_OVER);
      cairo_paint (cr);

      cairo_surface_destroy (image);
    }

  glPixelStorei (GL_PACK_ALIGNMENT, 4);
  glPixelStorei (GL_PACK_ROW_LENGTH, cairo_image_surface_get_stride (image) / 4);

  /* The implicit format conversion is going to make this path slower */
  if (!gdk_gl_context_get_use_es (paint_context))
    glReadPixels (x, y, width, height, GL_BGRA, GL_UNSIGNED_INT_8_8_8_8_REV,
                  cairo_image_surface_get_data (image));
  else
    glReadPixels (x, y, width, height, es_use_bgra ? GL_BGRA : GL_RGBA, GL_UNSIGNED_BYTE,
                  cairo_image_surface_get_data (image));

  glPixelStorei (GL_PACK_ROW_LENGTH, 0);

  glBindFramebuffer (GL_FRAMEBUFFER, 0);

  glDeleteFramebuffers (1, &framebuffer);

  cairo_surface_mark_dirty (image);

  cairo_set_source_surface (cr, image, 0, 0);
  cairo_set_operator (cr, CAIRO_OPERATOR_OVER);
  cairo_paint (cr);

  cairo_surface_destroy (image);
}
