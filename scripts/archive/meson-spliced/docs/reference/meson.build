toml_conf = configuration_data()
toml_conf.set('version', meson.project_version())

gidocgen = find_program('gi-docgen', required: get_option('documentation'), native: true)

gidocgen_common_args = [
  '--quiet',
  '--no-namespace-dir',
]

if get_option('werror')
  gidocgen_common_args += ['--fatal-warnings']
endif

docs_dir = bobgui_datadir / 'doc'

if get_option('documentation') and not build_gir
  error('API reference requires introspection.')
endif

subdir('gdk')
subdir('gsk')
subdir('bobgui')
if get_option('gtk_doc')
    glib_prefix = dependency('glib-2.0').get_pkgconfig_variable('prefix')
    glib_docpath = join_paths(glib_prefix, 'share', 'gtk-doc', 'html')

    cairo_prefix = dependency('cairo-gobject').get_pkgconfig_variable('prefix')
    cairo_docpath = join_paths(cairo_prefix, 'share', 'gtk-doc', 'html', 'cairo')

    gdkpixbuf_prefix = dependency('gdk-pixbuf-2.0').get_pkgconfig_variable('prefix')
    gdkpixbuf_docpath = join_paths(gdkpixbuf_prefix, 'share', 'gtk-doc', 'html', 'gdk-pixbuf')

    docpath = join_paths(gtk_datadir, 'gtk-doc', 'html')

    version_conf = configuration_data()
    version_conf.set('GTK_VERSION', meson.project_version())

    src_dir_conf = configuration_data()
    src_dir_conf.set('abs_top_srcdir', meson.source_root())
endif

subdir('gdk')
subdir('gtk')
subdir('libgail-util')
