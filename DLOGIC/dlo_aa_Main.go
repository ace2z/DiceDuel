package DLOGIC

import (
	// "flag"
	// . "local/_API_MBOUM"
	// . "github.com/ace2z/HP_COMMON"
	"github.com/fatih/color"
)

type RULE_OBJ struct {
	NAME    string
	ENABLED bool
	COLOR   *color.Color
	LEVEL   int // This is the LEVEL/PRIORITY of the rule. Higher Level rules take precedence over lower ones

}

var RULES = []RULE_OBJ{
	{"RED_INC", true, color.New(color.FgRed), 1},
	{"RED_DROP", true, color.New(color.FgRed), 1},
	{"HAVE_RED_6", true, color.New(color.FgRed), 1},
	{"HAVE_RED_1", true, color.New(color.FgRed), 1},
	{"RED_by_1", true, color.New(color.FgRed), 1},

	{"BLUE_INC", true, color.New(color.FgRed), 1},
	{"BLUE_DROP", true, color.New(color.FgRed), 1},
	{"BLUE_RED_6", true, color.New(color.FgRed), 1},
	{"BLUE_RED_1", true, color.New(color.FgRed), 1},
	{"BLUE_by_1", true, color.New(color.FgRed), 1},
}

type GAME_OBJ struct {
	RED_B int // B is the most recent RED dice value
	RED_A int // A is the previous RED dice value

	BLUE_B int // B is the most recent BLUE dice value
	BLUE_A int // A is the previous BLUE dice value

	RB_DIFF int // The difference between the most recent RED and BLUE dice values (B)

	ALL_EVEN bool // if both red and blue dice are EVEN
	ALL_ODD  bool // if both red and blue dice are ODD

	WINNER    string
	WIN_COLOR *color.Color
}
