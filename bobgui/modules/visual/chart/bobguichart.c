#include "bobguichart.h"

struct _BobguiChartView {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiChartView, bobgui_chart_view, G_TYPE_OBJECT)

static void bobgui_chart_view_init (BobguiChartView *self) {
}

static void bobgui_chart_view_class_init (BobguiChartViewClass *klass) {
}

BobguiChartView * bobgui_chart_view_new (BobguiChartType type) { return NULL; }
