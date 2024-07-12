package main

import (
	"fmt"
	"os"
	"testing"
)

func Test_deck_saveToFile(t *testing.T) {
	d := newDeck()
	tests := []struct {
		name     string
		d        deck
		wantErr  bool
		filename string
	}{
		{"testing saving to file", d, false, "test_all_cards"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.saveToFile(tt.filename); (err != nil) != tt.wantErr {
				t.Errorf("deck.saveToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := tt.d.saveToFile(tt.filename); (err != nil) == tt.wantErr {
				os.Remove("test_all_cards")
			}

		},
		)
	}
}

func Test_newDeck2(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"Test deck length", 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newDeck()
			if len(got) != tt.want {
				t.Errorf("newDeck() length = %v, want %v", len(got), tt.want)
			}
		})
	}
}

func Test_newDeck3(t *testing.T) {
	tests := []struct {
		name      string
		wantLen   int
		wantFirst string
		wantLast  string
	}{
		{"Test deck properties", 16, "Ace of Spades", "Four of Diamonds"},
		// TODO: Add more test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newDeck()
			if len(got) != tt.wantLen {
				t.Errorf("newDeck() length = %v, want %v", len(got), tt.wantLen)
			}
			if got[0] != tt.wantFirst {
				t.Errorf("newDeck() first card = %v, want %v", got[0], tt.wantFirst)
			}
			if got[len(got)-1] != tt.wantLast {
				t.Errorf("newDeck() last card = %v, want %v", got[len(got)-1], tt.wantLast)
			}
		})
	}
}
func Test_newDeck4(t *testing.T) {
	tests := []struct {
		name  string
		param string
		want  string
	}{
		{"Test deck length", "length", "16"},
		{"Test first card", "first card", "Ace of Spades"},
		{"Test last card", "last card", "Four of Diamonds"},
		// TODO: Add more test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newDeck()
			switch tt.param {
			case "length":
				if gotVal := fmt.Sprintf("%v", len(got)); gotVal != tt.want {
					t.Errorf("newDeck() %v = %v, want %v", tt.param, gotVal, tt.want)
				}
			case "first card":
				if gotVal := got[0]; gotVal != tt.want {
					t.Errorf("newDeck() %v = %v, want %v", tt.param, gotVal, tt.want)
				}
			case "last card":
				if gotVal := got[len(got)-1]; gotVal != tt.want {
					t.Errorf("newDeck() %v = %v, want %v", tt.param, gotVal, tt.want)
				}
			}
		})
	}
}

func Test_newDeck5(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length 16, but got %v", len(d))
	}
}
