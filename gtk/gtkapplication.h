/*
 * Copyright © 2010 Codethink Limited
 * Copyright (C) 2010 Red Hat, Inc.
 *
 * This library is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 2 of the License, or (at your option) any later version.
 *
 * This library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.	 See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public
 * License along with this library; if not, write to the
 * Free Software Foundation, Inc., 59 Temple Place - Suite 330,
 * Boston, MA 02111-1307, USA.
 *
 * Author: Ryan Lortie <desrt@desrt.ca>
 *         Colin Walters <walters@verbum.org>
 */

#if defined(GTK_DISABLE_SINGLE_INCLUDES) && !defined (__GTK_H_INSIDE__) && !defined (GTK_COMPILATION)
#error "Only <gtk/gtk.h> can be included directly."
#endif

#ifndef __GTK_APPLICATION_H__
#define __GTK_APPLICATION_H__

#include <gtk/gtkactiongroup.h>
#include <gtk/gtkwidget.h>
#include <gio/gio.h>

G_BEGIN_DECLS

#define GTK_TYPE_APPLICATION            (gtk_application_get_type ())
#define GTK_APPLICATION(obj)            (G_TYPE_CHECK_INSTANCE_CAST ((obj), GTK_TYPE_APPLICATION, GtkApplication))
#define GTK_APPLICATION_CLASS(klass)    (G_TYPE_CHECK_CLASS_CAST ((klass), GTK_TYPE_APPLICATION, GtkApplicationClass))
#define GTK_IS_APPLICATION(obj)         (G_TYPE_CHECK_INSTANCE_TYPE ((obj), GTK_TYPE_APPLICATION))
#define GTK_IS_APPLICATION_CLASS(klass) (G_TYPE_CHECK_CLASS_TYPE ((klass), GTK_TYPE_APPLICATION))
#define GTK_APPLICATION_GET_CLASS(obj)  (G_TYPE_INSTANCE_GET_CLASS ((obj), GTK_TYPE_APPLICATION, GtkApplicationClass))

typedef struct _GtkApplication        GtkApplication;
typedef struct _GtkApplicationClass   GtkApplicationClass;
typedef struct _GtkApplicationPrivate GtkApplicationPrivate;

struct _GtkApplication
{
  GApplication parent;

  /*< private >*/
  GtkApplicationPrivate *priv;
};

struct _GtkApplicationClass
{
  GApplicationClass parent_class;

  /*< vfuncs >*/
  GtkWindow *(* create_window) (GtkApplication *application);

  /**
   * GtkApplicationClass::save_state:
   * @state: a dictionary where to store the application's state
   *
   * Class closure for the [signal@Application::save-state] signal.
   *
   * Returns: true to stop further handlers from running
   *
   * Since: 4.22
   */
  gboolean (* save_state)    (GtkApplication   *application,
                              GVariantDict     *state);

  /**
   * GtkApplicationClass::restore_state:
   * @reason: the reason for restoring state
   * @state: a dictionary containing the application state to restore
   *
   * Class closure for the [signal@Application::restore-state] signal.
   *
   * Returns: true to stop further handlers from running
   *
   * Since: 4.22
   */
  gboolean (* restore_state) (GtkApplication   *application,
                              GtkRestoreReason  reason,
                              GVariant         *state);

  /**
   * GtkApplicationClass::restore_window:
   * @reason: the reason this window is restored
   * @state: (nullable): the state to restore, as saved by a
   *   [signal@Gtk.ApplicationWindow::save-state] handler
   *
   * Class closure for the [signal@Application::restore-window] signal.
   *
   * Since: 4.22
   */
  void     (*restore_window) (GtkApplication   *application,
                              GtkRestoreReason  reason,
                              GVariant         *state);

  /*< private >*/
  gpointer padding[12];
};

GType                   gtk_application_get_type                        (void) G_GNUC_CONST;

GtkApplication*         gtk_application_new                             (const gchar       *appid,
                                                                         GApplicationFlags  flags,
                                                                         GtkActionGroup    *action_group);

GtkActionGroup *        gtk_application_get_action_group                (GtkApplication    *application);
void                    gtk_application_set_action_group                (GtkApplication    *application,
                                                                         GtkActionGroup    *action_group);

GtkWindow *             gtk_application_get_default_window              (GtkApplication    *application);
void                    gtk_application_set_default_window              (GtkApplication    *application,
                                                                         GtkWindow         *window);

G_END_DECLS

#endif /* __GTK_APPLICATION_H__ */

