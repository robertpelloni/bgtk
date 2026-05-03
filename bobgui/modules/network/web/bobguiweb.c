#include "bobguiweb.h"

struct _BobguiWebView {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiWebView, bobgui_web_view, G_TYPE_OBJECT)

static void bobgui_web_view_init (BobguiWebView *self) {
}

static void bobgui_web_view_class_init (BobguiWebViewClass *klass) {
}

BobguiWebView * bobgui_web_view_new (void) { return g_object_new (BOBGUI_TYPE_WEB_VIEW, NULL); }
