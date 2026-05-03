#include "bobguiiot.h"

struct _BobguiIotManager {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiIotManager, bobgui_iot_manager, G_TYPE_OBJECT)

static void bobgui_iot_manager_init (BobguiIotManager *self) {
}

static void bobgui_iot_manager_class_init (BobguiIotManagerClass *klass) {
}
