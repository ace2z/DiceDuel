package EVENTS

import (
	// "flag"
	. "local/CORE"
	// . "github.com/ace2z/HP_COMMON"
	//"encoding/json"
	. "github.com/ace2z/GOGO/Gadgets"
	// "github.com/fatih/color"
)

func DL_Events_Engine(HND *HAND_OBJ, HIST *[]HAND_OBJ) {
	PLACEHOLDER()

	//1. Determine the winner of the Hand
	if HND.RED_B > HND.BLUE_B {
		HND.WINNER = "RED"

	} else if HND.RED_B < HND.BLUE_B {
		HND.WINNER = "BLUE"

	} else {
		HND.WINNER = "TIE"
	}

	//1. Go through all the EVENTs and see if any of the enabled ones are occuring in this Hand/Roll
	for _, ev := range EVENT_LIST {

		// IF the handler determines the event was FOUND, we know to save this event to this hand
		// NOTE: Some handlers may be saving to the event list themselves
		// In these cases, they should always return FALSE (to avoid duplicates)
		found := ev.Handler(HND, ev, HIST)
		if found == true {
			HND.EVENTS = append(HND.EVENTS, ev)
		}
	}

	//2. Now we need to determine the LAST_SEEN value for each event
	find_LAST_SEEN(HND, HIST)

}
