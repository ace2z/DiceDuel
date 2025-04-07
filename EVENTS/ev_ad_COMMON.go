package EVENTS

import (
	// "flag"
	. "local/CORE"
	// . "github.com/ace2z/HP_COMMON"
	//"encoding/json"
	. "github.com/ace2z/GOGO/Gadgets"
	// "github.com/fatih/color"
)

/*
For saving an Event from a HANDLER (rather than the main DLEvents_Engine loop)

	Pass either 2 INTS and/or a STRING (for the SIGNATURE)
*/
func save_INLINE_Event(name string, HND *HAND_OBJ, SHOW_ME bool, ALL_PARAMS ...interface{}) {
	PLACEHOLDER()
	//pm := Ingest_FUNC_PARAMS(ALL_PARAMS...)

	var val = 8675309
	var prevval = NULLV
	var siggy = ""

	// if pm.IsINT(0) {
	// 	val = pm.INT
	// }
	// if pm.IsINT(1) {
	// 	prevval = pm.INT
	// }
	// // Get the first string we find
	// if pm.IsString() {
	// 	siggy = pm.STRING
	// }

	// Save to Events
	HND.EVENTS = append(HND.EVENTS, EVENT_OBJ{name, 1, NULV, "", "", val, prevval, siggy, SHOW_ME, nil, nil})
}
