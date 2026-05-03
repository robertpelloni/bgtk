#pragma once

#include <bobgui/bobgui.h>
#include <bobgui/cpp/object_handle.hpp>
#include <bobgui/cpp/widget.hpp>
#include <string>
#include <memory>

namespace bobtk {
namespace network {

// ---------------------------------------------------------
// BASIC HTTP / WEB SOCKET SUBSYSTEM
// ---------------------------------------------------------

class NetworkAccessManager : public core::ObjectHandle<BobguiNetworkAccessManager> {
public:
    NetworkAccessManager()
        : core::ObjectHandle<BobguiNetworkAccessManager>(bobgui_network_access_manager_new()) {}

    void get(const std::string& url, GAsyncReadyCallback callback, gpointer userData) {
        bobgui_network_get(handle(), url.c_str(), callback, userData);
    }

    void post(const std::string& url, const std::string& data, GAsyncReadyCallback callback, gpointer userData) {
        bobgui_network_post(handle(), url.c_str(), data.c_str(), callback, userData);
    }
};

class WebSocket : public core::ObjectHandle<BobguiWebSocket> {
public:
    explicit WebSocket(const std::string& url)
        : core::ObjectHandle<BobguiWebSocket>(bobgui_web_socket_new(url.c_str())) {}

    void send(const std::string& message) {
        bobgui_web_socket_send(handle(), message.c_str());
    }
};

// ---------------------------------------------------------
// REMOTE / VNC SUBSYSTEM
// ---------------------------------------------------------

class RemoteServer : public core::ObjectHandle<BobguiRemoteServer> {
public:
    RemoteServer(int port, const std::string& codec)
        : core::ObjectHandle<BobguiRemoteServer>(bobgui_remote_server_new(port, codec.c_str())) {}

    void start(Widget& uiRoot) {
        bobgui_remote_server_start(handle(), BOBGUI_WIDGET(uiRoot.handle()));
    }

    void setQuality(int bitRate, int frameRate) {
        bobgui_remote_server_set_quality(handle(), bitRate, frameRate);
    }
};

class RemoteClient : public core::ObjectHandle<BobguiRemoteClient> {
public:
    RemoteClient(const std::string& host, int port)
        : core::ObjectHandle<BobguiRemoteClient>(bobgui_remote_client_new(host.c_str(), port)) {}

    void connect(Widget& localContainer) {
        bobgui_remote_client_connect(handle(), BOBGUI_WIDGET(localContainer.handle()));
    }
};

// ---------------------------------------------------------
// WEB VIEW SUBSYSTEM
// ---------------------------------------------------------

class WebView : public Widget {
public:
    WebView() : Widget(bobgui_web_view_new()) {}

    void loadUri(const std::string& uri) {
        bobgui_web_view_load_uri(BOBGUI_WEB_VIEW(handle()), uri.c_str());
    }

    void loadHtml(const std::string& html, const std::string& baseUri) {
        bobgui_web_view_load_html(BOBGUI_WEB_VIEW(handle()), html.c_str(), baseUri.c_str());
    }

    void exposeObject(const std::string& name, GObject* object) {
        bobgui_web_view_expose_object(BOBGUI_WEB_VIEW(handle()), name.c_str(), object);
    }

    void runJavascript(const std::string& script, GAsyncReadyCallback callback, gpointer userData) {
        bobgui_web_view_run_javascript(BOBGUI_WEB_VIEW(handle()), script.c_str(), callback, userData);
    }
};

// ---------------------------------------------------------
// SUPABASE / FIREBASE CLOUD SUBSYSTEM
// ---------------------------------------------------------

class CloudContext : public core::ObjectHandle<BobguiCloudContext> {
public:
    CloudContext(const std::string& apiUrl, const std::string& apiKey)
        : core::ObjectHandle<BobguiCloudContext>(bobgui_cloud_context_new(apiUrl.c_str(), apiKey.c_str())) {}

    void authSignUp(const std::string& email, const std::string& pass, GAsyncReadyCallback callback) {
        bobgui_cloud_auth_sign_up(handle(), email.c_str(), pass.c_str(), callback);
    }

    void authSignIn(const std::string& email, const std::string& pass, GAsyncReadyCallback callback) {
        bobgui_cloud_auth_sign_in(handle(), email.c_str(), pass.c_str(), callback);
    }

    void dbSync(const std::string& table, GAsyncReadyCallback callback) {
        bobgui_cloud_db_sync(handle(), table.c_str(), callback);
    }

    void uploadAsset(const std::string& path, GAsyncReadyCallback callback) {
        bobgui_cloud_upload_asset(handle(), path.c_str(), callback);
    }
};

// ---------------------------------------------------------
// WEB3 / BLOCKCHAIN SUBSYSTEM
// ---------------------------------------------------------

class BlockchainManager : public core::ObjectHandle<BobguiBlockchainManager> {
public:
    static std::shared_ptr<BlockchainManager> getDefault() {
        return std::shared_ptr<BlockchainManager>(new BlockchainManager(bobgui_blockchain_manager_get_default()));
    }

    void connectWallet(const std::string& provider) {
        bobgui_blockchain_connect_wallet(handle(), provider.c_str());
    }

    void signMessage(const std::string& message, GAsyncReadyCallback callback) {
        bobgui_blockchain_sign_message(handle(), message.c_str(), callback);
    }

    void callContract(const std::string& address, const std::string& abi, const std::string& method, const std::string& args, GAsyncReadyCallback callback) {
        bobgui_blockchain_call_contract(handle(), address.c_str(), abi.c_str(), method.c_str(), args.c_str(), callback);
    }

    void uploadIpfs(const std::string& filePath, GAsyncReadyCallback callback) {
        bobgui_blockchain_upload_ipfs(handle(), filePath.c_str(), callback);
    }

private:
    BlockchainManager(BobguiBlockchainManager* handle) : core::ObjectHandle<BobguiBlockchainManager>(handle) {}
};

// ---------------------------------------------------------
// MULTIPLAYER / SYNC ENGINE SUBSYSTEM
// ---------------------------------------------------------

class SyncEngine : public core::ObjectHandle<BobguiSyncEngine> {
public:
    explicit SyncEngine(BobguiStateMachine* stateMachine)
        : core::ObjectHandle<BobguiSyncEngine>(bobgui_sync_engine_new(stateMachine)) {}

    void connectTo(const std::string& peerAddress) {
        bobgui_sync_engine_connect(handle(), peerAddress.c_str());
    }

    void broadcast(GBytes* delta) {
        bobgui_sync_engine_broadcast(handle(), delta);
    }

    void trackCursor(const std::string& clientId, float x, float y) {
        bobgui_sync_engine_track_cursor(handle(), clientId.c_str(), x, y);
    }
};

} // namespace network
} // namespace bobtk
