package main

import (
	"testing"

	"golang.org/x/text/language"
)

func TestDecimal_String(t *testing.T) {
	type args struct {
		t language.Tag
	}
	en := args{language.English}
	de := args{language.German}
	fr := args{language.French}
	tests := []struct {
		name string
		n    Decimal
		args args
		want string
	}{
		// bytes
		{"0 en", 0, en, "0"},
		{"-0 en", -0, en, "0"},
		{"1 en", 1, en, "1"},
		{"-1 en", -1, en, "-1"},
		// Kilobytes
		{"1k en", 1_000, en, "1.0\u00A0kB"},
		{"1k de", 1_000, de, "1,0\u00A0kB"},
		{"-1k en", -1_000, en, "-1.0\u00A0kB"},
		{"-1k de", -1_000, de, "-1,0\u00A0kB"},
		// Megabytes
		{"1M en", 1_234_567, en, "1.23\u00A0MB"},
		{"1M de", 1_234_567, de, "1,23\u00A0MB"},
		// Gigabytes
		{"1G en", 1_234_567_890, en, "1.23\u00A0GB"},
		{"1G de", 1_234_567_890, de, "1,23\u00A0GB"},
		// Terabytes
		{"10T en", 9_843_234_567_890, en, "9.84\u00A0TB"},
		{"10T de", 9_843_234_567_890, de, "9,84\u00A0TB"},
		// Petabytes
		{"1P en", 1_000_234_567_890_000, en, "1.00\u00A0PB"},
		{"1P de", 1_000_234_567_890_000, de, "1,00\u00A0PB"},
		// Exabytes
		{"5E en", 5_000_000_234_567_890_000, en, "5.00\u00A0EB"},
		{"5E de", 5_000_000_234_567_890_000, de, "5,00\u00A0EB"},
		{"max int en", 9_223_372_036_854_775_807, en, "9.22\u00A0EB"},
		{"max int de", 9_223_372_036_854_775_807, de, "9,22\u00A0EB"},
		{"-max int en", -9_223_372_036_854_775_807, en, "-9.22\u00A0EB"},
		{"-max int de", -9_223_372_036_854_775_807, de, "-9,22\u00A0EB"},
		{"float en", 9_223_372_036_854_775_807_000_000, en, "9,223,372.04\u00A0EB"},
		{"float de", 9_223_372_036_854_775_807_000_000, de, "9.223.372,04\u00A0EB"},
		{"float fr", 9_223_372_036_854_775_807_000_000, fr, "9\u00A0223\u00A0372,04\u00A0EB"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.String(tt.args.t); got != tt.want {
				t.Errorf("Decimal.String() = %v, want %v", []rune(got), []rune(tt.want))
			}
		})
	}
}

func TestBinary_String(t *testing.T) {
	type args struct {
		t language.Tag
	}
	en := args{language.English}
	de := args{language.German}
	fr := args{language.French}
	tests := []struct {
		name string
		n    Binary
		args args
		want string
	}{
		// bytes
		{"0 en", 0, en, "0"},
		{"-0 en", -0, en, "0"},
		{"1 en", 1, en, "1"},
		{"-1 en", -1, en, "-1"},
		// Kibibytes
		{"1k en", 1_024, en, "1.0\u00A0KiB"},
		{"1k de", 1_024, de, "1,0\u00A0KiB"},
		{"-1k en", -1_024, en, "-1.0\u00A0KiB"},
		{"-1k de", -1_024, de, "-1,0\u00A0KiB"},
		// Mebibytes
		{"1M en", 1_234_567, en, "1.18\u00A0MiB"},
		{"1M de", 1_234_567, de, "1,18\u00A0MiB"},
		// Gibibytes
		{"1G en", 1_234_567_890, en, "1.15\u00A0GiB"},
		{"1G de", 1_234_567_890, de, "1,15\u00A0GiB"},
		// Tebibytes
		{"10T en", 9_843_234_567_890, en, "8.95\u00A0TiB"},
		{"10T de", 9_843_234_567_890, de, "8,95\u00A0TiB"},
		// Pebibytes
		{"1P en", 1_200_234_567_890_000, en, "1.07\u00A0PiB"},
		{"1P de", 1_200_234_567_890_000, de, "1,07\u00A0PiB"},
		// Exbibytes
		{"4E en", 5_000_000_234_567_890_000, en, "4.34\u00A0EiB"},
		{"4E de", 5_000_000_234_567_890_000, de, "4,34\u00A0EiB"},
		{"max int en", 9_223_372_036_854_775_807, en, "8.00\u00A0EiB"},
		{"max int de", 9_223_372_036_854_775_807, de, "8,00\u00A0EiB"},
		{"-max int en", -9_223_372_036_854_775_807, en, "-8.00\u00A0EiB"},
		{"-max int de", -9_223_372_036_854_775_807, de, "-8,00\u00A0EiB"},
		{"float en", 9_223_372_036_854_775_807_000_000, en, "8,000,000.00\u00A0EiB"},
		{"float de", 9_223_372_036_854_775_807_000_000, de, "8.000.000,00\u00A0EiB"},
		{"float fr", 9_223_372_036_854_775_807_000_000, fr, "8\u00A0000\u00A0000,00\u00A0EiB"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.String(tt.args.t); got != tt.want {
				t.Errorf("Binary.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
