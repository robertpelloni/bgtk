#include "bobguishader.h"

struct _BobguiEffectShader {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiEffectShader, bobgui_effect_shader, G_TYPE_OBJECT)

static void bobgui_effect_shader_init (BobguiEffectShader *self) {
}

static void bobgui_effect_shader_class_init (BobguiEffectShaderClass *klass) {
}

BobguiEffectShader * bobgui_effect_shader_new_from_source (const char *source) { return NULL; }
void bobgui_effect_shader_set_uniform_float(BobguiEffectShader *self, const char* name, float value) {}
void bobgui_effect_shader_set_uniform_vec3(BobguiEffectShader *self, const char* name, float x, float y, float z) {}
