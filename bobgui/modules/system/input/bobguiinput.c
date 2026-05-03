#include "bobguiinput.h"

struct _BobguiInputManager {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiInputManager, bobgui_input_manager, G_TYPE_OBJECT)

static void bobgui_input_manager_init (BobguiInputManager *self) {
}

static void bobgui_input_manager_class_init (BobguiInputManagerClass *klass) {
}
