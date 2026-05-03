#include "bobguibio.h"

struct _BobguiBioManager {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiBioManager, bobgui_bio_manager
bobgui_bio_viewer, G_TYPE_OBJECT)

static void bobgui_bio_manager
bobgui_bio_viewer_init (BobguiBioManager *self) {
}

static void bobgui_bio_manager
bobgui_bio_viewer_class_init (BobguiBioManagerClass *klass) {
}

G_DECLARE_FINAL_TYPE (BobguiBioManager, bobgui_bio_manager, BOBGUI, BIO_MANAGER, GObject) { return NULL; }
