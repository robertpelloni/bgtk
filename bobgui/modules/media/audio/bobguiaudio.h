#ifndef BOBGUI_AUDIO_H
#define BOBGUI_AUDIO_H

#include <glib-object.h>

G_BEGIN_DECLS

#define BOBGUI_TYPE_AUDIO_PROCESSOR (bobgui_audio_processor_get_type ())
G_DECLARE_DERIVABLE_TYPE (BobguiAudioProcessor, bobgui_audio_processor, BOBGUI, AUDIO_PROCESSOR, GObject)

struct _BobguiAudioProcessorClass {
  GObjectClass parent_class;
};

typedef struct _BobguiAudioDeviceManager BobguiAudioDeviceManager;

BobguiAudioDeviceManager * bobgui_audio_device_manager_new (void);

G_END_DECLS

#endif
void bobgui_audio_device_manager_init_devices(BobguiAudioDeviceManager *self, int input_channels, int output_channels);
void bobgui_audio_device_manager_set_sample_rate(BobguiAudioDeviceManager *self, double sample_rate);
void bobgui_audio_device_manager_play(BobguiAudioDeviceManager *self);
void bobgui_audio_device_manager_stop(BobguiAudioDeviceManager *self);
BobguiAudioProcessor * bobgui_audio_processor_new(const char *name);
void bobgui_audio_processor_set_parameter(BobguiAudioProcessor *self, int index, float value);
float bobgui_audio_processor_get_parameter(BobguiAudioProcessor *self, int index);
void bobgui_audio_processor_process_block(BobguiAudioProcessor *self, float** in_channels, float** out_channels, int num_samples);
