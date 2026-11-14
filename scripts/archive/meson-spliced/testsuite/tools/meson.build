run_command = find_program('run-command', dirs: meson.current_source_dir())

builder_tool_tests = [
  { 'name': 'simplify',
    'executable': get_variable('bobgui4_builder_tool').full_path(),
     'command': '@0@ simplify @1@',
    'inputdir': '@0@/simplify'.format(meson.current_source_dir()),
    'outputdir': '@0@/simplify'.format(meson.current_build_dir()),
    'tests': [
      'test1',
      'test2',
      'test3',
      'test4',
      'test5',
      'test6',
      'test7',
      'test8',
      'test9',
      'test10',
      'test11',
      'test12',
      'test13',
    ],
  },
  { 'name': 'simplify-3to4',
    'executable': get_variable('bobgui4_builder_tool').full_path(),
    'command': '@0@ simplify --3to4 @1@',
    'inputdir': '@0@/simplify-3to4'.format(meson.current_source_dir()),
    'outputdir': '@0@/simplify-3to4'.format(meson.current_build_dir()),
    'tests': [
      'actionbar',
      'assistant',
      'bin',
      'box',
      'checkbutton',
      'dialog',
      'fixed',
      'grid',
      'bobguifilechooserdialog',
      'headerbar',
      'image',
      'nm-crash',
      'notebook',
      'office-runner',
      'overlay',
      'paned',
      'popover',
      'scale',
      'separator',
      'stack',
      'toolbar',
    ],
  },
  { 'name': 'validate',
    'executable': get_variable('bobgui4_builder_tool').full_path(),
    'command': '@0@ validate --deprecations @1@',
    'inputdir': '@0@/validate'.format(meson.current_source_dir()),
    'outputdir': '@0@/validate'.format(meson.current_build_dir()),
    'tests': [
      'invalid1',
      'invalid2',
      'invalid3',
      'invalid4',
      'invalid5',
      'valid1',
      'valid2',
      'valid3',
      'valid4',
      'valid5',
    ],
  },
  { 'name': 'enumerate',
    'executable': get_variable('bobgui4_builder_tool').full_path(),
    'command': '@0@ enumerate --callbacks @1@',
    'inputdir': '@0@/enumerate'.format(meson.current_source_dir()),
    'outputdir': '@0@/enumerate'.format(meson.current_build_dir()),
    'tests': [
      'widget-factory',
      'widget-factory2',
      'widget-factory3',
    ],
  },
  { 'name': 'settings',
    'executable': get_variable('bobgui4_query_settings').full_path(),
    'command': '@0@',
    'inputdir': '@0@/settings'.format(meson.current_source_dir()),
    'outputdir': '@0@/settings'.format(meson.current_build_dir()),
    'tests': [
      'all',
    ],
  }
]

foreach t : builder_tool_tests
  foreach i : t.get('tests')
    inputfile = '@0@/@1@.ui'.format(t.get('inputdir'), i)
    test(t.get('name') + ' ' + i,
         run_command,
         workdir: meson.current_build_dir(),
         protocol: 'exitcode',
         args: [ t.get('command').format(t.get('executable'), inputfile),
                 i,
                 t.get('inputdir'),
                 t.get('outputdir')
               ],
         env: [
          'BOBGUI_A11Y=test',
          'GSETTINGS_BACKEND=memory',
          'GDK_DEBUG=default-settings,no-portals',
         ],
         suite: 'tools',
         )
  endforeach
endforeach
testexecdir = join_paths(installed_test_bindir)
testdatadir = join_paths(installed_test_datadir)

bash = find_program('bash', required : false)
if bash.found()
  test_env = environment()

  foreach t : ['test-simplify', 'test-settings']
    if get_option('installed_tests')
      configure_file(output: t,
                     input: '@0@.in'.format(t),
                     copy: true,
                     install_dir: testexecdir)
    else
      configure_file(output: t,
                     input: '@0@.in'.format(t),
                     copy: true)
    endif
    test(t, bash,
         args: t,
         workdir: meson.current_build_dir(),
         env: [ 'GIO_USE_VOLUME_MONITOR=unix',
                'GSETTINGS_BACKEND=memory',
                'GTK_CSD=1',
                'G_ENABLE_DIAGNOSTIC=0',
                'G_TEST_SRCDIR=@0@'.format(meson.current_source_dir()),
                'G_TEST_BUILDDIR=@0@'.format(meson.current_build_dir()),
                'GTK_BUILDER_TOOL=@0@'.format(gtk_builder_tool.full_path()),
                'GTK_QUERY_SETTINGS=@0@'.format(gtk_query_settings.full_path())
              ],
         suite: 'tools')
  endforeach
endif

if get_option('installed_tests')
  foreach t : ['test-simplify', 'test-settings']
    test_conf = configuration_data()
    test_conf.set('testexecdir', testexecdir)
    test_conf.set('test', t)
    configure_file(output: '@0@.test'.format(t),
                   input: 'tools.test.in',
                   configuration: test_conf,
                   install_dir: testdatadir)
  endforeach

  install_subdir('simplify', install_dir: testexecdir)
endif
