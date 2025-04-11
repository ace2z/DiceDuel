package EVENTS

import (
	// "flag"
	. "local/CORE"
	// . "github.com/ace2z/HP_COMMON"
	//"encoding/json"
	. "github.com/ace2z/GOGO/Gadgets"
	// "github.com/fatih/color"
	"github.com/charmbracelet/lipgloss"
)

/*
For saving an Event from a HANDLER (rather than the main DLEvents_Engine loop)

	** Pass 2 Ints (first will be val, second will be prevval )
	** Pass any other string and it will be interpreted as the signature
	** Pass true false (to show or HIDE this event.. default is to SHOW so pass false to hide
	** pass an LG COLOR_TYPE and it will be used to color/style the event when they are displayed
*/
func save_INLINE_Event(name string, HND *HAND_OBJ, ALL_PARAMS ...interface{}) {
	PLACEHOLDER()
	pm := Ingest_FUNC_PARAMS(ALL_PARAMS...)

	var COLOR = LG_WHITE
	var SHOW_ME = true
	var PRIORITY = 1
	var val = NULLV
	var prevval = NULLV
	var siggy = ""

	if Param_GENERIC[lipgloss.Style](pm) {
		COLOR = pm.GENERIC.(lipgloss.Style)
	}

	if pm.HAVE_Param("-hidden", "-hide") {
		SHOW_ME = false
	}
	if pm.HAVE_Param("-priority") {
		PRIORITY = pm.INT
	}

	if pm.IsINT() {
		val = pm.INT
	}
	if pm.IsINT(2) {
		prevval = pm.INT
	}

	// Alternate means of hidding this event
	if pm.IsBOOL() {
		SHOW_ME = pm.BOOL
	}

	if pm.IsString() {
		siggy = pm.STRING
	}

	// We also need to generate a COLOR id. We use this to search the matrix and find the color to use when showing this event
	color_id := GEN_UNIQUE_ID(COLOR, "-prefix|LG_COLOR_ID_")

	// Save to the COLOR_MATRIX so we can access this later
	var exists = false
	for _, col := range COLOR_MATRIX {
		if col.ID == color_id {
			exists = true
			break
		}
	}
	if exists == false {
		COLOR_MATRIX = append(COLOR_MATRIX, LG_COL_OBJ{color_id, COLOR})
	}

	var inline_EV = EVENT_OBJ{name, PRIORITY, NULV, "", "", val, prevval, siggy, SHOW_ME, nil, nil, COLOR, color_id}

	// Finally, save to the events
	HND.EVENTS = append(HND.EVENTS, inline_EV)

	// SHOW_STRUCT(HND.EVENTS)
	// SHOW_STRUCT(COLOR_MATRIX)
	// PressAny()

}
