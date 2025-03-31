package DLOGIC

import (
	// "flag"
	// . "local/_API_MBOUM"
	// . "github.com/ace2z/HP_COMMON"
	. "github.com/ace2z/GOGO/Gadgets"
)

// func get_winner(GM GAME_OBJ) string {
// 	// var name = event.NAME
// 	// var save_event = false

// 	if GM.RED_B > GM.BLUE_B {
// 		return "RED"
// 	} else if GM.BLUE_B > GM.RED_B {
// 		return "BLUE"
// 	}

// 	return "TIE"
// }

const HI_NUM = 4

const at_least = 2

func find_event_helper(lookfor string, tmp_events []EVENT_OBJ) bool {
	// See if we have lookfor in the list of events
	for _, event := range tmp_events {
		if CONTAINS(event.NAME, lookfor) {
			return true
		}
	}
	return false

}

// looks for Red or Blue winning last 3 in row (by way of 4 or higher values)
func detect_3x_and_2x(GM *GAME_OBJ, event EVENT_OBJ, GHIST *[]GAME_OBJ) bool {
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
	var hand_c GAME_OBJ

	have_3rd := false
	if hlen > 1 {

		hand_c = (*GHIST)[hlen-2]
		c_winner = hand_c.WINNER
		have_3rd = true
	}

	// Now see if we have the same winner for each of the last 3 hands
	no_empty := a_winner != "" && b_winner != ""
	if have_3rd {
		no_empty = a_winner != "" && b_winner != "" && c_winner != ""
	}

	win_3x := no_empty && a_winner == b_winner && a_winner == c_winner
	win2x := no_empty && a_winner == b_winner

	// Now lets see if each that won was by 4 or more
	red_high_3x := hand_a.RED_B >= HI_NUM && hand_a.RED_A >= HI_NUM && hand_b.RED_B >= HI_NUM
	red_high2x := hand_a.RED_B >= HI_NUM && hand_a.RED_A >= HI_NUM

	blue_high_3x := hand_a.BLUE_B >= HI_NUM && hand_a.BLUE_A >= HI_NUM && hand_b.BLUE_B >= HI_NUM
	blue_high2x := hand_a.BLUE_B >= HI_NUM && hand_a.BLUE_A >= HI_NUM

	/*
		C.Println(" win_3x: ", win_3x)
		C.Println(" win2x: ", win2x)
		C.Println(" No Empty: ", no_empty)
		C.Println(" red_high_3x: ", red_high_3x)
		C.Println(" red_high2x: ", red_high2x)
		C.Println(" blue_high_3x: ", blue_high_3x)
		C.Println(" blue_high2x: ", blue_high2x)
		W.Println("")
	*/

	// Now make decisions
	if CONTAINS(name, "3x_RED") && win_3x && red_high_3x {
		return true
	}
	if CONTAINS(name, "2x_RED") && win2x && red_high2x {

		return true
	}

	if CONTAINS(name, "3x_BLUE") && win_3x && blue_high_3x {
		return true
	}
	if CONTAINS(name, "2x_BLUE") && win2x && blue_high2x {
		return true
	}

	return false

}
