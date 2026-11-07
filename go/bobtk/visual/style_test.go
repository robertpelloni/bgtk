package visual

import (
	"testing"
)

func TestStyleSheetRuleManagement(t *testing.T) {
	ss := NewStyleSheet()

	decl := StyleDeclaration{
		BackgroundColor: 0xFF112233,
		BorderColor:     0xFFAABBCC,
		BorderWidth:     2,
		CornerRadius:    6,
		Padding:         8,
		TextColor:       0xFFFFFFFF,
	}

	ss.SetRule("Button", decl)

	retrieved, ok := ss.GetRule("Button")
	if !ok {
		t.Fatalf("expected style rule for Button to exist")
	}

	if retrieved.BackgroundColor != 0xFF112233 {
		t.Errorf("expected bg color 0xFF112233, got 0x%X", retrieved.BackgroundColor)
	}
	if retrieved.BorderWidth != 2 {
		t.Errorf("expected border width 2, got %d", retrieved.BorderWidth)
	}
	if retrieved.CornerRadius != 6 {
		t.Errorf("expected corner radius 6, got %d", retrieved.CornerRadius)
	}
}

func TestStyleSheetParseSimpleCSS(t *testing.T) {
	ss := NewStyleSheet()

	css := `
		Button {
			background-color: #336699;
			border-color: #112233;
			border-width: 3px;
			corner-radius: 8px;
			padding: 10px;
			color: #FFFFFF;
		}
		Label {
			color: #FF0000;
			padding: 4px;
		}
	`

	err := ss.ParseSimpleCSS(css)
	if err != nil {
		t.Fatalf("unexpected error parsing CSS: %v", err)
	}

	btnStyle, ok := ss.GetRule("Button")
	if !ok {
		t.Fatalf("expected Button rule to be parsed")
	}

	if btnStyle.BorderWidth != 3 {
		t.Errorf("expected Button border width 3, got %d", btnStyle.BorderWidth)
	}
	if btnStyle.CornerRadius != 8 {
		t.Errorf("expected Button corner radius 8, got %d", btnStyle.CornerRadius)
	}
	if btnStyle.Padding != 10 {
		t.Errorf("expected Button padding 10, got %d", btnStyle.Padding)
	}
	if btnStyle.BackgroundColor != 0xFF336699 {
		t.Errorf("expected Button background color 0xFF336699, got 0x%X", btnStyle.BackgroundColor)
	}

	lblStyle, ok := ss.GetRule("Label")
	if !ok {
		t.Fatalf("expected Label rule to be parsed")
	}
	if lblStyle.TextColor != 0xFFFF0000 {
		t.Errorf("expected Label text color 0xFFFF0000, got 0x%X", lblStyle.TextColor)
	}
	if lblStyle.Padding != 4 {
		t.Errorf("expected Label padding 4, got %d", lblStyle.Padding)
	}
}
