/* GTK - The GIMP Toolkit
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
 * Modified by the GTK+ Team and others 1997-2000.  See the AUTHORS
 * file for a list of people on the GTK+ Team.  See the ChangeLog
 * files for a list of changes.  These files are distributed with
 * GTK+ at ftp://ftp.gtk.org/pub/gtk/. 
 */

#include "config.h"

#include "gtktogglebutton.h"

#include "gtkbuttonprivate.h"
#include "gtktogglebuttonprivate.h"
#include "gtklabel.h"
#include "gtkmain.h"
#include "gtkmarshalers.h"
#include "deprecated/gtktoggleaction.h"
#include "deprecated/gtkactivatable.h"
#include "gtkprivate.h"
#include "gtkintl.h"
#include "a11y/gtktogglebuttonaccessible.h"


/**
 * SECTION:gtktogglebutton
 * @Short_description: Create buttons which retain their state
 * @Title: GtkToggleButton
 * @See_also: #GtkButton, #GtkCheckButton, #GtkCheckMenuItem
 *
 * A #GtkToggleButton is a #GtkButton which will remain “pressed-in” when
 * clicked. Clicking again will cause the toggle button to return to its
 * normal state.
 *
 * A toggle button is created by calling either gtk_toggle_button_new() or
 * gtk_toggle_button_new_with_label(). If using the former, it is advisable to
 * pack a widget, (such as a #GtkLabel and/or a #GtkImage), into the toggle
 * button’s container. (See #GtkButton for more information).
 *
 * The state of a #GtkToggleButton can be set specifically using
 * gtk_toggle_button_set_active(), and retrieved using
 * gtk_toggle_button_get_active().
 *
 * To simply switch the state of a toggle button, use gtk_toggle_button_toggled().
 *
 * # CSS nodes
 *
 * GtkToggleButton has a single CSS node with name button. To differentiate
 * it from a plain #GtkButton, it gets the .toggle style class.
 *
 * ## Creating two #GtkToggleButton widgets.
 *
 * |[<!-- language="C" -->
 * static void output_state (GtkToggleButton *source, gpointer user_data) {
 *   printf ("Active: %d\n", gtk_toggle_button_get_active (source));
 * }
 *
 * void make_toggles (void) {
 *   GtkWidget *window, *toggle1, *toggle2;
 *   GtkWidget *box;
 *   const char *text;
 *
 *   window = gtk_window_new (GTK_WINDOW_TOPLEVEL);
 *   box = gtk_box_new (GTK_ORIENTATION_VERTICAL, 12);
 *
 *   text = "Hi, I’m a toggle button.";
 *   toggle1 = gtk_toggle_button_new_with_label (text);
 *
 *   // Makes this toggle button invisible
 *   gtk_toggle_button_set_mode (GTK_TOGGLE_BUTTON (toggle1),
 *                               TRUE);
 *
 *   g_signal_connect (toggle1, "toggled",
 *                     G_CALLBACK (output_state),
 *                     NULL);
 *   gtk_container_add (GTK_CONTAINER (box), toggle1);
 *
 *   text = "Hi, I’m a toggle button.";
 *   toggle2 = gtk_toggle_button_new_with_label (text);
 *   gtk_toggle_button_set_mode (GTK_TOGGLE_BUTTON (toggle2),
 *                               FALSE);
 *   g_signal_connect (toggle2, "toggled",
 *                     G_CALLBACK (output_state),
 *                     NULL);
 *   gtk_container_add (GTK_CONTAINER (box), toggle2);
 *
 *   gtk_container_add (GTK_CONTAINER (window), box);
 *   gtk_widget_show_all (window);
 * }
 * ]|
 */


enum {
  PROP_0,
  PROP_ACTIVE,
  PROP_INCONSISTENT,
  PROP_DRAW_INDICATOR,
  NUM_PROPERTIES
};

static GParamSpec *toggle_button_props[NUM_PROPERTIES] = { NULL, };

static gboolean gtk_toggle_button_mnemonic_activate  (GtkWidget            *widget,
                                                      gboolean              group_cycling);
static void gtk_toggle_button_clicked       (GtkButton            *button);
static void gtk_toggle_button_set_property  (GObject              *object,
                                             guint                 prop_id,
                                             const GValue         *value,
                                             GParamSpec           *pspec);
