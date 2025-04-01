package DLOGIC

import (
	// "flag"
	// . "local/_API_MBOUM"
	// . "github.com/ace2z/HP_COMMON"
	//"encoding/json"
	. "github.com/ace2z/GOGO/Gadgets"
	"github.com/fatih/color"
)

const NULV = -999
const NULLV = NULV

type EVENT_OBJ struct {
	NAME     string
	PRIORITY int // Used for Prediction/Analysis. This is the PRIORITY of the event. Higher Priority will take precedence

	LAST_SEEN         int    // This is the number of hands/games BACK when we last saw this EXACT event
	LAST_SEEN_WINNER  string // When Last Seen, this is the WINNER of the hand
	LAST_SEEN_NEXTWIN string // When Last Seen, this is the WINNER of the NEXT hand that follows this one

	// Optional. If there is a value associated with this event, we save it here
	// (as well as the PREVIOUS value) for this in relation to the HAND)
	VAL       int
	PREVVAL   int
	SIGNATURE string // If applicable, this is a SIGNATURE of this event. Something that makes it searchable (not always unique)

	// this is the HANDLER function that does the LOGIC for detecting this particular event
	// Name must be lowercase or underscore
	handler func(GM *HAND_OBJ, event EVENT_OBJ, HIST *[]HAND_OBJ) bool
	COLOR   *color.Color // Color style for displaying the event
}

var red_by_ONE_color = color.New(color.FgHiYellow, color.BgMagenta, color.Bold, color.Underline, color.BlinkSlow)
var blu_by_ONE_color = color.New(color.FgHiBlue, color.BgWhite, color.Bold, color.Underline, color.BlinkSlow)

var just_BLUE_1 = color.New(color.FgBlue, color.BgCyan)
var just_RED_1 = color.New(color.FgYellow, color.BgMagenta, color.Bold) //color.New(color.FgHiRed, color.BgYellow)

var red_hinum = color.New(color.FgYellow, color.BgMagenta, color.Bold)
var blue_hinum = color.New(color.FgBlue, color.BgCyan)

var redside_col = color.New(color.FgBlue, color.BgHiMagenta)
var blueside_col = color.New(color.FgYellow, color.BgHiBlue) //color.New(color.FgHiRed, color.BgYellow)

var EVENT_LIST = []EVENT_OBJ{
	{"RB_DIFF", 1, NULV, "", "", NULV, NULV, "", save_RED_Blue_DIFF, nil},

	{"RED_INC++", 1, NULV, "", "", NULV, NULV, "", detect_INC_DROP, BOLD_MAGENTA},
	{"BLUE_INC++", 1, NULV, "", "", NULV, NULV, "", detect_INC_DROP, BOLD_CYAN},

	{" RED_1 ", 1, NULV, "", "", NULV, NULV, "", detect_HAVE_1, just_RED_1},
	{" BLUE_1 ", 1, NULV, "", "", NULV, NULV, "", detect_HAVE_1, just_BLUE_1},

	{" redB1 ", 1, NULV, "", "", NULV, NULV, "", detect_diceSIDES, redside_col},
	{" blueB1 ", 1, NULV, "", "", NULV, NULV, "", detect_diceSIDES, blueside_col},

	{" RED_FULL ", 1, NULV, "", "", NULV, NULV, "", detect_diceSIDES, redside_col},
	{" BLUE_FULL ", 1, NULV, "", "", NULV, NULV, "", detect_diceSIDES, blueside_col},

	{" RED_HINUM ", 1, NULV, "", "", NULV, NULV, "", detect_HI_NUM, red_hinum},
	{" BLUE_HINUM ", 1, NULV, "", "", NULV, NULV, "", detect_HI_NUM, blue_hinum},
}

type HAND_OBJ struct {
	RED_B int // B is the most recent RED dice value
	RED_A int // A is the previous RED dice value

	BLUE_B int // B is the most recent BLUE dice value
	BLUE_A int // A is the previous BLUE dice value

	EVENTS []EVENT_OBJ

	WINNER string // This is the WINNER of this hand

	// This stores the winner of the NEXT hand (when we have one)
	NEXT_WINNER string

	GAME_SESSION string // Unique ID of the ENTIRE game session and all the hands
	HAND_ID      string // ID of THIS PARTICULAR hand (within this GAME SESSION)
}
