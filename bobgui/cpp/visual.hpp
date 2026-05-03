#pragma once

#include <bobgui/bobgui.h>
#include <bobgui/cpp/object_handle.hpp>
#include <bobgui/cpp/widget.hpp>
#include <string>
#include <memory>
#include <vector>

namespace bobtk {
namespace visual {

// ---------------------------------------------------------
// WORKBENCH (IDE-STYLE LAYOUT FRAMEWORK)
// ---------------------------------------------------------

class Workbench : public core::ObjectHandle<BobguiWorkbench> {
public:
    explicit Workbench(BobguiApplication* application)
        : core::ObjectHandle<BobguiWorkbench>(bobgui_workbench_new(application)) {}

    Widget getWindow() const { return Widget(BOBGUI_WIDGET(bobgui_workbench_get_window(handle()))); }
    Widget getContent() const { return Widget(bobgui_workbench_get_content(handle())); }
    Widget getLeftSidebar() const { return Widget(bobgui_workbench_get_left_sidebar(handle())); }
    Widget getRightSidebar() const { return Widget(bobgui_workbench_get_right_sidebar(handle())); }
    Widget getCentral() const { return Widget(bobgui_workbench_get_central(handle())); }

    void setTitle(const std::string& title) {
        bobgui_workbench_set_title(handle(), title.c_str());
    }

    void setCentral(Widget& widget) {
        bobgui_workbench_set_central(handle(), BOBGUI_WIDGET(widget.handle()));
    }

    void setLeftSidebar(Widget& widget) {
        bobgui_workbench_set_left_sidebar(handle(), BOBGUI_WIDGET(widget.handle()));
    }

    void setRightSidebar(Widget& widget) {
        bobgui_workbench_set_right_sidebar(handle(), BOBGUI_WIDGET(widget.handle()));
    }

    void setStatus(Widget& widget) {
        bobgui_workbench_set_status(handle(), BOBGUI_WIDGET(widget.handle()));
    }

    void addCommand(const std::string& id, const std::string& label, GCallback callback, gpointer userData) {
        bobgui_workbench_add_command(handle(), id.c_str(), label.c_str(), callback, userData);
    }

    void present() {
        bobgui_workbench_present(handle());
    }
};

// ---------------------------------------------------------
// RENDERING HARDWARE INTERFACE (RHI - Vulkan/Metal/D3D)
// ---------------------------------------------------------

class RhiDevice : public core::ObjectHandle<BobguiRhiDevice> {
public:
    explicit RhiDevice(BobguiRhiBackend backend)
        : core::ObjectHandle<BobguiRhiDevice>(bobgui_rhi_device_new(backend)) {}

    void submitCommands(gpointer cmdList) {
        bobgui_rhi_submit_commands(handle(), cmdList);
    }
};

class RhiBuffer {
public:
    explicit RhiBuffer(BobguiRhiBuffer* handle) : handle_(handle) {}
    BobguiRhiBuffer* handle() const { return handle_; }
private:
    BobguiRhiBuffer* handle_;
};

// ---------------------------------------------------------
// IMGUI INTEGRATION
// ---------------------------------------------------------

class ImguiContext : public core::ObjectHandle<BobguiImguiContext> {
public:
    ImguiContext() : core::ObjectHandle<BobguiImguiContext>(bobgui_imgui_context_new()) {}
};

// ---------------------------------------------------------
// NODE GRAPH VIEW
// ---------------------------------------------------------

class GraphView : public Widget {
public:
    GraphView() : Widget(bobgui_graph_view_new()) {}
};

// ---------------------------------------------------------
// COMMAND PALETTE (Ctrl+Shift+P)
// ---------------------------------------------------------

class CommandPalette : public Widget {
public:
    CommandPalette() : Widget(bobgui_command_palette_new()) {}

    void addCommand(const std::string& id, const std::string& title, const std::string& shortcut) {
        bobgui_command_palette_add_command(BOBGUI_COMMAND_PALETTE(handle()), id.c_str(), title.c_str(), shortcut.c_str());
    }

    void present() {
        bobgui_command_palette_present(BOBGUI_COMMAND_PALETTE(handle()));
    }
};

// ---------------------------------------------------------
// DOCKING MANAGER (Tear-away tabs and panels)
// ---------------------------------------------------------

class DockManager : public core::ObjectHandle<BobguiDockManager> {
public:
    explicit DockManager(BobguiWindow* mainWindow)
        : core::ObjectHandle<BobguiDockManager>(bobgui_dock_manager_new(mainWindow)) {}

    void addWidget(Widget& widget, const std::string& title, BobguiDockPosition position) {
        bobgui_dock_manager_add_widget(handle(), BOBGUI_WIDGET(widget.handle()), title.c_str(), position);
    }

    void saveLayout(const std::string& path) {
        bobgui_dock_manager_save_layout(handle(), path.c_str());
    }

    void loadLayout(const std::string& path) {
        bobgui_dock_manager_load_layout(handle(), path.c_str());
    }
};

// ---------------------------------------------------------
// DESIGN SYSTEM (Global Token / Theming Framework)
// ---------------------------------------------------------

class DesignSystem : public core::ObjectHandle<BobguiDesignSystem> {
public:
    static std::shared_ptr<DesignSystem> getDefault() {
        return std::shared_ptr<DesignSystem>(new DesignSystem(bobgui_design_system_get_default()));
    }
private:
    DesignSystem(BobguiDesignSystem* handle) : core::ObjectHandle<BobguiDesignSystem>(handle) {}
};

} // namespace visual
} // namespace bobtk
