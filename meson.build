project('bobgui', 'c',
        version: '4.23.0',
project('gtk', 'c',
        version: '4.14.4',
        default_options: [
          'buildtype=debugoptimized',
          'warning_level=1',
          # We only need c99, but glib needs GNU-specific features
          # https://github.com/mesonbuild/meson/issues/2289
          'c_std=gnu11',
          'cpp_std=c++11',
        ],
        meson_version : '>= 1.5.0',
        license: 'LGPL-2.1-or-later')

glib_major_req = 2
glib_minor_req = 84

pango_major_req = 1
pango_minor_req = 56

# keep these numbers in sync with wrap files where there exist
glib_req           = '>= @0@.@1@'.format(glib_major_req, glib_minor_req)
pango_req          = '>= @0@.@1@'.format(pango_major_req, pango_minor_req)
harfbuzz_req       = '>= 8.4.0'
fribidi_req        = '>= 1.0.6'
cairo_req          = '>= 1.18.2'
gdk_pixbuf_req     = '>= 2.30.0'
wayland_proto_req  = '>= 1.44'
wayland_req        = '>= 1.24.0'
graphene_req       = '>= 1.10.0'
epoxy_req          = '>= 1.4'
cloudproviders_req = '>= 0.3.1'
xkbcommon_req      = '>= 0.2.0'
sysprof_req        = '>= 3.38.0'
vulkan_req         = '>= 1.3'
gstreamer_req      = '>= 1.24.0'
gi_req             = '>= 1.84'

fs = import('fs')
gnome = import('gnome')
pkg_config = import('pkgconfig')

add_project_arguments([
    '-DG_LOG_USE_STRUCTURED=1',
    '-DGLIB_VERSION_MIN_REQUIRED=GLIB_VERSION_@0@_@1@'.format(glib_major_req, glib_minor_req),
    '-DGLIB_VERSION_MAX_ALLOWED=GLIB_VERSION_@0@_@1@'.format(glib_major_req, glib_minor_req),
    '-DPANGO_VERSION_MIN_REQUIRED=G_ENCODE_VERSION(@0@,@1@)'.format(pango_major_req, pango_minor_req),
    '-DPANGO_VERSION_MAX_ALLOWED=G_ENCODE_VERSION(@0@,@1@)'.format(pango_major_req, pango_minor_req),
  ],
  language: 'c'
)

# Making releases:
#  1. new development cycle:
#    a. bobgui_minor_version += 1
#    b. bobgui_micro_version = 0
#  2. new stable cycle:
#    a. bobgui_minor_version += 1
#    b. bobgui_micro_version = 0
#  3. every new release:
#    a. bobgui_micro_version += 1
bobgui_version       = meson.project_version()
bobgui_major_version = bobgui_version.split('.')[0].to_int()
bobgui_minor_version = bobgui_version.split('.')[1].to_int()
bobgui_micro_version = bobgui_version.split('.')[2].to_int()

# Interface age gets reset during development cycles, when
# we add new API; new micro versions of the same minor
# stable cycle will have the same interface age
#
# If new API is added during a stable cycle, reset to 0
bobgui_interface_age = bobgui_minor_version.is_odd() ? 0 : bobgui_micro_version

add_project_arguments('-DBOBGUI_VERSION="@0@"'.format(bobgui_version), language: 'c')

add_project_arguments('-D_GNU_SOURCE', language: 'c')

if host_machine.system() == 'windows'
  add_project_arguments(['-DUNICODE',
                         '-D_UNICODE',
                         '-DCOBJMACROS',
                         '-DWINVER=_WIN32_WINNT_WIN10',
                         '-D_WIN32_WINNT=_WIN32_WINNT_WIN10'], language: 'c')
endif

# Use debug/optimization flags to determine whether to enable debug or disable
# cast checks
bobgui_debug_cflags = []
debug = get_option('debug')
optimization = get_option('optimization')
if debug
  bobgui_debug_cflags += '-DG_ENABLE_DEBUG'
  if optimization in ['0', 'g']
    bobgui_debug_cflags += '-DG_ENABLE_CONSISTENCY_CHECKS'
  endif
elif optimization in ['2', '3', 's']
  bobgui_debug_cflags += ['-DG_DISABLE_CAST_CHECKS', '-DG_DISABLE_ASSERT']
endif

add_project_arguments(bobgui_debug_cflags, language: 'c')
project('gtk+-3.0', 'c',
        version: '3.24.31',
        default_options: [
          'buildtype=debugoptimized',
          'warning_level=1'
        ],
        meson_version : '>= 0.48.0',
        license: 'LGPLv2.1+')

glib_major_req = 2
glib_minor_req = 57
glib_micro_req = 2

if glib_minor_req.is_odd()
  glib_min_required = 'GLIB_VERSION_@0@_@1@'.format(glib_major_req, glib_minor_req - 1)
else
  glib_min_required = 'GLIB_VERSION_@0@_@1@'.format(glib_major_req, glib_minor_req)
endif

if glib_minor_req.is_odd()
  glib_max_allowed = 'GLIB_VERSION_@0@_@1@'.format(glib_major_req, glib_minor_req + 1)
else
  glib_max_allowed = 'GLIB_VERSION_@0@_@1@'.format(glib_major_req, glib_minor_req)
endif

glib_req           = '>= @0@.@1@.@2@'.format(glib_major_req, glib_minor_req, glib_micro_req)
pango_req          = '>= 1.41.0'
fribidi_req        = '>= 0.19.7'
atk_req            = '>= 2.35.1'
at_spi2_atk_req    = '>= 2.15.1'
cairo_req          = '>= 1.14.0'
gdk_pixbuf_req     = '>= 2.30.0'
introspection_req  = '>= 1.39.0'
wayland_proto_req  = '>= 1.17'
wayland_req        = '>= 1.14.91'
epoxy_req          = '>= 1.4'
cloudproviders_req = '>= 0.2.5'
xkbcommon_req      = '>= 0.2.0'

gnome = import('gnome')

add_project_arguments('-DG_LOG_USE_STRUCTURED=1', language: 'c')
add_project_arguments('-DGLIB_MIN_REQUIRED_VERSION=' + glib_min_required, language: 'c')
add_project_arguments('-DGLIB_MAX_ALLOWED_VERSION=' + glib_max_allowed, language: 'c')

# Making releases:
#  1. gtk_micro_version += 1;
#  2. gtk_interface_age += 1;
#  3. if any functions have been added, set gtk_interface_age to 0.
#  4. if backwards compatibility has been broken, we're in trouble
gtk_version       = meson.project_version()
gtk_major_version = gtk_version.split('.')[0].to_int()
gtk_minor_version = gtk_version.split('.')[1].to_int()
gtk_micro_version = gtk_version.split('.')[2].to_int()
gtk_interface_age = 27
add_project_arguments('-DGTK_VERSION="@0@"'.format(meson.project_version()), language: 'c')

add_project_arguments('-D_GNU_SOURCE', language: 'c')

gtk_debug_cflags = []
debug = get_option('debug')
optimization = get_option('optimization')
if debug
  gtk_debug_cflags += '-DG_ENABLE_DEBUG'
  if optimization in ['0', 'g']
    gtk_debug_cflags += '-DG_ENABLE_CONSISTENCY_CHECKS'
  endif
elif optimization in ['2', '3', 's']
  gtk_debug_cflags += '-DG_DISABLE_CAST_CHECKS'
endif

add_project_arguments(gtk_debug_cflags, language: 'c')

# Define a string for the earliest version that this release has
# backwards binary compatibility with for all interfaces a module
# might. Unless we add module-only API with lower stability
# guarantees, this should be unchanged until we break binary compat
# for BOBGUI+.
bobgui_binary_version = '4.0.0'

bobgui_binary_age = 100 * bobgui_minor_version + bobgui_micro_version

bobgui_soversion = '1'
bobgui_library_version = '1.@0@.@1@'.format(bobgui_binary_age - bobgui_interface_age, bobgui_interface_age)

bobgui_api_version = '4.0'

x11_enabled      = get_option('x11-backend')
wayland_enabled  = get_option('wayland-backend')
broadway_enabled = get_option('broadway-backend')
macos_enabled    = get_option('macos-backend')
win32_enabled    = get_option('win32-backend')
android_enabled  = get_option('android-backend')
vulkan_enabled   = get_option('vulkan')
# for GTK+.
gtk_binary_version = '3.0.0'

gtk_binary_age = 100 * gtk_minor_version + gtk_micro_version

cc = meson.get_compiler('c')

if cc.get_id() == 'msvc'
  vsver = 0
  mscver = cc.version().split('.')[0].to_int()

  # pre-Visual Studio 2015 (18.xx.xxxxx or earlier): just subtract 6 from major
  # version of cc.version() to get the Visual Studio version
  if mscver < 19
    vsver = mscver - 6
  else
    # Visual Studio 2015 and later (19.xx.xxxxx or later): look at the minor version.
    # If minor version < 10: Visual Studio 2015,
    #    10 < minor version < 20: Visual Studio 2017,
    #    20 < minor version < 30: Visual Studio 2019,
    #    30 < minor version: Visual Studio 2022
    mscsubver = cc.version().split('.')[1].to_int()
    if mscsubver < 10
      vsver = 14
    elif mscsubver < 20
      vsver = 15
    elif mscsubver < 30
      vsver = 16
    else
      vsver = 17
    endif
  endif

  gtk_soversion = 'vs@0@'.format(vsver)
else
  gtk_soversion = '0'
endif

# Maintain compatibility with the Autotools build

gtk_library_version = ''
gail_library_version = ''
if cc.get_id() == 'msvc'
  gtk_library_version = gtk_major_version.to_string()
  gail_library_version = gtk_major_version.to_string()
else
  current = gtk_binary_age - gtk_interface_age
  gtk_library_version = '@0@.@1@.@2@'.format(gtk_soversion, current, gtk_interface_age)
  gail_library_version = '0.0.0'
endif

osx_current = gtk_binary_age - gtk_interface_age + 1
gtk_osxversions = [osx_current, '@0@.@1@.0'.format(osx_current, gtk_interface_age)]

gtk_api_version = '@0@.0'.format(gtk_major_version)

x11_enabled            = get_option('x11_backend')
wayland_enabled        = get_option('wayland_backend')
broadway_enabled       = get_option('broadway_backend')
quartz_enabled         = get_option('quartz_backend')
win32_enabled          = get_option('win32_backend')

os_unix   = false
os_linux  = false
os_win32  = false
os_darwin = false
os_android = false

cc = meson.get_compiler('c')

# Some windowing system backends depend on the platform we're
# building for, so we need to ensure they are disabled; in other
# cases, they are the only windowing system available, so we need
# to ensure they are enabled
if host_machine.system() == 'darwin'

  if not cc.compiles('''
    #include <Availability.h>
    #if __MAC_OS_X_VERSION_MIN_REQUIRED < 101500L
    # error message "Minimal macOS SDK version not met"
    #endif
    ''',
    name: 'macOS SDK version >= 10.15',
  )
    error('macOS SDK 10.15 or newer is required to build BOBGUI')
  endif

  os_darwin = true
elif host_machine.system() == 'windows'
  os_win32 = true
elif host_machine.system() == 'android'
  os_android = true
  os_darwin = true
elif host_machine.system() == 'windows'
  os_win32 = true
elif host_machine.system() == 'linux'
  os_linux = true
endif
os_unix = not os_win32

if os_darwin
  wayland_enabled = false
  x11_enabled = false
else
  macos_enabled = false
  quartz_enabled = false
endif

if os_win32
  wayland_enabled = false
  x11_enabled = false
else
  win32_enabled = false
endif

if os_android
  wayland_enabled = false
  x11_enabled = false
else
  android_enabled = false
endif

bobgui_prefix = get_option('prefix')
bobgui_includedir = join_paths(bobgui_prefix, get_option('includedir'))
bobgui_libdir = join_paths(bobgui_prefix, get_option('libdir'))
bobgui_datadir = join_paths(bobgui_prefix, get_option('datadir'))
bobgui_localedir = join_paths(bobgui_prefix, get_option('localedir'))
bobgui_sysconfdir = join_paths(bobgui_prefix, get_option('sysconfdir'))
bobgui_bindir = join_paths(bobgui_prefix, get_option('bindir'))
bobgui_applicationsdir = join_paths(bobgui_datadir, 'applications')
bobgui_schemasdir = join_paths(bobgui_datadir, 'glib-2.0/schemas')
bobgui_appdatadir = join_paths(bobgui_datadir, 'metainfo')

cdata = configuration_data()
cdata.set_quoted('PACKAGE_VERSION', meson.project_version())
cdata.set_quoted('BOBGUI_LOCALEDIR', bobgui_localedir)
cdata.set_quoted('BOBGUI_DATADIR', bobgui_datadir)
cdata.set_quoted('BOBGUI_LIBDIR', bobgui_libdir)
cdata.set_quoted('BOBGUI_SYSCONFDIR', bobgui_sysconfdir)
cdata.set_quoted('GETTEXT_PACKAGE', 'bobgui40')

check_headers = [
  'dev/evdev/input.h',
  'ftw.h',
  'linux/input.h',
  'sys/mman.h',
  'sys/sysmacros.h',
  'sys/time.h',
gtk_prefix = get_option('prefix')
gtk_includedir = join_paths(gtk_prefix, get_option('includedir'))
gtk_libdir = join_paths(gtk_prefix, get_option('libdir'))
gtk_datadir = join_paths(gtk_prefix, get_option('datadir'))
gtk_localedir = join_paths(gtk_prefix, get_option('localedir'))
gtk_sysconfdir = join_paths(gtk_prefix, get_option('sysconfdir'))
gtk_applicationsdir = join_paths(gtk_datadir, 'applications')
gtk_schemasdir = join_paths(gtk_datadir, 'glib-2.0/schemas')
gtk_appdatadir = join_paths(gtk_datadir, 'metainfo')

# multipress: is mplocaledir defined somewhere?
gtk_mplocaledir = join_paths(gtk_prefix, get_option('localedir'))

cdata = configuration_data()
cdata.set_quoted('PACKAGE_VERSION', meson.project_version())
cdata.set_quoted('GTK_LOCALEDIR', gtk_localedir)
cdata.set_quoted('GTK_DATADIR', gtk_datadir)
cdata.set_quoted('GTK_LIBDIR', gtk_libdir)
cdata.set_quoted('GTK_SYSCONFDIR', gtk_sysconfdir)
cdata.set_quoted('GETTEXT_PACKAGE', 'gtk30')
cdata.set('GTK_MAJOR_VERSION', gtk_major_version)
cdata.set('GTK_MINOR_VERSION', gtk_minor_version)
cdata.set('GTK_MICRO_VERSION', gtk_micro_version)
cdata.set('GTK_BINARY_AGE', gtk_binary_age)
cdata.set('GTK_INTERFACE_AGE', gtk_interface_age)

cdata.set_quoted('PACKAGE_URL', '')
cdata.set_quoted('PACKAGE_NAME', 'gtk+')
cdata.set_quoted('PACKAGE_TARNAME', 'gtk+')
cdata.set_quoted('PACKAGE_STRING', 'gtk+ @0@'.format(meson.project_version()))
cdata.set_quoted('PACKAGE_BUGREPORT', 'http://bugzilla.gnome.org/enter_bug.cgi?product=gtk%2B')

cdata.set('ENABLE_NLS', 1)

check_headers = [
  'crt/externs.h',
  'dlfcn.h',
  'ftw.h',
  'inttypes.h',
  'linux/memfd.h',
  'linux/input.h',
  'dev/evdev/input.h',
  'locale.h',
  'memory.h',
  'stdint.h',
  'stdlib.h',
  'string.h',
  'strings.h',
  'sys/mman.h',
  'sys/param.h',
  'sys/stat.h',
  'sys/sysinfo.h',
  'sys/systeminfo.h',
  'sys/time.h',
  'sys/types.h',
  'unistd.h',
]

foreach h : check_headers
  if cc.has_header(h)
    cdata.set('HAVE_' + h.underscorify().to_upper(), 1)
  endif
endforeach

# Maths functions might be implemented in libm
libm = cc.find_library('m', required: false)

check_functions = [
  'getpagesize',
  'getresuid',
  'madvise',
  'memfd_create',
  'mkostemp',
  'mlock',
  'mmap',
  'posix_fallocate',
  'sincos',
  'sincosf',
  'memmem',
]

foreach func : check_functions
  if cc.has_function(func,
                     args: '-D_GNU_SOURCE',
                     prefix:
                       '#include <stdlib.h>\n' +
                       '#include <string.h>\n' +
                       '#include <unistd.h>\n' +
                       '#include <sys/mman.h>\n' +
                       '#include <fcntl.h>\n' +
                       '#include <math.h>',
                     dependencies: libm)
  'dcgettext',
  'getpagesize',
  'getresuid',
  'lstat',
  'mmap',
  'nearbyint',
  'posix_fallocate',
  '_lock_file',
  'flockfile',
  'mkstemp',
  'mallinfo',
  'round',
  'rint',
  'log2',
  'exp2',
  'sincos',
  'trunc',
  'localtime_r',
  'fmin',
]

foreach func : check_functions
  if cc.has_function(func, dependencies: libm)
    cdata.set('HAVE_' + func.underscorify().to_upper(), 1)
  endif
endforeach

# We use links() because sigsetjmp() is often a macro hidden behind other macros
cdata.set('HAVE_SIGSETJMP',
  cc.links('''#define _POSIX_SOURCE
              #include <setjmp.h>
              int main (void) {
                sigjmp_buf env;
                sigsetjmp (env, 0);
                return 0;
              }''', name: 'sigsetjmp'),
)

# Check for __uint128_t (gcc) by checking for 128-bit division
uint128_t_src = '''int main() {
static __uint128_t v1 = 100;
static __uint128_t v2 = 10;
static __uint128_t u;
u = v1 / v2;
}'''
if cc.compiles(uint128_t_src, name : '__uint128_t available')
  cdata.set('HAVE_UINT128_T', 1)
endif

cdata.set('HAVE_BACKTRACE', cc.has_function('backtrace', prefix: '#include <execinfo.h>'))

# Disable deprecation checks for all libraries we depend on on stable branches.
# This is so newer versions of those libraries don't cause more warnings with
# a stable BOBGUI version.
#
# We don't ever want to turn off deprecation warnings for development cycles,
# however, because that's where we get rid of deprecated API we use.
if bobgui_minor_version.is_even()
cdata.set('HAVE_DECL_ISINF', cc.has_header_symbol('math.h', 'isinf') ? 1 : false)
cdata.set('HAVE_DECL_ISNAN', cc.has_header_symbol('math.h', 'isnan') ? 1 : false)

# Disable deprecation checks for all libraries we depend on on stable branches.
# This is so newer versions of those libraries don't cause more warnings with
# a stable GTK version.
# We don't ever want to turn off deprecation warnings for master however, because
# that's where we get rid of deprecated API we use.
if gtk_minor_version.is_even()
  cdata.set('GLIB_DISABLE_DEPRECATION_WARNINGS', 1)
endif

# Compiler flags
if cc.get_id() == 'msvc'
  required_cflags = []
  required_debug_cflags = '-Zc:preprocessor'
  test_cflags = [
    '-FI@0@/build-aux/msvc/msvc_recommended_pragmas.h'.format(meson.project_source_root()),
    '-D_USE_MATH_DEFINES',
    required_debug_cflags
  ]
  # Extra warning flags and those that should be disabled because they are too common or
  # break features.
  test_cflags += [
    '/wd4018', # signed/unsigned mismatch, gcc/clang don't warn about it so the codes littered with it
    '/wd4068', # unknown pragma, triggers for all gcc/clang pragmas
    '/wd4090', # VC2017 is way too aggressive with this
    '/wd4100', # same as -Wunused-parameter
    '/wd4116', # breaks G_ALIGNOF()
    '/wd4127', # conditional expression is constant - also triggers in macros...
    '/wd4146', # unary minus operator applied to unsigned type, result still unsigned
    '/wd4152', # nonstandard function to data covnversion - g_task_set_source_tag() uses it
    '/wd4189', # MSVC can't handle G_GNUC_UNUSED
    '/wd4200', # zero-sized array in struct/union
    '/wd4201', # nameless struct/union is required with BOBGUI
    '/wd4206', # empty file - gdkkeys-win32-impl-wow64.c qualifies
    '/wd4232', # non-static DLL imports - we do that in lots of places with functions, particularly get_type()
    '/wd4245', # conversion from signed to unsigned (no gcc warning, so BOBGUI's littered with it)
    '/wd4267', # conversion from $64BIT to $32BIT, possible loss of data (no gcc warning, so BOBGUI's littered with it)
    '/wd4324', # structure was padded due to alignment specifier - why are you warning about doing what I told you?!
    '/wd4334', # '<<': result of 32-bit shift implicitly converted to 64 bits (was 64-bit shift intended?)
    '/wd4389', # comparison between signed and unsigned (no gcc warning, so BOBGUI's littered with it)
    '/wd4701', # potentially uninitialized variable - triggers all the time in debug builds
    '/wd4702', # unreachable code, triggers after assert_not_reached() which can be compiled out
  ]

  msvc_supported_cflags = cc.get_supported_arguments(test_cflags)
  if debug
    required_cflags += required_debug_cflags
  endif
  foreach f : required_cflags
    if not msvc_supported_cflags.contains(f)
      error('@0@ required for this build'.format(f))
    endif
  endforeach
  add_project_arguments(msvc_supported_cflags, language: 'c')

  add_languages('cpp')
  cxx = meson.get_compiler('cpp')
  if cxx.get_id() == 'msvc'
    add_project_arguments(cxx.get_supported_arguments(test_cflags), language: 'cpp')
  endif
elif cc.get_id() == 'gcc' or cc.get_id() == 'clang'
  test_cflags = [
    '-fno-strict-aliasing',
    '-Wno-c++11-extensions',
    '-Wno-missing-include-dirs',
    '-Wno-typedef-redefinition',
    '-Wno-tautological-constant-out-of-range-compare',
    '-Wduplicated-branches',
    '-Wduplicated-cond',
    '-Wformat=2',
    '-Wformat-nonliteral',
    '-Wformat-security',
    '-Wignored-qualifiers',
    '-Wimplicit-function-declaration',
    '-Wlogical-op',
    '-Wmisleading-indentation',
    '-Wmissing-format-attribute',
    '-Wmissing-include-dirs',
    '-Wmissing-noreturn',
    '-Wnested-externs',
    '-Wold-style-definition',
    '-Wpointer-arith',
    '-Wshadow',
    '-Wstrict-prototypes',
    '-Wswitch-default',
    '-Wswitch-enum',
    '-Wundef',
    '-Wuninitialized',
    '-Wunused',
  ]

  extra_warnings = [
    'address',
    'array-bounds',
    'empty-body',
    'enum-int-mismatch',
    'implicit',
    'implicit-fallthrough',   # For non-gcc
    'implicit-fallthrough=5', # For GCC, only recognize the attribute and no comments
    'init-self',
    'int-to-pointer-cast',
    'main',
    'missing-braces',
    'missing-declarations',
    'missing-prototypes',
    'nonnull',
    'override-init',
    'pointer-to-int-cast',
    'redundant-decls',
    'return-type',
    'sequence-point',
    'trigraphs',
    'vla',
    'write-strings',
  ]


  if get_option('buildtype').startswith('debug')
    foreach warning: extra_warnings
      test_cflags += '-Werror=@0@'.format(warning)
    endforeach
  else
    foreach warning: extra_warnings
      test_cflags += '-W@0@'.format(warning)
    endforeach
  endif

  if cc.get_id() == 'gcc'
    test_cflags += ['-Wcast-align'] # This warns too much on clang
  endif

  if not bobgui_debug_cflags.contains('-DG_DISABLE_ASSERT')
    test_cflags += ['-Wnull-dereference'] # Too noisy when assertions are disabled
  endif
  # Compiler options taken from msvc_recommended_pragmas.h
  # in GLib, based on _Win32_Programming_ by Rector and Newcomer
  test_cflags = [
    '-FImsvc_recommended_pragmas.h',
    '-D_USE_MATH_DEFINES',
    '-utf-8',
  ]
  add_project_arguments(cc.get_supported_arguments(test_cflags), language: 'c')
elif cc.get_id() == 'gcc' or cc.get_id() == 'clang'
  test_cflags = [
    '-fno-strict-aliasing',
    '-Wpointer-arith',
    '-Wimplicit-function-declaration',
    '-Wformat=2',
    '-Wformat-security',
    '-Wnested-externs',
    '-Wold-style-definition',
    '-Wundef',
    '-Wunused',
    '-Wcast-align',
    '-Wmissing-noreturn',
    '-Wmissing-format-attribute',
    '-Wmissing-include-dirs',
    '-Wlogical-op',
    '-Wignored-qualifiers',
    '-Wno-discarded-qualifiers',
    '-Werror=implicit',
    '-Werror=nonnull',
    '-Werror=init-self',
    '-Werror=main',
    '-Werror=missing-braces',
    '-Werror=sequence-point',
    '-Werror=return-type',
    '-Werror=trigraphs',
    '-Werror=array-bounds',
    '-Werror=write-strings',
    '-Werror=address',
    '-Werror=int-to-pointer-cast',
    '-Werror=pointer-to-int-cast',
    '-Werror=empty-body',
    '-Werror=write-strings',
  ]
else
  test_cflags = []
endif

common_cflags = cc.get_supported_arguments(test_cflags)

# Symbol visibility
if get_option('default_library') != 'static'
  if os_win32
    cdata.set('DLL_EXPORT', true)
  else
    test_cflags += ['-fvisibility=hidden']
  endif
endif

common_cflags = cc.get_supported_arguments(test_cflags)

common_ldflags = cc.get_supported_link_arguments([
  '-Wl,-Bsymbolic',
  '-Wl,-z,relro',
  '-Wl,-z,now',
])

confinc = include_directories('.')
gdkinc = include_directories('gdk')
gskinc = include_directories('gsk')
bobguiinc = include_directories('bobgui')
testinc = include_directories('tests')

harfbuzz_fallback_options = [
  'coretext=enabled',
  'directwrite=enabled',
  'gdi=enabled',
  'freetype=disabled',
  'cairo=disabled',
  'docs=disabled',
  'utilities=disabled',
  'tests=disabled',
]

# Enable FreeType support in HarfBuzz if building it as a
# subproject and FontConfig support is required or requested
if os_linux
  harfbuzz_fallback_options += 'freetype=enabled'
else
endif

# Dependencies
glib_dep       = dependency('glib-2.0', version: glib_req)
gio_dep        = dependency('gio-2.0', version: glib_req)
gi_dep         = dependency('gobject-introspection-1.0', version: gi_req, required: get_option('introspection').enabled())
gobject_dep    = dependency('gobject-2.0', version: glib_req)
if os_win32
  giowin32_dep = dependency('gio-windows-2.0', version: glib_req, required: win32_enabled)
endif
if os_unix
  giounix_dep  = dependency('gio-unix-2.0', version: glib_req, required: false)
endif
gmodule_dep    = dependency('gmodule-2.0', version: glib_req)
pango_dep      = dependency('pango', version: pango_req)
cairo_dep      = dependency('cairo', version: cairo_req,
                            default_options: ['zlib=enabled', 'tests=disabled'])
cairogobj_dep  = dependency('cairo-gobject', version: cairo_req)
fribidi_dep    = dependency('fribidi', version: fribidi_req,
                            default_options: ['docs=false'])
harfbuzz_dep   = dependency('harfbuzz', version: harfbuzz_req,
                            default_options: harfbuzz_fallback_options)
hb_subset_dep  = dependency('harfbuzz-subset', version: harfbuzz_req)

# Require PangoFT2 if on X11 or wayland
pangoft_dep    = dependency('pangoft2', version: pango_req,
                            required: wayland_enabled or x11_enabled)

if win32_enabled
  # for BOBGUI_IM_CONTEXT_IME
  pangowin32_dep = dependency('pangowin32')
endif

pangocairo_dep    = dependency('pangocairo', version: pango_req)
pixbuf_dep        = dependency('gdk-pixbuf-2.0', version: gdk_pixbuf_req,
                               default_options: ['png=enabled', 'jpeg=enabled', 'gif=enabled', 'glycin=disabled', 'builtin_loaders=all', 'man=false', 'documentation=false'])
png_dep           = dependency('libpng', 'png')
tiff_dep          = dependency('libtiff-4', 'tiff')
jpeg_dep          = dependency('libjpeg', 'jpeg')
epoxy_dep         = dependency('epoxy', version: epoxy_req)
xkbdep            = dependency('xkbcommon', version: xkbcommon_req, required: wayland_enabled)
graphene_dep      = dependency('graphene-gobject-1.0', version: graphene_req,
                               default_options: ['tests=false', 'gobject_types=true'])
soup_dep          = dependency('libsoup-3.0')
sqlite_dep        = dependency('sqlite3')
iso_codes_dep     = dependency('iso-codes', required: false)
gstplay_dep       = dependency('gstreamer-play-1.0', version: gstreamer_req,
                               required: get_option('media-gstreamer'),
                               default_options: [
                                 'python=disabled',
                               ])
gstgl_dep         = dependency('gstreamer-gl-1.0', version: gstreamer_req,
                               required: get_option('media-gstreamer'))
gstallocators_dep = dependency('gstreamer-allocators-1.0', version: gstreamer_req,
                               required: get_option('media-gstreamer'))
if os_win32
  gstd3d12_dep    = dependency('gstreamer-d3d12-1.0', version: gstreamer_req,
                               required: get_option('media-gstreamer'))
else
  gstd3d12_dep = dependency('', required: false)
endif

if gstplay_dep.found() and gstgl_dep.found() and gstallocators_dep.found()
  cdata.set('HAVE_GSTREAMER', 1)
endif
gst_deps = [ gstplay_dep, gstgl_dep, gstallocators_dep, gstd3d12_dep ]

fontconfig_dep = [] # only used in x11 backend
    if cc.get_id() == 'msvc'
      cdata.set('_GDK_EXTERN', '__declspec(dllexport) extern')
    else
      cdata.set('_GDK_EXTERN', '__attribute__((visibility("default"))) __declspec(dllexport) extern')
      common_cflags += ['-fvisibility=hidden']
    endif
  else
    cdata.set('_GDK_EXTERN', '__attribute__((visibility("default"))) extern')
    common_cflags += ['-fvisibility=hidden']
  endif
endif

common_ldflags = []

if os_unix and not os_darwin
  foreach ldflag: [ '-Wl,-Bsymbolic', '-Wl,-z,relro', '-Wl,-z,now', ]
    if cc.links('int main () { return 0; }', name: ldflag, args: ldflag)
      common_ldflags += [ ldflag ]
    endif
  endforeach
endif

confinc = include_directories('.')
gdkinc = include_directories('gdk')
gtkinc = include_directories('gtk')
libgailutilinc = include_directories('libgail-util')
testinc = include_directories('tests')

# Dependencies
glib_dep       = dependency('glib-2.0', version: glib_req,
                            fallback : ['glib', 'libglib_dep'])
gobject_dep    = dependency('gobject-2.0', version: glib_req,
                            fallback : ['glib', 'libgobject_dep'])
if os_win32
  giowin32_dep = dependency('gio-windows-2.0', version: glib_req, required: win32_enabled,
                            fallback : ['glib', 'libgio_dep'])
endif
if os_unix
  giounix_dep  = dependency('gio-unix-2.0', version: glib_req, required: false,
                            fallback : ['glib', 'libgio_dep'])
endif
gmodule_dep    = dependency('gmodule-2.0', version: glib_req,
                            fallback : ['glib', 'libgmodule_dep'])

# Update when we can find depedencies in a declarative manner
cairo_dep      = dependency('cairo', version: cairo_req, required: cc.get_id() != 'msvc')
cairogobj_dep  = dependency('cairo-gobject', version: cairo_req, required: cc.get_id() != 'msvc')

# Remove when we can find depedencies in a declarative manner
if not cairo_dep.found() and cc.get_id() == 'msvc'
  if cc.has_header('cairo.h') and cc.has_header('cairo-win32.h')
    cairo_dep = cc.find_library('cairo', required: false)
  endif
endif

if not cairo_dep.found()
  cairo_dep      = dependency('cairo', version: cairo_req,
                              fallback : ['cairo', 'libcairo_dep'])
endif

# Remove when we can find depedencies in a declarative manner
if not cairogobj_dep.found() and cc.get_id() == 'msvc'
  if cc.has_header('cairo-gobject.h')
    cairogobj_dep = cc.find_library('cairo-gobject', required: false)
  endif
endif

if not cairogobj_dep.found()
  cairogobj_dep  = dependency('cairo-gobject', version: cairo_req,
                              fallback : ['cairo', 'libcairogobject_dep'])
endif

pango_dep      = dependency('pango', version: pango_req,
                            fallback : ['pango', 'libpango_dep'])
fribidi_dep    = dependency('fribidi', version: fribidi_req,
                            fallback : ['fribidi', 'libfribidi_dep'])

# Require PangoFT2 if on X11 or wayland
require_pangoft2 = wayland_enabled or x11_enabled
if require_pangoft2
  pangoft_dep    = dependency('pangoft2', fallback : ['pango', 'libpangoft2_dep'])
else
  pangoft_dep    = dependency('pangoft2', required: false)
endif

if pangoft_dep.found()
  # Need at least 2.7.1 for FT_Get_Var_Design_Coordinates()
  # We get the dependency itself from pango, but pango doesn't care
  # about ft2 version, so an extra check is needed.

  # Update once Meson can have deps declared in a declarative manner
  ft2_dep      = dependency('freetype2', version: '>= 2.7.1', required: false)

  if not ft2_dep.found() and cc.get_id() == 'msvc'
    if cc.has_header('ft2build.h') and cc.has_header('freetype/ftmm.h')
      ft_lib = cc.find_library('freetype', required: false)
      if ft_lib.found()
        if cc.has_function('FT_Get_Var_Design_Coordinates', dependencies: ft_lib,
                           prefix: '''#include <ft2build.h>
                                      #include FT_MULTIPLE_MASTERS_H''')
          ft2_dep = ft_lib
        endif
      endif
    endif
  endif

  if not ft2_dep.found()
    ft2_dep      = dependency('freetype2', version: '>= 2.7.1', fallback : ['freetype2', 'freetype_dep'])
  endif
endif

if win32_enabled
  # for GTK_IM_CONTEXT_IME
  pangowin32_dep  = dependency('pangowin32', fallback : ['pango', 'libpangowin32_dep'])
endif

pangocairo_dep = dependency('pangocairo', version: cairo_req,
                            fallback : ['pango', 'libpangocairo_dep'])
pixbuf_dep     = dependency('gdk-pixbuf-2.0', version: gdk_pixbuf_req,
                            fallback : ['gdk-pixbuf', 'gdkpixbuf_dep'])
epoxy_dep      = dependency('epoxy', version: epoxy_req,
                            fallback: ['libepoxy', 'libepoxy_dep'])
atk_dep        = dependency('atk', version: atk_req,
                            fallback : ['atk', 'libatk_dep'])

# Update once Meson can have deps declared in a declarative manner or can
# find deps properly with CMake again
harfbuzz_dep   = dependency('harfbuzz', version: '>= 0.9', required: false)
xkbdep         = dependency('xkbcommon', version: xkbcommon_req, required: wayland_enabled)
iso_codes_dep  = dependency('iso-codes', required: false)

fontconfig_dep = [] # only used in x11 backend
atkbridge_dep  = [] # only used in x11 backend

if os_win32
  platform_gio_dep = giowin32_dep
endif
if os_unix
  platform_gio_dep = giounix_dep
endif

tracker3_dep = dependency('tracker-sparql-3.0', required: get_option('tracker'))
cdata.set('HAVE_TRACKER3', tracker3_dep.found())

colord_dep = dependency('colord', version: '>= 0.1.9', required: get_option('colord'))
cdata.set('HAVE_COLORD', colord_dep.found())

if not os_win32
  if iso_codes_dep.found()
    cdata.set_quoted('ISO_CODES_PREFIX', iso_codes_dep.get_variable(pkgconfig: 'prefix'))
  else
    cdata.set_quoted('ISO_CODES_PREFIX', '/usr')
  endif
endif

cairo_backends = []
foreach backend: [ ['cairo-xlib', cairo_req, x11_enabled],
                   ['cairo-win32', cairo_req, win32_enabled],
                   ['cairo-quartz', cairo_req, macos_enabled],
                   ['cairo', cairo_req, broadway_enabled or wayland_enabled], ]
 backend_enabled = backend.get(2)
 cairo_backend_req = backend.get(1)
 cairo_backend = backend.get(0)
 if backend_enabled
   if dependency(cairo_backend, version: cairo_backend_req).found()
     cairo_backends += [ cairo_backend ]
   endif
 endif
# Remove once Meson is fixed to find deps properly with CMake or
# gains the ability to declare deps in a declarative manner
# Fallback for HarfBuzz
if cc.get_id() == 'msvc' and not harfbuzz_dep.found()
  if cc.has_header('hb.h')
    harfbuzz_dep = cc.find_library('harfbuzz', required : false)
  endif
endif

require_harfbuzz = wayland_enabled or x11_enabled
if require_harfbuzz and not harfbuzz_dep.found()
  harfbuzz_dep   = dependency('harfbuzz', version: '>= 0.9', required: false,
                              fallback: ['harfbuzz', 'libharfbuzz_dep'])
endif

tracker3_enabled = get_option('tracker3')
if tracker3_enabled
  tracker3_dep = dependency('tracker-sparql-3.0', required: false)
  if tracker3_dep.found()
    cdata.set('HAVE_TRACKER3', tracker3_dep.found())
  else
    error('Tracker3 not found, but was explicitly requested.')
  endif
endif

if iso_codes_dep.found()
  cdata.set_quoted('ISO_CODES_PREFIX', iso_codes_dep.get_pkgconfig_variable('prefix'))
else
  cdata.set_quoted('ISO_CODES_PREFIX', get_option('prefix'))
endif

backend_immodules = []

pc_gdk_extra_libs = []

cairo_found_type = cairo_dep.type_name()
if cairo_found_type == 'internal'
  cairo_features = subproject('cairo').get_variable('built_features')
endif

cairogobj_found_type = cairogobj_dep.type_name()

cairo_backends = []
foreach backend: [ ['cairo-xlib', cairo_req, x11_enabled],
                   ['cairo-win32', cairo_req, win32_enabled],
                   ['cairo-quartz', cairo_req, quartz_enabled],
                   ['cairo', cairo_req, broadway_enabled or wayland_enabled], ]
  backend_enabled = backend.get(2)
  cairo_backend_req = backend.get(1)
  cairo_backend = backend.get(0)
  if backend_enabled
    if cairo_found_type == 'pkgconfig'
      if dependency(cairo_backend, version: cairo_backend_req, required : cc.get_id() != 'msvc').found()
        cairo_backends += cairo_backend
      endif
    elif cairo_found_type == 'internal'
      foreach f: cairo_features
        if f['name'] == cairo_backend
          cairo_backends += cairo_backend
        endif
      endforeach
    endif
  endif
endforeach

cairo_pkg_found = false
cairogobj_pkg_found = false

if cairo_dep.found()
  cairo_pkg_found = true
endif
if cairogobj_dep.found()
  cairogobj_pkg_found = true
endif

cairo_csi_dep = dependency('cairo-script-interpreter', required: false)
if not cairo_csi_dep.found()
  cairo_csi_dep = cc.find_library('cairo-script-interpreter', required: get_option('build-tests'))
endif

have_egl = epoxy_dep.get_variable(
  pkgconfig: 'epoxy_has_egl',
  internal: 'epoxy_has_egl',
  default_value: '0') == '1'

cdata.set('HAVE_CAIRO_SCRIPT_INTERPRETER', cairo_csi_dep.found())
if have_egl
  cdata.set('HAVE_EGL', 1)
endif
cdata.set('HAVE_HARFBUZZ', harfbuzz_dep.found())
cdata.set('HAVE_PANGOFT', pangoft_dep.found())
if win32_enabled
  cdata.set('HAVE_PANGOWIN32', pangowin32_dep.found())
  cdata.set('CONST_VTABLE', 1)
endif

wayland_pkgs = []
if wayland_enabled
  wlclientdep    = dependency('wayland-client', version:  wayland_req,
                              default_options: ['documentation=false', 'tests=false'])
  wlprotocolsdep = dependency('wayland-protocols', version: wayland_proto_req,
                              default_options: ['tests=false'])
  wlegldep       = dependency('wayland-egl')

  wayland_pkgs = [
    'wayland-client @0@'.format(wayland_req),
    'xkbcommon @0@'.format(xkbcommon_req),
if cairo_found_type == 'internal' or cairo_found_type == 'pkgconfig'
  cairo_pkg_found = true
endif
if cairogobj_found_type == 'internal' or cairogobj_found_type == 'pkgconfig'
  cairogobj_pkg_found = true
endif

cairo_libs = []

cdata.set('HAVE_HARFBUZZ', harfbuzz_dep.found() ? 1 : false)
cdata.set('HAVE_PANGOFT', pangoft_dep.found() ? 1 : false)

atk_pkgs = ['atk']

wayland_pkgs = []
if wayland_enabled
  wlclientdep    = dependency('wayland-client', version:  wayland_req)
  wlprotocolsdep = dependency('wayland-protocols', version: wayland_proto_req)
  wlcursordep    = dependency('wayland-cursor', version: wayland_req)
  wlegldep       = dependency('wayland-egl')
  backend_immodules += ['wayland']

  wayland_pkgs = [
    'wayland-client', wayland_req,
    'xkbcommon', xkbcommon_req,
    'wayland-cursor', wayland_req,
    'wayland-egl',
  ]
endif

x11_pkgs = []
if x11_enabled
  xrandr_dep     = dependency('xrandr', version: '>= 1.2.99')
  xrandr15_dep   = dependency('xrandr', version: '>= 1.5', required: false)
  x11_dep        = dependency('x11')
  xrender_dep    = dependency('xrender')
  xi_dep         = dependency('xi')
  xext_dep       = dependency('xext')
  xcursor_dep    = dependency('xcursor')
  xdamage_dep    = dependency('xdamage')
  xfixes_dep     = dependency('xfixes')
  fontconfig_dep = dependency('fontconfig')

  x11_pkgs = ['fontconfig', 'x11', 'xext', 'xi', 'xrandr', 'xcursor', 'xdamage', 'xfixes', 'xinerama']

  cdata.set('HAVE_XCURSOR', 1)
  cdata.set('HAVE_XDAMAGE', 1)
  cdata.set('HAVE_XFIXES', 1)

  if not cc.has_function('XkbQueryExtension', dependencies: x11_dep,
                         prefix : '#include <X11/XKBlib.h>')
    error('X11 backend enabled, but Xkb not found.')
  endif
  cdata.set('HAVE_XKB', 1)

  if not cc.has_function('XSyncQueryExtension', dependencies: xext_dep,
                         prefix: '''#include <X11/Xlib.h>
                                    #include <X11/extensions/sync.h>''')
    error('X11 backend enabled, but Xsync not found.')
  endif
  cdata.set('HAVE_XSYNC', 1)

  if not cc.has_function('XGetEventData', dependencies: x11_dep)
    error('X11 backend enabled, but no generic event support.')
  endif
  cdata.set('HAVE_XGENERICEVENTS', 1)

  if not xi_dep.found() or not cc.has_header('X11/extensions/XInput2.h', dependencies: xi_dep)
    error('X11 backend enabled, but XInput2 not found.')
  endif

  # Note that we also check that the XIScrollClassInfo struct is defined,
  # because at least Ubuntu Oneiric seems to have XIAllowTouchEvents(),
  #   but not the XIScrollClassInfo struct
  has_allow_touch_events = cc.has_function('XIAllowTouchEvents', dependencies: xi_dep)
  has_scroll_class_info = cc.has_member('XIScrollClassInfo', 'number', dependencies: xi_dep,
                                        prefix: '''#include <X11/Xlib.h>
                                                   #include <X11/extensions/XInput2.h>''')
  if not has_allow_touch_events or not has_scroll_class_info
    error('X11 backend enabled, but XInput2.2 not found.')
  endif
  cdata.set('XINPUT_2_2', 1)
  xcursor_dep    = dependency('xcursor', required: false)
  xdamage_dep    = dependency('xdamage', required: false)
  xfixes_dep     = dependency('xfixes', required: false)
  xcomposite_dep = dependency('xcomposite', required: false)
  fontconfig_dep = dependency('fontconfig', fallback: ['fontconfig', 'fontconfig_dep'])
  atkbridge_dep  = dependency('atk-bridge-2.0', version: at_spi2_atk_req)

  backend_immodules += ['xim']

  x11_pkgs = ['fontconfig', 'x11', 'xext', 'xi', 'xrandr']

  if xcursor_dep.found()
    x11_pkgs += ['xcursor']
  endif
  if xfixes_dep.found()
    x11_pkgs += ['xfixes']
  endif
  if xcomposite_dep.found()
    x11_pkgs += ['xcomposite']
  endif
  if xdamage_dep.found()
    x11_pkgs += ['xdamage']
  endif

  atk_pkgs += ['atk-bridge-2.0']

  cdata.set('HAVE_XDAMAGE', xdamage_dep.found() ? 1 : false)
  cdata.set('HAVE_XCURSOR', xcursor_dep.found() ? 1 : false)
  cdata.set('HAVE_XCOMPOSITE', xcomposite_dep.found() ? 1 : false)
  cdata.set('HAVE_XFIXES', xfixes_dep.found() ? 1 : false)

  if cc.has_function('XkbQueryExtension', dependencies: x11_dep,
                     prefix : '#include <X11/XKBlib.h>')
    cdata.set('HAVE_XKB', 1)
  endif

  if cc.has_function('XSyncQueryExtension', dependencies: xext_dep,
                     prefix: '''#include <X11/Xlib.h>
                                #include <X11/extensions/sync.h>''')
    cdata.set('HAVE_XSYNC', 1)
  endif

  if cc.has_function('XGetEventData', dependencies: x11_dep)
    cdata.set('HAVE_XGENERICEVENTS', 1)
  endif

  if xi_dep.found() and cc.has_header('X11/extensions/XInput2.h', dependencies: xi_dep)
    cdata.set('XINPUT_2', 1)
    # Note that we also check that the XIScrollClassInfo struct is defined,
    # because at least Ubuntu Oneiric seems to have XIAllowTouchEvents(),
    # but not the XIScrollClassInfo struct
    has_allow_touch_evens = cc.has_function('XIAllowTouchEvents', dependencies: xi_dep)
    has_scroll_class_info = cc.has_member('XIScrollClassInfo', 'number', dependencies: xi_dep,
                                          prefix: '''#include <X11/Xlib.h>
                                                     #include <X11/extensions/XInput2.h>''')
    if has_allow_touch_evens and has_scroll_class_info
      cdata.set('XINPUT_2_2', 1)
    endif
  endif

  has_gesture_pinch_event = cc.has_member('XIGesturePinchEvent', 'type', dependencies: xi_dep,
                                          prefix: '''#include <X11/Xlib.h>
                                                     #include <X11/extensions/XInput2.h>''')
  if has_gesture_pinch_event
    cdata.set('XINPUT_2_4', 1)
  endif

  xinerama_dep = dependency('xinerama')
  if not cc.has_header_symbol('X11/extensions/Xinerama.h', 'XineramaQueryExtension', dependencies: xinerama_dep)
    error('X11 backend enabled, but Xinerama extension does not work.')
  endif
  cdata.set('HAVE_XFREE_XINERAMA', 1)

  cdata.set('HAVE_RANDR', xrandr_dep.found())
  cdata.set('HAVE_RANDR15', xrandr15_dep.found())

  enable_xinerama = get_option('xinerama')
  if enable_xinerama != 'no'
    want_xinerama = enable_xinerama == 'yes'
    xinerama_dep = dependency('xinerama', required: want_xinerama)
    if xinerama_dep.found() and cc.has_header_symbol('X11/extensions/Xinerama.h', 'XineramaQueryExtension', dependencies: xinerama_dep)
      cdata.set('HAVE_XFREE_XINERAMA', 1)
      cdata.set('HAVE_XINERAMA', 1)
      x11_pkgs += ['xinerama']
    endif
  else
    xinerama_dep = []
  endif

  cdata.set('HAVE_RANDR', xrandr_dep.found() ? 1 : false)
  cdata.set('HAVE_RANDR15', xrandr15_dep.found() ? 1 : false)
endif

if broadway_enabled
  pc_gdk_extra_libs += ['-lz']
  backend_immodules += ['broadway']
endif

if quartz_enabled
  pc_gdk_extra_libs += ['-Wl,-framework,Cocoa', '-Wl,-framework,Carbon', '-Wl,-framework,CoreGraphics']
  backend_immodules += ['quartz']
endif

extra_demo_ldflags = []
if win32_enabled
  pc_gdk_extra_libs += ['-lgdi32', '-limm32', '-lshell32', '-lole32']
  if cc.get_id() == 'msvc'
    # Since the demo programs are now built as pure GUI programs, we
    # need to pass in /entry:mainCRTStartup so that they will properly
    # link on Visual Studio builds
    extra_demo_ldflags = ['/entry:mainCRTStartup']
  endif

  # Check whether libepoxy is built with EGL support on Windows
endif

# Check for bind_textdomain_codeset, including -lintl if GLib brings it in by
# doing the same check as glib. We can't check that by linking to glib because
# it might be a subproject and hence not built yet.
libintl_prefix = '#include <libintl.h>'
if cc.has_function('ngettext', prefix: libintl_prefix)
  else
    pc_gdk_extra_libs += ['-Wl,-luuid']
  endif
  pc_gdk_extra_libs += ['-lwinmm', '-ldwmapi', '-lsetupapi', '-lcfgmgr32', '-lhid']
  backend_immodules += ['ime']

  # Check whether libepoxy is built with EGL support on Windows
  win32_has_egl = epoxy_dep.get_variable(
    pkgconfig: 'epoxy_has_egl',
    internal: 'epoxy_has_egl',
    default_value: '0') == '1'
endif

# IMModules stuff-unfortunately we need to put items here
# as they can be built as modules or built directly into GTK
builtin_all_immodules = false

builtin_modules_opt = get_option('builtin_immodules')
builtin_immodules = []

if builtin_modules_opt == ''
  if os_win32
    # Current MSVC projects build all immodules directly into GTK by default,
    # as does the mingw autotools build
    message('IMModules are built into GTK for MSVC builds by default')
    builtin_immodules = ['all']
  else
    builtin_immodules = ['none']
  endif
else
  builtin_immodules = builtin_modules_opt.split(',')
endif

if builtin_immodules.contains('none')
  builtin_immodules = []
elif builtin_immodules.contains('all')
  builtin_all_immodules = true
elif builtin_immodules.contains('backend')
  builtin_immodules += backend_immodules
endif

proto_sources = [
  'text-input-unstable-v3',
  'gtk-text-input',
]
proto_sources_outputs = []

foreach proto_name: proto_sources
  if wayland_enabled

    wayland_scanner = find_program('wayland-scanner')
    input_path = join_paths('modules', 'input', proto_name + '.xml')
    sources = []
    sources += custom_target('@0@ client header'.format(proto_name),
                             input: input_path,
                             output: '@0@-client-protocol.h'.format(proto_name),
                             command: [
                               wayland_scanner,
                               'client-header',
                               '@INPUT@', '@OUTPUT@',
                            ])

    sources += custom_target('@0@ source'.format(proto_name),
                             input: input_path,
                             output: '@0@-protocol.c'.format(proto_name),
                             command: [
                               wayland_scanner,
                               'private-code',
                               '@INPUT@', '@OUTPUT@',
                             ])
    proto_sources_outputs += [sources]
  else
    proto_sources_outputs += [[]]
  endif
endforeach

# Unfortunately it seems that we need to include the input module sources here, and build
# them in $(buildrootdir) if we are to build them included in GTK...
# Make sure the source listings are the same as in modules/input/Makefile.am
immodule_srcdir = 'modules/input/'
immodules = [
  [ 'am-et', files([ immodule_srcdir + 'imam-et.c' ]) ],
  [ 'cedilla', files([ immodule_srcdir + 'imcedilla.c' ]) ],
  [ 'cyrillic-translit', files([ immodule_srcdir + 'imcyrillic-translit.c' ]) ],
  [ 'inuktitut', files([ immodule_srcdir + 'iminuktitut.c' ]) ],
  [ 'ipa', files([ immodule_srcdir + 'imipa.c' ]) ],
  [ 'multipress', files( [ immodule_srcdir + 'gtkimcontextmultipress.c', immodule_srcdir + 'immultipress.c' ]), true,
                    ['-DMULTIPRESS_LOCALEDIR="@0@"'.format(gtk_mplocaledir),
                     '-DMULTIPRESS_CONFDIR="@0@/gtk-3.0"'.format(gtk_sysconfdir)] ],
  [ 'thai', files([ immodule_srcdir + 'thai-charprop.c', immodule_srcdir + 'gtkimcontextthai.c', immodule_srcdir + 'imthai.c' ]) ],
  [ 'ti-er', files([ immodule_srcdir + 'imti-er.c' ]) ],
  [ 'ti-et', files([ immodule_srcdir + 'imti-et.c' ]) ],
  [ 'viqr', files([ immodule_srcdir + 'imviqr.c' ]) ],
  [ 'broadway', files([ immodule_srcdir + 'imbroadway.c' ]), broadway_enabled ],
  [ 'ime', files([ immodule_srcdir + 'gtkimcontextime.c', immodule_srcdir + 'imime.c' ]), win32_enabled ],
  [ 'quartz', files([ immodule_srcdir + 'imquartz.c' ]), quartz_enabled, [ '-xobjective-c' ] ],
  [ 'wayland', files([ immodule_srcdir + 'imwayland.c' ]) + proto_sources_outputs[0], wayland_enabled ],
  [ 'waylandgtk', files([ immodule_srcdir + 'imwaylandgtk.c'])  + proto_sources_outputs[1], wayland_enabled ],
  [ 'xim', files([ immodule_srcdir + 'gtkimcontextxim.c', immodule_srcdir + 'imxim.c' ]), x11_enabled ],
]

install_data(join_paths(immodule_srcdir, 'im-multipress.conf'), install_dir: join_paths(gtk_sysconfdir, 'gtk-3.0'))

# Check for bind_textdomain_codeset, including -lintl if GLib brings it in by
# doing the same check as glib. We can't check that by linking to glib because
# it might be a subproject and hence not built yet.
if cc.has_function('ngettext')
  libintl_dep = []
  cdata.set('HAVE_BIND_TEXTDOMAIN_CODESET', 1)
else
  libintl_dep = cc.find_library('intl', required : false)
  if cc.has_function('bind_textdomain_codeset', prefix: libintl_prefix, dependencies: libintl_dep)
  if cc.has_function('bind_textdomain_codeset', dependencies: libintl_dep)
    cdata.set('HAVE_BIND_TEXTDOMAIN_CODESET', 1)
  else
    # Don't use subproject('proxy-libintl').get_variable('intl_dep') because that
    # makes the dependency unconditional. This way, people have the option of
    # either not providing the subproject or disabling it entirely with
    # --wrap-mode=nodownload or nofallback.
    libintl_dep = dependency('', required : false,
                             fallback: ['proxy-libintl', 'intl_dep'])
    if libintl_dep.found()
      cdata.set('HAVE_BIND_TEXTDOMAIN_CODESET', 1)
    endif
  endif
endif

if os_unix
  cdata.set('HAVE_GIO_UNIX', giounix_dep.found())
endif

# Check for Vulkan support
# Uses meson's custom vulkan dependency searching. Set the VULKAN_SDK env var
# to use a custom path for the Vulkan SDK. Bugs that are found with it should
# be reported upstream and fixed.
vulkan_dep = dependency('vulkan',
                        version: vulkan_req,
                        required: vulkan_enabled)
glslc = find_program('glslc', required: vulkan_enabled)
if vulkan_dep.found() and glslc.found()
  have_vulkan = true
  vulkan_pkg_found = vulkan_dep.type_name() == 'pkgconfig'
else
  have_vulkan = false
  vulkan_pkg_found = false
endif

libdrm_dep = dependency('libdrm', required: os_linux)
# We only care about drm_fourcc.h for all the fourccs,
# but not about linking to libdrm
libdrm_dep = libdrm_dep.partial_dependency(includes: true, compile_args: true)
cdata.set('HAVE_DRM_FOURCC_H', libdrm_dep.found())

if not libdrm_dep.found()
  have_drm_fourcc = cc.has_header('drm/drm_fourcc.h', required: os_android)
  cdata.set('HAVE_DRM_DRM_FOURCC_H', have_drm_fourcc)
endif

cdata.set('HAVE_DMABUF', cc.has_header('linux/dma-buf.h'))
cdata.set('HAVE_UDMABUF', cc.has_header('linux/udmabuf.h'))

cloudproviders_dep = dependency('cloudproviders',
  required: get_option('cloudproviders'),
  version: cloudproviders_req,
  default_options: [
    'vapigen=false',
  ])
cdata.set('HAVE_CLOUDPROVIDERS', cloudproviders_dep.found())

accesskit_dep = dependency('accesskit-c-0.18', required: get_option('accesskit').enabled())

cdata.set('HAVE_ACCESSKIT', accesskit_dep.found())

# libsysprof-capture support
if not get_option('sysprof').disabled()
  libsysprof_capture_dep = dependency('sysprof-capture-4', version: sysprof_req,
    required: get_option('sysprof'),
    default_options: [
      'agent=false',
      'examples=false',
      'bobgui=false',
      'tests=false',
      'tools=false',
      'libsysprof=true',
      'sysprofd=none',
      'help=false',
    ],
  )
  cdata.set('HAVE_SYSPROF', libsysprof_capture_dep.found())
  profiler_enabled = true
else
  libsysprof_capture_dep = disabler()
  profiler_enabled = false
endif

graphene_dep_type = graphene_dep.type_name()
if graphene_dep_type == 'pkgconfig'
  graphene_has_sse2 = graphene_dep.get_variable(pkgconfig: 'graphene_has_sse2') == '1'
  graphene_has_gcc = graphene_dep.get_variable(pkgconfig: 'graphene_has_gcc') == '1'
else
  graphene_simd = subproject('graphene').get_variable('graphene_simd')
  graphene_has_sse2 = graphene_simd.contains('sse2')
  graphene_has_gcc = graphene_simd.contains('gcc')
endif

if graphene_has_sse2 or graphene_has_gcc
  message('Need aligned memory due to the use of SSE2 or GCC vector instructions')

  if os_win32 and cc.get_id() == 'gcc'
    add_project_arguments(['-mstackrealign'], language: 'c')
  endif
endif

f16c_cflags = []
if get_option('f16c').enabled()
  f16c_prog = '''
#if defined(__GNUC__)
# if !defined(__amd64__) && !defined(__x86_64__)
#   error "F16C intrinsics are only available on x86_64"
# endif
#if !defined __has_attribute
# error "No __has_attribute() support"
#endif
#if !__has_attribute(ifunc)
# error "ifunc not supported"
#endif
static int resolved_ifunc (void) { return 0; }
static void *resolve_ifunc (void) { return resolved_ifunc; }
int ifunc_test (void) __attribute__((ifunc ("resolve_ifunc")));
#else
int ifunc_test (void) { return 0; }
#endif
#if defined(__SSE__) || defined(_MSC_VER)
# include <immintrin.h>
#else
# error "No F16C intrinsics available"
#endif

int main () {
  float f[4] = { 0, };
  unsigned short h[4] = { 0, };
  __m128 s = _mm_loadu_ps (f);
  __m128i i = _mm_cvtps_ph (s, 0);
  _mm_storel_epi64 ((__m128i*)h, i);

#if defined (__GNUC__) || defined (__clang__)
  __builtin_cpu_init ();
  __builtin_cpu_supports ("f16c");
#endif

    return ifunc_test ();
}'''
  if cc.get_id() != 'msvc'
    test_f16c_cflags = [ '-mf16c' ]
  else
    test_f16c_cflags = []
  endif

  if cc.compiles(f16c_prog, args: test_f16c_cflags, name: 'F16C intrinsics')
    cdata.set('HAVE_F16C', 1)
    f16c_cflags = test_f16c_cflags
  else
    f16c_cflags = []
  endif
endif

if os_unix
  cpdb_dep = dependency('cpdb-frontend', version : '>=2.0', required: get_option('print-cpdb'))
  cups_dep = dependency('cups', version : ['>=2.0', '<3.0'], required: false)
  if not cups_dep.found()
    cups_dep = dependency('cups3', version : '>=3.0', required: false)
    if not cups_dep.found() and get_option('print-cups').enabled()
      error('CUPS >= 2.0 required')
    endif
  endif
endif

# Shell completion
bash = find_program('bash', required : false)

if bash.found()
  bash_comp_dep = dependency('bash-completion', version: '>=2.0', required: false)
  bash_comp_inst_dir = ''
  if bash_comp_dep.found()
    bash_comp_dir_override = bash_comp_dep.version().version_compare('>= 2.10') ? ['datadir', get_option('datadir')] : ['prefix', get_option('prefix')]
    bash_comp_inst_dir = bash_comp_dep.get_variable('completionsdir', pkgconfig_define: bash_comp_dir_override)
  endif

  if bash_comp_inst_dir == ''
    message('Found bash-completion but the .pc file did not set \'completionsdir\', fallback to a predefined path')
    bash_comp_inst_dir = join_paths(get_option('datadir'), 'bash-completion/completions')
  endif
endif

# Introspection
gir = find_program('g-ir-scanner', required : get_option('introspection'))

if not gir.found() and get_option('introspection').enabled()
  error('Introspection enabled, but g-ir-scanner not found.')
endif

build_gir = gir.found() and (get_option('introspection').enabled() or
                             (get_option('introspection').allowed() and get_option('documentation')))

# Resource building
glib_compile_resources = find_program('glib-compile-resources')

project_build_root = meson.current_build_dir()

gen_visibility_macros = find_program('build-aux/meson/gen-visibility-macros.py')

gen_profile_conf = find_program('build-aux/meson/gen-profile-conf.py')

profile = get_option('profile')
if profile == 'auto'
  if bobgui_minor_version.is_even()
    profile = 'default'
  else
    profile = 'devel'
  endif
endif

profile_conf_h = declare_dependency(
  sources: custom_target('profile-conf',
    command: [gen_profile_conf, meson.project_source_root(), profile],
    capture: true,
    output: 'profile_conf.h',
    build_by_default: true,
    build_always_stale: true,
  )
)

subdir('gdk/version')
subdir('bobgui/css')
subdir('gdk')
subdir('gsk')
subdir('bobgui')

subdir('tools')
subdir('po')
subdir('docs/reference')
if get_option('build-demos')
  subdir('demos')
endif
if get_option('build-tests')
  subdir('tests')
endif
if get_option('build-testsuite')
  subdir('testsuite')
endif
if get_option('build-examples')
if os_win32
  if cc.has_header_symbol('windows.h', 'IPrintDialogCallback')
    cdata.set('HAVE_IPRINTDIALOGCALLBACK', 1)
  endif

  dvot_size = cc.sizeof('DISPLAYCONFIG_VIDEO_OUTPUT_TECHNOLOGY', prefix : '''
    #define _WIN32_WINNT 0x601
    #include <windows.h>
  ''')
  cdata.set('SIZEOF_DISPLAYCONFIG_VIDEO_OUTPUT_TECHNOLOGY', dvot_size)

  getdevprop_code = '''
    #define _WIN32_WINNT 0x0600
    #include <windows.h>
    #include <devpropdef.h>
    #include <setupapi.h>

    int main(int argc, char *argv[]) {
      return SetupDiGetDevicePropertyW(NULL, NULL, NULL, NULL, NULL, 0, NULL, 0);
    }
  '''
  result = cc.links(getdevprop_code, args: ['-lsetupapi'], name: 'has SetupDiGetDevicePropertyW')
  cdata.set('HAVE_SETUP_DI_GET_DEVICE_PROPERTY_W', result ? 1 : false)
endif

have_gio_unix = false
if os_unix
  have_gio_unix = giounix_dep.found()
  cdata.set('HAVE_GIO_UNIX', have_gio_unix ? 1 : false)
endif

cloudproviders_enabled = get_option('cloudproviders')
cloudproviders_packages = []
if cloudproviders_enabled
  cloudproviders_dep = dependency('cloudproviders', version: cloudproviders_req, required: true)
  if cloudproviders_dep.found()
    cloudproviders_packages += ['cloudproviders', cloudproviders_req]
    cdata.set('HAVE_CLOUDPROVIDERS', cloudproviders_dep.found() ? 1 : false)
  else
    error('Cloudproviders support not found, but was explicitly requested.')
  endif
endif

profiler_enabled = get_option('profiler')
if profiler_enabled
  profiler_dep = dependency('sysprof-capture-3', static: true, required: false)
  if not profiler_dep.found()
    profiler_dep = dependency('sysprof-capture-4', static: true, required: false)
  endif
  if profiler_dep.found()
    cdata.set('HAVE_SYSPROF_CAPTURE', profiler_dep.found())
  else
    error('Profiler support not found, but was explicitly requested.')
  endif
endif

module_suffix = []
if os_darwin
  module_suffix = 'so'
endif

build_gir = get_option('introspection')
subdir('gdk')
subdir('gtk')
subdir('modules')
subdir('libgail-util')
if get_option('demos')
  subdir('demos')
endif
if get_option('tests')
  subdir('tests')
  subdir('testsuite')
endif
if get_option('examples')
  subdir('examples')
endif

# config.h
configure_file(output: 'config.h',
               configuration: cdata)

# Requires
pango_pkgname = win32_enabled ? 'pangowin32' : 'pango'
gdk_packages = [
  '@0@ @1@'.format(pango_pkgname, pango_req),
  'pangocairo @0@'.format(pango_req),
  'gdk-pixbuf-2.0 @0@'.format(gdk_pixbuf_req),
]

if cairo_pkg_found
  gdk_packages += 'cairo @0@'.format(cairo_req)
endif
if cairogobj_pkg_found
  gdk_packages += 'cairo-gobject @0@'.format(cairo_req)
endif

if vulkan_pkg_found
  gdk_packages += ' vulkan'
endif

gsk_packages = [ 'graphene-gobject-1.0 @0@'.format(graphene_req) ]
bobgui_packages = [ 'gio-2.0 @0@'.format(glib_req) ]

enabled_backends = []
foreach backend: [ 'android', 'broadway', 'macos', 'wayland', 'win32', 'x11', ]
  if get_variable('@0@_enabled'.format(backend))
    enabled_backends += backend
  endif
endforeach

common_pc_variables = [
  'targets=@0@'.format(' '.join(enabled_backends)),
  'bobgui_binary_version=@0@'.format(bobgui_binary_version),
  'bobgui_host=@0@-@1@'.format(host_machine.cpu_family(), host_machine.system()), # FIXME
]

pkg_config.generate(libbobgui,
  filebase: 'bobgui4',
  unescaped_variables: common_pc_variables,
  name: 'BOBGUI',
  description: 'BOBGUI Graphical UI Library',
  requires: gdk_packages + gsk_packages + bobgui_packages,
  subdirs: ['bobgui-@0@'.format(bobgui_api_version)],
)
meson.override_dependency('bobgui4', libbobgui_dep)

foreach backend: enabled_backends
  pkg = 'bobgui4-@0@'.format(backend)
  pkg_config.generate(
    filebase: pkg,
    unescaped_variables: common_pc_variables,
    name: 'BOBGUI',
    description: 'BOBGUI Graphical UI Library',
    requires: ['bobgui4', get_variable('@0@_public_deps'.format(backend), [])],
  )
  meson.override_dependency(pkg, libbobgui_dep)
endforeach

foreach a11y_backend: bobgui_a11y_backends
  pkg = 'bobgui4-@0@'.format(a11y_backend)
  pkg_config.generate(
    filebase: pkg,
    unescaped_variables: common_pc_variables,
    name: 'BOBGUI',
    description: bobgui_a11y_backend_info[a11y_backend]['description'],
    requires: ['bobgui4', get_variable('@0@_public_deps'.format(a11y_backend), [])],
  )
  meson.override_dependency(pkg, libbobgui_dep)
endforeach

if os_unix
  pkg_config.generate(
    filebase: 'bobgui4-unix-print',
    unescaped_variables: common_pc_variables,
    name: 'BOBGUI',
    description: 'BOBGUI Unix print support',
    requires: 'bobgui4',
    subdirs: ['bobgui-@0@/unix-print'.format(bobgui_api_version)],
  )
  meson.override_dependency('bobgui4-unix-print', libbobgui_dep)
endif


if not meson.is_cross_build()
  gnome.post_install(
    glib_compile_schemas: true,
    gtk_update_icon_cache: get_option('build-demos'),
    update_mime_database: os_linux,
  )
else
  message('Not executing post-install steps automatically when cross compiling')
endif

if not meson.is_subproject()
  meson.add_dist_script('build-aux/meson/dist-data.py')
endif
configure_file(input: 'config.h.meson',
               output: 'config.h',
               configuration: cdata)

# pkg-config files - bit of a mess all of this
pkgconf = configuration_data()

pkgconf.set('prefix', get_option('prefix'))
pkgconf.set('exec_prefix', '${prefix}')
pkgconf.set('libdir', '${prefix}' / get_option('libdir'))
pkgconf.set('includedir', '${prefix}' / get_option('includedir'))
pkgconf.set('GTK_API_VERSION', gtk_api_version)
pkgconf.set('VERSION', meson.project_version())
pkgconf.set('GTK_BINARY_VERSION', gtk_binary_version)
pkgconf.set('host', '@0@-@1@'.format(host_machine.cpu_family(), host_machine.system())) # FIXME

# Requires
pango_pkgname = win32_enabled ? 'pangowin32' : 'pango'
gdk_packages = ' '.join([ pango_pkgname, pango_req,
                       'pangocairo', pango_req,
                       'gdk-pixbuf-2.0', gdk_pixbuf_req ])

cairo_packages = ''

if cairo_pkg_found
  cairo_packages += ' '.join([ ' cairo', cairo_req ])
elif cc.get_id() == 'msvc' and cairo_dep.found()
  cairo_libs += '-lcairo'
endif
if cairogobj_pkg_found
  cairo_packages += ' '.join([ ' cairo-gobject', cairo_req ])
elif cc.get_id() == 'msvc' and cairogobj_dep.found()
  cairo_libs += '-lcairo-gobject'
endif

gdk_packages += cairo_packages
pkgconf.set('GDK_PACKAGES', gdk_packages)

if have_gio_unix
  gio_packages = ['gio-unix-2.0', glib_req]
else
  gio_packages = ['gio-2.0', glib_req]
endif

pkgconf.set('GDK_PRIVATE_PACKAGES',
            ' '.join(gio_packages + x11_pkgs + wayland_pkgs + cairo_backends +
                     ['epoxy', epoxy_req] + cloudproviders_packages +
                     ['fribidi', fribidi_req]))

gtk_packages = ' '.join([
    atk_dep.name(), atk_req,
    cairo_packages,
    pixbuf_dep.name(), gdk_pixbuf_req,
    'gio-2.0', glib_req,
])
pkgconf.set('GTK_PACKAGES', gtk_packages)

# Requires.private
pc_gdk_extra_libs += cairo_libs

gtk_private_packages = atk_pkgs + wayland_pkgs + ['epoxy', epoxy_req, 'fribidi', fribidi_req]
if wayland_enabled or x11_enabled
  gtk_private_packages += ['pangoft2']
endif
gtk_private_packages += gio_packages

pkgconf.set('GTK_PRIVATE_PACKAGES', ' '.join(gtk_private_packages))

pkgconf.set('GDK_EXTRA_LIBS', ' '.join(pc_gdk_extra_libs))
pkgconf.set('GSK_EXTRA_LIBS', '')
pkgconf.set('GTK_EXTRA_LIBS', '')

pkgconf.set('GDK_EXTRA_CFLAGS', '')
pkgconf.set('GSK_EXTRA_CFLAGS', '')
pkgconf.set('GTK_EXTRA_CFLAGS', '')

pkg_install_dir = join_paths(get_option('libdir'), 'pkgconfig')

gtk_pcs = ['gtk+-3.0.pc']
gdk_pcs = ['gdk-3.0.pc']

pkg_targets = ''
disabled_backends = []
foreach backend: [ 'broadway', 'quartz', 'wayland', 'win32', 'x11']
  if get_variable('@0@_enabled'.format(backend))
    gtk_pcs += ['gtk+-@0@-3.0.pc'.format(backend)]
    gdk_pcs += ['gdk-@0@-3.0.pc'.format(backend)]
    pkg_targets += ' ' + backend
  else
    disabled_backends += [backend]
  endif
endforeach
pkgconf.set('GDK_BACKENDS', pkg_targets.strip())

configure_file(input: 'gail-3.0.pc.in',
               output: 'gail-3.0.pc',
               configuration: pkgconf,
               install_dir: pkg_install_dir)

foreach pkg: gdk_pcs
  configure_file(input: 'gdk-3.0.pc.in',
                 output: pkg,
                 configuration: pkgconf,
                 install_dir: pkg_install_dir)
endforeach

foreach pkg: gtk_pcs
  configure_file(input: 'gtk+-3.0.pc.in',
                 output: pkg,
                 configuration: pkgconf,
                 install_dir: pkg_install_dir)
endforeach

if os_unix
  configure_file(input: 'gtk+-unix-print-3.0.pc.in',
                 output: 'gtk+-unix-print-3.0.pc',
                 configuration: pkgconf,
                 install_dir: pkg_install_dir)
endif

subdir('po')
subdir('po-properties')
subdir('docs/tools')
subdir('docs/reference')

install_data('m4macros/gtk-3.0.m4', install_dir: join_paths(get_option('datadir'), 'aclocal'))

# Keep this in sync with post-install.py expected arguments
meson.add_install_script('build-aux/meson/post-install.py',
                         gtk_api_version,
                         gtk_binary_version,
                         join_paths(gtk_prefix, get_option('bindir')),
                         gtk_libdir,
                         gtk_datadir)

if host_machine.system() != 'windows'
  # Install Valgrind suppression files (except on Windows,
  # as Valgrind is currently not supported on Windows)
  install_data('bobgui.supp',
    install_dir : join_paths(bobgui_datadir, 'bobgui-4.0', 'valgrind'))
endif


#### Summary ####

summary('Display backends', enabled_backends, section: 'Components')
summary('Print backends', print_backends, section: 'Components')
summary('Media backends', media_backends, section: 'Components')

summary('Vulkan support', vulkan_dep.found(), section: 'Features')
summary('Cloud support', cloudproviders_dep.found(), section: 'Features')
summary('Sysprof support', libsysprof_capture_dep.found(), section: 'Features')
summary('Colord support', colord_dep.found(), section: 'Features')
summary('Tracker support', tracker3_dep.found(), section: 'Features')
summary('AccessKit support', accesskit_dep.found(), section: 'Features')

summary('Compiler', cc.get_id(), section: 'Toolchain')
summary('Linker', cc.get_linker_id(), section: 'Toolchain')

summary('Debugging', get_option('debug'), section: 'Build')
summary('Optimization', get_option('optimization'), section: 'Build')
summary('Introspection', build_gir, section: 'Build')
summary('Documentation', get_option('documentation'), section: 'Build')
summary('Man pages', get_option('man-pages'), section: 'Build')
summary('Testsuite', get_option('build-testsuite'), section: 'Build')
summary('Tests', get_option('build-tests'), section: 'Build')
summary('Demos', get_option('build-demos'), section: 'Build')
summary('Examples', get_option('build-examples'), section: 'Build')

summary('Prefix', bobgui_prefix, section: 'Directories')
summary('Includedir', bobgui_includedir, section: 'Directories')
summary('Libdir', bobgui_libdir, section: 'Directories')
summary('Datadir', bobgui_datadir, section: 'Directories')
  install_data('gtk.supp',
    install_dir : join_paths(gtk_datadir, 'gtk-3.0', 'valgrind'))
endif

test(
  'version-check',
  find_program('check-version.py'),
  args: [
    join_paths(meson.current_source_dir(), 'configure.ac'),
    join_paths(meson.current_source_dir(), 'meson.build'),
  ],
  suite: 'gtk',
)

summary = [
  '',
  '------',
  'GTK+ @0@ (@1@)'.format(gtk_version, gtk_api_version),
  '',
  '  Display backends: @0@ [disabled: @1@]'.format(pkg_targets.strip(), ' '.join(disabled_backends)),
  '    Print backends: @0@ [disabled: @1@]'.format(' '.join(print_backends), ' '.join(disabled_print_backends)),
  '     Cloud support: @0@'.format(get_option('cloudproviders')),
  '    Colord support: @0@'.format(get_option('colord')),
  '          Profiler: @0@'.format(get_option('profiler')),
  '     Introspection: @0@'.format(get_option('introspection')),
  '     Documentation: @0@'.format(get_option('gtk_doc')),
  '         Man pages: @0@'.format(get_option('man')),
  '       Build tests: @0@'.format(get_option('tests')),
  '   Installed tests: @0@'.format(get_option('installed_tests')),
  '             Demos: @0@'.format(get_option('demos')),
  '          Examples: @0@'.format(get_option('examples')),
  'Directories:',
  '            prefix: @0@'.format(gtk_prefix),
  '        includedir: @0@'.format(gtk_includedir),
  '            libdir: @0@'.format(gtk_libdir),
  '           datadir: @0@'.format(gtk_datadir),
  '------',
  ''
]

message('\n'.join(summary))
