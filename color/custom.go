package color

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

// Custom is a user-defined 24-bit (truecolor) terminal color.
// Use RGB or Hex to build one, or Register/Lookup for named custom colors.
type Custom struct {
	r, g, b uint8
	bg      bool
}

var (
	namedMu sync.RWMutex
	named   = map[string]Custom{}
)

// RGB returns a custom foreground color from 8-bit red, green, and blue components.
func RGB(r, g, b uint8) Custom {
	return Custom{r: r, g: g, b: b}
}

// Hex parses a hex color into a custom foreground color.
// Accepted forms: "#RGB", "#RRGGBB", "RGB", "RRGGBB" (case-insensitive).
func Hex(s string) (Custom, error) {
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, "#")
	if len(s) != 3 && len(s) != 6 {
		return Custom{}, fmt.Errorf("color: invalid hex color %q", s)
	}
	if len(s) == 3 {
		s = string([]byte{s[0], s[0], s[1], s[1], s[2], s[2]})
	}
	n, err := strconv.ParseUint(s, 16, 32)
	if err != nil {
		return Custom{}, fmt.Errorf("color: invalid hex color %q: %w", s, err)
	}
	return RGB(uint8(n>>16), uint8(n>>8), uint8(n)), nil
}

// MustHex is like Hex but panics if s is not a valid hex color.
func MustHex(s string) Custom {
	c, err := Hex(s)
	if err != nil {
		panic(err)
	}
	return c
}

// Background returns the same RGB values as a background color.
func (c Custom) Background() Custom {
	c.bg = true
	return c
}

// Foreground returns the same RGB values as a foreground (font) color.
func (c Custom) Foreground() Custom {
	c.bg = false
	return c
}

// Register associates name with a custom color for later Lookup.
// Names are case-insensitive and must be non-empty.
func Register(name string, c Custom) error {
	name = strings.ToLower(strings.TrimSpace(name))
	if name == "" {
		return fmt.Errorf("color: empty custom color name")
	}
	namedMu.Lock()
	named[name] = c
	namedMu.Unlock()
	return nil
}

// MustRegister is like Register but panics on error.
func MustRegister(name string, c Custom) {
	if err := Register(name, c); err != nil {
		panic(err)
	}
}

// Lookup returns a previously registered custom color by name.
func Lookup(name string) (Custom, bool) {
	name = strings.ToLower(strings.TrimSpace(name))
	namedMu.RLock()
	c, ok := named[name]
	namedMu.RUnlock()
	return c, ok
}

// Unregister removes a named custom color. It is a no-op if the name is unknown.
func Unregister(name string) {
	name = strings.ToLower(strings.TrimSpace(name))
	namedMu.Lock()
	delete(named, name)
	namedMu.Unlock()
}

// Sequence returns the ANSI escape sequence for this custom color (without reset).
func (c Custom) Sequence() string {
	if c.bg {
		return fmt.Sprintf("\x1b[48;2;%d;%d;%dm", c.r, c.g, c.b)
	}
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", c.r, c.g, c.b)
}

// Print writes colored text with no newline.
func (c Custom) Print(a string) {
	PrintCustom(c, a)
}

// Println writes colored text with a newline.
func (c Custom) Println(a string) {
	PrintlnCustom(c, a)
}

// Sprint returns colored text as a string.
func (c Custom) Sprint(a string) string {
	return SprintCustom(c, a)
}

func addCustomColor(c Custom, a string) string {
	return c.Sequence() + a + "\x1b[0m"
}
