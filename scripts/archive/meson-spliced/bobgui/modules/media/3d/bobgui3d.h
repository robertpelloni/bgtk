#ifndef BOBGUI_3D_H
#define BOBGUI_3D_H

#include <glib-object.h>

G_BEGIN_DECLS

#define BOBGUI_TYPE_3D_NODE (bobgui_3d_node_get_type ())
G_DECLARE_FINAL_TYPE (Bobgui3dNode, bobgui_3d_node, BOBGUI, 3D_NODE, GObject)

Bobgui3dNode * bobgui_3d_node_new (void);

G_END_DECLS

#endif
void bobgui_3d_node_set_translation(Bobgui3dNode *self, float x, float y, float z);
void bobgui_3d_node_set_rotation(Bobgui3dNode *self, float pitch, float yaw, float roll);
void bobgui_3d_node_set_scale(Bobgui3dNode *self, float scale);
void bobgui_3d_node_load_mesh(Bobgui3dNode *self, const char* path);
