package DLOGIC

import (
	// "flag"
	// . "local/_API_MBOUM"
	// . "github.com/ace2z/HP_COMMON"
	//"encoding/json"
	. "github.com/ace2z/GOGO/Gadgets"
	"github.com/fatih/color"
)

type EVENT_OBJ struct {
	NAME    string
	ENABLED bool
	COLOR   *color.Color
	LEVEL   int // This is the LEVEL/PRIORITY of the rule. Higher Level rules take precedence over lower ones

	// this is the HANDLER function that does the LOGIC for detecting this particular event
	// Name must be lowercase or underscore
	handler func(GM *GAME_OBJ, event EVENT_OBJ) bool
}

var red_by_ONE_color = color.New(color.FgHiYellow, color.BgMagenta, color.Bold, color.Underline, color.BlinkSlow)
var blu_by_ONE_color = color.New(color.FgHiBlue, color.BgWhite, color.Bold, color.Underline, color.BlinkSlow)

var just_BLUE_1 = color.New(color.FgBlue, color.BgCyan)
var just_RED_1 = color.New(color.FgHiRed, color.BgYellow)

var EVENT_LIST = []EVENT_OBJ{
	{"RED_INC", true, BOLD_MAGENTA, 1, detect_INC_DROP},
	{"RED_DROP", true, BOLD_MAGENTA, 1, detect_INC_DROP},
	{"BLUE_INC", true, BOLD_CYAN, 1, detect_INC_DROP},
	{"BLUE_DROP", true, BOLD_CYAN, 1, detect_INC_DROP},
	{"RED*by*ONE", true, red_by_ONE_color, 2, detect_by_ONE},
	{"BLUE*by*ONE", true, blu_by_ONE_color, 2, detect_by_ONE},

	{"RED_1", true, just_RED_1, 3, detect_HAVE_1},
	{"BLUE_1", true, just_BLUE_1, 3, detect_HAVE_1},
	{"RED_Blue_DIFF", true, WHITE, 1, detect_RED_Blue_DIFF},
}

type GAME_OBJ struct {
	RED_B int // B is the most recent RED dice value
	RED_A int // A is the previous RED dice value

	BLUE_B int // B is the most recent BLUE dice value
	BLUE_A int // A is the previous BLUE dice value

	RB_DIFF int // The difference between the most recent RED and BLUE dice values (B)

	EVENTS []EVENT_OBJ

	WINNER    string
	WIN_COLOR *color.Color
}

func DL_Events_Engine(GM *GAME_OBJ) {
	PLACEHOLDER()

	//1. Go through all the EVENTs and see if any of the enabled ones are occuring in this game
	for _, ev := range EVENT_LIST {

		if ev.ENABLED == false {
			continue
		}

		// IF the handler determines the event was FOUND, we know to save this event to this game
		found := ev.handler(GM, ev)
		if found == true {
			GM.EVENTS = append(GM.EVENTS, ev)
		}
	}

	//SHOW_STRUCT(GM.EVENTS)

	//return GM
}
