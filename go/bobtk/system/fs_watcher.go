package system

import (
	"fmt"
	"sync"
)

// FileWatcherEvent describes an OS file or directory mutation event.
type FileWatcherEvent struct {
	Path string
	Op   string
}

// FileSystemWatcher monitors files and directories for modification (QFileSystemWatcher parity).
type FileSystemWatcher struct {
	mu        sync.RWMutex
	watched   map[string]bool
	listeners []func(FileWatcherEvent)
}

// NewFileSystemWatcher instantiates a file system watcher.
func NewFileSystemWatcher() *FileSystemWatcher {
	return &FileSystemWatcher{
		watched:   make(map[string]bool),
		listeners: make([]func(FileWatcherEvent), 0),
	}
}

// AddPath adds a file or directory path to the watch list.
func (f *FileSystemWatcher) AddPath(path string) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.watched[path] {
		return fmt.Errorf("path %q is already being watched", path)
	}

	f.watched[path] = true
	return nil
}

// RemovePath stops watching a specific path.
func (f *FileSystemWatcher) RemovePath(path string) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	if !f.watched[path] {
		return fmt.Errorf("path %q not in watch list", path)
	}

	delete(f.watched, path)
	return nil
}

// WatchedPaths returns a slice of all monitored paths.
func (f *FileSystemWatcher) WatchedPaths() []string {
	f.mu.RLock()
	defer f.mu.RUnlock()

	paths := make([]string, 0, len(f.watched))
	for p := range f.watched {
		paths = append(paths, p)
	}
	return paths
}

// OnFileChanged registers an event listener for file changes.
func (f *FileSystemWatcher) OnFileChanged(cb func(FileWatcherEvent)) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.listeners = append(f.listeners, cb)
}

// TriggerEvent simulates a file mutation event dispatched to all listeners.
func (f *FileSystemWatcher) TriggerEvent(path, op string) {
	f.mu.RLock()
	if !f.watched[path] {
		f.mu.RUnlock()
		return
	}
	cbs := append([]func(FileWatcherEvent){}, f.listeners...)
	f.mu.RUnlock()

	ev := FileWatcherEvent{Path: path, Op: op}
	for _, cb := range cbs {
		cb(ev)
	}
}
