package media

import (
	"sync"
)

// UMPPacket represents a Universal MIDI Packet (MIDI 2.0 32-bit word packet).
type UMPPacket struct {
	Word0 uint32
	Word1 uint32
}

// MessageType extracts the 4-bit message type from Word0.
func (u UMPPacket) MessageType() uint8 {
	return uint8((u.Word0 >> 28) & 0x0F)
}

// Group extracts the 4-bit group number from Word0.
func (u UMPPacket) Group() uint8 {
	return uint8((u.Word0 >> 24) & 0x0F)
}

// Status extracts the MIDI 2.0 status code.
func (u UMPPacket) Status() uint8 {
	return uint8((u.Word0 >> 16) & 0xFF)
}

// NoteOn creates a MIDI 2.0 Note On packet with 16-bit high-resolution velocity.
func NewNoteOn(group, channel, note uint8, velocity16 uint16) UMPPacket {
	w0 := (uint32(0x4) << 28) | (uint32(group&0x0F) << 24) | (uint32(0x90|(channel&0x0F)) << 16) | (uint32(note) << 8)
	w1 := uint32(velocity16) << 16
	return UMPPacket{Word0: w0, Word1: w1}
}

// MidiBus handles dispatching UMP packets to listeners.
type MidiBus struct {
	mu        sync.RWMutex
	listeners []func(UMPPacket)
}

// NewMidiBus creates a MIDI 2.0 bus.
func NewMidiBus() *MidiBus {
	return &MidiBus{
		listeners: make([]func(UMPPacket), 0),
	}
}

// Subscribe registers a listener for incoming MIDI 2.0 events.
func (m *MidiBus) Subscribe(cb func(UMPPacket)) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.listeners = append(m.listeners, cb)
}

// Broadcast sends a UMP packet to all active subscribers.
func (m *MidiBus) Broadcast(pkt UMPPacket) {
	m.mu.RLock()
	cbs := append([]func(UMPPacket){}, m.listeners...)
	m.mu.RUnlock()

	for _, cb := range cbs {
		cb(pkt)
	}
}
