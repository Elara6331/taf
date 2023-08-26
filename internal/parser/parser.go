package parser

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

type AST struct {
	Items []*Item `@@*`
}

type Item struct {
	Pos          lexer.Position
	Time         *string       `( @Number "Z" WS`
	Valid        *ValidPair    `  @@`
	Probability  *Probability  `| @@`
	Change       *Change       `| @@`
	WindSpeed    *WindSpeed    `| @@`
	Visibility   *Visibility   `| @@`
	SkyCondition *SkyCondition `| @@`
	Vicinity     *Vicinity     `| @@`
	Weather      *Weather      `| @@`
	Temperature  *Temperature  `| @@`
	Flag         *Flag         `| @@`
	Remark       *string       `| @Remark`
	ID           *string       `| @Ident ) WS?`
}

type ValidPair struct {
	Pos   lexer.Position
	Start string `@Number "/"`
	End   string `@Number`
}

type Probability struct {
	Pos   lexer.Position
	Value string    `Prob @Number`
	Valid ValidPair `(WS @@)?`
}

type WindSpeed struct {
	Pos       lexer.Position
	WindShear string `("WS" @Number "/")?`
	Variable  bool   `@"VRB"?`
	Value     string `@Number`
	Gusts     string `("G" @Number)?`
	Unit      string `@("MPS"|"KMH"|"KT")`
}

type Visibility struct {
	Pos   lexer.Position
	Plus  bool   `@"P"?`
	Value string `@Number @WS? @(Number? "/" Number)?`
	Unit  string `@"SM"?`
}

type SkyCondition struct {
	Pos       lexer.Position
	Type      string `@("FEW"|"SCT"|"BKN"|"OVC"|"VV"|"SKC")`
	Altitude  string `@Number?`
	CloudType string `@("CB"|"TCU")?`
}

type Vicinity struct {
	Pos           lexer.Position
	Descriptor    string `"VC" ( @Descriptor`
	Precipitation string `     | @Precip )`
}

type Weather struct {
	Pos           lexer.Position
	Modifier      string `@Modifier?`
	Descriptor    string `@Descriptor?`
	Precipitation string ` ( @Precip`
	Obscuration   string ` | @Obscur`
	Other         string ` | @Phenom )`
}

type Change struct {
	Pos   lexer.Position
	Type  string     `@("FM"|"BECMG"|"TEMPO")`
	Time  string     `@Number?`
	Valid *ValidPair `WS? @@?`
}

type Temperature struct {
	Pos   lexer.Position
	Type  string `@("TX"|"TN")`
	Value string `@Number "/"`
	Time  string `@Number "Z"`
}

type Flag struct {
	Pos   lexer.Position
	CAVOK bool `@"CAVOK"`
}

var Parser = participle.MustBuild[AST](participle.Lexer(lex))
