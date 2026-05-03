#pragma once

#include <bobgui/cpp/module.hpp>
#include <bobgui/cpp/core_subsystems.hpp>
#include <memory>
#include <string>

namespace bobtk {
namespace module {

// The Core Pillar interface acts as the service registry for underlying logic, data, and performance engines
class CorePillar : public Module {
public:
    CorePillar() : Module("Core", "1.0.0") {}

    std::shared_ptr<core_subsystems::AutonomousController> getAutonomousController() {
        return core_subsystems::AutonomousController::getDefault();
    }

    std::shared_ptr<core_subsystems::BrainNode> createBrainNode(const std::string& modelPath) {
        return std::make_shared<core_subsystems::BrainNode>(modelPath);
    }

    std::shared_ptr<core_subsystems::ComputePipeline> createComputePipeline(const std::string& kernelSource) {
        return std::make_shared<core_subsystems::ComputePipeline>(kernelSource);
    }

    std::shared_ptr<core_subsystems::SqlDatabase> createSqlDatabase(const std::string& driver) {
        return std::make_shared<core_subsystems::SqlDatabase>(driver);
    }

    std::shared_ptr<core_subsystems::EntityManager> getEntityManager() {
        return core_subsystems::EntityManager::getDefault();
    }

    std::shared_ptr<core_subsystems::FinChart> createFinChart(BobguiChartType type) {
        return std::make_shared<core_subsystems::FinChart>(type);
    }

    std::shared_ptr<core_subsystems::Translator> createTranslator(const std::string& domain) {
        return std::make_shared<core_subsystems::Translator>(domain);
    }

    std::shared_ptr<core_subsystems::ModernLayout> createModernLayout(BobguiLayoutType type) {
        return std::make_shared<core_subsystems::ModernLayout>(type);
    }

    std::shared_ptr<core_subsystems::MetaOrchestrator> getMetaOrchestrator() {
        return core_subsystems::MetaOrchestrator::getDefault();
    }

    std::shared_ptr<core_subsystems::QuantumScheduler> getQuantumScheduler() {
        return core_subsystems::QuantumScheduler::getDefault();
    }

    std::shared_ptr<core_subsystems::RealtimeKernel> getRealtimeKernel() {
        return core_subsystems::RealtimeKernel::getInstance();
    }

    std::shared_ptr<core_subsystems::ReflectContext> getReflectContext() {
        return core_subsystems::ReflectContext::getDefault();
    }

    std::shared_ptr<core_subsystems::ScriptContext> createScriptContext(const std::string& engineType) {
        return std::make_shared<core_subsystems::ScriptContext>(engineType);
    }

    std::shared_ptr<core_subsystems::StateMachine> createStateMachine(const core_subsystems::StateNode& initial) {
        return std::make_shared<core_subsystems::StateMachine>(initial);
    }

    std::shared_ptr<core_subsystems::Stream> createStreamInterval(int intervalMs) {
        return core_subsystems::Stream::interval(intervalMs);
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
