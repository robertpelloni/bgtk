#include "bobguiworkbench.h"

struct _BobguiWorkbench {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiWorkbench, bobgui_workbench, G_TYPE_OBJECT)

static void bobgui_workbench_init (BobguiWorkbench *self) {
}

static void bobgui_workbench_class_init (BobguiWorkbenchClass *klass) {
}

BobguiWorkbench * bobgui_workbench_new               (BobguiApplication *application) { return NULL; }
