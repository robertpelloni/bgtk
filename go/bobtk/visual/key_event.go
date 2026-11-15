package visual

// KeyModifier bitmask (Qt::KeyboardModifier parity).
type KeyModifier int

const (
	NoModifier      KeyModifier = 0
	ShiftModifier   KeyModifier = 1 << 0
	ControlModifier KeyModifier = 1 << 1
	AltModifier     KeyModifier = 1 << 2
	MetaModifier    KeyModifier = 1 << 3
)

// KeyEvent represents a key press or release event (QKeyEvent parity).
type KeyEvent struct {
	Key       int
	Modifiers KeyModifier
	Text      string
	Accepted  bool
}

// Accept marks the key event as handled.
func (k *KeyEvent) Accept() {
	k.Accepted = true
}

// Ignore marks the key event as unhandled, propagating up the focus tree.
func (k *KeyEvent) Ignore() {
	k.Accepted = false
}

// HasModifier checks whether a specific modifier is active.
func (k *KeyEvent) HasModifier(mod KeyModifier) bool {
	return (k.Modifiers & mod) != 0
}
