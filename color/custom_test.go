package color

import (
	"bytes"
	"strings"
	"testing"
)

func TestHex(t *testing.T) {
	c, err := Hex("#FF8000")
	if err != nil {
		t.Fatalf("Hex: %v", err)
	}
	if c.r != 255 || c.g != 128 || c.b != 0 || c.bg {
		t.Fatalf("Hex RGB = %+v, want r=255 g=128 b=0 fg", c)
	}

	c, err = Hex("0af")
	if err != nil {
		t.Fatalf("Hex short: %v", err)
	}
	if c.r != 0 || c.g != 0xaa || c.b != 0xff {
		t.Fatalf("Hex short RGB = %+v, want r=0 g=170 b=255", c)
	}

	if _, err := Hex("zz"); err == nil {
		t.Fatal("Hex expected error for invalid input")
	}
	if _, err := Hex("#12345"); err == nil {
		t.Fatal("Hex expected error for wrong length")
	}
}

func TestMustHex(t *testing.T) {
	c := MustHex("#112233")
	if c.r != 0x11 || c.g != 0x22 || c.b != 0x33 {
		t.Fatalf("MustHex = %+v", c)
	}
	defer func() {
		if recover() == nil {
			t.Fatal("MustHex should panic on invalid input")
		}
	}()
	MustHex("nope")
}

func TestRGBSequenceAndSprint(t *testing.T) {
	fg := RGB(1, 2, 3)
	got := fg.Sprint("hi")
	wantPrefix := "\x1b[38;2;1;2;3mhi\x1b[0m"
	if got != wantPrefix {
		t.Fatalf("Sprint = %q, want %q", got, wantPrefix)
	}

	bg := RGB(10, 20, 30).Background()
	got = SprintCustom(bg, "bg")
	want := "\x1b[48;2;10;20;30mbg\x1b[0m"
	if got != want {
		t.Fatalf("SprintCustom bg = %q, want %q", got, want)
	}

	if fg.Background().Foreground().Sequence() != fg.Sequence() {
		t.Fatal("Foreground should clear background flag")
	}
}

func TestRegisterLookup(t *testing.T) {
	name := "brand-orange-test"
	Unregister(name)
	defer Unregister(name)

	if err := Register("", RGB(1, 1, 1)); err == nil {
		t.Fatal("Register empty name should fail")
	}

	c := MustHex("#FF6600")
	if err := Register(name, c); err != nil {
		t.Fatalf("Register: %v", err)
	}

	got, ok := Lookup("Brand-Orange-Test")
	if !ok {
		t.Fatal("Lookup failed")
	}
	if got.r != c.r || got.g != c.g || got.b != c.b {
		t.Fatalf("Lookup = %+v, want %+v", got, c)
	}

	Unregister(name)
	if _, ok := Lookup(name); ok {
		t.Fatal("Lookup should miss after Unregister")
	}
}

func TestFprintCustom(t *testing.T) {
	var buf bytes.Buffer
	n, err := FprintCustom(RGB(255, 0, 0), &buf, "red")
	if err != nil {
		t.Fatalf("FprintCustom: %v", err)
	}
	if n == 0 {
		t.Fatal("FprintCustom wrote 0 bytes")
	}
	if !strings.Contains(buf.String(), "\x1b[38;2;255;0;0mred\x1b[0m") {
		t.Fatalf("unexpected output %q", buf.String())
	}

	buf.Reset()
	_, err = FprintlnCustom(MustHex("#00ff00"), &buf, "green")
	if err != nil {
		t.Fatalf("FprintlnCustom: %v", err)
	}
	out := buf.String()
	if !strings.Contains(out, "\x1b[38;2;0;255;0mgreen\x1b[0m") || !strings.HasSuffix(out, "\n") {
		t.Fatalf("unexpected FprintlnCustom output %q", out)
	}
}

func TestPresetSprintUnchanged(t *testing.T) {
	got := Sprint(RedText, "x")
	want := "\x1b[0;31mx\x1b[0m"
	if got != want {
		t.Fatalf("preset Sprint = %q, want %q", got, want)
	}
}
