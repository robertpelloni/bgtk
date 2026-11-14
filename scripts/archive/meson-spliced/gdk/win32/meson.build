gdk_win32_public_sources = files([
  'gdkcursor-win32.c',
  'gdkd3d12texture.c',
  'gdkd3d12texturebuilder.c',
  'gdkdisplay-win32.c',
  'gdkdisplaymanager-win32.c',
  'gdkdrag-win32.c',
  'gdkglcontext-win32.c',
  'gdkglcontext-win32-wgl.c',
  'gdkwin32misc.c',
  'gdksurface-win32.c',
  'gdkevents-win32.c',
  'gdkmonitor-win32.c',
])

gdk_win32_sources = gdk_win32_public_sources + files([
  'gdkcairocontext-win32.c',
  'gdkclipboard-win32.c',
  'gdkclipdrop-win32.c',
  'gdkd3d12utils.c',
gdk_win32_sources = files(
  'gdkcursor-win32.c',
  'gdkdevicemanager-win32.c',
  'gdkdevice-virtual.c',
  'gdkdevice-win32.c',
  'gdkdevice-winpointer.c',
  'gdkdevice-wintab.c',
  'gdkdrop-win32.c',
  'gdkglcontext-win32-wgl-private.c',
  'gdkhdataoutputstream-win32.c',
  'gdkinput-dmanipulation.c',
  'gdkinput-winpointer.c',
  'gdkdisplay-win32.c',
  'gdkdisplaymanager-win32.c',
  'gdkdnd-win32.c',
  'gdkevents-win32.c',
  'gdkgeometry-win32.c',
  'gdkglcontext-win32.c',
  'gdkglobals-win32.c',
  'gdkkeys-win32.c',
  'gdkkeys-win32-impl.c',
  'gdkkeys-win32-impl-wow64.c',
  'gdkmain-win32.c',
  'gdkvulkancontext-win32.c',
  'gdkwin32cursor.h',
  'gdkwin32display.h',
  'gdkwin32messagesource.c',
  'gdkwin32keys.h',
])

gdk_win32_public_headers = files([
  'gdkd3d12texture.h',
  'gdkd3d12texturebuilder.h',
  'gdkmonitor-win32.c',
  'gdkproperty-win32.c',
  'gdkscreen-win32.c',
  'gdkselection-win32.c',
  'gdktestutils-win32.c',
  'gdkwin32id.c',
  'gdkwin32langnotification.c',
  'gdkwindow-win32.c',
)

gdk_win32_public_headers = files(
  'gdkwin32cursor.h',
  'gdkwin32display.h',
  'gdkwin32displaymanager.h',
  'gdkwin32dnd.h',
  'gdkwin32glcontext.h',
  'gdkwin32misc.h',
  'gdkwin32monitor.h',
  'gdkwin32surface.h',
])

install_headers(gdk_win32_public_headers, 'gdkwin32.h', subdir: 'bobgui-4.0/gdk/win32/')

if have_egl
  gdk_win32_sources += ['gdkglcontext-win32-egl.c']
endif

d3d12_dep = [
  cc.find_library('d3d11'),
  cc.find_library('d3d12'),
  cc.find_library('dcomp'),
  cc.find_library('dxgi'),
  cc.find_library('dxguid'),
  dependency('DirectX-Headers', version: '>= 1.611'),
]

gdk_win32_deps = [
  pangowin32_dep, # FIXME
  cc.find_library('hid'),
  cc.find_library('opengl32'),
  cc.find_library('shcore'),
  d3d12_dep,
]

libgdk_win32 = static_library('gdk-win32',
  sources: [ gdk_win32_sources, gdk_gen_headers ],
  include_directories: [ confinc, gdkinc ],
  c_args: [
    '-DBOBGUI_COMPILATION',
    '-DG_LOG_DOMAIN="Gdk"',
    '-DINSIDE_GDK_WIN32',
  ],
  dependencies: [ gdk_deps, gdk_win32_deps ],
)
  'gdkwin32keys.h',
  'gdkwin32misc.h',
  'gdkwin32monitor.h',
  'gdkwin32screen.h',
  'gdkwin32window.h',
)

GDK_WIN32_CFLAGS = ['-DGDK_COMPILATION', '-DG_LOG_DOMAIN="Gdk"', '-DINSIDE_GDK_WIN32']

if win32_has_egl
  GDK_WIN32_CFLAGS += '-DGDK_WIN32_ENABLE_EGL'
endif


install_headers(gdk_win32_public_headers, subdir: 'gtk-3.0/gdk/win32')
install_headers('gdkwin32.h', subdir: 'gtk-3.0/gdk')

gdk_win32_deps = [ # FIXME
    pangowin32_dep,
    meson.get_compiler('c').find_library('hid')
]

libgdk_win32 = static_library('gdk-win32',
  gdk_win32_sources, gdkconfig, gdkenum_h,
  include_directories: [confinc, gdkinc],
  c_args: GDK_WIN32_CFLAGS,
  dependencies: [gdk_deps, gdk_win32_deps])
