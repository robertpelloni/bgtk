gdk_wayland_sources = files([
  'gdkapplaunchcontext-wayland.c',
  'gdkcairocontext-wayland.c',
  'gdkclipboard-wayland.c',
  'gdkcursor-wayland.c',
  'gdkdevice-wayland.c',
  'gdkdevicepad-wayland.c',
  'gdkdisplay-wayland.c',
  'gdkdmabuf-wayland.c',
  'gdkdrag-wayland.c',
  'gdkdragsurface-wayland.c',
  'gdkdrop-wayland.c',
  'gdkeventsource.c',
  'gdkglcontext-wayland.c',
  'gdkkeymap-wayland.c',
  'gdkmonitor-wayland.c',
  'gdkprimary-wayland.c',
  'gdkseat-wayland.c',
  'gdksettings-wayland.c',
  'gdkshm.c',
  'gdksurface-wayland.c',
  'gdksubsurface-wayland.c',
  'gdktoplevel-wayland.c',
  'gdkpopup-wayland.c',
  'gdkvulkancontext-wayland.c',
  'gdkwaylandcolor.c',
  'gdkwaylandpresentationtime.c',
  'wm-button-layout-translation.c',
])

gdk_wayland_public_headers = files([
gdk_wayland_sources = files(
  'gdkapplaunchcontext-wayland.c',
  'gdkcursor-wayland.c',
  'gdkdevice-wayland.c',
  'gdkdisplay-wayland.c',
  'gdkdnd-wayland.c',
  'gdkeventsource.c',
  'gdkkeys-wayland.c',
  'gdkglcontext-wayland.c',
  'gdkmonitor-wayland.c',
  'gdkscreen-wayland.c',
  'gdkselection-wayland.c',
  'gdkwindow-wayland.c',
  'wm-button-layout-translation.c',
)

gdk_wayland_public_headers = files(
  'gdkwaylanddevice.h',
  'gdkwaylanddisplay.h',
  'gdkwaylandglcontext.h',
  'gdkwaylandmonitor.h',
  'gdkwaylandpopup.h',
  'gdkwaylandseat.h',
  'gdkwaylandsurface.h',
  'gdkwaylandtoplevel.h',
])

install_headers(gdk_wayland_public_headers, 'gdkwayland.h', subdir: 'bobgui-4.0/gdk/wayland/')
  'gdkwaylandselection.h',
  'gdkwaylandwindow.h',
)

install_headers(gdk_wayland_public_headers, subdir: 'gtk-3.0/gdk/wayland')
install_headers('gdkwayland.h', subdir: 'gtk-3.0/gdk')

gdk_wayland_deps = [
  shmlib,
  xkbdep,
  wlclientdep,
  wlprotocolsdep,
  wlegldep,
]

# Fields:
#  - name: protocol name
#  - stability: protocol stability ('private', 'stable' or 'unstable')
#  - version: protocol version
#  - required: wayland_protocols version check
proto_sources = [
  {
    'name': 'bobgui-shell',
    'stability': 'private',
  },
  {
    'name': 'primary-selection',
    'stability': 'unstable',
    'version': 1,
  },
  {
    'name': 'pointer-gestures',
    'stability': 'unstable',
    'version': 1,
  },
  {
    'name': 'viewporter',
    'stability': 'stable',
  },
  {
    'name': 'xdg-shell',
    'stability': 'unstable',
    'version': 6,
  },
  {
    'name': 'xdg-shell',
    'stability': 'stable',
  },
  {
    'name': 'xdg-foreign',
    'stability': 'unstable',
    'version': 1,
  },
  {
    'name': 'xdg-foreign',
    'stability': 'unstable',
    'version': 2,
  },
  {
    'name': 'tablet',
    'stability': 'stable',
    'version': 2,
  },
  {
    'name': 'keyboard-shortcuts-inhibit',
    'stability': 'unstable',
    'version': 1,
  },
  {
    'name': 'server-decoration',
    'stability': 'private',
  },
  {
    'name': 'xdg-output',
    'stability': 'unstable',
    'version': 1,
  },
  {
    'name': 'idle-inhibit',
    'stability': 'unstable',
    'version': 1,
  },
  {
    'name': 'xdg-activation',
    'stability': 'staging',
    'version': 1,
  },
  {
    'name': 'fractional-scale',
    'stability': 'staging',
    'version': 1,
  },
  {
    'name': 'linux-dmabuf',
    'stability': 'unstable',
    'version': 1,
  },
  {
    'name': 'presentation-time',
    'stability': 'stable',
  },
  {
    'name': 'single-pixel-buffer',
    'stability': 'staging',
    'version': 1,
  },
  {
    'name': 'xdg-dialog',
    'stability': 'staging',
    'version': 1,
  },
  {
    'name': 'color-management',
    'stability': 'staging',
    'version': 1,
  },
  {
    'name': 'xdg-system-bell',
    'stability': 'staging',
    'version': 1,
  },
  {
    'name': 'cursor-shape',
    'stability': 'staging',
    'version': 1,
  },
  {
    'name': 'xdg-toplevel-icon',
    'stability': 'staging',
    'version': 1,
  },
  {
    'name': 'color-representation',
    'stability': 'staging',
    'version': 1,
  },
  {
    'name': 'xdg-session-management',
    'stability': 'staging',
    'version': 1,
  },
  wlcursordep,
  wlegldep,
]

