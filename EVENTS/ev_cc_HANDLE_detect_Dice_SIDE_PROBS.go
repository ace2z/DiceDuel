package EVENTS

import (
	// "flag"
	. "local/CORE"
	// . "github.com/ace2z/HP_COMMON"
	. "github.com/ace2z/GOGO/Gadgets"
)

// This maps properties of the SIDES of a dice piece
type DSIDE_OBJ struct {
	SIDE int // This is the number of the side of the dice when facing UP

	// If the dice is turned by ONE space, here are all the possible NEW #s
	// that could be NEW side facing UP
	BY_1 []int

	// For when the dice is flipped completely over (360 degrees)
	FULL_FLIP int
}

var SIDE_MATRIX = []DSIDE_OBJ{
	// This is the number of the side of the dice when facing UP
	DSIDE_OBJ{1, []int{2, 3, 4, 5}, 6},
	DSIDE_OBJ{2, []int{1, 3, 4, 6}, 5},
	DSIDE_OBJ{3, []int{1, 2, 5, 6}, 4},
	DSIDE_OBJ{4, []int{1, 2, 5, 6}, 3},
	DSIDE_OBJ{5, []int{1, 3, 4, 6}, 2},
	DSIDE_OBJ{6, []int{2, 3, 4, 5}, 1},
}

// func init() {
// 	// This is the number of the side of the dice when facing UP
// 	var d1 =
// 	var d2 =

// 	SIDE_MATRIX = append(SIDE_MATRIX[:], d1)
// 	SIDE_MATRIX = append(SIDE_MATRIX[:], d2)
// 	SIDE_MATRIX = append(SIDE_MATRIX[:], d3)
// 	SIDE_MATRIX = append(SIDE_MATRIX[:], d4)
// 	SIDE_MATRIX = append(SIDE_MATRIX[:], d5)
// 	SIDE_MATRIX = append(SIDE_MATRIX[:], d6)
// }

func check_flip_action(val int, preval int, mode string) bool {

	for _, x := range SIDE_MATRIX {
		if val != x.SIDE {
			continue
		}

		if mode == "by1" {
			for _, by1 := range x.BY_1 {
				if preval == by1 {
					return true
				}
			} //inner for
			continue
		}

		// and check for a full rotation flip
		if mode == "full" {
			if preval == x.FULL_FLIP {
				return true
			}
			continue
		}

	} //end of for
	return false
}

// looks for Red or Blue winning last 3 in row (by way of 4 or higher values)
func Detect_diceSIDE_Probabilities(GM *HAND_OBJ, event EVENT_OBJ, GHIST *[]HAND_OBJ) bool {
	PLACEHOLDER()
	var name = event.NAME
	var for_redby1 = CONTAINS(name, "redB1")
	var for_red_full = CONTAINS(name, "RED_FULL")
	var for_blueby1 = CONTAINS(name, "blueB1")
	var for_blue_full = CONTAINS(name, "BLUE_FULL")

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
