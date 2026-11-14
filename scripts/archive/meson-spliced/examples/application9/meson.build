app9_resources = gnome.compile_resources('exampleapp9_resources',
  'exampleapp.gresource.xml',
  source_dir: meson.current_source_dir())

app9_schemas = gnome.compile_schemas()

executable('exampleapp9',
  'main.c',
  'exampleapp.c',
  'exampleappwin.c',
  'exampleappprefs.c',
  app9_resources,
  app9_schemas,
  dependencies: libbobgui_dep,
  c_args: common_cflags)

app9_resources = gnome.compile_resources(
  'exampleapp9 resources',
  'exampleapp.gresource.xml',
  source_dir: '.'
)

app9_schemas = gnome.compile_schemas(depend_files: files(['org.gtk.exampleapp.gschema.xml']))


app9 = executable(
  'exampleapp9',
  'main.c',
  'exampleapp.c',
  'exampleapp.h',
  'exampleappwin.c',
  'exampleappwin.h',
  'exampleappprefs.c',
  'exampleappprefs.h',
  app9_resources,
  app9_schemas,
  dependencies: libgtk_dep
)

install_data('org.gtk.exampleapp.gschema.xml', install_dir: gtk_schemasdir)
