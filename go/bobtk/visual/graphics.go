package visual

import (
	"fmt"
	"sync"
)

// Color represents an RGBA color value.
type Color struct {
	R, G, B, A uint8
}

// Predefined colors (juce::Colours & Qt::GlobalColor parity).
var (
	ColorTransparent = Color{0, 0, 0, 0}
	ColorBlack       = Color{0, 0, 0, 255}
	ColorWhite       = Color{255, 255, 255, 255}
	ColorRed         = Color{255, 0, 0, 255}
	ColorGreen       = Color{0, 255, 0, 255}
	ColorBlue        = Color{0, 0, 255, 255}
)

func NewRGB(r, g, b uint8) Color {
	return Color{R: r, G: g, B: b, A: 255}
}

func NewRGBA(r, g, b, a uint8) Color {
	return Color{R: r, G: g, B: b, A: a}
}

// Point represents a 2D integer coordinate.
type Point struct {
	X, Y int
}

// Size represents 2D dimensions.
type Size struct {
	Width, Height int
}

// PathCommand represents an individual segment in a vector path.
type PathCommand struct {
	Type string
	P1   Point
	P2   Point
}

// Path represents a 2D vector path (juce::Path & QPainterPath parity).
type Path struct {
	commands []PathCommand
}

func NewPath() *Path {
	return &Path{commands: make([]PathCommand, 0)}
}

func (p *Path) StartNewSubPath(x, y int) {
	p.commands = append(p.commands, PathCommand{
		Type: "MOVE_TO",
		P1:   Point{X: x, Y: y},
	})
}

func (p *Path) LineTo(x, y int) {
	p.commands = append(p.commands, PathCommand{
		Type: "LINE_TO",
		P1:   Point{X: x, Y: y},
	})
}

func (p *Path) CloseSubPath() {
	p.commands = append(p.commands, PathCommand{
		Type: "CLOSE",
	})
}

func (p *Path) CommandCount() int {
	return len(p.commands)
}

// DrawOp represents a recorded draw command.
type DrawOp struct {
	Type  string
	Color Color
	Rect  Rect
	Text  string
	Path  *Path
}

// GraphicsContext provides 2D canvas drawing methods (juce::Graphics & QPainter parity).
type GraphicsContext struct {
	mu          sync.Mutex
	currentColor Color
	clipRect    Rect
	operations  []DrawOp
}

// NewGraphicsContext creates a graphics context for rendering.
func NewGraphicsContext(bounds Rect) *GraphicsContext {
	return &GraphicsContext{
		currentColor: ColorWhite,
		clipRect:     bounds,
		operations:   make([]DrawOp, 0),
	}
}

func (g *GraphicsContext) SetColour(col Color) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.currentColor = col
}

func (g *GraphicsContext) DrawRect(x, y, w, h int) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.operations = append(g.operations, DrawOp{
		Type:  "DRAW_RECT",
		Color: g.currentColor,
		Rect:  Rect{X: x, Y: y, W: w, H: h},
	})
}

func (g *GraphicsContext) FillRect(x, y, w, h int) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.operations = append(g.operations, DrawOp{
		Type:  "FILL_RECT",
		Color: g.currentColor,
		Rect:  Rect{X: x, Y: y, W: w, H: h},
	})
}

func (g *GraphicsContext) DrawText(text string, x, y, w, h int) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.operations = append(g.operations, DrawOp{
		Type:  "DRAW_TEXT",
		Color: g.currentColor,
		Rect:  Rect{X: x, Y: y, W: w, H: h},
		Text:  text,
	})
}

func (g *GraphicsContext) StrokePath(p *Path) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.operations = append(g.operations, DrawOp{
		Type:  "STROKE_PATH",
		Color: g.currentColor,
		Path:  p,
	})
}

func (g *GraphicsContext) OperationCount() int {
	g.mu.Lock()
	defer g.mu.Unlock()
	return len(g.operations)
}

func (g *GraphicsContext) Clear() {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.operations = g.operations[:0]
}

func (g *GraphicsContext) LastOperation() (DrawOp, error) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if len(g.operations) == 0 {
		return DrawOp{}, fmt.Errorf("no recorded operations")
	}
	return g.operations[len(g.operations)-1], nil
}
