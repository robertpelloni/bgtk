/* GTK - The GIMP Toolkit
 * Copyright (C) 2011 Alberto Ruiz <aruiz@gnome.org>
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
 * License along with this library. If not, see <http://www.gnu.org/licenses/>
 *
 */

#include "config.h"

#include <stdlib.h>
#include <glib/gprintf.h>
#include <string.h>

#include <atk/atk.h>

#include "gtkfontchooserdialog.h"
#include "gtkfontchooser.h"
#include "gtkfontchooserwidget.h"
#include "gtkfontchooserwidgetprivate.h"
#include "gtkfontchooserutils.h"
#include "gtkbox.h"
#include "deprecated/gtkstock.h"
#include "gtkintl.h"
#include "gtkaccessible.h"
#include "gtkbuildable.h"
#include "gtkprivate.h"
#include "gtkwidget.h"
#include "gtksettings.h"
#include "gtkdialogprivate.h"
#include "gtktogglebutton.h"
#include "gtkheaderbar.h"
#include "gtkactionable.h"

struct _GtkFontChooserDialogPrivate
{
  GtkWidget *fontchooser;

  GtkWidget *select_button;
  GtkWidget *cancel_button;
  GtkWidget *tweak_button;
};

/**
 * SECTION:gtkfontchooserdialog
 * @Short_description: A dialog for selecting fonts
 * @Title: GtkFontChooserDialog
 * @See_also: #GtkFontChooser, #GtkDialog
 *
 * The #GtkFontChooserDialog widget is a dialog for selecting a font.
 * It implements the #GtkFontChooser interface.
 *
 * # GtkFontChooserDialog as GtkBuildable
 *
 * The GtkFontChooserDialog implementation of the #GtkBuildable
 * interface exposes the buttons with the names “select_button”
 * and “cancel_button”.
 *
 * Since: 3.2
 */

static void     gtk_font_chooser_dialog_buildable_interface_init     (GtkBuildableIface *iface);
static GObject *gtk_font_chooser_dialog_buildable_get_internal_child (GtkBuildable *buildable,
                                                                      GtkBuilder   *builder,
                                                                      const gchar  *childname);

G_DEFINE_TYPE_WITH_CODE (GtkFontChooserDialog, gtk_font_chooser_dialog, GTK_TYPE_DIALOG,
                         G_ADD_PRIVATE (GtkFontChooserDialog)
                         G_IMPLEMENT_INTERFACE (GTK_TYPE_FONT_CHOOSER,
                                                _gtk_font_chooser_delegate_iface_init)
                         G_IMPLEMENT_INTERFACE (GTK_TYPE_BUILDABLE,
                                                gtk_font_chooser_dialog_buildable_interface_init))

static GtkBuildableIface *parent_buildable_iface;

static void
gtk_font_chooser_dialog_set_property (GObject      *object,
                                      guint         prop_id,
                                      const GValue *value,
                                      GParamSpec   *pspec)
{
  GtkFontChooserDialog *dialog = GTK_FONT_CHOOSER_DIALOG (object);
  GtkFontChooserDialogPrivate *priv = dialog->priv;

  switch (prop_id)
    {
    default:
      g_object_set_property (G_OBJECT (priv->fontchooser), pspec->name, value);
      break;
    }
}

static void
gtk_font_chooser_dialog_get_property (GObject      *object,
                                      guint         prop_id,
                                      GValue       *value,
                                      GParamSpec   *pspec)
{
  GtkFontChooserDialog *dialog = GTK_FONT_CHOOSER_DIALOG (object);
  GtkFontChooserDialogPrivate *priv = dialog->priv;

  switch (prop_id)
    {
    default:
      g_object_get_property (G_OBJECT (priv->fontchooser), pspec->name, value);
      break;
    }
}

static void
font_activated_cb (GtkFontChooser *fontchooser,
                   const gchar    *fontname,
                   gpointer        user_data)
{
  GtkDialog *dialog = user_data;

  gtk_dialog_response (dialog, GTK_RESPONSE_OK);
}

