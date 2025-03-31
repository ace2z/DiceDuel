package DLOGIC

import (
	// "flag"
	// . "local/_API_MBOUM"
	// . "github.com/ace2z/HP_COMMON"
	//"encoding/json"
	. "github.com/ace2z/GOGO/Gadgets"
	CLR "github.com/fatih/color"
)

type EVENT_OBJ struct {
	NAME    string
	ENABLED bool
	COLOR   *CLR.Color
	LEVEL   int // This is the LEVEL/PRIORITY of the rule. Higher Level rules take precedence over lower ones

	// this is the HANDLER function that does the LOGIC for detecting this particular event
	// Name must be lowercase or underscore
	handler func(GM *GAME_OBJ, event EVENT_OBJ, HIST *[]GAME_OBJ) bool
}

var red_by_ONE_color = CLR.New(CLR.FgHiYellow, CLR.BgMagenta, CLR.Bold, CLR.Underline, CLR.BlinkSlow)
var blu_by_ONE_color = CLR.New(CLR.FgHiBlue, CLR.BgWhite, CLR.Bold, CLR.Underline, CLR.BlinkSlow)

var just_BLUE_1 = CLR.New(CLR.FgBlue, CLR.BgCyan)
var just_RED_1 = CLR.New(CLR.FgHiRed, CLR.BgYellow)

var EVENT_LIST = []EVENT_OBJ{
	{"3x_RED++", true, BOLD_MAGENTA, 1, detect_3x_and_2x},
	{"3x_BLUE++", true, BOLD_CYAN, 1, detect_3x_and_2x},

	{"2x_RED++", true, BOLD_MAGENTA, 1, detect_3x_and_2x},
	{"2x_BLUE++", true, BOLD_CYAN, 1, detect_3x_and_2x},

	{"RED_INC++", true, BOLD_MAGENTA, 1, detect_INC_DROP},

	{"BLUE_INC++", true, BOLD_CYAN, 1, detect_INC_DROP},
	{" RED*by*ONE ", true, red_by_ONE_color, 2, detect_by_ONE},
	{" BLUE*by*ONE ", true, blu_by_ONE_color, 2, detect_by_ONE},

	{" RED_1 ", true, just_RED_1, 3, detect_HAVE_1},
	{" BLUE_1 ", true, just_BLUE_1, 3, detect_HAVE_1},
	{"RED_Blue_DIFF", true, WHITE, 1, detect_RED_Blue_DIFF},

	//{"RED_DROP", true, BOLD_MAGENTA, 1, detect_INC_DROP},
	//{"BLUE_DROP", true, BOLD_CYAN, 1, detect_INC_DROP},

}

type GAME_OBJ struct {
	RED_B int // B is the most recent RED dice value
	RED_A int // A is the previous RED dice value

	BLUE_B int // B is the most recent BLUE dice value
	BLUE_A int // A is the previous BLUE dice value

	RB_DIFF int // The difference between the most recent RED and BLUE dice values (B)

	EVENTS []EVENT_OBJ

	WINNER    string // This is the WINNER of this hand
	WIN_COLOR *CLR.Color

	// This stores the winner of the NEXT hand (when we have one)
	FUTURE_WINNER string
}

func DL_Events_Engine(GM *GAME_OBJ, GHIST *[]GAME_OBJ) {
	PLACEHOLDER()

	//1. Determine the winner of the last Hand / Game
	if GM.RED_B > GM.BLUE_B {
		GM.WINNER = " R "
		GM.WIN_COLOR = WHITE_RED
	} else if GM.RED_B < GM.BLUE_B {
		GM.WINNER = " B "
		GM.WIN_COLOR = WHITE_BLUE
	} else {
		GM.WINNER = " T "
		GM.WIN_COLOR = BLUE_GREEN
	}

	//1. Go through all the EVENTs and see if any of the enabled ones are occuring in this Hand/Roll
	for _, ev := range EVENT_LIST {

		if ev.ENABLED == false {
			continue
		}

		// IF the handler determines the event was FOUND, we know to save this event to this game
		found := ev.handler(GM, ev, GHIST)
		if found == true {
			GM.EVENTS = append(GM.EVENTS, ev)
		}
	}

}
