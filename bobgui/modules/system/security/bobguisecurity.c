#include "bobguisecurity.h"

struct _BobguiSecureEntry {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiSecureEntry, bobgui_secure_entry, G_TYPE_OBJECT)

static void bobgui_secure_entry_init (BobguiSecureEntry *self) {
}

static void bobgui_secure_entry_class_init (BobguiSecureEntryClass *klass) {
}

BobguiSecureEntry * bobgui_secure_entry_new (void) { return g_object_new (BOBGUI_TYPE_SECURE_ENTRY, NULL); }
