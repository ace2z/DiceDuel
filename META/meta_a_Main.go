package META

import (
	// "flag"
	. "local/CORE"
	// . "local/_API_MBOUM"
	// . "github.com/ace2z/HP_COMMON"
	. "github.com/ace2z/GOGO/Gadgets"
)

// Put all the meta data generation functions here
var META_DATA_LIST = []METAGEN_OBJ{
	{"gen_RBDIFF", gen_RB_DIFF},
}

type METAGEN_OBJ struct {
	NAME    string
	handler func(HND *HAND_OBJ, HIST *[]HAND_OBJ)
}

func MetaData_Generate_ENGINE(HND *HAND_OBJ, HIST *[]HAND_OBJ) {

	//1. Go through all the EVENTs and see if any of the enabled ones are occuring in this Hand/Roll
	for _, ev := range META_DATA_LIST {

		// IF the handler determines the event was FOUND, we know to save this event to this hand
		// NOTE: Some handlers may be saving to the event list themselves
		// In these cases, they should always return FALSE (to avoid duplicates)
		// found := ev.handler(HND, HIST)
		// if found == true {
		// 	HND.EVENTS = append(HND.EVENTS, ev)
		// }

		// Execute the handler associated with this Meta Data
		ev.handler(HND, HIST)

	} //end of for
}
