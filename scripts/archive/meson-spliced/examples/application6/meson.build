app6_resources = gnome.compile_resources('exampleapp6_resources',
  'exampleapp.gresource.xml',
  source_dir: meson.current_source_dir())

app6_schemas = gnome.compile_schemas()

executable('exampleapp6',
  'exampleapp.c',
  'exampleappwin.c',
  'main.c',
  'exampleappprefs.c',
  app6_resources,
  app6_schemas,
  dependencies: libbobgui_dep,
  c_args: common_cflags)

app6_resources = gnome.compile_resources(
  'exampleapp6 resources',
  'exampleapp.gresource.xml',
  source_dir: '.'
)

app6_schemas = gnome.compile_schemas(depend_files: files(['org.gtk.exampleapp.gschema.xml']))


app6 = executable(
  'exampleapp6',
  'main.c',
  'exampleapp.c',
  'exampleapp.h',
  'exampleappwin.c',
  'exampleappwin.h',
  'exampleappprefs.c',
  'exampleappprefs.h',
  app6_resources,
  app6_schemas,
  dependencies: libgtk_dep
)

install_data('org.gtk.exampleapp.gschema.xml', install_dir: gtk_schemasdir)
