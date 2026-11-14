#include <bobgui/bobgui.h>

G_GNUC_BEGIN_IGNORE_DEPRECATIONS

static GdkContentProvider *
prepare (BobguiDragSource *source,
         double         x,
         double         y,
         BobguiWidget     *row)
{
  return gdk_content_provider_new_typed (BOBGUI_TYPE_LIST_BOX_ROW, row);
  GtkWidget *row;
  GtkAllocation alloc;
  cairo_surface_t *surface;
  cairo_t *cr;
  int x, y;
  double sx, sy;

  row = gtk_widget_get_ancestor (widget, GTK_TYPE_LIST_BOX_ROW);
  gtk_widget_get_allocation (row, &alloc);
  surface = cairo_image_surface_create (CAIRO_FORMAT_ARGB32, alloc.width, alloc.height);
  cr = cairo_create (surface);

  gtk_style_context_add_class (gtk_widget_get_style_context (row), "drag-icon");
  gtk_widget_draw (row, cr);
  gtk_style_context_remove_class (gtk_widget_get_style_context (row), "drag-icon");

  gtk_widget_translate_coordinates (widget, row, 0, 0, &x, &y);
  cairo_surface_get_device_scale (surface, &sx, &sy);
  cairo_surface_set_device_offset (surface, -x * sy, -y * sy);
  gtk_drag_set_icon_surface (context, surface);

  cairo_destroy (cr);
  cairo_surface_destroy (surface);

  g_object_set_data (G_OBJECT (gtk_widget_get_parent (row)), "drag-row", row);
  gtk_style_context_add_class (gtk_widget_get_style_context (row), "drag-row");
}

static void
drag_begin (BobguiDragSource *source,
            GdkDrag       *drag,
            BobguiWidget      *widget)
{
  BobguiWidget *row;
  BobguiAllocation alloc;
  GdkPaintable *paintable;
  double x, y;

  row = bobgui_widget_get_ancestor (widget, BOBGUI_TYPE_LIST_BOX_ROW);
  bobgui_widget_get_allocation (row, &alloc);

  paintable = bobgui_widget_paintable_new (row);
  bobgui_widget_translate_coordinates (widget, row, 0, 0, &x, &y);
  bobgui_drag_source_set_icon (source, paintable, -x, -y);

  g_object_unref (paintable);
}

static gboolean
drag_drop (BobguiDropTarget *dest,
           const GValue  *value,
           double         x,
           double         y,
           gpointer       data)
{
  BobguiWidget *target = data;
  BobguiWidget *source;
  int pos;

  source = g_value_get_object (value);
  if (source == NULL)
    return FALSE;

  pos = bobgui_list_box_row_get_index (BOBGUI_LIST_BOX_ROW (target));
  if (source == target)
    return FALSE;

  g_object_ref (source);
  bobgui_box_remove (BOBGUI_BOX (bobgui_widget_get_parent (source)), source);
  bobgui_list_box_insert (BOBGUI_LIST_BOX (bobgui_widget_get_parent (target)), source, pos);
  g_object_unref (source);

  return TRUE;
}

static BobguiWidget *
create_row (const char *text)
{
  BobguiWidget *row, *box, *label, *image;
  BobguiDragSource *source;
  BobguiDropTarget *dest;

  row = bobgui_list_box_row_new (); 
  image = bobgui_image_new_from_icon_name ("open-menu-symbolic");
  box = bobgui_box_new (BOBGUI_ORIENTATION_HORIZONTAL, 10);
  g_object_set (box, "margin-start", 10, "margin-end", 10, NULL);
  label = bobgui_label_new (text);
  bobgui_list_box_row_set_child (BOBGUI_LIST_BOX_ROW (row), box);
  bobgui_widget_set_hexpand (label, TRUE);
  bobgui_box_append (BOBGUI_BOX (box), label);
  bobgui_box_append (BOBGUI_BOX (box), image);

  source = bobgui_drag_source_new ();
  bobgui_drag_source_set_actions (source, GDK_ACTION_MOVE);
  g_signal_connect (source, "drag-begin", G_CALLBACK (drag_begin), image);
  g_signal_connect (source, "prepare", G_CALLBACK (prepare), row);
  bobgui_widget_add_controller (image, BOBGUI_EVENT_CONTROLLER (source));

  dest = bobgui_drop_target_new (BOBGUI_TYPE_LIST_BOX_ROW, GDK_ACTION_MOVE);
  g_signal_connect (dest, "drop", G_CALLBACK (drag_drop), row);
  bobgui_widget_add_controller (BOBGUI_WIDGET (row), BOBGUI_EVENT_CONTROLLER (dest));

  return row;
}

static void
on_row_activated (BobguiListBox *self,
                  BobguiWidget  *child)
{
  const char *id;
  id = g_object_get_data (G_OBJECT (bobgui_list_box_row_get_child (BOBGUI_LIST_BOX_ROW (child))), "id");
  g_message ("Row activated %p: %s", child, id);
}

static void
on_selected_children_changed (BobguiListBox *self)
{
  g_message ("Selection changed");
}

static void
selection_mode_changed (BobguiComboBox *combo, gpointer data)
{
  BobguiListBox *list = data;

  bobgui_list_box_set_selection_mode (list, bobgui_combo_box_get_active (combo));
}

static const char *css =
  ".during-dnd { "
  "  background: white; "
  "  border: 1px solid black; "
  "}";

