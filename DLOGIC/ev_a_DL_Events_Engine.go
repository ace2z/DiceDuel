package DLOGIC

import (
	// "flag"
	// . "local/_API_MBOUM"
	// . "github.com/ace2z/HP_COMMON"
	//"encoding/json"
	. "github.com/ace2z/GOGO/Gadgets"
	"github.com/fatih/color"
)


func find_LAST_SEEN(GM *HAND_OBJ, GHIST *[]HAND_OBJ) {
	// This function is called for each event to determine the LAST_SEEN value
	// The event object itself is then updated

	hlen := len(*GHIST)
	start_at := hlen - 1

	// Loop in reverse. Remember the current GM is NOT yet in history
	var cnt = -1
	for i := start_at; i >= 0; i--{
		cnt++ // This is the HOW FAR BACK counter

		hist := (*GHIST)[i]
		tmp_events := hist.EVENTS

		// Lets see if we can get the FUTURE hand
		n := i + 1
		have_next := false
		if n < start_at {
			have_next = true
		}


		for _, tmpev := range tmp_events {
			// Now go through the GM.EVENTS and see if any of them match
			for x, GMEV := range GM.EVENTS {
				if GMEV.NAME != tmpev.NAME {
					continue
				}

				// We found a match. So lets update the LAST_SEEN value
				GMEV.LAST_SEEN = cnt
				GM.EVENTS[x] = GMEV // Update the GM.EVENTS with the new LAST_SEEN value

				break
			} //end of GMEV loop
		} //end of tmp_events loop
	} //end of history loop

			


		ev := GM.EVENTS[i]
		ev.LAST_SEEN = 0 // Reset the last seen value

		for j := 0; j < len(*GHIST); j++ {
			hist := (*GHIST)[j]

			if hist.EVENTS[i].NAME == ev.NAME {
				ev.LAST_SEEN = j
				break
			}
		}
	}
}

func DL_Events_Engine(GM *HAND_OBJ, GHIST *[]HAND_OBJ) {
	PLACEHOLDER()

	//1. Determine the winner of the Hand
	if GM.RED_B > GM.BLUE_B {
		GM.WINNER = " R "

	} else if GM.RED_B < GM.BLUE_B {
		GM.WINNER = " B "

	} else {
		GM.WINNER = " T "
	}

	//1. Go through all the EVENTs and see if any of the enabled ones are occuring in this Hand/Roll
	for _, ev := range EVENT_LIST {

		// IF the handler determines the event was FOUND, we know to save this event to this hand
		// NOTE: Some handlers may be saving to the event list themselves
		// In these cases, they should always return FALSE (to avoid duplicates)
		found := ev.handler(GM, ev, GHIST)
		if found == true {
			GM.EVENTS = append(GM.EVENTS, ev)
		}

		//2. Determine Last Seen for each event.
		// The event object itself is then updated
		find_LAST_SEEN(GM, GHIST)
	}

}
