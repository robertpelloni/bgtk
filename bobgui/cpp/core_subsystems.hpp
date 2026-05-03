#pragma once

#include <bobgui/bobgui.h>
#include <bobgui/cpp/object_handle.hpp>
#include <bobgui/cpp/widget.hpp>
#include <string>
#include <memory>
#include <vector>

namespace bobtk {
namespace core_subsystems {

// ---------------------------------------------------------
// AUTONOMOUS / AI CONTROL
// ---------------------------------------------------------

class AutonomousController : public core::ObjectHandle<BobguiAutonomousController> {
public:
    static std::shared_ptr<AutonomousController> getDefault() {
        return std::shared_ptr<AutonomousController>(new AutonomousController(bobgui_autonomous_controller_get_default()));
    }

    void enableForWidget(Widget& widget, const std::string& model) {
        bobgui_autonomous_enable_for_widget(handle(), BOBGUI_WIDGET(widget.handle()), model.c_str());
    }

private:
    AutonomousController(BobguiAutonomousController* handle) : core::ObjectHandle<BobguiAutonomousController>(handle) {}
};

// ---------------------------------------------------------
// BRAIN / AI INFERENCE (LLMs)
// ---------------------------------------------------------

class BrainNode : public core::ObjectHandle<BobguiBrainNode> {
public:
    explicit BrainNode(const std::string& modelPath)
        : core::ObjectHandle<BobguiBrainNode>(bobgui_brain_node_new(modelPath.c_str())) {}

    void inferAsync(const std::string& prompt, GAsyncReadyCallback callback, gpointer userData) {
        bobgui_brain_infer_async(handle(), prompt.c_str(), callback, userData);
    }
};

class TranslationContext : public core::ObjectHandle<BobguiTranslationContext> {
public:
    TranslationContext() : core::ObjectHandle<BobguiTranslationContext>(bobgui_translation_context_new()) {}
};

// ---------------------------------------------------------
// COMPUTE (GPU / OpenCL / Compute Shaders)
// ---------------------------------------------------------

class ComputePipeline : public core::ObjectHandle<BobguiComputePipeline> {
public:
    explicit ComputePipeline(const std::string& kernelSource)
        : core::ObjectHandle<BobguiComputePipeline>(bobgui_compute_pipeline_new(kernelSource.c_str())) {}

    void dispatch(int x, int y, int z) {
        bobgui_compute_dispatch(handle(), x, y, z);
    }
};

// ---------------------------------------------------------
// DATA (SQL / Properties)
// ---------------------------------------------------------

class SqlDatabase : public core::ObjectHandle<BobguiSqlDatabase> {
public:
    explicit SqlDatabase(const std::string& driver)
        : core::ObjectHandle<BobguiSqlDatabase>(bobgui_sql_database_add(driver.c_str())) {}
};

class PropertyFloat : public core::ObjectHandle<BobguiProperty> {
public:
    explicit PropertyFloat(float initialValue)
        : core::ObjectHandle<BobguiProperty>(bobgui_property_new_float(initialValue)) {}
};

// ---------------------------------------------------------
// ENTITY (ECS SYSTEM)
// ---------------------------------------------------------

class EntityManager : public core::ObjectHandle<BobguiEntityManager> {
public:
    static std::shared_ptr<EntityManager> getDefault() {
        return std::shared_ptr<EntityManager>(new EntityManager(bobgui_entity_manager_get_default()));
    }

    uint64_t createEntity() {
        return bobgui_entity_create(handle());
    }

private:
    EntityManager(BobguiEntityManager* handle) : core::ObjectHandle<BobguiEntityManager>(handle) {}
};

// ---------------------------------------------------------
// FIN (QUANTITATIVE / FINANCIAL UI)
// ---------------------------------------------------------

class FinChart : public Widget {
public:
    explicit FinChart(BobguiChartType type)
        : Widget(bobgui_fin_chart_new(type)) {}
};

class OrderBook : public Widget {
public:
    OrderBook() : Widget(bobgui_order_book_new()) {}
};

class TradingTerminal : public Widget {
public:
    TradingTerminal() : Widget(bobgui_trading_terminal_new()) {}
};

// ---------------------------------------------------------
// I18N (INTERNATIONALIZATION)
// ---------------------------------------------------------

class Translator : public core::ObjectHandle<BobguiTranslator> {
public:
    explicit Translator(const std::string& domain)
        : core::ObjectHandle<BobguiTranslator>(bobgui_translator_new(domain.c_str())) {}
};

// ---------------------------------------------------------
// LAYOUT MANAGERS (Modern flex/grid)
// ---------------------------------------------------------

class ModernLayout : public core::ObjectHandle<BobguiLayoutManager> {
public:
    explicit ModernLayout(BobguiLayoutType type)
        : core::ObjectHandle<BobguiLayoutManager>(bobgui_layout_manager_new(type)) {}

