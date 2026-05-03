#include "bobguigraph.h"

struct _BobguiGraphView {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiGraphView, bobgui_graph_view, G_TYPE_OBJECT)

static void bobgui_graph_view_init (BobguiGraphView *self) {
}

static void bobgui_graph_view_class_init (BobguiGraphViewClass *klass) {
}

BobguiGraphView * bobgui_graph_view_new (void) { return g_object_new (BOBGUI_TYPE_GRAPH_VIEW, NULL); }
