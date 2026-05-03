#pragma once

#include <bobgui/bobgui.h>
#include <bobgui/cpp/object_handle.hpp>
#include <bobgui/cpp/widget.hpp>
#include <string>
#include <memory>
#include <vector>

namespace bobtk {
namespace tools_subsystems {

// ---------------------------------------------------------
// FORGE CONTEXT (Build / Compilation Tooling)
// ---------------------------------------------------------

class ForgeContext : public core::ObjectHandle<BobguiForgeContext> {
public:
    explicit ForgeContext(const std::string& projectDir)
        : core::ObjectHandle<BobguiForgeContext>(bobgui_forge_context_new(projectDir.c_str())) {}

    void optimizeAssets(GAsyncReadyCallback callback) {
        bobgui_forge_optimize_assets(handle(), callback);
    }

    void buildTarget(const std::string& targetName, GAsyncReadyCallback callback) {
        bobgui_forge_build_target(handle(), targetName.c_str(), callback);
    }

    void compileShaders() {
        bobgui_forge_compile_shaders(handle());
    }

    void generateBoilerplate(const std::string& schemaPath) {
        bobgui_forge_generate_boilerplate(handle(), schemaPath.c_str());
    }
};

// ---------------------------------------------------------
// REPORT DESIGNER & PREVIEW
// ---------------------------------------------------------

class ReportDesigner : public core::ObjectHandle<BobguiReportDesigner> {
public:
    ReportDesigner() : core::ObjectHandle<BobguiReportDesigner>(bobgui_report_designer_new()) {}

    void setTemplate(Widget& templateRoot) {
        bobgui_report_designer_set_template(handle(), BOBGUI_WIDGET(templateRoot.handle()));
    }

    void setData(GList* dataRecords) {
        bobgui_report_designer_set_data(handle(), dataRecords);
    }

    void exportPdf(const std::string& path, GAsyncReadyCallback callback) {
        bobgui_report_export_pdf(handle(), path.c_str(), callback);
    }

    void exportSvg(const std::string& path) {
        bobgui_report_export_svg(handle(), path.c_str());
    }

    void print(BobguiPrintOperation* op) {
        bobgui_report_print(handle(), op);
    }
};

class ReportPreview : public Widget {
public:
    explicit ReportPreview(ReportDesigner& designer)
        : Widget(bobgui_report_preview_new(designer.handle())) {}
};

// ---------------------------------------------------------
// STUDIO MANAGER (IDE, Inspector, Profiler)
// ---------------------------------------------------------

class StudioManager : public core::ObjectHandle<BobguiStudioManager> {
public:
    static std::shared_ptr<StudioManager> getDefault() {
        return std::shared_ptr<StudioManager>(new StudioManager(bobgui_studio_manager_get_default()));
    }

    void showPerformanceOverlay(bool show) {
        bobgui_studio_show_performance_overlay(handle(), show ? TRUE : FALSE);
    }

    void logMetric(const std::string& name, double value) {
        bobgui_studio_log_metric(handle(), name.c_str(), value);
    }

    void inspectWidget(Widget& widget) {
        bobgui_studio_inspect_widget(handle(), BOBGUI_WIDGET(widget.handle()));
    }

    void enableLiveEdit(bool enable) {
        bobgui_studio_enable_live_edit(handle(), enable ? TRUE : FALSE);
    }

    void showResourceUsage() {
        bobgui_studio_show_resource_usage(handle());
    }

private:
    StudioManager(BobguiStudioManager* handle) : core::ObjectHandle<BobguiStudioManager>(handle) {}
};

// ---------------------------------------------------------
// TEST RUNNER
// ---------------------------------------------------------

class TestRunner : public core::ObjectHandle<BobguiTestRunner> {
public:
    explicit TestRunner(Widget& rootWidget)
        : core::ObjectHandle<BobguiTestRunner>(bobgui_test_runner_new(BOBGUI_WIDGET(rootWidget.handle()))) {}
};

} // namespace tools_subsystems
} // namespace bobtk