static void gtk_toggle_button_get_property  (GObject              *object,
					     guint                 prop_id,
					     GValue               *value,
					     GParamSpec           *pspec);


static void gtk_toggle_button_activatable_interface_init (GtkActivatableIface  *iface);
static void gtk_toggle_button_update                 (GtkActivatable       *activatable,
                                                      GtkAction            *action,
                                                      const gchar          *property_name);
static void gtk_toggle_button_sync_action_properties (GtkActivatable       *activatable,
                                                      GtkAction            *action);

static GtkActivatableIface *parent_activatable_iface;

G_GNUC_BEGIN_IGNORE_DEPRECATIONS;
G_DEFINE_TYPE_WITH_CODE (GtkToggleButton, gtk_toggle_button, GTK_TYPE_BUTTON,
                         G_ADD_PRIVATE (GtkToggleButton)
			 G_IMPLEMENT_INTERFACE (GTK_TYPE_ACTIVATABLE,
						gtk_toggle_button_activatable_interface_init))
G_GNUC_END_IGNORE_DEPRECATIONS;

static void
gtk_toggle_button_class_init (GtkToggleButtonClass *class)
{
  GObjectClass *gobject_class;
  GtkWidgetClass *widget_class;
  GtkButtonClass *button_class;

  gobject_class = G_OBJECT_CLASS (class);
  widget_class = (GtkWidgetClass*) class;
  button_class = (GtkButtonClass*) class;

  gobject_class->set_property = gtk_toggle_button_set_property;
  gobject_class->get_property = gtk_toggle_button_get_property;

  widget_class->mnemonic_activate = gtk_toggle_button_mnemonic_activate;

  button_class->clicked = gtk_toggle_button_clicked;

  class->toggled = NULL;

  toggle_button_props[PROP_ACTIVE] =
      g_param_spec_boolean ("active",
                            P_("Active"),
                            P_("If the toggle button should be pressed in"),
                            FALSE,
                            GTK_PARAM_READWRITE|G_PARAM_EXPLICIT_NOTIFY);

  toggle_button_props[PROP_INCONSISTENT] =
      g_param_spec_boolean ("inconsistent",
                            P_("Inconsistent"),
                            P_("If the toggle button is in an \"in between\" state"),
                            FALSE,
                            GTK_PARAM_READWRITE|G_PARAM_EXPLICIT_NOTIFY);

  toggle_button_props[PROP_DRAW_INDICATOR] =
      g_param_spec_boolean ("draw-indicator",
                            P_("Draw Indicator"),
                            P_("If the toggle part of the button is displayed"),
                            FALSE,
                            GTK_PARAM_READWRITE);

  g_object_class_install_properties (gobject_class, NUM_PROPERTIES, toggle_button_props);

  /**
   * GtkToggleButton::toggled:
   * @togglebutton: the object which received the signal.
   *
   * Should be connected if you wish to perform an action whenever the
   * #GtkToggleButton's state is changed.
   */
  toggle_button_signals[TOGGLED] =
    g_signal_new (I_("toggled"),
		  G_OBJECT_CLASS_TYPE (gobject_class),
		  G_SIGNAL_RUN_FIRST,
		  G_STRUCT_OFFSET (GtkToggleButtonClass, toggled),
		  NULL, NULL,
		  NULL,
		  G_TYPE_NONE, 0);

  gtk_widget_class_set_accessible_type (widget_class, GTK_TYPE_TOGGLE_BUTTON_ACCESSIBLE);
  gtk_widget_class_set_css_name (widget_class, "button");
}

static void
gtk_toggle_button_init (GtkToggleButton *toggle_button)
{
  GtkStyleContext *context;

  toggle_button->priv = gtk_toggle_button_get_instance_private (toggle_button);
  toggle_button->priv->active = FALSE;
  toggle_button->priv->draw_indicator = FALSE;

  context = gtk_widget_get_style_context (GTK_WIDGET (toggle_button));
  gtk_style_context_add_class (context, "toggle");
}

static void
gtk_toggle_button_activatable_interface_init (GtkActivatableIface *iface)
{
  parent_activatable_iface = g_type_interface_peek_parent (iface);
  iface->update = gtk_toggle_button_update;
  iface->sync_action_properties = gtk_toggle_button_sync_action_properties;
}

