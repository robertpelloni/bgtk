#include "bobguiipc.h"

struct _BobguiIpcManager {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiIpcManager, bobgui_ipc_manager, G_TYPE_OBJECT)

static void bobgui_ipc_manager_init (BobguiIpcManager *self) {
}

static void bobgui_ipc_manager_class_init (BobguiIpcManagerClass *klass) {
}

BobguiIpcManager * bobgui_ipc_manager_new (const char *namespace_id) { return NULL; }
