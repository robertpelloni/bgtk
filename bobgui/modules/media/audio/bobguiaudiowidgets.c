#include "bobguiaudiowidgets.h"

struct _BobguiAudioDial {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiAudioDial, bobgui_audio_dial
bobgui_audio_fader
bobgui_audio_scope
bobgui_audio_analyzer, G_TYPE_OBJECT)

static void bobgui_audio_dial
bobgui_audio_fader
bobgui_audio_scope
bobgui_audio_analyzer_init (BobguiAudioDial *self) {
}

static void bobgui_audio_dial
bobgui_audio_fader
bobgui_audio_scope
bobgui_audio_analyzer_class_init (BobguiAudioDialClass *klass) {
}

G_DECLARE_FINAL_TYPE (BobguiAudioDial, bobgui_audio_dial, BOBGUI, AUDIO_DIAL, BobguiWidget) { return NULL; }