static void
gtk_toggle_button_update (GtkActivatable *activatable,
                          GtkAction      *action,
                          const gchar    *property_name)
{
  GtkToggleButton *button;

  parent_activatable_iface->update (activatable, action, property_name);

  button = GTK_TOGGLE_BUTTON (activatable);

  if (strcmp (property_name, "active") == 0)
    {
      G_GNUC_BEGIN_IGNORE_DEPRECATIONS;
      gtk_action_block_activate (action);
      gtk_toggle_button_set_active (button, gtk_toggle_action_get_active (GTK_TOGGLE_ACTION (action)));
      gtk_action_unblock_activate (action);
      G_GNUC_END_IGNORE_DEPRECATIONS;
    }

}

static void
gtk_toggle_button_sync_action_properties (GtkActivatable *activatable,
                                          GtkAction      *action)
{
  GtkToggleButton *button;
  gboolean is_toggle_action;

  parent_activatable_iface->sync_action_properties (activatable, action);

  G_GNUC_BEGIN_IGNORE_DEPRECATIONS;
  is_toggle_action = GTK_IS_TOGGLE_ACTION (action);
  G_GNUC_END_IGNORE_DEPRECATIONS;

  if (!is_toggle_action)
    return;

  button = GTK_TOGGLE_BUTTON (activatable);

  G_GNUC_BEGIN_IGNORE_DEPRECATIONS;
  gtk_action_block_activate (action);
  gtk_toggle_button_set_active (button, gtk_toggle_action_get_active (GTK_TOGGLE_ACTION (action)));
  gtk_action_unblock_activate (action);
  G_GNUC_END_IGNORE_DEPRECATIONS;
}

/**
 * gtk_toggle_button_new:
 *
 * Creates a new toggle button. A widget should be packed into the button, as in gtk_button_new().
 *
 * Returns: a new toggle button.
 */
GtkWidget*
gtk_toggle_button_new (void)
{
  return g_object_new (GTK_TYPE_TOGGLE_BUTTON, NULL);
}

/**
 * gtk_toggle_button_new_with_label:
 * @label: a string containing the message to be placed in the toggle button.
 *
 * Creates a new toggle button with a text label.
 *
 * Returns: a new toggle button.
 */
GtkWidget*
gtk_toggle_button_new_with_label (const gchar *label)
{
  return g_object_new (GTK_TYPE_TOGGLE_BUTTON,
                       "label", label,
                       NULL);
}

/**
 * gtk_toggle_button_new_with_mnemonic:
 * @label: the text of the button, with an underscore in front of the
 *         mnemonic character
 *
 * Creates a new #GtkToggleButton containing a label. The label
 * will be created using gtk_label_new_with_mnemonic(), so underscores
 * in @label indicate the mnemonic for the button.
 *
 * Returns: a new #GtkToggleButton
 */
GtkWidget*
gtk_toggle_button_new_with_mnemonic (const gchar *label)
{
  return g_object_new (GTK_TYPE_TOGGLE_BUTTON, 
		       "label", label, 
		       "use-underline", TRUE, 
		       NULL);
}

static void
gtk_toggle_button_set_property (GObject      *object,
				guint         prop_id,
				const GValue *value,
				GParamSpec   *pspec)
{
  GtkToggleButton *tb;

  tb = GTK_TOGGLE_BUTTON (object);

  switch (prop_id)
    {
    case PROP_DRAW_INDICATOR:
      gtk_button_set_indicator_style (GTK_BUTTON (tb),
                                      g_value_get_boolean (value)
                                              ? GTK_INDICATOR_STYLE_CHECK
                                              : GTK_INDICATOR_STYLE_PLAIN);
      break;
    default:
      G_OBJECT_WARN_INVALID_PROPERTY_ID (object, prop_id, pspec);
      break;
    }
}

static void
gtk_toggle_button_get_property (GObject      *object,
                                guint         prop_id,
                                GValue       *value,
                                GParamSpec   *pspec)
{
  switch (prop_id)
    {
    case PROP_DRAW_INDICATOR:
      g_value_set_boolean (value, gtk_button_get_indicator_style (GTK_BUTTON (object)) != GTK_INDICATOR_STYLE_PLAIN);
      break;
    default:
      G_OBJECT_WARN_INVALID_PROPERTY_ID (object, prop_id, pspec);
      break;
    }
}

