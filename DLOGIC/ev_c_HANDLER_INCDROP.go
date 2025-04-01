package DLOGIC

import (
	// "flag"

	. "github.com/ace2z/GOGO/Gadgets"
)

func detect_INC_DROP(GM *HAND_OBJ, event EVENT_OBJ, GHIST *[]HAND_OBJ) bool {
	PLACEHOLDER()

	var name = event.NAME
	//var save_event = false

	red_nozero := GM.RED_B > 0 && GM.RED_A > 0
	blue_nozero := GM.BLUE_B > 0 && GM.BLUE_A > 0

	//C.Println(red_nozero, blue_nozero)

	if CONTAINS(name, "RED_INC") {
		if GM.RED_B > GM.RED_A && red_nozero {
			return true
		}
	}
	if CONTAINS(name, "BLUE_INC") {
		if GM.BLUE_B > GM.BLUE_A && blue_nozero {
			return true
		}
	}

	return false
}
