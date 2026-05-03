#include "bobguitimeline.h"

struct _BobguiTimeline {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiTimeline, bobgui_timeline, G_TYPE_OBJECT)

static void bobgui_timeline_init (BobguiTimeline *self) {
}

static void bobgui_timeline_class_init (BobguiTimelineClass *klass) {
}

BobguiTimeline * bobgui_timeline_new (void) { return g_object_new (BOBGUI_TYPE_TIMELINE, NULL); }
void bobgui_timeline_play(BobguiTimeline *self) {}
void bobgui_timeline_pause(BobguiTimeline *self) {}
void bobgui_timeline_seek(BobguiTimeline *self, double timeSeconds) {}