/**
 * gtk_toggle_button_set_mode:
 * @button: a #GtkToggleButton
 * @draw_indicator: if %TRUE, draw the button as a separate indicator
 * and label; if %FALSE, draw the button like a normal button
 *
 * Sets whether the button is displayed as a separate indicator and label.
 * You can call this function on a checkbutton or a radiobutton with
 * @draw_indicator = %FALSE to make the button look like a normal button.
 *
 * This can be used to create linked strip of buttons that work like
 * a #GtkStackSwitcher.
 *
 * This function only affects instances of classes like #GtkCheckButton
 * and #GtkRadioButton that derive from #GtkToggleButton,
 * not instances of #GtkToggleButton itself.
 *
 * Deprecated:3.0: Use gtk_button_set_indicator_style() instead
 */
void
gtk_toggle_button_set_mode (GtkToggleButton *button,
                            gboolean         draw_indicator)
{
  g_return_if_fail (GTK_IS_TOGGLE_BUTTON (button));

  g_return_if_fail (GTK_IS_TOGGLE_BUTTON (toggle_button));

  priv = toggle_button->priv;

  draw_indicator = draw_indicator ? TRUE : FALSE;

  if (priv->draw_indicator != draw_indicator)
    {
      priv->draw_indicator = draw_indicator;

      if (gtk_widget_get_visible (GTK_WIDGET (toggle_button)))
	gtk_widget_queue_resize (GTK_WIDGET (toggle_button));

      g_object_notify_by_pspec (G_OBJECT (toggle_button), toggle_button_props[PROP_DRAW_INDICATOR]);
    }
}

/**
 * gtk_toggle_button_get_mode:
 * @button: a #GtkToggleButton
 *
 * Retrieves whether the button is displayed as a separate indicator
 * and label. See gtk_toggle_button_set_mode().
 *
 * Returns: %TRUE if the togglebutton is drawn as a separate indicator
 *   and label.
 *
 * Deprecated:3.0: Use gtk_button_get_indicator_style() instead
 **/
gboolean
gtk_toggle_button_get_mode (GtkToggleButton *button)
{
  g_return_val_if_fail (GTK_IS_TOGGLE_BUTTON (button), FALSE);

  return gtk_button_get_indicator_style (GTK_BUTTON (button)) != GTK_INDICATOR_STYLE_PLAIN;
}

/**
 * gtk_toggle_button_set_active:
 * @toggle_button: a #GtkToggleButton.
 * @is_active: %TRUE or %FALSE.
 *
 * Sets the status of the toggle button. Set to %TRUE if you want the
 * GtkToggleButton to be “pressed in”, and %FALSE to raise it.
 * This action causes the #GtkToggleButton::toggled signal and the
 * #GtkButton::clicked signal to be emitted.
 */
void
gtk_toggle_button_set_active (GtkToggleButton *button,
			      gboolean         is_active)
{
  g_return_if_fail (GTK_IS_TOGGLE_BUTTON (button));

  g_return_if_fail (GTK_IS_TOGGLE_BUTTON (toggle_button));

  priv = toggle_button->priv;

  is_active = is_active != FALSE;

  if (priv->active != is_active)
    {
      gtk_button_clicked (GTK_BUTTON (toggle_button));
      g_object_notify_by_pspec (G_OBJECT (toggle_button), toggle_button_props[PROP_ACTIVE]);
    }
}

void
_gtk_toggle_button_set_active (GtkToggleButton *toggle_button,
                               gboolean         is_active)
{
  toggle_button->priv->active = is_active;

  if (is_active)
    gtk_widget_set_state_flags (GTK_WIDGET (toggle_button), GTK_STATE_FLAG_CHECKED, FALSE);
  else
    gtk_widget_unset_state_flags (GTK_WIDGET (toggle_button), GTK_STATE_FLAG_CHECKED);

}

/**
 * gtk_toggle_button_get_active:
 * @toggle_button: a #GtkToggleButton.
 *
 * Queries a #GtkToggleButton and returns its current state. Returns %TRUE if
 * the toggle button is pressed in and %FALSE if it is raised.
 *
 * Returns: a #gboolean value.
 */
gboolean
gtk_toggle_button_get_active (GtkToggleButton *button)
{
  g_return_val_if_fail (GTK_IS_TOGGLE_BUTTON (button), FALSE);

  return gtk_button_get_active (GTK_BUTTON (button));
}

