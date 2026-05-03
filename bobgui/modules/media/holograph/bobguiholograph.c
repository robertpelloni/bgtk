#include "bobguiholograph.h"

struct _BobguiHologram {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiHologram, bobgui_hologram
bobgui_holograph_view, G_TYPE_OBJECT)

static void bobgui_hologram
bobgui_holograph_view_init (BobguiHologram *self) {
}

static void bobgui_hologram
bobgui_holograph_view_class_init (BobguiHologramClass *klass) {
}

G_DECLARE_FINAL_TYPE (BobguiHologram, bobgui_hologram, BOBGUI, HOLOGRAM, GObject) { return NULL; }
