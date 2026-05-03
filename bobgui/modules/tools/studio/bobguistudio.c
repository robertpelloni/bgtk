#include "bobguistudio.h"

struct _BobguiStudioManager {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiStudioManager, bobgui_studio_manager, G_TYPE_OBJECT)

static void bobgui_studio_manager_init (BobguiStudioManager *self) {
}

static void bobgui_studio_manager_class_init (BobguiStudioManagerClass *klass) {
}
