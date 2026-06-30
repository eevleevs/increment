package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		input  string
		amount int
		want   string
	}{
		// ── Decimal ──────────────────────────────────────────────
		{"42", 1, "43"},
		{"42", -1, "41"},
		{"0", 1, "1"},
		{"0", -1, "-1"},
		{"1", -1, "0"},
		{"-5", 1, "-4"},
		{"-5", -1, "-6"},
		{"9999999999999999", 1, "10000000000000000"},
		{"100", 0, "100"},

		// ── Hex ─────────────────────────────────────────────────
		{"0xFF", 1, "0x100"},
		{"0xFF", -1, "0xFE"},
		{"0x00", 1, "0x1"},
		{"0x100", -1, "0xFF"},
		{"0xFE", 1, "0xFF"},
		{"0x1", -1, "0x0"},
		{"0xA", 10, "0x14"},
		{"0xAB", 1, "0xAC"},
		{"0x1F", 1, "0x20"},

		// ── Binary ──────────────────────────────────────────────
		{"0b1010", 1, "0b1011"},
		{"0b1010", -1, "0b1001"},
		{"0b0000", 1, "0b1"},
		{"0b1", -1, "0b0"},
		{"0b1111", 1, "0b10000"},
		{"0b0", -1, "-0b1"},

		// ── Octal ───────────────────────────────────────────────
		{"0o77", 1, "0o100"},
		{"0o77", -1, "0o76"},
		{"0o00", 1, "0o1"},
		{"0o7", 1, "0o10"},
		{"0o10", -1, "0o7"},

		// ── Toggle (odd amounts) ────────────────────────────────
		{"true", 1, "false"},
		{"false", 1, "true"},
		{"True", 1, "False"},
		{"False", 1, "True"},
		{"TRUE", 1, "FALSE"},
		{"FALSE", 1, "TRUE"},
		{"yes", 1, "no"},
		{"no", 1, "yes"},
		{"Yes", 1, "No"},
		{"No", 1, "Yes"},
		{"YES", 1, "NO"},
		{"NO", 1, "YES"},
		{"on", 1, "off"},
		{"off", 1, "on"},
		{"On", 1, "Off"},
		{"Off", 1, "On"},
		{"ON", 1, "OFF"},
		{"OFF", 1, "ON"},
		{"enable", 1, "disable"},
		{"disable", 1, "enable"},
		{"enabled", 1, "disabled"},
		{"disabled", 1, "enabled"},

		// toggle with -1 (odd)
		{"true", -1, "false"},
		{"false", -1, "true"},
		{"enable", -1, "disable"},

		// ── Even amounts do NOT toggle ─────────────────────────
		{"true", 2, "true"},
		{"false", 0, "false"},
		{"enable", 2, "enable"},
		{"yes", -2, "yes"},

		// ── Unknown values left unchanged ─────────────────────
		{"hello", 1, "hello"},
		{"hello", -1, "hello"},
		{"42abc", 1, "42abc"},
		{"0xyz", 1, "0xyz"},

		// ── Empty input ───────────────────────────────────────
		{"", 1, ""},
		{"", -1, ""},

		// ── Large toggle amounts (odd=toggle, even=no-op) ────
		{"true", 3, "false"},
		{"true", 5, "false"},
		{"true", 4, "true"},
		{"true", -5, "false"},
	}

	for _, tc := range tests {
		got := Run(tc.input, tc.amount)
		if got != tc.want {
			t.Errorf("Run(%q, %d) = %q, want %q", tc.input, tc.amount, got, tc.want)
		}
	}
}
