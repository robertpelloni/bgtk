#ifndef BOBGUI_SHADER_H
#define BOBGUI_SHADER_H

#include <glib-object.h>

G_BEGIN_DECLS

#define BOBGUI_TYPE_EFFECT_SHADER (bobgui_effect_shader_get_type ())
G_DECLARE_FINAL_TYPE (BobguiEffectShader, bobgui_effect_shader, BOBGUI, EFFECT_SHADER, GObject)

BobguiEffectShader * bobgui_effect_shader_new_from_source (const char *source);

G_END_DECLS

#endif
void bobgui_effect_shader_set_uniform_float(BobguiEffectShader *self, const char* name, float value);
void bobgui_effect_shader_set_uniform_vec3(BobguiEffectShader *self, const char* name, float x, float y, float z);
