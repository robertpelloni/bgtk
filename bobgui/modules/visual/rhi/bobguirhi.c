#include "bobguirhi.h"

struct _BobguiRhiDevice {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiRhiDevice, bobgui_rhi_device, G_TYPE_OBJECT)

static void bobgui_rhi_device_init (BobguiRhiDevice *self) {
}

static void bobgui_rhi_device_class_init (BobguiRhiDeviceClass *klass) {
}

BobguiRhiDevice * bobgui_rhi_device_new (BobguiRhiBackend backend) { return NULL; }
