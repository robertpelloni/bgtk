/* gdkmacosglcontext-private.h
 *
 * Copyright (C) 2020 Red Hat, Inc.
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

#pragma once

#include "gdkmacosglcontext.h"

#include "gdkglcontextprivate.h"
#include "gdkdisplayprivate.h"
#include "gdksurface.h"

#include "gdkmacosdisplay.h"
#include "gdkmacossurface.h"

#import <epoxy/gl.h>

#import <Cocoa/Cocoa.h>

G_BEGIN_DECLS

struct _GdkMacosGLContext
{
  GdkGLContext parent_instance;

  G_GNUC_BEGIN_IGNORE_DEPRECATIONS
  CGLContextObj cgl_context;
  G_GNUC_END_IGNORE_DEPRECATIONS

  GLuint texture;
  GLuint target;
  GLuint fbo;

  guint last_opaque : 1;
};

struct _GdkMacosGLContextClass
{
  GdkGLContextClass parent_class;
};

void       _gdk_quartz_device_manager_register_device_for_ns_event (GdkDeviceManager *device_manager,
                                                                    NSEvent          *nsevent);

GdkDevice *_gdk_quartz_device_manager_core_device_for_ns_event (GdkDeviceManager *device_manager,
                                                                NSEvent          *ns_event);

G_END_DECLS

