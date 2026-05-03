#include "bobguilayout.h"

struct _BobguiModernLayout {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiModernLayout, bobgui_modern_layout, G_TYPE_OBJECT)
static void bobgui_modern_layout_init (BobguiModernLayout *s) {}
static void bobgui_modern_layout_class_init (BobguiModernLayoutClass *k) {}
BobguiModernLayout * bobgui_modern_layout_new (BobguiLayoutType t) { return g_object_new (BOBGUI_TYPE_MODERN_LAYOUT, NULL); }

void bobgui_modern_layout_set_flex_direction (BobguiModernLayout *self, const char *direction) {}
void bobgui_modern_layout_set_justify_content (BobguiModernLayout *self, const char *justify) {}
void bobgui_modern_layout_set_align_items (BobguiModernLayout *self, const char *align) {}

void bobgui_modern_layout_set_grid_template_columns (BobguiModernLayout *self, const char *template) {}
void bobgui_modern_layout_set_grid_template_rows    (BobguiModernLayout *self, const char *template) {}

void bobgui_container_set_layout_manager (BobguiWidget *container, BobguiModernLayout *mgr) {}
