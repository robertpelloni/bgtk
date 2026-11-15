#ifndef BOBGUI_CPP_SYSTEM_TRAY_HPP
#define BOBGUI_CPP_SYSTEM_TRAY_HPP

#include <functional>
#include <map>
#include <memory>
#include <string>
#include <vector>

namespace bobgui {
namespace cpp {

/**
 * ServerProcessController:
 * Provides central registration, status monitoring, graceful shutdown,
 * and forceful termination of running background servers (such as Broadway,
 * IPC daemons, live-reloaders, or network services).
 */
class ServerProcessController
{
public:
  struct ServerInfo
  {
    std::string id;
    std::string name;
    std::string description;
    int port;
    bool is_running;
    std::function<void()> stop_handler;
  };

  ServerProcessController () = default;

  void register_server (const std::string &id,
                        const std::string &name,
                        const std::string &description,
                        int                port,
                        std::function<void()> stop_handler)
  {
    ServerInfo info;
    info.id = id;
    info.name = name;
    info.description = description;
    info.port = port;
    info.is_running = true;
    info.stop_handler = std::move (stop_handler);
    servers_[id] = std::move (info);
  }

  bool stop_server (const std::string &id)
  {
    auto it = servers_.find (id);
    if (it != servers_.end () && it->second.is_running)
      {
        if (it->second.stop_handler)
          {
            it->second.stop_handler ();
          }
        it->second.is_running = false;
        return true;
      }
    return false;
  }

  void stop_all_servers ()
  {
    for (auto &pair : servers_)
      {
        if (pair.second.is_running)
          {
            if (pair.second.stop_handler)
              {
                pair.second.stop_handler ();
              }
            pair.second.is_running = false;
          }
      }
  }

  std::vector<ServerInfo> get_all_servers () const
  {
    std::vector<ServerInfo> list;
    for (const auto &pair : servers_)
      {
        list.push_back (pair.second);
      }
    return list;
  }

  std::size_t running_server_count () const
  {
    std::size_t count = 0;
    for (const auto &pair : servers_)
      {
        if (pair.second.is_running)
          count++;
      }
    return count;
  }

private:
  std::map<std::string, ServerInfo> servers_;
};

/**
 * SystemTrayMenuItem:
 * Represents an entry in the system tray popup context menu.
 */
struct SystemTrayMenuItem
{
  std::string label;
  std::string tooltip;
  bool is_separator = false;
  std::function<void()> on_triggered;
};

/**
 * SystemTrayIcon:
 * Cross-platform abstraction for system tray / status area notification icon.
 * Synthesizes concepts from JUCE SystemTrayIconComponent and GTK StatusIcon.
 * Features server management integration and application termination.
 */
class SystemTrayIcon
{
public:
  typedef std::function<void()> QuitCallback;

  explicit SystemTrayIcon (const std::string &icon_name = "application-default-icon")
  : icon_name_ (icon_name),
    tooltip_ ("bgtk Application Status"),
    visible_ (true),
    server_controller_ (std::make_shared<ServerProcessController> ())
  {
  }

  void set_icon_name (const std::string &icon_name)
  {
    icon_name_ = icon_name;
  }

  const std::string &icon_name () const
  {
    return icon_name_;
  }

  void set_tooltip (const std::string &tooltip)
  {
    tooltip_ = tooltip;
  }

  const std::string &tooltip () const
  {
    return tooltip_;
  }

  void set_visible (bool visible)
  {
    visible_ = visible;
  }

  bool is_visible () const
  {
    return visible_;
  }

  std::shared_ptr<ServerProcessController> server_controller ()
  {
    return server_controller_;
  }

  void add_menu_item (const std::string &label,
                      const std::string &tooltip,
                      std::function<void()> on_triggered)
  {
    SystemTrayMenuItem item;
    item.label = label;
    item.tooltip = tooltip;
    item.is_separator = false;
    item.on_triggered = std::move (on_triggered);
    menu_items_.push_back (std::move (item));
  }

  void add_separator ()
  {
    SystemTrayMenuItem item;
    item.is_separator = true;
    menu_items_.push_back (std::move (item));
  }

  const std::vector<SystemTrayMenuItem> &menu_items () const
  {
    return menu_items_;
  }

  void set_quit_handler (QuitCallback handler)
  {
    quit_handler_ = std::move (handler);
  }

  void quit_application ()
  {
    if (server_controller_)
      {
        server_controller_->stop_all_servers ();
      }
    if (quit_handler_)
      {
        quit_handler_ ();
      }
  }

private:
  std::string icon_name_;
  std::string tooltip_;
  bool visible_;
  std::vector<SystemTrayMenuItem> menu_items_;
  std::shared_ptr<ServerProcessController> server_controller_;
  QuitCallback quit_handler_;
};

} /* namespace cpp */
} /* namespace bobgui */

#endif /* BOBGUI_CPP_SYSTEM_TRAY_HPP */
