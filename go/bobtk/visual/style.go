package visual

import (
	"strconv"
	"strings"
	"sync"
)

// StyleDeclaration holds concrete CSS-like visual properties for widgets.
type StyleDeclaration struct {
	BackgroundColor uint32
	BorderColor     uint32
	BorderWidth     int
	CornerRadius    int
	Padding         int
	TextColor       uint32
}

// StyleSheet manages cascading CSS-like style rules matching widget selectors.
type StyleSheet struct {
	mu    sync.RWMutex
	rules map[string]StyleDeclaration
}

// NewStyleSheet instantiates an empty stylesheet engine.
func NewStyleSheet() *StyleSheet {
	return &StyleSheet{
		rules: make(map[string]StyleDeclaration),
	}
}

// SetRule associates a style rule with a selector (e.g. "Button", "Label", "#submit-btn", ".danger").
func (ss *StyleSheet) SetRule(selector string, style StyleDeclaration) {
	ss.mu.Lock()
	defer ss.mu.Unlock()
	ss.rules[strings.TrimSpace(selector)] = style
}

// GetRule resolves the style declaration for a given selector.
func (ss *StyleSheet) GetRule(selector string) (StyleDeclaration, bool) {
	ss.mu.RLock()
	defer ss.mu.RUnlock()
	rule, exists := ss.rules[strings.TrimSpace(selector)]
	return rule, exists
}

// ParseSimpleCSS parses basic CSS declarations such as:
// "Button { background-color: #336699; border-width: 2; padding: 4; }"
func (ss *StyleSheet) ParseSimpleCSS(cssText string) error {
	ss.mu.Lock()
	defer ss.mu.Unlock()

	blocks := strings.Split(cssText, "}")
	for _, block := range blocks {
		block = strings.TrimSpace(block)
		if block == "" {
			continue
		}
		parts := strings.SplitN(block, "{", 2)
		if len(parts) != 2 {
			continue
		}
		selector := strings.TrimSpace(parts[0])
		body := strings.TrimSpace(parts[1])

		decl := ss.rules[selector]
		lines := strings.Split(body, ";")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}
			kv := strings.SplitN(line, ":", 2)
			if len(kv) != 2 {
				continue
			}
			k := strings.ToLower(strings.TrimSpace(kv[0]))
			v := strings.TrimSpace(kv[1])

			switch k {
			case "background-color":
				decl.BackgroundColor = parseHexColor(v)
			case "border-color":
				decl.BorderColor = parseHexColor(v)
			case "text-color", "color":
				decl.TextColor = parseHexColor(v)
			case "border-width":
				decl.BorderWidth = parseIntValue(v)
			case "corner-radius":
				decl.CornerRadius = parseIntValue(v)
			case "padding":
				decl.Padding = parseIntValue(v)
			}
		}
		ss.rules[selector] = decl
	}
	return nil
}

func parseHexColor(hexStr string) uint32 {
	hexStr = strings.TrimPrefix(hexStr, "#")
	if len(hexStr) == 6 {
		hexStr = "FF" + hexStr
	}
	val, err := strconv.ParseUint(hexStr, 16, 32)
	if err != nil {
		return 0
	}
	return uint32(val)
}

func parseIntValue(numStr string) int {
	numStr = strings.TrimSuffix(numStr, "px")
	val, err := strconv.Atoi(strings.TrimSpace(numStr))
	if err != nil {
		return 0
	}
	return val
}
