package EVENTS

import (
	// "flag"
	. "local/CORE"
	// . "local/_API_MBOUM"
	// . "github.com/ace2z/HP_COMMON"
	. "github.com/ace2z/GOGO/Gadgets"
)

const HI_NUM = 5 // This is the start of the HIGHEST numbers on the dice

// looks for Red or Blue winning last 3 in row (by way of 4 or higher values)
func detect_HI_NUM(GM *HAND_OBJ, event EVENT_OBJ, GHIST *[]HAND_OBJ) bool {
	PLACEHOLDER()
	var name = event.NAME
	//hlen := len(*GHIST)

	if CONTAINS(name, "RED_HINUM") {
		if GM.RED_B >= HI_NUM {
			return true
		}
	}

	if CONTAINS(name, "BLUE_HINUM") {
		if GM.BLUE_B >= HI_NUM {
			return true
		}
	}

	return false

}
