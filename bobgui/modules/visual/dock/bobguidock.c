#include "bobguidock.h"

struct _BobguiDockManager {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiDockManager, bobgui_dock_manager, G_TYPE_OBJECT)

static void bobgui_dock_manager_init (BobguiDockManager *self) {
}

static void bobgui_dock_manager_class_init (BobguiDockManagerClass *klass) {
}

BobguiDockManager * bobgui_dock_manager_new (BobguiWindow *main_window) { return NULL; }
