#include "bobguireport.h"

struct _BobguiReportDesigner {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiReportDesigner, bobgui_report_designer
bobgui_report_preview, G_TYPE_OBJECT)

static void bobgui_report_designer
bobgui_report_preview_init (BobguiReportDesigner *self) {
}

static void bobgui_report_designer
bobgui_report_preview_class_init (BobguiReportDesignerClass *klass) {
}

G_DECLARE_FINAL_TYPE (BobguiReportDesigner, bobgui_report_designer, BOBGUI, REPORT_DESIGNER, GObject) { return NULL; }
