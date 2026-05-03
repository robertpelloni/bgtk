#include "bobguipackage.h"

struct _BobguiPackageManager {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiPackageManager, bobgui_package_manager, G_TYPE_OBJECT)

static void bobgui_package_manager_init (BobguiPackageManager *self) {
}

static void bobgui_package_manager_class_init (BobguiPackageManagerClass *klass) {
}

