app8_resources = gnome.compile_resources('exampleapp8_resources',
  'exampleapp.gresource.xml',
  source_dir: meson.current_source_dir())

app8_schemas = gnome.compile_schemas()

executable('exampleapp8',
  'main.c',
  'exampleapp.c',
  'exampleappwin.c',
  'exampleappprefs.c',
  app8_resources,
  app8_schemas,
  dependencies: libbobgui_dep,
  c_args: common_cflags)

app8_resources = gnome.compile_resources(
  'exampleapp8 resources',
  'exampleapp.gresource.xml',
  source_dir: '.'
)

app8_schemas = gnome.compile_schemas(depend_files: files(['org.gtk.exampleapp.gschema.xml']))


app8 = executable(
  'exampleapp8',
  'main.c',
  'exampleapp.c',
  'exampleapp.h',
  'exampleappwin.c',
  'exampleappwin.h',
  'exampleappprefs.c',
  'exampleappprefs.h',
  app8_resources,
  app8_schemas,
  dependencies: libgtk_dep
)

install_data('org.gtk.exampleapp.gschema.xml', install_dir: gtk_schemasdir)
