package main

import (
	// = = = = = Native Libraries
	//. "github.com/ace2z/GOGO/REFGADS/BLOX"
	//. "github.com/ace2z/GOGO/REFGADS/ColorOPS"
	// // "os"
	// // "path/filepath"
	// //. "github.com/ace2z/GOGO/REFGADS/BLOX"
	// //"github.com/fatih/color"
	// //. "github.com/ace2z/GOGO/REFGADS/PRETTY"
	//. "github.com/ace2z/GOGO/REFGADS/ColorOPS"
	//"github.com/fatih/color"
	// //"reflect"
	// //"sort"

	//"slices"
	"strings"
)

// *************************** MAIN AREA ****************************

var name_OVERRIDES = map[string]string{
	"hi_white":        "hiWHI",
	"brown":           "BRN",
	"hi_green":        "hiGREEN",
	"magenta":         "MAGE",
	"BLUE":            "hiBLUE",
	"BLACK_on_PURPLE": "BLACK_on_MAGENTA",
	"BLACK_on_TEAL":   "BLACK_on_CYAN",
}

func generate_COLOR_ID(input string) string {
	input = strings.ToLower(input)

	//1. See if we have any overrides for the name
	for color, override := range name_OVERRIDES {
		if strings.Contains(input, color) {
			result := strings.ReplaceAll(input, color, override)
			return result
		}
	}

	// If no overrides, the default action is to make the name uppercase
	result := strings.ToUpper(input)
	return result

}

func GEN_Color_IDs() {

	for n, cm := range COLOR_MATRIX {

		// Only generate color ID's for the items that are BLANK
		if cm.ID == "" {
			cm.ID = generate_COLOR_ID(cm.DESC)
			// update COLOR MATRIX
			COLOR_MATRIX[n] = cm
		}
	}
}
