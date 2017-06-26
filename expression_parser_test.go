//go:generate ragel -Z scanner.rl
//go:generate goyacc expression_parser.y

package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// var lexerTests = []struct{}{
// 	{"{{var}}", "value"},
// 	{"{{x}}", "1"},
// }

func ScanExpression(data string) ([]yySymType, error) {
	l := newLexer([]byte(data))
	var symbols []yySymType
	var s yySymType
	for {
		t := l.Lex(&s)
		if t == 0 {
			break
		}
		symbols = append(symbols, s)
	}
	return symbols, nil
}

func TestExpressionScanner(t *testing.T) {
	tokens, err := ScanExpression("abc > 123")
	require.NoError(t, err)
	require.Len(t, tokens, 3)
}
