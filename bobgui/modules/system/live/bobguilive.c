#include "bobguilive.h"
struct _BobguiLiveContext {
  int dummy;
};
BobguiLiveContext * bobgui_live_context_new (const char *path) { return g_new0 (BobguiLiveContext, 1); }
