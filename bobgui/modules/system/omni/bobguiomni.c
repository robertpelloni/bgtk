#include "bobguiomni.h"

struct _BobguiOmniRuntime {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiOmniRuntime, bobgui_omni_runtime, G_TYPE_OBJECT)

static void bobgui_omni_runtime_init (BobguiOmniRuntime *self) {
}

static void bobgui_omni_runtime_class_init (BobguiOmniRuntimeClass *klass) {
}
