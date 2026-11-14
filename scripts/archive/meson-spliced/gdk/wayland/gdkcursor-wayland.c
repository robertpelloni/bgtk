/* GDK - The GIMP Drawing Kit
 * Copyright (C) 1995-1997 Peter Mattis, Spencer Kimball and Josh MacDonald
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

/*
 * Modified by the BOBGUI+ Team and others 1997-2000.  See the AUTHORS
 * file for a list of people on the BOBGUI+ Team.  See the ChangeLog
 * files for a list of changes.  These files are distributed with
 * BOBGUI+ at ftp://ftp.bobgui.org/pub/bobgui/.
 */

#include "config.h"

#define GDK_PIXBUF_ENABLE_BACKEND

#include <string.h>

#include "gdkcursor-wayland.h"
#include "gdkcursorprivate.h"
#include "gdkdisplay-wayland.h"
#include "gdkshm-private.h"


typedef struct
{
  const char *name;
  GdkTexture *texture;
  int hotspot_x;
  int hotspot_y;
} CursorTexture;

static CursorTexture textures[] = {
  { "default",       NULL,  5,  5 },
  { "context-menu",  NULL,  5,  5 },
  { "help",          NULL, 16, 27 },
  { "pointer",       NULL, 14,  9 },
  { "progress",      NULL,  5,  4 },
  { "wait",          NULL, 11, 11 },
  { "cell",          NULL, 15, 15 },
  { "crosshair",     NULL, 15, 15 },
  { "text",          NULL, 14, 15 },
  { "vertical-text", NULL, 16, 15 },
  { "alias",         NULL, 12, 11 },
  { "copy",          NULL, 12, 11 },
  { "move",          NULL, 12, 11 },
  { "no-drop",       NULL, 12, 11 },
  { "not-allowed",   NULL, 12, 11 },
  { "grab",          NULL, 10,  6 },
  { "grabbing",      NULL, 15, 14 },
  { "e-resize",      NULL, 25, 17 },
  { "n-resize",      NULL, 17,  7 },
  { "ne-resize",     NULL, 20, 13 },
  { "nw-resize",     NULL, 13, 13 },
  { "s-resize",      NULL, 17, 23 },
  { "se-resize",     NULL, 19, 19 },
  { "sw-resize",     NULL, 13, 19 },
  { "w-resize",      NULL,  8, 17 },
  { "ew-resize",     NULL, 16, 15 },
  { "ns-resize",     NULL, 15, 17 },
  { "nesw-resize",   NULL, 14, 14 },
  { "nwse-resize",   NULL, 14, 14 },
  { "col-resize",    NULL, 16, 15 },
  { "row-resize",    NULL, 15, 17 },
  { "all-scroll",    NULL, 15, 15 },
  { "zoom-in",       NULL, 14, 13 },
  { "zoom-out",      NULL, 14, 13 },
  { "dnd-ask",       NULL,  5,  5 },
  { "all-resize",    NULL, 12, 11 },
};

static GdkTexture *
get_cursor_texture (const char *name,
                    int        *hotspot_x,
                    int        *hotspot_y)
{
  for (gsize i = 0; i < G_N_ELEMENTS (textures); i++)
    {
      if (strcmp (name, textures[i].name) == 0)
        {
          if (textures[i].texture == NULL)
            {
              char path[256];

              g_snprintf (path, sizeof (path), "/org/bobgui/libgdk/cursor/%s", name);
              if (g_resources_get_info (path, 0, NULL, NULL, NULL))
                textures[i].texture = gdk_texture_new_from_resource (path);
            }

          *hotspot_x = textures[i].hotspot_x;
          *hotspot_y = textures[i].hotspot_y;

          return textures[i].texture;
        }
    }

  return NULL;
}

