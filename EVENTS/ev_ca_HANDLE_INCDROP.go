package EVENTS

import (
	// "flag"
	. "github.com/ace2z/GOGO/Gadgets"
	. "local/CORE"
	"strings"
)

func detect_INC_DROP(GM *HAND_OBJ, event EVENT_OBJ, GHIST *[]HAND_OBJ) bool {
	var name = strings.TrimSpace(event.NAME)

	red_nozero := GM.RED_B > 0 && GM.RED_A > 0
	blue_nozero := GM.BLUE_B > 0 && GM.BLUE_A > 0

	if CONTAINS(name, RED_INC) {
		if red_nozero && GM.RED_B > GM.RED_A {
			return true
		}
	}
	if CONTAINS(name, BLUE_INC) {
		if blue_nozero && GM.BLUE_B > GM.BLUE_A {
			return true
		}
	}

	return false
}
