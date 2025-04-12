package EVENTS

import (
	// "flag"
	. "github.com/ace2z/GOGO/Gadgets"
	. "local/CORE"
	"strings"
)

var skip_list = []string{
	CONFLICT_Finder,
	COMMON_NUMS,
}

type matrix_ITEM struct {
	NAME           string
	CONFLICTS_with string
}

var matrix_list = []matrix_ITEM{
	{RED_INC, BLUE_INC},
	{RED_FULL, BLUE_FULL},
	{RED_1, BLUE_1},
	{RED_HINUM, BLUE_HINUM},
}

func in_skipList(name string) bool {
	for _, skip := range skip_list {
		if name == skip {
			return true
		}
	}
	return false
}

func in_MATRIX(lookfor string) (bool, matrix_ITEM) {
	for _, item := range matrix_list {
		if item.NAME == lookfor {
			return true, item
		}
	}
	return false, matrix_ITEM{}
}
func conflict_was_Found(evname, conflict_with string, events []EVENT_OBJ) bool {
	for _, evl := range events {
		tmpname := evl.NAME

		if tmpname == conflict_with {
			return true
		}
	}
	return false
}

var LG_YELLOW = LG_NewColor(HEX_YELLOW)

func detect_CONFLICTS(HND *HAND_OBJ, event EVENT_OBJ, GHIST *[]HAND_OBJ) bool {
	var name = strings.TrimSpace(event.NAME)

	// Make sure name passed matches the event we want
	if CONTAINS(name, CONFLICT_Finder) == false {
		return false
	}

	C.Println(" About to iterate EEEVENTS")

	for _, evl := range HND.EVENTS {
		evname := evl.NAME

		//1. If its in the skip list, we ignore
		if in_skipList(evname) {
			continue
		}

		// If its in the correlation matrix, we need to check it
		exists, mrx := in_MATRIX(evname)
		if exists == false {
			continue
		}

		// We have the EVNAME itself and what it conflicts with
		// Scan the rest of the events and see if we HAVE that conflict
		mrx_conflict := mrx.CONFLICTS_with

		// If found, we save_inline the conflict event
		// We only need to track the FIRST conflict we find
		// So we can break afterwards
		if conflict_was_Found(evname, mrx_conflict, HND.EVENTS) {
			save_INLINE_Event(name, HND, LG_YELLOW)
			break
		}
	} //end of for

	// Always return false.. we are saving this event INLINE
	return false
}
