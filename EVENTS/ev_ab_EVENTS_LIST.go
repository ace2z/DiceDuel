package EVENTS

import (
	// "flag"
	. "local/CORE"
	// . "github.com/ace2z/HP_COMMON"
	//"encoding/json"
	//. "github.com/ace2z/GOGO/Gadgets"
	"github.com/fatih/color"
)

var red_by_ONE_color = color.New(color.FgHiYellow, color.BgMagenta, color.Bold, color.Underline, color.BlinkSlow)
var blu_by_ONE_color = color.New(color.FgHiBlue, color.BgWhite, color.Bold, color.Underline, color.BlinkSlow)

var just_BLUE_1 = color.New(color.FgBlue, color.BgCyan)
var just_RED_1 = color.New(color.FgYellow, color.BgMagenta, color.Bold) //color.New(color.FgHiRed, color.BgYellow)

var red_hinum = color.New(color.FgYellow, color.BgMagenta, color.Bold)
var blue_hinum = color.New(color.FgBlue, color.BgCyan)

var redside_col = color.New(color.FgBlue, color.BgHiMagenta)
var blueside_col = color.New(color.FgYellow, color.BgHiBlue) //color.New(color.FgHiRed, color.BgYellow)

var EVENT_LIST = []EVENT_OBJ{
	{"DiceFlip", 1, NULV, "", "", NULV, NULV, "", true, Detect_diceSIDE_Probabilities, nil},

	// {"RED_INC++", 1, NULV, "", "", NULV, NULV, "", detect_INC_DROP, BOLD_MAGENTA},
	// {"BLUE_INC++", 1, NULV, "", "", NULV, NULV, "", detect_INC_DROP, BOLD_CYAN},

	// {" RED_1 ", 1, NULV, "", "", NULV, NULV, "", detect_HAVE_1, just_RED_1},
	// {" BLUE_1 ", 1, NULV, "", "", NULV, NULV, "", detect_HAVE_1, just_BLUE_1},

	// {" redB1 ", 1, NULV, "", "", NULV, NULV, "", detect_diceSIDES, redside_col},
	// {" blueB1 ", 1, NULV, "", "", NULV, NULV, "", detect_diceSIDES, blueside_col},

	// {" RED_FULL ", 1, NULV, "", "", NULV, NULV, "", detect_diceSIDES, redside_col},
	// {" BLUE_FULL ", 1, NULV, "", "", NULV, NULV, "", detect_diceSIDES, blueside_col},

	// {" RED_HINUM ", 1, NULV, "", "", NULV, NULV, "", detect_HI_NUM, red_hinum},
	// {" BLUE_HINUM ", 1, NULV, "", "", NULV, NULV, "", detect_HI_NUM, blue_hinum},
}
