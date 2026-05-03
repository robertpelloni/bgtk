#include "bobguicompute.h"

struct _BobguiComputeContext {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiComputeContext, bobgui_compute_context, G_TYPE_OBJECT)

static void bobgui_compute_context_init (BobguiComputeContext *self) {
}

static void bobgui_compute_context_class_init (BobguiComputeContextClass *klass) {
}

BobguiComputeContext * bobgui_compute_context_new (const char *backend) { return NULL; }
