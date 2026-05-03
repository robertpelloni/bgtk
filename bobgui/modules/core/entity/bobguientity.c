#include "bobguientity.h"

struct _BobguiEntityWorld {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiEntityWorld, bobgui_entity_world, G_TYPE_OBJECT)

static void bobgui_entity_world_init (BobguiEntityWorld *self) {
}

static void bobgui_entity_world_class_init (BobguiEntityWorldClass *klass) {
}

BobguiEntityWorld * bobgui_entity_world_new (void) { return g_object_new (BOBGUI_TYPE_ENTITY_WORLD, NULL); }
