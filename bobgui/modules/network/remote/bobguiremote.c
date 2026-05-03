#include "bobguiremote.h"

struct _BobguiRemoteServer {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiRemoteServer, bobgui_remote_server, G_TYPE_OBJECT)

static void bobgui_remote_server_init (BobguiRemoteServer *self) {
}

static void bobgui_remote_server_class_init (BobguiRemoteServerClass *klass) {
}

BobguiRemoteServer * bobgui_remote_server_new (int port, const char *codec) { return NULL; }
