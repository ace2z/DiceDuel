package DLOGIC

import (
	// "flag"
	// . "local/_API_MBOUM"
	// . "github.com/ace2z/HP_COMMON"
	. "github.com/ace2z/GOGO/Gadgets"
)

// This is simple. Just calculate the Red Blue diff... no events.. this is for a property of the game
func save_RED_Blue_DIFF(GM *HAND_OBJ, event EVENT_OBJ, GHIST *[]HAND_OBJ) bool {
	PLACEHOLDER()
	//var name = event.NAME

	GM.RB_DIFF = INT_GetDiff(GM.RED_B, GM.BLUE_B)
	GM.PREV_RBDIFF = INT_GetDiff(GM.RED_A, GM.BLUE_A)

	return false

}
