package EVENTS

import (
	// "flag"
	// . "local/_API_MBOUM"
	"strings"

	. "github.com/ace2z/GOGO/Gadgets"
	. "local/CORE"
)

func detect_HAVE_1(GM *HAND_OBJ, event EVENT_OBJ, GHIST *[]HAND_OBJ) bool {
	PLACEHOLDER()
	var name = strings.TrimSpace(event.NAME)
	var for_red = CONTAINS(name, RED_1)
	var for_blue = CONTAINS(name, BLUE_1)

	if for_red && GM.RED_B == 1 {
		return true
	}
	if for_blue && GM.BLUE_B == 1 {
		return true
	}

	return false
}
