package service

import "testing"

func TestNormalizeColor(t *testing.T) {
	got, err := normalizeColor("#7c5cff")
	if err != nil {
		t.Fatalf("normalizeColor returned error: %v", err)
	}

	if got != 8150271 {
		t.Fatalf("expected 8150271, got %d", got)
	}
}

func TestNormalizeColorRejectsInvalidHex(t *testing.T) {
	if _, err := normalizeColor("#not-a-color"); err == nil {
		t.Fatal("expected error for invalid color")
	}
}
