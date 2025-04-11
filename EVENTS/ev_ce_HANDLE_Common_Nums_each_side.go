package EVENTS

import (
	// "flag"
	. "github.com/ace2z/GOGO/Gadgets"
	. "local/CORE"
	"strings"
)

var NUMS = []int{1, 2, 3, 4, 5, 6}

func compare_from_SOURCE(NUM int, FIRST int, opp_a int, opp_b int) bool {
	if FIRST == NUM {
		if opp_a == NUM {
			return true
		}
		if opp_b == NUM {
			return true
		}
	}
	return false
}

func detect_COMMON_NUMS(HND *HAND_OBJ, event EVENT_OBJ, GHIST *[]HAND_OBJ) bool {
	var name = strings.TrimSpace(event.NAME)

	// Make sure name passed matches the event we want
	if CONTAINS(name, COMMON_NUMS) == false {
		return false
	}

	//Otherwise, we are looking for the common numbers on the blue and green sids for this hand

	var found_comm_num = 0
	var found = false
	for _, NUM := range NUMS {

		if compare_from_SOURCE(NUM, HND.RED_B, HND.BLUE_A, HND.BLUE_B) {
			found = true
			found_comm_num = NUM
			break
		}
		if compare_from_SOURCE(NUM, HND.RED_A, 0, HND.BLUE_B) {
			found = true
			found_comm_num = NUM
			break
		}

	} //end of for

	LG_GREEN := LG_NewColor("#afd700")
	// Finally save this event to the events for this hand
	if found {

		var new_name = name + "_" + INT_to_STRING(found_comm_num)
		//M.Println("Saving event: ", new_name)
		save_INLINE_Event(new_name, HND, LG_GREEN)
	}
	// Always return false.. we are saving this event INLINE
	return false
}
