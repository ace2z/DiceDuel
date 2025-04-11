package META

import (
	// "flag"
	. "local/CORE"
	// . "local/_API_MBOUM"
	// . "github.com/ace2z/HP_COMMON"
	. "github.com/ace2z/GOGO/Gadgets"
)

// Put all the meta data generation functions here
type HAND_META_OBJ struct {
	NAME    string
	handler func(HND *HAND_OBJ, HIST *[]HAND_OBJ)
}

var META_LIST = []HAND_META_OBJ{
	{"gen_RBDIFF", gen_RB_DIFF},
}

func MetaData_Generation_ENGINE(HND *HAND_OBJ, HIST *[]HAND_OBJ) {
	PLACEHOLDER()

	//1. Go through all the EVENTs and see if any of the enabled ones are occuring in this Hand/Roll
	for _, ev := range META_LIST {

		// Execute the handler associated with this Meta Data
		ev.handler(HND, HIST)

	} //end of for
}
