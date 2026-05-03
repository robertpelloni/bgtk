#include "bobguisync.h"

struct _BobguiSyncEngine {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiSyncEngine, bobgui_sync_engine, G_TYPE_OBJECT)

static void bobgui_sync_engine_init (BobguiSyncEngine *self) {
}

static void bobgui_sync_engine_class_init (BobguiSyncEngineClass *klass) {
}

BobguiSyncEngine * bobgui_sync_engine_new (BobguiStateMachine *state_machine) { return NULL; }
