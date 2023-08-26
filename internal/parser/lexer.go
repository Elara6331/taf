package parser

import (
	"github.com/alecthomas/participle/v2/lexer"
)

var lex = lexer.MustSimple([]lexer.SimpleRule{
	{Name: "header", Pattern: `TAF (AMD|COR)? ?`},
	{Name: "Remark", Pattern: `RMK[^\n]*`},
	{Name: "Number", Pattern: `\d+`},
	{Name: "Modifier", Pattern: `[+-]|VC`},
	{Name: "Prob", Pattern: "PROB"},
	{Name: "Slash", Pattern: `/`},
	{Name: "Descriptor", Pattern: `MI|BC|DR|BL|SH|TS|FZ|PR`},
	{Name: "Precip", Pattern: `DZ|RA|SN|SG|IC|PL|GR|GS|UP`},
	{Name: "Obscur", Pattern: "BR|FG|FU|DU|SA|HZ|PY|VA"},
	{Name: "Phenom", Pattern: "PO|SQ|FC|SS|DS"},
	{Name: "Ident", Pattern: `[A-Z]+`},
	{Name: "WS", Pattern: `[ \t\n\r]+`},
})
