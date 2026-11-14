executable('exampleapp',
  'exampleapp.c', 'exampleappwin.c', 'main.c',
  dependencies: libbobgui_dep,
  c_args: common_cflags)


app1 = executable(
  'exampleapp',
  'main.c',
  'exampleapp.c',
  'exampleapp.h',
  'exampleappwin.c',
  'exampleappwin.h',
  dependencies: libgtk_dep
)
