#include "bobguiplugin.h"

struct _BobguiPluginManager {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiPluginManager, bobgui_plugin_manager, G_TYPE_OBJECT)

static void bobgui_plugin_manager_init (BobguiPluginManager *self) {
}

static void bobgui_plugin_manager_class_init (BobguiPluginManagerClass *klass) {
}

BobguiPluginManager * bobgui_plugin_manager_new (const char *search_path) { return NULL; }
