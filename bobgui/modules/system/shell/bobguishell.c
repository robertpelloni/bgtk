#include "bobguishell.h"

struct _BobguiShellManager {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiShellManager, bobgui_shell_manager, G_TYPE_OBJECT)

static void bobgui_shell_manager_init (BobguiShellManager *self) {
}

static void bobgui_shell_manager_class_init (BobguiShellManagerClass *klass) {
}

