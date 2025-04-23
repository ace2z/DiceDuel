package main

import (
	// = = = = = Native Libraries
	. "github.com/ace2z/GOGO/REFGADS/BLOX"
	. "github.com/ace2z/GOGO/REFGADS/ColorOPS"
	//"github.com/fatih/color"
	//"sort"
	//"strconv"
	//"strings"
	"fmt"
	"strings"
)

// *************************** MAIN AREA ****************************

const (
	hex_esc    string = "\x1b["
	fg_256mode string = "38;5;"
	bg_256mode string = "48;5;"
	ec_RESET   string = hex_esc + "0m"
	ec_CLOSE   string = "m"

	escBOLD    string = "1;"
	escUNDER   string = "4;"
	escREVERSE string = "7;"
	escBLINK   string = "5;"
)

func get_ATTRIB(att_param string, color_formatted string) string {
	PLACEHOLDER()
	COLOR_PLACEHOLDER()

	// MUST BE A STRING

	// Otherwise lets split the attribs on the __
	parts := strings.Split(att_param, "|")
	if len(parts) > 1 {
		attr_val := parts[1]
		color_formatted += attr_val
	}
	return color_formatted
}

func standard_ColorCode_Adjust(color_code int, need_BRIGHT bool, for_background bool) int {

	if for_background {
		color_code += 40
	} else {
		color_code += 30
	}

	// IF we want BRIGHT colors for 16 color mode
	if need_BRIGHT {
		color_code += 60
	}
	return color_code
}

// Whole new type of Print ln.. leverages the matrix

// redBoldGreenBackground := "\x1b[1;31;42m"
// Returns the Formatted ESC string, the original # for FG and BG
func ESC_Code_Generator(ALL_PARAMS ...interface{}) (string, string) {

	// Attribs
	// var use_BOLD = false
	// var use_REVERSE = false
	// var use_BLINK = false
	// var use_UNDER = false

	var codes []int

	//1. Start with the HEX escape sequence
	var color_formatted = hex_esc

	var att_notes = ""
	var flag_fg_BRIGHT = false
	var flag_BG_BRIGHT = false
	for _, param := range ALL_PARAMS {
		str_val, isSTR := param.(string)
		int_val, isINT := param.(int)
		if isINT {
			codes = append(codes, int_val)
			continue
		}
		if isSTR == false {
			continue
		}

		//If they passed bright
		if is_FuncParam(param, C_FG_BRIGHT) {
			flag_fg_BRIGHT = true
			att_notes += C_FG_BRIGHT + " "
			//			G.Println(" FORE is BRIGHT!!")
			continue
		}
		if is_FuncParam(param, C_BG_BRIGHT) {
			flag_BG_BRIGHT = true
			att_notes += C_BG_BRIGHT + " "
			//			Y.Println(" BACK is BRIGHT!!!!!")
			//DO_EXIT()
			continue
		}

		// Any of the supported attributes
		if is_FuncParam(param, C_BOLD) {
			color_formatted = get_ATTRIB(str_val, color_formatted)
			att_notes += str_val + " "
			continue
		}
		if is_FuncParam(param, C_UNDER) {
			color_formatted = get_ATTRIB(str_val, color_formatted)
			att_notes += str_val + " "
			continue
		}

		if is_FuncParam(param, C_REVERSE) {
			color_formatted = get_ATTRIB(str_val, color_formatted)
			att_notes += str_val + " "
			continue
		}

		if is_FuncParam(param, C_BLINK) {
			color_formatted = get_ATTRIB(str_val, color_formatted)
			att_notes += str_val + " "
			continue
		}

	}
	var fgnum = 7 // always have a defualt.. so the string formats properly
	var bgnum = 0

	if len(codes) > 0 {
		fgnum = codes[0]
	}
	if len(codes) > 1 {
		bgnum = codes[1]
	}

	orig_notes := "FG: " + fmt.Sprintf("%d", fgnum) + " | BG: " + fmt.Sprintf("%d", bgnum) + " | " + att_notes

	// SHOW_STRUCT(codes)
	// C.Println("fgnum: ", fgnum)
	// Y.Println("bgnum: ", bgnum)
	// PressAny()

	// Determine what "mode" we need to display in
	var fg_standard = fgnum < 8
	var BG_standard = bgnum < 8

	//3. Adjust the color codes as needed
	var cc_text = ""

	if fg_standard {
		fgnum = standard_ColorCode_Adjust(fgnum, flag_fg_BRIGHT, false)
	}
	cc_text = fmt.Sprintf("%d", fgnum)
	if fg_standard {
		color_formatted += cc_text
		// If we are in 256 mode, we need to activate the sequence for it
	} else {
		color_formatted += fg_256mode + cc_text
	}
	color_formatted += ";"

	// For BACKGROUND
	var bg_text = ""
	if BG_standard {
		bgnum = standard_ColorCode_Adjust(bgnum, flag_BG_BRIGHT, true)
	}
	bg_text = fmt.Sprintf("%d", bgnum)

	if BG_standard {
		color_formatted += bg_text

	} else {
		color_formatted += bg_256mode + bg_text
	}

	//5. Finally close it off with the m
	color_formatted += ec_CLOSE

	// C.Println("fgnum: ", fgnum)
	// Y.Println("bgnum: ", bgnum)

	// SHOW_STRUCT(color_formatted, GREEN)
	// PressAny()

	return color_formatted, orig_notes

}
