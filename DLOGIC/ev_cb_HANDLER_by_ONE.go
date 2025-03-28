package DLOGIC

import (
	// "flag"
	// . "local/_API_MBOUM"
	// . "github.com/ace2z/HP_COMMON"
	. "github.com/ace2z/GOGO/Gadgets"
)

func detect_by_ONE(GM *GAME_OBJ, event EVENT_OBJ) bool {
	PLACEHOLDER()
	var name = event.NAME

	if CONTAINS(name, "RED") {
		var diff = INT_GetDiff(GM.RED_B, GM.RED_A)
		if diff == 1 {
			return true
		}
	}
	if CONTAINS(name, "BLUE") {
		var diff = INT_GetDiff(GM.BLUE_B, GM.BLUE_A)
		if diff == 1 {
			return true
		}
	}

	return false
}
