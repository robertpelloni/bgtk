#include "bobguitest.h"

struct _BobguiTestRunner {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiTestRunner, bobgui_test_runner, G_TYPE_OBJECT)

static void bobgui_test_runner_init (BobguiTestRunner *self) {
}

static void bobgui_test_runner_class_init (BobguiTestRunnerClass *klass) {
}

BobguiTestRunner * bobgui_test_runner_new (BobguiWidget *root) { return NULL; }
