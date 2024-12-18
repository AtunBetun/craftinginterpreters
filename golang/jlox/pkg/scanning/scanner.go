package scanning

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"

	"atunbetun.jlox.com/pkg/keywords"
)

func LineTokens(line string, lineNumber int, logger *slog.Logger) ([]keywords.Token, error) {
	logger.Debug(fmt.Sprintf("analyzing line: %s", line))
	// var current int = 0
	// var start int = 0
	tokens := make([]keywords.Token, 0, 10)

	runes := []rune(line)
	for i := 0; i < len(runes); i++ {
		// fmt.Printf("Rune %v is '%c'\n", i, runes[i])
		switch runes[i] {
		case '(':
			tokens = append(tokens, keywords.Token{TokenType: "(", Line: lineNumber})
		case ')':
			tokens = append(tokens, keywords.Token{TokenType: ")", Line: lineNumber})
		case '{':
			tokens = append(tokens, keywords.Token{TokenType: "{", Line: lineNumber})
		case '}':
			tokens = append(tokens, keywords.Token{TokenType: "}", Line: lineNumber})
		case ',':
			tokens = append(tokens, keywords.Token{TokenType: ",", Line: lineNumber})
		case '.':
			tokens = append(tokens, keywords.Token{TokenType: ".", Line: lineNumber})
		case '-':
			tokens = append(tokens, keywords.Token{TokenType: "-", Line: lineNumber})
		case '+':
			tokens = append(tokens, keywords.Token{TokenType: "+", Line: lineNumber})
		case ';':
			tokens = append(tokens, keywords.Token{TokenType: ";", Line: lineNumber})
		case '*':
			tokens = append(tokens, keywords.Token{TokenType: "*", Line: lineNumber})
		}

	}

	return tokens, nil

}

// file must be valid when passed here
func FileTokens(file *os.File, logger *slog.Logger) ([]keywords.Token, error) {

	scanner := bufio.NewScanner(file)
	tokens := make([]keywords.Token, 0, 10)

	lineNumber := 1
	for scanner.Scan() {
		logger.Debug(fmt.Sprintf("line: %d", lineNumber))

		lineTokens, err := LineTokens(scanner.Text(), lineNumber, logger)
		if err != nil {
			panic(fmt.Sprintf("line does not contain correct token: %s", scanner.Text()))
		}

		tokens = append(tokens, lineTokens...)

		lineNumber = lineNumber + 1
	}
	logger.Debug("finished scanning file")

	logger.Debug(fmt.Sprintf("%+v", tokens))

	return make([]keywords.Token, 0, 10), nil

}
