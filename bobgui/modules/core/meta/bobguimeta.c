#include "bobguimeta.h"

struct _BobguiMetaOrchestrator {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiMetaOrchestrator, bobgui_meta_orchestrator, G_TYPE_OBJECT)

static void bobgui_meta_orchestrator_init (BobguiMetaOrchestrator *self) {
}

static void bobgui_meta_orchestrator_class_init (BobguiMetaOrchestratorClass *klass) {
}
