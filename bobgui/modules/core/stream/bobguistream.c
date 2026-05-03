#include "bobguistream.h"

struct _BobguiStream {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiStream, bobgui_stream, G_TYPE_OBJECT)

static void bobgui_stream_init (BobguiStream *self) {
}

static void bobgui_stream_class_init (BobguiStreamClass *klass) {
}
