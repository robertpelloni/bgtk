#include "bobgui3d.h"

struct _Bobgui3dNode {
  GObject parent_instance;
};

G_DEFINE_TYPE (Bobgui3dNode, bobgui_3d_node, G_TYPE_OBJECT)

static void bobgui_3d_node_init (Bobgui3dNode *self) {
}

static void bobgui_3d_node_class_init (Bobgui3dNodeClass *klass) {
}

Bobgui3dNode * bobgui_3d_node_new (void) { return g_object_new (BOBGUI_TYPE_3D_NODE, NULL); }
void bobgui_3d_node_set_translation(Bobgui3dNode *self, float x, float y, float z) {}
void bobgui_3d_node_set_rotation(Bobgui3dNode *self, float pitch, float yaw, float roll) {}
void bobgui_3d_node_set_scale(Bobgui3dNode *self, float scale) {}
void bobgui_3d_node_load_mesh(Bobgui3dNode *self, const char* path) {}
