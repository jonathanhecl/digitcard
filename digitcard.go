package digitcard

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type DigitCard struct {
	Card map[string]string
}

func (card DigitCard) GetDigit(code string) string {
	if _, ok := card.Card[code]; !ok {
		return "?"
	}

	return card.Card[code]
}

func LoadCardFromString(data string) (DigitCard, error) {
	card := make(map[string]string)

	scanner := bufio.NewScanner(strings.NewReader(data))

	var lineNum int
	var rows []string

	for scanner.Scan() {
		line := scanner.Text()
		if lineNum == 0 {
			rows = strings.Split(line, "|")
		} else {
			cells := strings.Split(line, "|")
			for i := 1; i < len(cells); i++ {
				key := strings.TrimSpace(rows[i]) + strings.TrimSpace(cells[0])
				value := strings.TrimSpace(cells[i])
				card[key] = value
			}
		}
		lineNum++
	}

	return DigitCard{Card: card}, nil
}

func LoadCardFromFile(path string) (DigitCard, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var content string
	// get file content
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	return LoadCardFromString(content)
}
