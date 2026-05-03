#include "bobguireflect.h"

struct _BobguiReflectContext {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiReflectContext, bobgui_reflect_context, G_TYPE_OBJECT)

static void bobgui_reflect_context_init (BobguiReflectContext *self) {
}

static void bobgui_reflect_context_class_init (BobguiReflectContextClass *klass) {
}
