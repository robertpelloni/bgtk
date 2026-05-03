#include "bobguiactionregistry.h"

struct _BobguiActionRegistry {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiActionRegistry, bobgui_action_registry, G_TYPE_OBJECT)

static void bobgui_action_registry_init (BobguiActionRegistry *self) {
}

static void bobgui_action_registry_class_init (BobguiActionRegistryClass *klass) {
}

BobguiActionRegistry * bobgui_action_registry_new              (void) { return g_object_new (BOBGUI_TYPE_ACTION_REGISTRY, NULL); }
