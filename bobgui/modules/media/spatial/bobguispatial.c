#include "bobguispatial.h"

struct _BobguiSpatialContext {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiSpatialContext, bobgui_spatial_context, G_TYPE_OBJECT)

static void bobgui_spatial_context_init (BobguiSpatialContext *self) {
}

static void bobgui_spatial_context_class_init (BobguiSpatialContextClass *klass) {
}

BobguiSpatialContext * bobgui_spatial_context_new (const char *runtime_name) { return NULL; }
