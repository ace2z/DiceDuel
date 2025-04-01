package DLOGIC

import (
	// "flag"
	// . "local/_API_MBOUM"

	. "github.com/ace2z/GOGO/Gadgets"
	"github.com/fatih/color"
)

// {"RED_Blue_DIFF", 1, 0, detect_RED_Blue_DIFF, nil},

var sigred_col = color.New(color.FgBlue, color.BgHiMagenta)
var sigblue_col = color.New(color.FgYellow, color.BgHiBlue) //color.New(color.FgHiRed, color.BgYellow)

func save_RollSig(GM *HAND_OBJ, event EVENT_OBJ, GHIST *[]HAND_OBJ) bool {
	PLACEHOLDER()
	// var name = event.NAME
	// var for_redby1 = CONTAINS(name, "redB1")
	// var for_red_full = CONTAINS(name, "RED_FULL")

	redname := " " + INT_to_STRING(GM.RED_A) + "_" + INT_to_STRING(GM.RED_B) + " "
	bluename := " " + INT_to_STRING(GM.BLUE_A) + "_" + INT_to_STRING(GM.BLUE_B) + " "

	evred := EVENT_OBJ{redname, 1, -999, "", nil, sigred_col}
	evblue := EVENT_OBJ{bluename, 1, -999, "", nil, sigblue_col}

	// Add the Signatures to the Events
	GM.EVENTS = append(GM.EVENTS, GM.SIG_RED)

	red_by_one := for_redby1 && check_flip_action(GM.RED_B, GM.RED_A, "by1")
	blue_by_one := for_blueby1 && check_flip_action(GM.BLUE_B, GM.BLUE_A, "by1")

	redfull_flip := for_red_full && check_flip_action(GM.RED_B, GM.RED_A, "full")
	bluefull_flip := for_blue_full && check_flip_action(GM.BLUE_B, GM.BLUE_A, "full")

	// If the most recent flipped BY_1
	if red_by_one {
		return true
	}
	if blue_by_one {
		return true
	}

	// If the most recent flipped FULL
	if redfull_flip {
		return true
	}
	if bluefull_flip {
		return true
	}

	return false

}
