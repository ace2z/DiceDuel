package CORE

import (
	// "flag"
	// . "local/_API_MBOUM"
	// . "github.com/ace2z/HP_COMMON"
	//"encoding/json"
	//. "github.com/ace2z/GOGO/Gadgets"
	"github.com/fatih/color"
)

const NULV = -999
const NULLV = NULV

var HISTORY []HAND_OBJ

type EVT_CM_OBJ struct {
	NAME  string
	COLOR *color.Color // Color style for displaying the event
}

// Alternate wway to get colors for events
var EVT_COLOR_MATRIX = []EVT_CM_OBJ{
	{"DUMMY_COLOR", color.New(color.FgHiYellow, color.BgMagenta, color.Bold, color.Underline, color.BlinkSlow)},
}

type EVENT_OBJ struct {
	NAME     string
	PRIORITY int // Used for Prediction/Analysis. This is the PRIORITY of the event. Higher Priority will take precedence

	LAST_SEEN            int    // This is the number of hands/games BACK when we last saw this EXACT event
	LAST_SEEN_WINNER     string // When Last Seen, this is the WINNER of the hand
	LAST_SEEN_NEXTWINNER string // When Last Seen, this is the WINNER of the NEXT hand that follows this one

	// Optional. If there is a value associated with this event, we save it here
	// (as well as the PREVIOUS value) for this in relation to the HAND)
	VAL       int
	PREVVAL   int
	SIGNATURE string // If applicable, this is a SIGNATURE of this event. Something that makes it searchable (not always unique)

	// this is the HANDLER function that does the LOGIC for detecting this particular event
	// Name must be lowercase or underscore
	SHOW_ME bool // If FALSE, this event is hidden from the display
	Handler func(GM *HAND_OBJ, event EVENT_OBJ, HIST *[]HAND_OBJ) bool
	COLOR   *color.Color // Color style for displaying the event

}

type META_OBJ struct {
	RB_DIFF      int //
	PREV_RB_DIFF int // Previous RB_DIFF value
}

type HAND_OBJ struct {
	RED_B int // B is the most recent RED dice value
	RED_A int // A is the previous RED dice value

	BLUE_B int // B is the most recent BLUE dice value
	BLUE_A int // A is the previous BLUE dice value

	EVENTS []EVENT_OBJ
	META   META_OBJ // This is the meta data for this hand

	WINNER      string // This is the WINNER of this hand
	NEXT_WINNER string // This is the FUTURE/ Following Winner of the NEXT hand (when we have it)

	GAME_SESSION string // Unique ID of the ENTIRE game session and all the hands
	HID          string // ID of THIS PARTICULAR hand (within this GAME SESSION)
}