/**
 * gtk_toggle_button_toggled:
 * @toggle_button: a #GtkToggleButton.
 *
 * Emits the #GtkToggleButton::toggled signal on the
 * #GtkToggleButton. There is no good reason for an
 * application ever to call this function.
 */
void
gtk_toggle_button_toggled (GtkToggleButton *button)
{
  g_return_if_fail (GTK_IS_TOGGLE_BUTTON (button));

  gtk_button_toggled (GTK_BUTTON (button));
}

/**
 * gtk_toggle_button_set_inconsistent:
 * @button: a #GtkToggleButton
 * @is_inconsistent: %TRUE if state is inconsistent
 *
 * If the user has selected a range of elements (such as some text or
 * spreadsheet cells) that are affected by a toggle button, and the
 * current values in that range are inconsistent, you may want to
 * display the toggle in an “in between” state. This function turns on
 * “in between” display.  Normally you would turn off the inconsistent
 * state again if the user toggles the toggle button. This has to be
 * done manually, gtk_toggle_button_set_inconsistent() only affects
 * visual appearance, it doesn’t affect the semantics of the button.
 * 
 **/
void
gtk_toggle_button_set_inconsistent (GtkToggleButton *button,
                                    gboolean         is_inconsistent)
{
  GtkToggleButtonPrivate *priv;
  GtkToggleButtonPrivate *priv = gtk_toggle_button_get_instance_private (toggle_button);
  GtkToggleButtonPrivate *group_priv;

  g_return_if_fail (GTK_IS_TOGGLE_BUTTON (toggle_button));

  priv = toggle_button->priv;

  setting = setting != FALSE;

  if (setting != priv->inconsistent)
    {
      priv->inconsistent = setting;

      if (setting)
        gtk_widget_set_state_flags (GTK_WIDGET (toggle_button), GTK_STATE_FLAG_INCONSISTENT, FALSE);
      else
        gtk_widget_unset_state_flags (GTK_WIDGET (toggle_button), GTK_STATE_FLAG_INCONSISTENT);

      g_object_notify_by_pspec (G_OBJECT (toggle_button), toggle_button_props[PROP_INCONSISTENT]);
    }

  if (priv->group_next == group)
    return;

  group_priv = gtk_toggle_button_get_instance_private (group);

  priv->group_prev = NULL;
  if (group_priv->group_prev)
    {
      GtkToggleButtonPrivate *prev = gtk_toggle_button_get_instance_private (group_priv->group_prev);

      prev->group_next = toggle_button;
      priv->group_prev = group_priv->group_prev;
    }

  group_priv->group_prev = toggle_button;
  priv->group_next = group;

  g_object_notify_by_pspec (G_OBJECT (toggle_button), toggle_button_props[PROP_GROUP]);
}

/**
 * gtk_toggle_button_get_inconsistent:
 * @button: a #GtkToggleButton
 * 
 * Gets the value set by gtk_toggle_button_set_inconsistent().
 * 
 * Returns: %TRUE if the button is displayed as inconsistent, %FALSE otherwise
 **/
gboolean
gtk_toggle_button_get_inconsistent (GtkToggleButton *button)
{
  g_return_val_if_fail (GTK_IS_TOGGLE_BUTTON (button), FALSE);

  return gtk_button_get_inconsistent (GTK_BUTTON (button));
}

static void
gtk_toggle_button_get_preferred_width (GtkWidget *widget,
                                       gint      *minimum,
                                       gint      *natural)
{
  if (gtk_button_get_indicator_style (GTK_BUTTON (widget)) != GTK_INDICATOR_STYLE_PLAIN)
    {
      GtkWidget *child;
      gint indicator_size;
      gint indicator_spacing;
      gint focus_width;
      gint focus_pad;
      guint border_width;

      border_width = gtk_container_get_border_width (GTK_CONTAINER (widget));

      gtk_widget_style_get (GTK_WIDGET (widget),
                            "focus-line-width", &focus_width,
                            "focus-padding", &focus_pad,
                            "indicator-size", &indicator_size,
                            "indicator-spacing", &indicator_spacing,
                            NULL);
      *minimum = 2 * border_width;
      *natural = 2 * border_width;

      child = gtk_bin_get_child (GTK_BIN (widget));
      if (child && gtk_widget_get_visible (child))
        {
          gint child_min, child_nat;

          gtk_widget_get_preferred_width (child, &child_min, &child_nat);

          *minimum += child_min + indicator_spacing;
          *natural += child_nat + indicator_spacing;
        }

      *minimum += (indicator_size + indicator_spacing * 2 + 2 * (focus_width + focus_pad));
      *natural += (indicator_size + indicator_spacing * 2 + 2 * (focus_width + focus_pad));
    }
  else
    GTK_WIDGET_CLASS (gtk_toggle_button_parent_class)->get_preferred_width (widget, minimum, natural);
}

