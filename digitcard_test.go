package digitcard

import (
	"testing"
)

func TestLoadCardFromString(t *testing.T) {
	data := ` |A |B |C |D |E |F |G |H |I |J
1|10|11|12|13|14|15|16|18|19|20
2|30|31|32|33|34|35|36|38|39|40
3|50|51|52|53|54|55|56|58|59|60
4|70|71|72|73|74|75|76|78|79|80
5|90|91|92|93|94|95|96|98|99|00`

	expectedCard := DigitCard{
		Card: map[string]string{
			"A1": "10", "B1": "11", "C1": "12", "D1": "13", "E1": "14", "F1": "15", "G1": "16", "H1": "18", "I1": "19", "J1": "20",
			"A2": "30", "B2": "31", "C2": "32", "D2": "33", "E2": "34", "F2": "35", "G2": "36", "H2": "38", "I2": "39", "J2": "40",
			"A3": "50", "B3": "51", "C3": "52", "D3": "53", "E3": "54", "F3": "55", "G3": "56", "H3": "58", "I3": "59", "J3": "60",
			"A4": "70", "B4": "71", "C4": "72", "D4": "73", "E4": "74", "F4": "75", "G4": "76", "H4": "78", "I4": "79", "J4": "80",
			"A5": "90", "B5": "91", "C5": "92", "D5": "93", "E5": "94", "F5": "95", "G5": "96", "H5": "98", "I5": "99", "J5": "00",
		},
	}

	card, err := LoadCardFromString(data)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(card.Card) != len(expectedCard.Card) {
		t.Errorf("Expected card length %d, got %d", len(expectedCard.Card), len(card.Card))
	}

	for key, value := range expectedCard.Card {
		if card.Card[key] != value {
			t.Errorf("Expected value %s for key %s, got %s", value, key, card.Card[key])
		}
	}
}

func TestLoadCardFromFile(t *testing.T) {
	path := "test_data/digitcard.txt"
	expectedCard := DigitCard{
		Card: map[string]string{
			"A1": "10", "B1": "11", "C1": "12", "D1": "13", "E1": "14", "F1": "15", "G1": "16", "H1": "18", "I1": "19", "J1": "20",
			"A2": "30", "B2": "31", "C2": "32", "D2": "33", "E2": "34", "F2": "35", "G2": "36", "H2": "38", "I2": "39", "J2": "40",
			"A3": "50", "B3": "51", "C3": "52", "D3": "53", "E3": "54", "F3": "55", "G3": "56", "H3": "58", "I3": "59", "J3": "60",
			"A4": "70", "B4": "71", "C4": "72", "D4": "73", "E4": "74", "F4": "75", "G4": "76", "H4": "78", "I4": "79", "J4": "80",
			"A5": "90", "B5": "91", "C5": "92", "D5": "93", "E5": "94", "F5": "95", "G5": "96", "H5": "98", "I5": "99", "J5": "00",
		},
	}

	card, err := LoadCardFromFile(path)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(card.Card) != len(expectedCard.Card) {
		t.Errorf("Expected card length %d, got %d", len(expectedCard.Card), len(card.Card))
	}

	for key, value := range expectedCard.Card {
		if card.Card[key] != value {
			t.Errorf("Expected value %s for key %s, got %s", value, key, card.Card[key])
		}
	}
}

func TestDigitCard_GetDigit(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "GetDigit A4",
			args: args{code: "A4"},
			want: "70",
		},
		{
			name: "GetDigit B2",
			args: args{code: "B2"},
			want: "31",
		},
		{
			name: "GetDigit C5",
			args: args{code: "C5"},
			want: "92",
		},
	}

	card, _ := LoadCardFromFile("test_data/digitcard.txt")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := card.GetDigit(tt.args.code)

			if got != tt.want {
				t.Errorf("GetDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
