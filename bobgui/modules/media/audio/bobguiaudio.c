#include "bobguiaudio.h"
G_DEFINE_ABSTRACT_TYPE (BobguiAudioProcessor, bobgui_audio_processor, G_TYPE_OBJECT)
static void bobgui_audio_processor_class_init (BobguiAudioProcessorClass *klass) {}
static void bobgui_audio_processor_init (BobguiAudioProcessor *self) {}
BobguiAudioDeviceManager * bobgui_audio_device_manager_new (void) { return g_new0 (BobguiAudioDeviceManager, 1); }
void bobgui_audio_device_manager_init_devices(BobguiAudioDeviceManager *self, int input_channels, int output_channels) {}
void bobgui_audio_device_manager_set_sample_rate(BobguiAudioDeviceManager *self, double sample_rate) {}
void bobgui_audio_device_manager_play(BobguiAudioDeviceManager *self) {}
void bobgui_audio_device_manager_stop(BobguiAudioDeviceManager *self) {}
BobguiAudioProcessor * bobgui_audio_processor_new(const char *name) { return NULL; }
void bobgui_audio_processor_set_parameter(BobguiAudioProcessor *self, int index, float value) {}
float bobgui_audio_processor_get_parameter(BobguiAudioProcessor *self, int index) { return 0.0f; }
void bobgui_audio_processor_process_block(BobguiAudioProcessor *self, float** in_channels, float** out_channels, int num_samples) {}