static void
gtk_toggle_button_get_preferred_height (GtkWidget *widget,
                                        gint      *minimum,
                                        gint      *natural)
{
  if (gtk_button_get_indicator_style (GTK_BUTTON (widget)) != GTK_INDICATOR_STYLE_PLAIN)
    {
      GtkWidget *child;
      gint temp;
      gint indicator_size;
      gint indicator_spacing;
      gint focus_width;
      gint focus_pad;
      guint border_width;

      border_width = gtk_container_get_border_width (GTK_CONTAINER (widget));

      gtk_widget_style_get (GTK_WIDGET (widget),
                            "focus-line-width", &focus_width,
                            "focus-padding", &focus_pad,
                            "indicator-size", &indicator_size,
                            "indicator-spacing", &indicator_spacing,
                            NULL);

      *minimum = border_width * 2;
      *natural = border_width * 2;

      child = gtk_bin_get_child (GTK_BIN (widget));
      if (child && gtk_widget_get_visible (child))
        {
          gint child_min, child_nat;

          gtk_widget_get_preferred_height (child, &child_min, &child_nat);

          *minimum += child_min;
          *natural += child_nat;
        }

      temp = indicator_size + indicator_spacing * 2;
      *minimum = MAX (*minimum, temp) + 2 * (focus_width + focus_pad);
      *natural = MAX (*natural, temp) + 2 * (focus_width + focus_pad);
    }
  else
    GTK_WIDGET_CLASS (gtk_toggle_button_parent_class)->get_preferred_height (widget, minimum, natural);
}

static void
gtk_toggle_button_size_allocate (GtkWidget     *widget,
                                 GtkAllocation *allocation)
{
  GtkAllocation child_allocation;

  if (gtk_button_get_indicator_style (GTK_BUTTON (widget)) != GTK_INDICATOR_STYLE_PLAIN)
    {
      GtkWidget *child;
      gint indicator_size;
      gint indicator_spacing;
      gint focus_width;
      gint focus_pad;

      gtk_widget_style_get (widget,
                            "focus-line-width", &focus_width,
                            "focus-padding", &focus_pad,
                            "indicator-size", &indicator_size,
                            "indicator-spacing", &indicator_spacing,
                            NULL);

      gtk_widget_set_allocation (widget, allocation);

      if (gtk_widget_get_realized (widget))
        gdk_window_move_resize (gtk_button_get_event_window (GTK_BUTTON (widget)),
                                allocation->x, allocation->y,
                                allocation->width, allocation->height);

      child = gtk_bin_get_child (GTK_BIN (widget));
      if (child && gtk_widget_get_visible (child))
        {
          GtkRequisition child_requisition;
          guint border_width = gtk_container_get_border_width (GTK_CONTAINER (widget));

          gtk_widget_get_preferred_size (child, &child_requisition, NULL);
          child_allocation.width = MIN (child_requisition.width,
                                        allocation->width -
                                        ((border_width + focus_width + focus_pad) * 2
                                        + indicator_size + indicator_spacing * 3));
          child_allocation.width = MAX (child_allocation.width, 1);

          child_allocation.height = MIN (child_requisition.height,
                                         allocation->height - (border_width + focus_width + focus_pad) * 2);
          child_allocation.height = MAX (child_allocation.height, 1);

          child_allocation.x = (border_width + indicator_size + indicator_spacing * 3 +
                                allocation->x + focus_width + focus_pad);
          child_allocation.y = allocation->y + (allocation->height - child_allocation.height) / 2;

          if (gtk_widget_get_direction (widget) == GTK_TEXT_DIR_RTL)
            child_allocation.x = allocation->x + allocation->width
              - (child_allocation.x - allocation->x + child_allocation.width);

          gtk_widget_size_allocate (child, &child_allocation);
        }
    }
  else
    GTK_WIDGET_CLASS (gtk_toggle_button_parent_class)->size_allocate (widget, allocation);
}