# wayland protocols
proto_dir = dependency('wayland-protocols').get_pkgconfig_variable('pkgdatadir')
assert(proto_dir != '', 'Could not get pkgdatadir from wayland-protocols.pc')

wayland_scanner = find_program('wayland-scanner')

# Format:
#  - protocol name
#  - protocol stability ('private', 'stable' or 'unstable')
#  - protocol version (if stability is 'unstable')
proto_sources = [
  ['gtk-shell', 'private', ],
  ['gtk-primary-selection', 'private', ],
  ['pointer-gestures', 'unstable', 'v1', ],
  ['xdg-shell', 'unstable', 'v6', ],
  ['xdg-shell', 'stable', ],
  ['xdg-foreign', 'unstable', 'v1', ],
  ['tablet', 'unstable', 'v2', ],
  ['keyboard-shortcuts-inhibit', 'unstable', 'v1', ],
  ['server-decoration', 'private' ],
  ['xdg-output', 'unstable', 'v1', ],
  ['primary-selection', 'unstable', 'v1', ],
]

gdk_wayland_gen_headers = []

wlmod = import('unstable-wayland')
foreach p: proto_sources
  if wlprotocolsdep.version().version_compare(p.get('required', '>=0'))
    if p.get('stability') == 'private'
      if (p.has_key('version'))
        proto_file = 'protocol/@0@-v@1@.xml'.format(p.get('name'), p.get('version'))
      else
        proto_file = 'protocol/@0@.xml'.format(p.get('name'))
      endif
    elif p.get('stability') == 'stable'
      if (p.has_key('version'))
        proto_file = wlmod.find_protocol(p.get('name'),
          state: p.get('stability'),
          version: p.get('version'),
        )
      else
        proto_file = wlmod.find_protocol(p.get('name'),
          state: p.get('stability'),
        )
      endif
    else
      proto_file = wlmod.find_protocol(p.get('name'),
        state: p.get('stability'),
        version: p.get('version'),
      )
    endif

    # Returns a list [.c, .h]
    gen = wlmod.scan_xml(proto_file)
    assert(gen.length() == 2)
    gdk_wayland_sources += gen[0]
    gdk_wayland_gen_headers += gen[1]
  endif
endforeach

libgdk_wayland = static_library('gdk-wayland',
  sources: [ gdk_wayland_sources, gdk_wayland_gen_headers, gdk_gen_headers ],
  include_directories: [ confinc, gdkinc, ],
  c_args: [
    '-DBOBGUI_COMPILATION',
    '-DG_LOG_DOMAIN="Gdk"',
  ] + common_cflags,
  dependencies: [ gdk_deps, gdk_wayland_deps ],
)

# Used to generate pkg-config Requires
wayland_public_deps = []

if wlclientdep.type_name() != 'internal'
  wayland_public_deps += [wlclientdep]
endif
foreach p: proto_sources
  proto_name = p.get(0)
  proto_stability = p.get(1)

  if proto_stability == 'stable'
    output_base = proto_name
    input = join_paths(proto_dir, '@0@/@1@/@2@.xml'.format(proto_stability, proto_name, output_base))
  elif proto_stability == 'private'
    output_base = proto_name
    input = 'protocol/@0@.xml'.format(proto_name)
  else
    proto_version = p.get(2)
    output_base = '@0@-@1@-@2@'.format(proto_name, proto_stability, proto_version)
    input = join_paths(proto_dir, '@0@/@1@/@2@.xml'.format(proto_stability, proto_name, output_base))
  endif

  gdk_wayland_gen_headers += custom_target('@0@ client header'.format(output_base),
                                           input: input,
                                           output: '@0@-client-protocol.h'.format(output_base),
                                           command: [
                                             wayland_scanner,
                                             'client-header',
                                             '@INPUT@', '@OUTPUT@',
                                           ])

  gdk_wayland_sources += custom_target('@0@ source'.format(output_base),
                                       input: input,
                                       output: '@0@-protocol.c'.format(output_base),
                                       command: [
                                         wayland_scanner,
                                         'private-code',
                                         '@INPUT@', '@OUTPUT@',
                                       ])
endforeach

libgdk_wayland = static_library('gdk-wayland',
                                gdk_wayland_sources, gdk_wayland_gen_headers, gdkconfig, gdkenum_h,
                                include_directories: [ confinc, gdkinc, ],
                                c_args: [
                                  '-DGDK_COMPILATION',
                                  '-DG_LOG_DOMAIN="Gdk"',
                                ] + common_cflags,
                                link_args: common_ldflags,
                                dependencies: [ gdk_deps, gdk_wayland_deps, ])
