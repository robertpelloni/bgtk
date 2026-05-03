#include "bobguibrain.h"

struct _BobguiBrainModel {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiBrainModel, bobgui_brain_model, G_TYPE_OBJECT)

static void bobgui_brain_model_init (BobguiBrainModel *self) {
}

static void bobgui_brain_model_class_init (BobguiBrainModelClass *klass) {
}
