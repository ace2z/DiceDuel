package main

import (
	// = = = = = Native Libraries
	// . "github.com/ace2z/GOGO/REFGADS/BLOX"
	. "github.com/ace2z/GOGO/REFGADS/ColorOPS"
	// // "os"
	// // "path/filepath"
	// //. "github.com/ace2z/GOGO/REFGADS/BLOX"
	// //"github.com/fatih/color"
	// . "github.com/ace2z/GOGO/REFGADS/PRETTY"
	// . "github.com/ace2z/GOGO/REFGADS/ColorOPS"
	// "github.com/fatih/color"
	
	// //"reflect"
	// //"sort"
	// "sort"
	// "slices"
	"strings"
)

// *************************** MAIN AREA ****************************

// These are the colors that are ignored for one reason or another (ie they dont work well together.. or they are redundant)
var IGNORED_LIST = []string{
	"GREY_on_BLACK",
	"RED_on_GREY",
	"LIME_on_GREY",
	"RED_on_BLACK",
	"BLACK_on_GREY",
	"hiWHI_on_GREY",
	"BLACK_on_BRN",
	"BLACK_on_BLUE",
}

func save_To_Cache(COMP color_OBJ, cache *[]color_OBJ) {
	// Check if the item already exists in the cache
	var already_exists = false
	// for _, item := range *cache {
	// 	if item.ID == COMP.ID {
	// 		//			already_exists = true
	// 		//			return
	// 	}
	// }

	// If we dont have a match already, we need to compare the attribs to what was saved earlier
	// Also check to see if the ATTRIBUTES that were passed were already used by something in the cache
	// var hard_exit = false
	// total_found := 0
	for _, cache := range *cache {
		cache_compare := cache.comp_test

		COMPtest := COMP.comp_test
		if COMPtest == cache_compare {
			return
			// already_exists = true
			// return
		}

	} //end of for
	already_exists = false

	if already_exists {
		return
	}

	// otherwise SAVE IT!!
	*cache = append(*cache, COMP)
}

func removeDUPS(slice []string) []string {
	seen := make(map[string]bool)
	result := []string{}

	for _, str := range slice {
		if _, exists := seen[str]; !exists {
			seen[str] = true
			result = append(result, str)
		}
	}
	return result
}

func Generate_Compound_Colors() {
	C.Println("  Generating Compound Colors")

	var cache = []color_OBJ{}

	for n, cm := range COLOR_MATRIX {
		//var fgname = cm.DESC
		var fgID = cm.ID

		// Now do an inner loop to find all the background colors
		for p, bg := range COLOR_MATRIX {
			if p == n {
				continue
			}

			//var bgname = bg.DESC
			var bgID = bg.ID

			// Generate a color ID based on fgname and bgname
			NEW_COLOR_ID := fgID + "_on_" + bgID
			found := false
			// see if this is Ignored
			for _, tmp := range IGNORED_LIST {
				if NEW_COLOR_ID == tmp {
					found = true
					break
				}
			}
			if found {
				continue
			}

			// Now save this new item to the matrix.
			// The DESC will be the same as the COLOR_ID.. also we will set the compund flags
			var all_attrs = []string{}

			var text_attrs = []string{}

			//We want to make sure the attributes are in proper order.. #'s first, then the text attributes
			// We will also remove any duplicates
			for _, attr := range cm.ATTRS {
				if isNumber(attr) == false {
					text_attrs = append(text_attrs, attr)
					continue
				}
				// Otherwise its a # and we add it to the list
				all_attrs = append(all_attrs, attr)
			}
			for _, attr := range bg.ATTRS {
				if isNumber(attr) == false {
					text_attrs = append(text_attrs, attr)
					continue
				}
				// Otherwise its a # and we add it to the list
				all_attrs = append(all_attrs, attr)
			}
			//C.Println(" test_attrs: ", text_attrs)
			text_attrs = removeDUPS(text_attrs)
			//Y.Println(" AFTER: ", text_attrs)
			all_attrs = append(all_attrs, text_attrs...)

			// Create attr_id STRING
			attr_id := strings.Join(all_attrs, "|")
			//G.Println(" ATTR_ID: ", attr_id)
			//W.Println("")

			// Now add the text_attrs to the end of the all_attrs list
			all_attrs = append(all_attrs, text_attrs...)

			// SHOW_STRUCT(all_attrs)
			// W.Println("")

			// create new item
			var COMP = color_OBJ{
				ID:            NEW_COLOR_ID,
				DESC:          NEW_COLOR_ID,
				ATTRS:         all_attrs,
				IS_16color:    true,
				IS_COMP_Color: true,
				comp_test:     attr_id,
			}

			// save to the CACHE (if it doesnt already exist)
			save_To_Cache(COMP, &cache)
		}

	} //end of for

	// Sort the cache

	//8. FINALLY... Now add anythign in CACHE to the COLOR_MATRIX
	COLOR_MATRIX = append(COLOR_MATRIX, cache...)

}
