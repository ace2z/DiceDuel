package DLOGIC

import (
	// "flag"

	. "github.com/ace2z/GOGO/Gadgets"
)

func detect_INC_DROP(GM *GAME_OBJ, event EVENT_OBJ) bool {
	PLACEHOLDER()

	var name = event.NAME
	//var save_event = false

	have_RED_INC := CONTAINS(name, "RED_INC") && GM.RED_B > GM.RED_A
	have_RED_DROP := CONTAINS(name, "RED_DROP") && GM.RED_B < GM.RED_A
	have_BLUE_INC := CONTAINS(name, "BLUE_INC") && GM.BLUE_B > GM.BLUE_A
	have_BLUE_DROP := CONTAINS(name, "BLUE_DROP") && GM.BLUE_B < GM.BLUE_A

	if have_RED_INC || have_RED_DROP || have_BLUE_INC || have_BLUE_DROP {
		return true
	}

	return false
}
