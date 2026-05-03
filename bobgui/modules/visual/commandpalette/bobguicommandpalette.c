#include "bobguicommandpalette.h"

struct _BobguiCommandPalette {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiCommandPalette, bobgui_command_palette, G_TYPE_OBJECT)

static void bobgui_command_palette_init (BobguiCommandPalette *self) {
}

static void bobgui_command_palette_class_init (BobguiCommandPaletteClass *klass) {
}

BobguiCommandPalette * bobgui_command_palette_new               (BobguiApplication *application) { return NULL; }