    void setFlexDirection(const std::string& dir) {
        bobgui_layout_manager_set_flex_direction(handle(), dir.c_str());
    }
};

// ---------------------------------------------------------
// META ORCHESTRATOR (Dynamic Forms)
// ---------------------------------------------------------

class MetaOrchestrator : public core::ObjectHandle<BobguiMetaOrchestrator> {
public:
    static std::shared_ptr<MetaOrchestrator> getDefault() {
        return std::shared_ptr<MetaOrchestrator>(new MetaOrchestrator(bobgui_meta_orchestrator_get_default()));
    }

private:
    MetaOrchestrator(BobguiMetaOrchestrator* handle) : core::ObjectHandle<BobguiMetaOrchestrator>(handle) {}
};

// ---------------------------------------------------------
// QUANTUM SCHEDULER
// ---------------------------------------------------------

class QuantumScheduler : public core::ObjectHandle<BobguiQuantumScheduler> {
public:
    static std::shared_ptr<QuantumScheduler> getDefault() {
        return std::shared_ptr<QuantumScheduler>(new QuantumScheduler(bobgui_quantum_scheduler_get_default()));
    }

private:
    QuantumScheduler(BobguiQuantumScheduler* handle) : core::ObjectHandle<BobguiQuantumScheduler>(handle) {}
};

// ---------------------------------------------------------
// REALTIME KERNEL (Low latency execution)
// ---------------------------------------------------------

class RealtimeKernel : public core::ObjectHandle<BobguiRealtimeKernel> {
public:
    static std::shared_ptr<RealtimeKernel> getInstance() {
        return std::shared_ptr<RealtimeKernel>(new RealtimeKernel(bobgui_realtime_kernel_get_instance()));
    }

    void reservePriority(int priorityLevel) {
        bobgui_realtime_reserve_priority(handle(), priorityLevel);
    }

    void checkDeadlines() {
        bobgui_realtime_check_deadlines(handle());
    }

private:
    RealtimeKernel(BobguiRealtimeKernel* handle) : core::ObjectHandle<BobguiRealtimeKernel>(handle) {}
};

// ---------------------------------------------------------
// REFLECTION CONTEXT (Serialization)
// ---------------------------------------------------------

class ReflectContext : public core::ObjectHandle<BobguiReflectContext> {
public:
    static std::shared_ptr<ReflectContext> getDefault() {
        return std::shared_ptr<ReflectContext>(new ReflectContext(bobgui_reflect_context_get_default()));
    }

private:
    ReflectContext(BobguiReflectContext* handle) : core::ObjectHandle<BobguiReflectContext>(handle) {}
};

// ---------------------------------------------------------
// SCRIPT CONTEXT (JS/Lua embedding)
// ---------------------------------------------------------

class ScriptContext : public core::ObjectHandle<BobguiScriptContext> {
public:
    explicit ScriptContext(const std::string& engineType)
        : core::ObjectHandle<BobguiScriptContext>(bobgui_script_context_new(engineType.c_str())) {}
};

// ---------------------------------------------------------
// STATE MACHINE
// ---------------------------------------------------------

class StateNode {
public:
    explicit StateNode(BobguiStateNode* handle) : handle_(handle) {}
    BobguiStateNode* handle() const { return handle_; }
private:
    BobguiStateNode* handle_;
};

class StateMachine : public core::ObjectHandle<BobguiStateMachine> {
public:
    explicit StateMachine(const StateNode& initialState)
        : core::ObjectHandle<BobguiStateMachine>(bobgui_state_machine_new(initialState.handle())) {}
};

// ---------------------------------------------------------
// STREAM (FRP / Rx-style events)
// ---------------------------------------------------------

class Stream : public core::ObjectHandle<BobguiStream> {
public:
    static std::shared_ptr<Stream> interval(int intervalMs) {
        return std::shared_ptr<Stream>(new Stream(bobgui_stream_interval(intervalMs)));
    }

    void subscribe(GCallback callback, gpointer userData) {
        bobgui_stream_subscribe(handle(), callback, userData);
    }

private:
    Stream(BobguiStream* handle) : core::ObjectHandle<BobguiStream>(handle) {}
};

} // namespace core_subsystems
} // namespace bobtk
