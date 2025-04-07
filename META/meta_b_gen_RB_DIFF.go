package META

import (
	// "flag"
	. "local/CORE"
	// . "local/_API_MBOUM"
	// . "github.com/ace2z/HP_COMMON"
	. "github.com/ace2z/GOGO/Gadgets"
)

// Just calculate the Red Blue diff... no events.. this is meta data for the current HAND
func gen_RB_DIFF(HND *HAND_OBJ, HIST *[]HAND_OBJ) {
	PLACEHOLDER()

	// Get the differences
	rbdiff := INT_GetDiff(HND.RED_B, HND.BLUE_B)
	prev_diff := INT_GetDiff(HND.RED_A, HND.BLUE_A)

	HND.META.RB_DIFF = rbdiff
	HND.META.PREV_RB_DIFF = prev_diff

}
