examples = [
  'builder',
  'drawing',
  'grid-packing',
  'plugman',
  'search-bar',
  'sunny',
  'window-default',
]

foreach ex : examples
  # TODO: name clash with testsuite/bobgui/builder, rename one of the two
  example_name = ex == 'builder' ? 'builder_example' : ex
  executable(example_name, '@0@.c'.format(ex),
             c_args: [
               '-DBOBGUI_SRCDIR="@0@"'.format(meson.current_source_dir()),
             ] + common_cflags,
             dependencies: libbobgui_dep)
endforeach
builder_example = executable(
  'builder_example',
  'builder.c',
  dependencies: libgtk_dep
)

drawing = executable(
  'drawing',
  'drawing.c',
  dependencies: libgtk_dep
)

grid_packing = executable(
  'grid-packing',
  'grid-packing.c',
  dependencies: libgtk_dep
)

hello_world = executable(
  'hello-world',
  'hello-world.c',
  dependencies: libgtk_dep
)

plugman = executable(
  'plugman',
  'plugman.c',
  dependencies: libgtk_dep
)

search_bar = executable(
  'search-bar',
  'search-bar.c',
  dependencies: libgtk_dep
)

sunny = executable(
  'sunny',
  'sunny.c',
  dependencies: libgtk_dep
)

window_default = executable(
  'window-default',
  'window-default.c',
  dependencies: libgtk_dep
)


subdir('bp')
subdir('application1')
subdir('application2')
subdir('application3')
subdir('application4')
subdir('application5')
subdir('application6')
subdir('application7')
subdir('application8')
subdir('application9')
subdir('application10')
