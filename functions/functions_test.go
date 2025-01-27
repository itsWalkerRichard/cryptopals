package functions_test

import (
	"cryptopals/functions"
	"testing"
)

func TestHexToBase64(t *testing.T) {
	t.Run("Test HexToBase64", func(t *testing.T) {
		got, err := functions.HexToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
		if err != nil {
			t.Errorf("got error %q", err)
		}
		want := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
		if *got != want {
			t.Errorf("got %q, want %q", *got, want)
		}
	})
}

func TestFixedXOR(t *testing.T) {
	t.Run("Test FixedXOR", func(t *testing.T) {
		got, err := functions.FixedXOR("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
		if err != nil {
			t.Errorf("got error %q", err)
			return
		}
		if got == nil {
			t.Errorf("got nil")
			return
		}
		want := "746865206b696420646f6e277420706c6179"
		if *got != want {
			t.Errorf("got %q, want %q", *got, want)
		}
	})
}
