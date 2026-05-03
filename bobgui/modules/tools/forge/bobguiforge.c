#include "bobguiforge.h"

struct _BobguiForgeContext {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiForgeContext, bobgui_forge_context, G_TYPE_OBJECT)

static void bobgui_forge_context_init (BobguiForgeContext *self) {
}

static void bobgui_forge_context_class_init (BobguiForgeContextClass *klass) {
}

BobguiForgeContext * bobgui_forge_context_new (const char *project_dir) { return NULL; }
