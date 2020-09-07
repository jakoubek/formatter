package formatter

import (
	"fmt"
	"testing"
)

func TestFormatInteger(t *testing.T) {
	var cases = []struct {
		number int64
		wanted string
	}{
		{55, "55"},
		{102, "102"},
		{999, "999"},
		{1000, "1.000"},
		{10001, "10.001"},
		{1000000, "1.000.000"},
		{2500000000, "2.500.000.000"},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%d", c.number), func(t *testing.T) {
			got := FormatInteger(c.number)
			want := c.wanted
			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}
		})
	}
}

func TestFormatDecimal(t *testing.T) {
	var cases = []struct {
		number float64
		wanted string
	}{
		{55.0, "55,00"},
		{55.555, "55,55"},
		{102.55, "102,55"},
		{999.9, "999,90"},
		{1001.25, "1.001,25"},
		{10001, "10.001,00"},
		{1000000, "1.000.000,00"},
		{2500000000, "2.500.000.000,00"},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%f", c.number), func(t *testing.T) {
			got := FormatDecimal(c.number)
			want := c.wanted
			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}
		})
	}
}