struct wl_buffer *
_gdk_wayland_cursor_get_buffer (GdkWaylandDisplay *display,
                                GdkCursor         *cursor,
                                double             desired_scale,
                                int               *hotspot_x,
                                int               *hotspot_y,
                                int               *width,
                                int               *height,
                                double            *scale)
{
  GdkTexture *texture;
  const char *name;
  GdkCursor *fallback;
  GdkWaylandCursor *wayland_cursor = GDK_WAYLAND_CURSOR (cursor);

  if (wayland_cursor->wl_cursor)
    {
      struct wl_cursor_image *image;

      if (image_index >= wayland_cursor->wl_cursor->image_count)
        {
          g_warning (G_STRLOC " out of bounds cursor image [%d / %d]",
                     image_index,
                     wayland_cursor->wl_cursor->image_count - 1);
          image_index = 0;
        }

      image = wayland_cursor->wl_cursor->images[image_index];

      *hotspot_x = image->hotspot_x / wayland_cursor->scale;
      *hotspot_y = image->hotspot_y / wayland_cursor->scale;

      *w = image->width / wayland_cursor->scale;
      *h = image->height / wayland_cursor->scale;
      *scale = wayland_cursor->scale;

      return wl_cursor_image_get_buffer (image);
    }
  else if (wayland_cursor->name == NULL) /* From surface */
    {
      *hotspot_x =
        wayland_cursor->surface.hotspot_x / wayland_cursor->surface.scale;
      *hotspot_y =
        wayland_cursor->surface.hotspot_y / wayland_cursor->surface.scale;

      *w = wayland_cursor->surface.width / wayland_cursor->surface.scale;
      *h = wayland_cursor->surface.height / wayland_cursor->surface.scale;
      *scale = wayland_cursor->surface.scale;

      cairo_surface_reference (wayland_cursor->surface.cairo_surface);

      if (wayland_cursor->surface.cairo_surface)
        return _gdk_wayland_shm_surface_get_wl_buffer (wayland_cursor->surface.cairo_surface);
    }
  else
    {
      *hotspot_x = 0;
      *hotspot_y = 0;
      *w = 0;
      *h = 0;
      *scale = 1;
    }

  return NULL;
}

guint
_gdk_wayland_cursor_get_next_image_index (GdkCursor *cursor,
                                          guint      current_image_index,
                                          guint     *next_image_delay)
{
  struct wl_cursor *wl_cursor = GDK_WAYLAND_CURSOR (cursor)->wl_cursor;

  if (wl_cursor && wl_cursor->image_count > 1)
    {
      if (current_image_index >= wl_cursor->image_count)
        {
          g_warning (G_STRLOC " out of bounds cursor image [%d / %d]",
                     current_image_index, wl_cursor->image_count - 1);
          current_image_index = 0;
        }

      /* Return the time to next image */
      if (next_image_delay)
        *next_image_delay = wl_cursor->images[current_image_index]->delay;

      return (current_image_index + 1) % wl_cursor->image_count;
    }
  else
    return current_image_index;
}

void
_gdk_wayland_cursor_set_scale (GdkCursor *cursor,
                               guint      scale)
{
  GdkWaylandDisplay *display_wayland =
    GDK_WAYLAND_DISPLAY (gdk_cursor_get_display (cursor));
  GdkWaylandCursor *wayland_cursor = GDK_WAYLAND_CURSOR (cursor);

  if (scale > GDK_WAYLAND_MAX_THEME_SCALE)
    {
      g_warning (G_STRLOC ": cursor theme size %u too large", scale);
      scale = GDK_WAYLAND_MAX_THEME_SCALE;
    }

  if (wayland_cursor->scale == scale)
    return;

  wayland_cursor->scale = scale;

  /* Blank cursor case */
  if (g_strcmp0 (wayland_cursor->name, "none") == 0)
    return;

  _gdk_wayland_cursor_update (display_wayland, wayland_cursor);
}

static void
_gdk_wayland_cursor_class_init (GdkWaylandCursorClass *wayland_cursor_class)
{
  GdkCursorClass *cursor_class = GDK_CURSOR_CLASS (wayland_cursor_class);
  GObjectClass *object_class = G_OBJECT_CLASS (wayland_cursor_class);

  object_class->finalize = gdk_wayland_cursor_finalize;

  cursor_class->get_surface = gdk_wayland_cursor_get_surface;
}

static void
_gdk_wayland_cursor_init (GdkWaylandCursor *cursor)
{
}

static GdkCursor *
_gdk_wayland_display_get_cursor_for_name_with_scale (GdkDisplay  *display,
                                                     const gchar *name,
                                                     guint        scale)
{
  GdkWaylandCursor *wayland_cursor;
  GdkWaylandDisplay *display_wayland = GDK_WAYLAND_DISPLAY (display);

  g_return_val_if_fail (GDK_IS_DISPLAY (display), NULL);

  wayland_cursor = g_hash_table_lookup (display_wayland->cursor_cache, name);
  if (wayland_cursor && wayland_cursor->scale == scale)
    return GDK_CURSOR (g_object_ref (wayland_cursor));

  wayland_cursor = g_object_new (GDK_TYPE_WAYLAND_CURSOR,
                                 "cursor-type", GDK_CURSOR_IS_PIXMAP,
                                 "display", display,
                                 NULL);

  /* Blank cursor case */
  if (!name || g_str_equal (name, "none") || g_str_equal (name, "blank_cursor"))
    {
      wayland_cursor->name = g_strdup ("none");
      wayland_cursor->scale = scale;

      return GDK_CURSOR (wayland_cursor);
    }

  wayland_cursor->name = g_strdup (name);
  wayland_cursor->scale = scale;

  if (!_gdk_wayland_cursor_update (display_wayland, wayland_cursor))
    {
      g_object_unref (wayland_cursor);
      return NULL;
    }

  /* Insert into cache. */
  g_hash_table_replace (display_wayland->cursor_cache,
                        wayland_cursor->name,
                        g_object_ref (wayland_cursor));
  return GDK_CURSOR (wayland_cursor);
}

