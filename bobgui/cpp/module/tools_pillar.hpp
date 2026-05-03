#pragma once

#include <bobgui/cpp/module.hpp>
#include <bobgui/cpp/tools_subsystems.hpp>
#include <memory>
#include <string>

namespace bobtk {
namespace module {

// The Tools Pillar interface acts as the service registry for editors, compilation toolchains, reports, and tests
class ToolsPillar : public Module {
public:
    ToolsPillar() : Module("Tools", "1.0.0") {}

    // Forge (Build and compilation tools)
    std::shared_ptr<tools_subsystems::ForgeContext> createForgeContext(const std::string& projectDir) {
        return std::make_shared<tools_subsystems::ForgeContext>(projectDir);
    }

    // Reports (PDF/SVG report builders)
    std::shared_ptr<tools_subsystems::ReportDesigner> createReportDesigner() {
        return std::make_shared<tools_subsystems::ReportDesigner>();
    }

    std::shared_ptr<tools_subsystems::ReportPreview> createReportPreview(tools_subsystems::ReportDesigner& designer) {
        return std::make_shared<tools_subsystems::ReportPreview>(designer);
    }

    // Studio (Editor, IDE, Inspection integrations)
    std::shared_ptr<tools_subsystems::StudioManager> getStudioManager() {
        return tools_subsystems::StudioManager::getDefault();
    }

    // Testing (Unit/Integration runners)
    std::shared_ptr<tools_subsystems::TestRunner> createTestRunner(Widget& rootWidget) {
        return std::make_shared<tools_subsystems::TestRunner>(rootWidget);
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
