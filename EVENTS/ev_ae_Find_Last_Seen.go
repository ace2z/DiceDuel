package EVENTS

import (
	// "flag"
	// . "local/_API_MBOUM"
	// . "github.com/ace2z/HP_COMMON"
	//"encoding/json"
	//. "github.com/ace2z/GOGO/Gadgets"

	. "local/CORE"
	// "github.com/fatih/color"
)

func find_LAST_SEEN(HND *HAND_OBJ, GHIST *[]HAND_OBJ) {
	// This function is called for each event to determine the LAST_SEEN value
	// The event object itself is then updated

	hlen := len(*GHIST)
	start_at := hlen - 1

	// Loop in reverse. Remember the current HND is NOT yet in history
	var cnt = 0
	for i := start_at; i >= 0; i-- {
		cnt++ // This is the HOW FAR BACK counter

		hist := (*GHIST)[i]
		hist_events := hist.EVENTS

		// Lets see if we can get the FUTURE hand
		n := i + 1
		var have_next = false
		var next HAND_OBJ
		if n < start_at {
			have_next = true
			next = (*GHIST)[n]
		}

		// iterate through the events in this HAND
		for e, evHND := range HND.EVENTS {
			// now lets go through the tmp_events (of the history hand)
			found := false
			for _, HISTev := range hist_events {
				all_match_found := HISTev.NAME == evHND.NAME && HISTev.VAL == evHND.VAL && HISTev.SIGNATURE == evHND.SIGNATURE
				if all_match_found == false {
					continue
				}

				//1. We jave a match, so lets update the last_SEEN properties
				evHND.LAST_SEEN = cnt
				evHND.LAST_SEEN_WINNER = hist.WINNER
				//and if we have a NEXT hand, we need to update the LAST_SEEN_NEXTWINNER
				if have_next {
					evHND.LAST_SEEN_NEXTWINNER = next.WINNER
				}
				found = true
				break
			} //end of HISTev loop

			if found {
				HND.EVENTS[e] = evHND // Update the HND.EVENTS so the LAST_SEEN data is saved
			}
		} //end of events loop
	} //end of history loop
}
