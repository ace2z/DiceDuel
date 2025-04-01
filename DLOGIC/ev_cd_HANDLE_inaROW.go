package DLOGIC

import (
	// "flag"
	// . "local/_API_MBOUM"
	// . "github.com/ace2z/HP_COMMON"
	. "github.com/ace2z/GOGO/Gadgets"
)

// func get_winner(GM HAND_OBJ) string {
// 	// var name = event.NAME
// 	// var save_event = false

// 	if GM.RED_B > GM.BLUE_B {
// 		return "RED"
// 	} else if GM.BLUE_B > GM.RED_B {
// 		return "BLUE"
// 	}

// 	return "TIE"
// }
// const at_least = 2

// func find_event_helper(lookfor string, tmp_events []EVENT_OBJ) bool {
// 	// See if we have lookfor in the list of events
// 	for _, event := range tmp_events {
// 		if CONTAINS(event.NAME, lookfor) {
// 			return true
// 		}
// 	}
// 	return false

// }

// looks for Red or Blue winning last 3 in row (by way of 4 or higher values)
func detect_in_a_ROW(GM *HAND_OBJ, event EVENT_OBJ, GHIST *[]HAND_OBJ) bool {
	PLACEHOLDER()
	var name = event.NAME
	hlen := len(*GHIST)

	if hlen < 1 {

		return false
	}

	// Get the last 3 hands// remember GM isnt in HISTORY yet. So it is the MOST RECENT hand
	hand_a := *GM
	hand_b := (*GHIST)[hlen-1]

	a_winner := hand_a.WINNER
	b_winner := hand_b.WINNER

	// We might have LESS than 3 hands in the history
	var c_winner = ""
	var hand_c HAND_OBJ

	have_3rd := false
	if hlen > 1 {

		hand_c = (*GHIST)[hlen-2]
		c_winner = hand_c.WINNER
		have_3rd = true
	}

	// Now see if we have the same winner for each of the last 3 hands
	no_2empty := a_winner != "" && b_winner != ""
	no3_empty := false
	if have_3rd {
		no3_empty = a_winner != "" && b_winner != "" && c_winner != ""
	}

	is_red := CONTAINS(a_winner, "R")
	is_blue := CONTAINS(a_winner, "B")

	win_3x := no3_empty && a_winner == b_winner && a_winner == c_winner
	win2x := no_2empty && a_winner == b_winner

	if CONTAINS(name, "3x_RED") && is_red && win_3x {
		return true
	}
	if CONTAINS(name, "2x_RED") && is_red && win2x {
		return true
	}
	if CONTAINS(name, "3x_BLUE") && is_blue && win_3x {
		return true
	}
	if CONTAINS(name, "2x_BLUE") && is_blue && win2x {
		return true
	}

	return false

}
