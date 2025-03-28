package DLOGIC

import (
	// "flag"
	// . "local/_API_MBOUM"
	// . "github.com/ace2z/HP_COMMON"
	. "github.com/ace2z/GOGO/Gadgets"
)

// This is simple. Just calculate the Red Blue diff... no events.. this is for a property of the game
func detect_RED_Blue_DIFF(GM *GAME_OBJ, event EVENT_OBJ) bool {
	PLACEHOLDER()
	//var name = event.NAME

	diff := INT_GetDiff(GM.RED_B, GM.BLUE_B)
	GM.RB_DIFF = diff

	return false

}
