#include "bobguidsl.h"
struct _BobguiDslEngine {
  int dummy;
};
BobguiDslEngine * bobgui_dsl_engine_new (void) { return g_new0 (BobguiDslEngine, 1); }
