app2_resources = gnome.compile_resources('exampleapp2_resources',
  'exampleapp.gresource.xml',
  source_dir: meson.current_source_dir())

executable('exampleapp2',
  'exampleapp.c', 'exampleappwin.c', 'main.c', app2_resources,
  dependencies: libbobgui_dep,
  c_args: common_cflags)

app2_resources = gnome.compile_resources(
  'exampleapp2 resources',
  'exampleapp.gresource.xml',
  source_dir: '.'
)

app2 = executable(
  'exampleapp2',
  'main.c',
  'exampleapp.c',
  'exampleapp.h',
  'exampleappwin.c',
  'exampleappwin.h',
  app2_resources,
  dependencies: libgtk_dep
)