static gboolean
gtk_font_chooser_dialog_key_press_event (GtkWidget   *dialog,
                                         GdkEventKey *event)
{
  GtkFontChooserDialog *fdialog = GTK_FONT_CHOOSER_DIALOG (dialog);
  gboolean handled = FALSE;

  handled = GTK_WIDGET_CLASS (gtk_font_chooser_dialog_parent_class)->key_press_event (dialog, event);

  if (!handled)
    handled = gtk_font_chooser_widget_handle_event (fdialog->priv->fontchooser, event);

  return handled;
}

static void
update_tweak_button (GtkFontChooserDialog *dialog)
{
  GtkFontChooserLevel level;

  if (!dialog->priv->tweak_button)
    return;

  g_object_get (dialog->priv->fontchooser, "level", &level, NULL);
  if ((level & (GTK_FONT_CHOOSER_LEVEL_VARIATIONS | GTK_FONT_CHOOSER_LEVEL_FEATURES)) != 0)
    gtk_widget_show (dialog->priv->tweak_button);
  else
    gtk_widget_hide (dialog->priv->tweak_button);
}

static void
setup_tweak_button (GtkFontChooserDialog *dialog)
{
  gboolean use_header;

  if (dialog->priv->tweak_button)
    return;

  g_object_get (dialog, "use-header-bar", &use_header, NULL);
  if (use_header)
    {
      GtkWidget *button;
      GtkWidget *image;
      GtkWidget *header;
      GActionGroup *actions;

      actions = G_ACTION_GROUP (g_simple_action_group_new ());
      g_action_map_add_action (G_ACTION_MAP (actions), gtk_font_chooser_widget_get_tweak_action (dialog->priv->fontchooser));
      gtk_widget_insert_action_group (GTK_WIDGET (dialog), "font", actions);
      g_object_unref (actions);

      button = gtk_toggle_button_new ();
      gtk_actionable_set_action_name (GTK_ACTIONABLE (button), "font.tweak");
      gtk_widget_set_focus_on_click (button, FALSE);
      gtk_widget_set_valign (button, GTK_ALIGN_CENTER);

      image = gtk_image_new_from_icon_name ("emblem-system-symbolic", GTK_ICON_SIZE_BUTTON);
      gtk_widget_show (image);
      gtk_container_add (GTK_CONTAINER (button), image);

      header = gtk_dialog_get_header_bar (GTK_DIALOG (dialog));
      gtk_header_bar_pack_end (GTK_HEADER_BAR (header), button);

      dialog->priv->tweak_button = button;
      update_tweak_button (dialog);
    }
}

static void
gtk_font_chooser_dialog_map (GtkWidget *widget)
{
  GtkFontChooserDialog *dialog = GTK_FONT_CHOOSER_DIALOG (widget);

  setup_tweak_button (dialog);

  GTK_WIDGET_CLASS (gtk_font_chooser_dialog_parent_class)->map (widget);
}

static void
gtk_font_chooser_dialog_class_init (GtkFontChooserDialogClass *klass)
{
  GObjectClass *gobject_class = G_OBJECT_CLASS (klass);
  GtkWidgetClass *widget_class = GTK_WIDGET_CLASS (klass);

  gobject_class->get_property = gtk_font_chooser_dialog_get_property;
  gobject_class->set_property = gtk_font_chooser_dialog_set_property;

  widget_class->key_press_event = gtk_font_chooser_dialog_key_press_event;
  widget_class->map = gtk_font_chooser_dialog_map;

  _gtk_font_chooser_install_properties (gobject_class);

  /* Bind class to template
   */
  gtk_widget_class_set_template_from_resource (widget_class,
                                               "/org/gtk/libgtk/ui/gtkfontchooserdialog.ui");

  gtk_widget_class_bind_template_child_private (widget_class, GtkFontChooserDialog, fontchooser);
  gtk_widget_class_bind_template_child_private (widget_class, GtkFontChooserDialog, select_button);
  gtk_widget_class_bind_template_child_private (widget_class, GtkFontChooserDialog, cancel_button);
  gtk_widget_class_bind_template_callback (widget_class, font_activated_cb);
}

