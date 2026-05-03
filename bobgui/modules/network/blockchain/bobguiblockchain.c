#include "bobguiblockchain.h"

struct _BobguiBlockchainManager {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiBlockchainManager, bobgui_blockchain_manager, G_TYPE_OBJECT)

static void bobgui_blockchain_manager_init (BobguiBlockchainManager *self) {
}

static void bobgui_blockchain_manager_class_init (BobguiBlockchainManagerClass *klass) {
}
