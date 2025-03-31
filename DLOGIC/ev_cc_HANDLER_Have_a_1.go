package DLOGIC

import (
	// "flag"
	// . "local/_API_MBOUM"
	// . "github.com/ace2z/HP_COMMON"
	. "github.com/ace2z/GOGO/Gadgets"
)

func detect_HAVE_1(GM *GAME_OBJ, event EVENT_OBJ, GHIST *[]GAME_OBJ) bool {
	PLACEHOLDER()
	var name = event.NAME

	if CONTAINS(name, "RED") {
		if GM.RED_B == 1 {
			return true
		}
	}
	if CONTAINS(name, "BLUE") {
		if GM.BLUE_B == 1 {
			return true
		}
	}

	return false
}
