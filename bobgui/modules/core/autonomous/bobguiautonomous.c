#include "bobguiautonomous.h"

struct _BobguiAutonomousEngine {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiAutonomousEngine, bobgui_autonomous_engine, G_TYPE_OBJECT)

static void bobgui_autonomous_engine_init (BobguiAutonomousEngine *self) {
}

static void bobgui_autonomous_engine_class_init (BobguiAutonomousEngineClass *klass) {
}
