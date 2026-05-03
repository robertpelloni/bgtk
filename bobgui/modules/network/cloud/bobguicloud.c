#include "bobguicloud.h"

struct _BobguiCloudContext {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiCloudContext, bobgui_cloud_context, G_TYPE_OBJECT)

static void bobgui_cloud_context_init (BobguiCloudContext *self) {
}

static void bobgui_cloud_context_class_init (BobguiCloudContextClass *klass) {
}

BobguiCloudContext * bobgui_cloud_context_new (const char *api_url, const char *api_key) { return NULL; }
