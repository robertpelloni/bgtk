#include "bobguifin.h"

struct _BobguiFinEngine {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiFinEngine, bobgui_fin_engine
bobgui_fin_chart, G_TYPE_OBJECT)

static void bobgui_fin_engine
bobgui_fin_chart_init (BobguiFinEngine *self) {
}

static void bobgui_fin_engine
bobgui_fin_chart_class_init (BobguiFinEngineClass *klass) {
}

G_DECLARE_FINAL_TYPE (BobguiFinEngine, bobgui_fin_engine, BOBGUI, FIN_ENGINE, GObject) { return NULL; }
