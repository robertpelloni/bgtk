app3_resources = gnome.compile_resources('exampleapp3_resources',
  'exampleapp.gresource.xml',
  source_dir: meson.current_source_dir())

executable('exampleapp3',
  'exampleapp.c', 'exampleappwin.c', 'main.c', app3_resources,
  dependencies: libbobgui_dep,
  c_args: common_cflags)

app3_resources = gnome.compile_resources(
  'exampleapp3 resources',
  'exampleapp.gresource.xml',
  source_dir: '.'
)

app3 = executable(
  'exampleapp3',
  'main.c',
  'exampleapp.c',
  'exampleapp.h',
  'exampleappwin.c',
  'exampleappwin.h',
  app3_resources,
  dependencies: libgtk_dep
)
