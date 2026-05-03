#include "bobguiscript.h"

struct _BobguiScriptContext {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiScriptContext, bobgui_script_context, G_TYPE_OBJECT)

static void bobgui_script_context_init (BobguiScriptContext *self) {
}

static void bobgui_script_context_class_init (BobguiScriptContextClass *klass) {
}

BobguiScriptContext * bobgui_script_context_new (const char *engine_type) { return NULL; }