static void
update_button (GtkFontChooserDialog *dialog)
{
  GtkFontChooserDialogPrivate *priv = dialog->priv;
  PangoFontDescription *desc;

  desc = gtk_font_chooser_get_font_desc (GTK_FONT_CHOOSER (priv->fontchooser));

  gtk_widget_set_sensitive (priv->select_button, desc != NULL);

  if (desc)
    pango_font_description_free (desc);
}

static void
gtk_font_chooser_dialog_init (GtkFontChooserDialog *fontchooserdiag)
{
  GtkFontChooserDialogPrivate *priv;

  fontchooserdiag->priv = gtk_font_chooser_dialog_get_instance_private (fontchooserdiag);
  priv = fontchooserdiag->priv;

  gtk_widget_init_template (GTK_WIDGET (fontchooserdiag));
  gtk_dialog_set_use_header_bar_from_setting (GTK_DIALOG (fontchooserdiag));

G_GNUC_BEGIN_IGNORE_DEPRECATIONS
  gtk_dialog_set_alternative_button_order (GTK_DIALOG (fontchooserdiag),
                                           GTK_RESPONSE_OK,
                                           GTK_RESPONSE_CANCEL,
                                           -1);
G_GNUC_END_IGNORE_DEPRECATIONS
  _gtk_font_chooser_set_delegate (GTK_FONT_CHOOSER (fontchooserdiag),
                                  GTK_FONT_CHOOSER (priv->fontchooser));

  g_signal_connect_swapped (priv->fontchooser, "notify::font-desc",
                            G_CALLBACK (update_button), fontchooserdiag);
  update_button (fontchooserdiag);

  g_signal_connect_swapped (priv->fontchooser, "notify::level",
                            G_CALLBACK (update_tweak_button), fontchooserdiag);
}

/**
 * gtk_font_chooser_dialog_new:
 * @title: (allow-none): Title of the dialog, or %NULL
 * @parent: (allow-none): Transient parent of the dialog, or %NULL
 *
 * Creates a new #GtkFontChooserDialog.
 *
 * Returns: a new #GtkFontChooserDialog
 *
 * Since: 3.2
 */
GtkWidget*
gtk_font_chooser_dialog_new (const gchar *title,
                             GtkWindow   *parent)
{
  GtkFontChooserDialog *dialog;

  dialog = g_object_new (GTK_TYPE_FONT_CHOOSER_DIALOG,
                         "title", title,
                         "transient-for", parent,
                         NULL);

  return GTK_WIDGET (dialog);
}

static void
gtk_font_chooser_dialog_buildable_interface_init (GtkBuildableIface *iface)
{
  parent_buildable_iface = g_type_interface_peek_parent (iface);
  iface->get_internal_child = gtk_font_chooser_dialog_buildable_get_internal_child;
}

static GObject *
gtk_font_chooser_dialog_buildable_get_internal_child (GtkBuildable *buildable,
                                                      GtkBuilder   *builder,
                                                      const gchar  *childname)
{
  GtkFontChooserDialogPrivate *priv;

  priv = GTK_FONT_CHOOSER_DIALOG (buildable)->priv;

  if (g_strcmp0 (childname, "select_button") == 0)
    return G_OBJECT (priv->select_button);
  else if (g_strcmp0 (childname, "cancel_button") == 0)
    return G_OBJECT (priv->cancel_button);

  return parent_buildable_iface->get_internal_child (buildable, builder, childname);
}


static void
cancelled_cb (GCancellable *cancellable,
              GtkDialog    *dialog)
{
  gtk_dialog_response (dialog, GTK_RESPONSE_CANCEL);
}

