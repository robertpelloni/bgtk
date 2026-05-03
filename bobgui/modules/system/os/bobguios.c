#include "bobguios.h"

struct _BobguiVirtualOs {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiVirtualOs, bobgui_virtual_os, G_TYPE_OBJECT)

static void bobgui_virtual_os_init (BobguiVirtualOs *self) {
}

static void bobgui_virtual_os_class_init (BobguiVirtualOsClass *klass) {
}

BobguiVirtualOs * bobgui_virtual_os_new (void) { return g_object_new (BOBGUI_TYPE_VIRTUAL_OS, NULL); }
