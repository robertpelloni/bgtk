#pragma once

#include <bobgui/cpp/module.hpp>
#include <bobgui/cpp/network.hpp>
#include <memory>
#include <string>

namespace bobtk {
namespace module {

// The Network Pillar interface acts as the service registry for all connectivity modules
class NetworkPillar : public Module {
public:
    NetworkPillar() : Module("Network", "1.0.0") {}

    // HTTP / WebSocket Subsystem
    std::shared_ptr<network::NetworkAccessManager> getAccessManager() {
        if (!accessManager) {
            accessManager = std::make_shared<network::NetworkAccessManager>();
        }
        return accessManager;
    }

    std::shared_ptr<network::WebSocket> createWebSocket(const std::string& url) {
        return std::make_shared<network::WebSocket>(url);
    }

    // Remote / VNC Subsystem
    std::shared_ptr<network::RemoteServer> createRemoteServer(int port, const std::string& codec) {
        return std::make_shared<network::RemoteServer>(port, codec);
    }

    std::shared_ptr<network::RemoteClient> createRemoteClient(const std::string& host, int port) {
        return std::make_shared<network::RemoteClient>(host, port);
    }

    // Web View Subsystem
    std::shared_ptr<network::WebView> createWebView() {
        return std::make_shared<network::WebView>();
    }

    // Cloud Context Subsystem
    std::shared_ptr<network::CloudContext> createCloudContext(const std::string& apiUrl, const std::string& apiKey) {
        return std::make_shared<network::CloudContext>(apiUrl, apiKey);
    }

    // Blockchain Subsystem
    std::shared_ptr<network::BlockchainManager> getBlockchainManager() {
        return network::BlockchainManager::getDefault();
    }

    // Multiplayer / Sync Engine Subsystem
    std::shared_ptr<network::SyncEngine> createSyncEngine(BobguiStateMachine* stateMachine) {
        return std::make_shared<network::SyncEngine>(stateMachine);
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

private:
    std::shared_ptr<network::NetworkAccessManager> accessManager;
};

} // namespace module
} // namespace bobtk