static void
response_cb (GtkDialog *dialog,
             int        response,
             GTask     *task)
{
  GCancellable *cancellable = g_task_get_cancellable (task);

  if (cancellable)
    g_signal_handlers_disconnect_by_func (cancellable, cancelled_cb, dialog);

  if (response == GTK_RESPONSE_OK)
    g_task_return_boolean (task, TRUE);
  else
    g_task_return_new_error (task, G_IO_ERROR, G_IO_ERROR_CANCELLED, "Cancelled");

  g_object_unref (task);
  gtk_window_destroy (GTK_WINDOW (dialog));
}

/**
 * gtk_choose_font:
 * @parent: (nullable): parent window
 * @title: title for the font chooser
 * @cancellable: (nullable): a `GCancellable` to cancel the operation
 * @callback: (scope async): callback to call when the action is complete
 * @user_data: (closure callback): data to pass to @callback
 *
 * This function presents a font chooser to let the user
 * pick a font.
 *
 * The @callback will be called when the dialog is closed.
 * It should call [function@Gtk.choose_font_finish] to
 * find out whether the operation was completed successfully,
 * and use [class@Gtk.FontChooser] API to obtain the results.
 */
void
gtk_choose_font (GtkWindow           *parent,
                 const char          *title,
                 GCancellable        *cancellable,
                 GAsyncReadyCallback  callback,
                 gpointer             user_data)
{
  gtk_choose_font_full (parent, title, NULL, NULL, cancellable, callback, user_data);
}

/**
 * gtk_choose_font_full:
 * @parent: (nullable): parent window
 * @title: title for the font chooser
 * @prepare: (nullable) (scope call): callback to set up the font chooser
 * @prepare_data: (closure prepare): data to pass to @prepare
 * @cancellable: (nullable): a `GCancellable` to cancel the operation
 * @callback: (scope async): callback to call when the action is complete
 * @user_data: (closure callback): data to pass to @callback
 *
 * This function presents a font chooser to let the user
 * choose a font.
 *
 * In addition to [function@Gtk.choose_font], this function takes
 * a @prepare callback that lets you set up the font chooser according
 * to your needs.
 *
 * The @callback will be called when the dialog is closed.
 * It should use [function@Gtk.choose_font_finish] to find
 * out whether the operation was completed successfully,
 * and use [class@Gtk.FontChooser] API to obtain the results.
 */
void
gtk_choose_font_full (GtkWindow                      *parent,
                      const char                     *title,
                      GtkFontChooserPrepareCallback   prepare,
                      gpointer                        prepare_data,
                      GCancellable                   *cancellable,
                      GAsyncReadyCallback             callback,
                      gpointer                        user_data)
{
  GtkWidget *dialog;
  GTask *task;

  dialog = gtk_font_chooser_dialog_new (title, parent);
  if (prepare)
    prepare (GTK_FONT_CHOOSER (dialog), prepare);

  if (cancellable)
    g_signal_connect (cancellable, "cancelled", G_CALLBACK (cancelled_cb), dialog);

  task = g_task_new (dialog, cancellable, callback, user_data);
  g_task_set_source_tag (task, gtk_choose_font_full);

  g_signal_connect (dialog, "response", G_CALLBACK (response_cb), task);

  gtk_window_present (GTK_WINDOW (dialog));
}

/**
 * gtk_choose_font_finish:
 * @chooser: the `GtkFontChooser`
 * @result: `GAsyncResult` that was passed to @callback
 * @error: return location for an error
 *
 * Finishes a gtk_choose_font() or gtk_choose_font_full() call
 * and returns whether the operation was successful.
 *
 * If this function returns `TRUE`, you can use
 * [class@Gtk.FontChooser] API to get the results.
 *
 * Returns: `TRUE` if the operation was successful
 */
gboolean
gtk_choose_font_finish (GtkFontChooser  *chooser,
                        GAsyncResult     *result,
                        GError          **error)
{
  return g_task_propagate_boolean (G_TASK (result), error);
}
