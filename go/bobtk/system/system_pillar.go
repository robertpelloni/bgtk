package system

import "fmt"

// SystemPillar manages low-level OS integrations, servers, system tray, and file watchers.
type SystemPillar struct {
	initialized bool
	Servers     *ServerManager
	Tray        *SystemTray
	FSWatcher   *FileSystemWatcher
}

func NewSystemPillar() *SystemPillar {
	sm := NewServerManager()
	return &SystemPillar{
		initialized: false,
		Servers:     sm,
		Tray:        NewSystemTray("utilities-system-monitor", sm),
		FSWatcher:   NewFileSystemWatcher(),
	}
}

func (s *SystemPillar) Initialize() {
	if s.initialized {
		return
	}
	_ = s.FSWatcher.AddPath("/app/config")
	fmt.Println("SystemPillar initialized.")
	s.initialized = true
}

func (s *SystemPillar) ServerManager() *ServerManager {
	return s.Servers
}

func (s *SystemPillar) SystemTray() *SystemTray {
	return s.Tray
}

func (s *SystemPillar) FileWatcher() *FileSystemWatcher {
	return s.FSWatcher
}
