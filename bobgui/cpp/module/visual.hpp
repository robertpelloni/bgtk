#pragma once

#include <bobgui/cpp/module.hpp>
#include <bobgui/cpp/visual.hpp>
#include <memory>
#include <string>

namespace bobtk {
namespace module {

// The Visual Pillar interface acts as the service registry for UI, rendering, and window layout frameworks
class VisualPillar : public Module {
public:
    VisualPillar() : Module("Visual", "1.0.0") {}

    // Workbench IDE UI
    std::shared_ptr<visual::Workbench> createWorkbench(BobguiApplication* app) {
        return std::make_shared<visual::Workbench>(app);
    }

    // RHI Renderer
    std::shared_ptr<visual::RhiDevice> createRhiDevice(BobguiRhiBackend backend) {
        return std::make_shared<visual::RhiDevice>(backend);
    }

    // Imgui Direct Integration
    std::shared_ptr<visual::ImguiContext> createImguiContext() {
        return std::make_shared<visual::ImguiContext>();
    }

    // Graph Views
    std::shared_ptr<visual::GraphView> createGraphView() {
        return std::make_shared<visual::GraphView>();
    }

    // Command Palette
    std::shared_ptr<visual::CommandPalette> createCommandPalette() {
        return std::make_shared<visual::CommandPalette>();
    }

    // Docking
    std::shared_ptr<visual::DockManager> createDockManager(BobguiWindow* mainWindow) {
        return std::make_shared<visual::DockManager>(mainWindow);
    }

    // Global Design Tokens
    std::shared_ptr<visual::DesignSystem> getDesignSystem() {
        return visual::DesignSystem::getDefault();
    }

protected:
    bool onInitialize() override {
        // Initialize underlying C libraries
        return true;
    }

    bool onStart() override {
        return true;
    }

    bool onStop() override {
        return true;
    }
};

} // namespace module
} // namespace bobtk
