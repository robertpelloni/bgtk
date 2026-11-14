csstest_env = environment()
csstest_env.set('BOBGUI_A11Y', 'test')
csstest_env.set('GSK_RENDERER', 'cairo')
csstest_env.set('G_TEST_SRCDIR', meson.current_source_dir())
csstest_env.set('G_TEST_BUILDDIR', meson.current_build_dir())

subdir('parser')
subdir('nodes')
subdir('style')
subdir('change')

test_api = executable('api',
  sources: ['api.c'],
  c_args: common_cflags,
  dependencies: libbobgui_dep,
)

test('api', test_api,
  args: ['--tap', '-k' ],
  protocol: 'tap',
  env: csstest_env,
  suite: 'css',
)

test_data = executable('data',
  sources: ['data.c'],
  c_args: common_cflags + ['-DBOBGUI_COMPILATION'],
  include_directories: [confinc, ],
  dependencies: libbobgui_static_dep,
)

test('data', test_data,
  args: ['--tap', '-k' ],
  protocol: 'tap',
  env: csstest_env,
  suite: 'css',
)

transition = executable('transition',
  sources: ['transition.c'],
  c_args: common_cflags + ['-DBOBGUI_COMPILATION'],
  dependencies: libbobgui_static_dep,
)

test('transition', transition,
     args: [ '--tap', '-k' ],
     protocol: 'tap',
     env: csstest_env,
     suite: 'css'
)

color = executable('color',
  sources: ['color.c'],
  c_args: common_cflags + ['-DBOBGUI_COMPILATION'],
  dependencies: libbobgui_static_dep,
)

test('color', color,
     args: [ '--tap', '-k'],
     protocol: 'tap',
     env: csstest_env,
     suite: 'css'
)

compute = executable('compute',
  sources: ['compute.c'],
  c_args: common_cflags + ['-DBOBGUI_COMPILATION'],
  dependencies: libbobgui_static_dep
)

test('compute', compute,
  args: [ '--tap', '-k'],
  protocol: 'tap',
  env: csstest_env,
  suite: 'css'
)
subdir('parser')
subdir('nodes')
subdir('style')

testexecdir = join_paths(installed_test_bindir, 'css')
testdatadir = join_paths(installed_test_datadir, 'css')

test_api = executable('api', 'api.c',
                      dependencies: libgtk_dep,
                      install: get_option('installed_tests'),
                      install_dir: testexecdir)
test('css/api', test_api)

if get_option('installed_tests')
  conf = configuration_data()
  conf.set('libexecdir', gtk_libexecdir)
  configure_file(input: 'api.test.in',
                 output: 'api.test',
                 configuration: conf,
                 install_dir: testdatadir)
endif