static void
gtk_toggle_button_draw_indicator (GtkButton *button,
                                  cairo_t         *cr)
{
  GtkWidget *widget = GTK_WIDGET (button);
  GtkWidget *child;
  GtkStateType state_type;
  GtkShadowType shadow_type;
  gint x, y;
  gint indicator_size;
  gint indicator_spacing;
  gint focus_width;
  gint focus_pad;
  guint border_width;
  gboolean interior_focus;

  gtk_widget_style_get (widget,
                        "interior-focus", &interior_focus,
                        "focus-line-width", &focus_width,
                        "focus-padding", &focus_pad,
                        "indicator-size", &indicator_size,
                        "indicator-spacing", &indicator_spacing,
                        NULL);

  border_width = gtk_container_get_border_width (GTK_CONTAINER (widget));

  x = indicator_spacing + border_width;
  y = (gtk_widget_get_allocated_height (widget) - indicator_size) / 2;

  child = gtk_bin_get_child (GTK_BIN (widget));
  if (!interior_focus || !(child && gtk_widget_get_visible (child)))
    x += focus_width + focus_pad;

  if (gtk_widget_get_direction (widget) == GTK_TEXT_DIR_RTL)
    x = gtk_widget_get_allocated_width (widget) - (indicator_size + x);

  if (gtk_button_get_inconsistent (button))
    shadow_type = GTK_SHADOW_ETCHED_IN;
  else if (gtk_button_get_active (button))
    shadow_type = GTK_SHADOW_IN;
  else
    shadow_type = GTK_SHADOW_OUT;

  if (GTK_BUTTON (widget)->priv->activate_timeout ||
      (GTK_BUTTON (widget)->priv->button_down && GTK_BUTTON (widget)->priv->in_button))
    state_type = GTK_STATE_ACTIVE;
  else if (GTK_BUTTON (widget)->priv->in_button)
    state_type = GTK_STATE_PRELIGHT;
  else if (!gtk_widget_is_sensitive (widget))
    state_type = GTK_STATE_INSENSITIVE;
  else
    state_type = GTK_STATE_NORMAL;

  if (gtk_widget_get_state (widget) == GTK_STATE_PRELIGHT)
    gtk_paint_flat_box (gtk_widget_get_style (widget), cr, GTK_STATE_PRELIGHT,
                        GTK_SHADOW_ETCHED_OUT,
                        widget, "checkbutton",
                        border_width, border_width,
                        gtk_widget_get_allocated_width (widget) - (2 * border_width),
                        gtk_widget_get_allocated_height (widget) - (2 * border_width));

  if (gtk_button_get_indicator_style (GTK_BUTTON (widget)) == GTK_INDICATOR_STYLE_CHECK)
    gtk_paint_check (gtk_widget_get_style (widget), cr,
                     state_type, shadow_type,
                     widget, "checkbutton",
                     x, y, indicator_size, indicator_size);
  else
    gtk_paint_option (gtk_widget_get_style (widget), cr,
                      state_type, shadow_type,
                      widget, "radiobutton",
                      x, y, indicator_size, indicator_size);
}

static gboolean
gtk_toggle_button_mnemonic_activate (GtkWidget *widget,
                                     gboolean   group_cycling)
{
  /*
   * We override the standard implementation in 
   * gtk_widget_real_mnemonic_activate() in order to focus the widget even
   * if there is no mnemonic conflict.
   */
  if (gtk_widget_get_can_focus (widget))
    gtk_widget_grab_focus (widget);

  if (!group_cycling)
    gtk_widget_activate (widget);

  return TRUE;
}

static void
gtk_toggle_button_clicked (GtkButton *button)
{
  GtkToggleButton *toggle_button = GTK_TOGGLE_BUTTON (button);
  GtkToggleButtonPrivate *priv = toggle_button->priv;

  _gtk_toggle_button_set_active (toggle_button, !priv->active);

  gtk_toggle_button_toggled (toggle_button);

  g_object_notify_by_pspec (G_OBJECT (toggle_button), toggle_button_props[PROP_ACTIVE]);

  if (GTK_BUTTON_CLASS (gtk_toggle_button_parent_class)->clicked)
    GTK_BUTTON_CLASS (gtk_toggle_button_parent_class)->clicked (button);
}

