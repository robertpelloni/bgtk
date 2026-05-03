#include "bobguidata.h"

struct _BobguiSqlDatabase {
  GObject parent_instance;
};

G_DEFINE_TYPE (BobguiSqlDatabase, bobgui_sql_database, G_TYPE_OBJECT)
static void bobgui_sql_database_init (BobguiSqlDatabase *s) {}
static void bobgui_sql_database_class_init (BobguiSqlDatabaseClass *k) {}

BobguiSqlDatabase * bobgui_sql_database_add (const char *d) { return g_object_new (BOBGUI_TYPE_SQL_DATABASE, NULL); }
