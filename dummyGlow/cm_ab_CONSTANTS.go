package main

import (
// = = = = = Native Libraries
// . "github.com/ace2z/GOGO/REFGADS/BLOX"
// . "github.com/ace2z/GOGO/REFGADS/ColorOPS"
// "github.com/fatih/color"
// "sort"
// "strconv"
// "strings"
// "fmt"
// "strings"
)

// *************************** MAIN AREA ****************************

// FORE: \x1b[38;5;<color_code>m
// BACK:  \x1b[48;5;<color_code>m
const (
	param_PREF string  = "-" // unique prefix for all parameters passed to functions
	nullv      int     = -69696969696969
	nullfloat  float64 = float64(nullv)

	// ANSI Escape Codes
	// escSEQ   string = "\x1b["
	// escFG    string = escSEQ + "38;5;"
	// escBG    string = escSEQ + "48;5;"
	// escRESET string = escSEQ + "0m"



	escEND string = "m" // add to the end after the fgcolor

	ADD_CR      string = param_PREF + "add_new_line"
	ADD_NEWLINE string = ADD_CR

	// Common Attributes. Be sure to add teh ___ suffix if you want to support any more in the future

	C_BOLD      string = param_PREF + "bold|" + escBOLD
	C_BLINK     string = param_PREF + "blink|" + escBLINK
	C_UNDERLINE string = param_PREF + "under|" + escUNDER
	C_REVERSE   string = param_PREF + "reverse|" + escREVERSE
	C_REV       string = C_REVERSE
	C_UNDER     string = C_UNDERLINE
	C_FG_BRIGHT string = param_PREF + "fgBright" // This is actually a FLAG and only affects 16color mode
	C_BG_BRIGHT string = param_PREF + "BGbright" // This is actually a FLAG and only affects 16color mode

	// Easy access color constants
	color_PREF  string = "**$clr$**"
	C_BLACK     string = color_PREF + "black"
	C_BROWN     string = color_PREF + "brown"
	C_RED       string = color_PREF + "red"
	C_drk_GREEN string = color_PREF + "drkgrn"
	C_drk_BLUE  string = color_PREF + "drkblue"
	C_drk_CYAN  string = color_PREF + "drkcyan"
	C_GREY      string = color_PREF + "grey"
	C_GREEN     string = color_PREF + "green"
	C_BLUE      string = color_PREF + "blue"
	C_PURPLE    string = color_PREF + "purple"

	C_OLIVE    string = color_PREF + "olive"
	C_YELLOW   string = color_PREF + "yellow"
	C_MAGENTA  string = color_PREF + "magenta"
	C_CYAN     string = color_PREF + "cyan"
	C_WHITE    string = color_PREF + "white"
	C_Hi_WHITE string = color_PREF + "hiwhite"
)

var (
	BW_MODE  bool = false // Set to true if you want to disable colors
	NO_COLOR bool = BW_MODE
)