GdkCursor *
_gdk_wayland_display_get_cursor_for_name (GdkDisplay  *display,
                                          const gchar *name)
{
  return _gdk_wayland_display_get_cursor_for_name_with_scale (display, name, 1);
}

GdkCursor *
_gdk_wayland_display_get_cursor_for_type_with_scale (GdkDisplay    *display,
                                                     GdkCursorType  cursor_type,
                                                     guint          scale)
{
  GEnumClass *enum_class;
  GEnumValue *enum_value;
  gchar *cursor_name;
  GdkCursor *result;

  enum_class = g_type_class_ref (GDK_TYPE_CURSOR_TYPE);
  enum_value = g_enum_get_value (enum_class, cursor_type);
  cursor_name = g_strdup (enum_value->value_nick);
  g_strdelimit (cursor_name, "-", '_');
  g_type_class_unref (enum_class);

  result = _gdk_wayland_display_get_cursor_for_name_with_scale (display,
                                                                cursor_name,
                                                                scale);

  g_free (cursor_name);

  return result;
}

GdkCursor *
_gdk_wayland_display_get_cursor_for_type (GdkDisplay    *display,
                                          GdkCursorType  cursor_type)
{
  return _gdk_wayland_display_get_cursor_for_type_with_scale (display,
                                                              cursor_type,
                                                              1);
}

static void
buffer_release_callback (void             *_data,
                         struct wl_buffer *wl_buffer)
{
  cairo_surface_t *cairo_surface = _data;

  cairo_surface_destroy (cairo_surface);
}

static const struct wl_buffer_listener buffer_listener = {
  buffer_release_callback
};

GdkCursor *
_gdk_wayland_display_get_cursor_for_surface (GdkDisplay *display,
					     cairo_surface_t *surface,
					     gdouble     x,
					     gdouble     y)
{
  GdkWaylandCursor *cursor;
  GdkWaylandDisplay *display_wayland = GDK_WAYLAND_DISPLAY (display);
  struct wl_buffer *buffer;

  *scale = 1;
  *hotspot_x = *hotspot_y = 0;
  *width = *height = 0;

  name = gdk_cursor_get_name (cursor);
  if (name)
    {
      if (g_str_equal (name, "none"))
        return NULL;

      texture = get_cursor_texture (name, hotspot_x, hotspot_y);
      if (texture)
        {
          g_object_ref (texture);
          *scale = (double) display->cursor_theme_size / gdk_texture_get_width (texture);
        }
    }
  else if (gdk_cursor_get_texture (cursor))
    {
      texture = g_object_ref (gdk_cursor_get_texture (cursor));

      *hotspot_x = gdk_cursor_get_hotspot_x (cursor);
      *hotspot_y = gdk_cursor_get_hotspot_y (cursor);
    }
  else
    {
      int w, h;

      texture = gdk_cursor_get_texture_for_size (cursor,
                                                 display->cursor_theme_size,
                                                 desired_scale,
                                                 &w,
                                                 &h,
                                                 hotspot_x,
                                                 hotspot_y);

      if (texture)
        {
          *scale = MAX (gdk_texture_get_width (texture) / (float) w,
                        gdk_texture_get_height (texture) / (float) h);
        }
    }

  if (texture)
    {
      *width = gdk_texture_get_width (texture);
      *height = gdk_texture_get_height (texture);

      buffer = _gdk_wayland_shm_texture_get_wl_buffer (display, texture);

      g_object_unref (texture);

      return buffer;
    }

  if (gdk_cursor_get_fallback (cursor))
    fallback = g_object_ref (gdk_cursor_get_fallback (cursor));
  else
    fallback = gdk_cursor_new_from_name ("default", NULL);

  buffer = _gdk_wayland_cursor_get_buffer (display,
                                           fallback,
                                           desired_scale,
                                           hotspot_x, hotspot_y,
                                           width, height,
                                           scale);

  g_object_unref (fallback);

  return buffer;
}
