#include "bobguivfs.h"

struct _BobguiVFS {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiVFS, bobgui_vfs, G_TYPE_OBJECT)

static void bobgui_vfs_init (BobguiVFS *self) {
}

static void bobgui_vfs_class_init (BobguiVFSClass *klass) {
}
