#include "bobguiphysics.h"

struct _BobguiPhysicsWorld {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiPhysicsWorld, bobgui_physics_world, G_TYPE_OBJECT)

static void bobgui_physics_world_init (BobguiPhysicsWorld *self) {
}

static void bobgui_physics_world_class_init (BobguiPhysicsWorldClass *klass) {
}

BobguiPhysicsWorld * bobgui_physics_world_new (void) { return g_object_new (BOBGUI_TYPE_PHYSICS_WORLD, NULL); }