static const char *css =
  ".row:not(:first-child) { "
  "  border-top: 1px solid alpha(gray,0.5); "
  "  border-bottom: 1px solid transparent; "
  "}"
  ".row:first-child { "
  "  border-top: 1px solid transparent; "
  "  border-bottom: 1px solid transparent; "
  "}"
  ".row:last-child { "
  "  border-top: 1px solid alpha(gray,0.5); "
  "  border-bottom: 1px solid alpha(gray,0.5); "
  "}"
  ".row.drag-icon { "
  "  background: white; "
  "  border: 1px solid black; "
  "}"
  ".row.drag-row { "
  "  color: gray; "
  "  background: alpha(gray,0.2); "
  "}"
  ".row.drag-row.drag-hover { "
  "  border-top: 1px solid #4e9a06; "
  "  border-bottom: 1px solid #4e9a06; "
  "}"
  ".row.drag-hover image, "
  ".row.drag-hover label { "
  "  color: #4e9a06; "
  "}"
  ".row.drag-hover-top {"
  "  border-top: 1px solid #4e9a06; "
  "}"
  ".row.drag-hover-bottom {"
  "  border-bottom: 1px solid #4e9a06; "
  "}"
;

int
main (int argc, char *argv[])
{
  BobguiWidget *window, *list, *sw, *row;
  BobguiWidget *hbox, *vbox, *combo, *button;
  int i;
  char *text;
  BobguiCssProvider *provider;

  bobgui_init ();

  provider = bobgui_css_provider_new ();
  bobgui_css_provider_load_from_data (provider, css, -1);
  bobgui_style_context_add_provider_for_display (gdk_display_get_default (), BOBGUI_STYLE_PROVIDER (provider), 800);
  window = bobgui_window_new ();
  bobgui_window_set_default_size (BOBGUI_WINDOW (window), -1, 300);

  hbox = bobgui_box_new (BOBGUI_ORIENTATION_HORIZONTAL, 12);
  bobgui_window_set_child (BOBGUI_WINDOW (window), hbox);
  vbox = bobgui_box_new (BOBGUI_ORIENTATION_VERTICAL, 6);
  bobgui_widget_set_margin_start (vbox, 12);
  bobgui_widget_set_margin_end (vbox, 12);
  bobgui_widget_set_margin_top (vbox, 12);
  bobgui_widget_set_margin_bottom (vbox, 12);
  bobgui_box_append (BOBGUI_BOX (hbox), vbox);

  list = bobgui_list_box_new ();
  bobgui_list_box_set_selection_mode (BOBGUI_LIST_BOX (list), BOBGUI_SELECTION_NONE);

  gtk_drag_dest_set (list, GTK_DEST_DEFAULT_MOTION | GTK_DEST_DEFAULT_DROP, entries, 1, GDK_ACTION_MOVE);
  g_signal_connect (list, "drag-data-received", G_CALLBACK (drag_data_received), NULL);
  g_signal_connect (list, "drag-motion", G_CALLBACK (drag_motion), NULL);
  g_signal_connect (list, "drag-leave", G_CALLBACK (drag_leave), NULL);

  g_signal_connect (list, "row-activated", G_CALLBACK (on_row_activated), NULL);
  g_signal_connect (list, "selected-rows-changed", G_CALLBACK (on_selected_children_changed), NULL);

  sw = bobgui_scrolled_window_new ();
  bobgui_widget_set_hexpand (sw, TRUE);
  bobgui_scrolled_window_set_policy (BOBGUI_SCROLLED_WINDOW (sw), BOBGUI_POLICY_NEVER, BOBGUI_POLICY_ALWAYS);
  bobgui_box_append (BOBGUI_BOX (hbox), sw);
  bobgui_scrolled_window_set_child (BOBGUI_SCROLLED_WINDOW (sw), list);

  button = bobgui_check_button_new_with_label ("Activate on single click");
  g_object_bind_property (list, "activate-on-single-click",
                          button, "active",
                          G_BINDING_BIDIRECTIONAL | G_BINDING_SYNC_CREATE);
  bobgui_box_append (BOBGUI_BOX (vbox), button);

  combo = bobgui_combo_box_text_new ();
  bobgui_combo_box_text_append_text (BOBGUI_COMBO_BOX_TEXT (combo), "None");
  bobgui_combo_box_text_append_text (BOBGUI_COMBO_BOX_TEXT (combo), "Single");
  bobgui_combo_box_text_append_text (BOBGUI_COMBO_BOX_TEXT (combo), "Browse");
  bobgui_combo_box_text_append_text (BOBGUI_COMBO_BOX_TEXT (combo), "Multiple");
  g_signal_connect (combo, "changed", G_CALLBACK (selection_mode_changed), list);
  bobgui_box_append (BOBGUI_BOX (vbox), combo);
  bobgui_combo_box_set_active (BOBGUI_COMBO_BOX (combo), bobgui_list_box_get_selection_mode (BOBGUI_LIST_BOX (list)));

  for (i = 0; i < 20; i++)
    {
      text = g_strdup_printf ("Row %d", i);
      row = create_row (text);
      bobgui_list_box_insert (BOBGUI_LIST_BOX (list), row, -1);
    }

  bobgui_window_present (BOBGUI_WINDOW (window));

  while (TRUE)
    g_main_context_iteration (NULL, TRUE);

  return 0;
}
