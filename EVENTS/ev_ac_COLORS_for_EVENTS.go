package EVENTS

import (
	// "flag"
	"fmt"
	//. "local/CORE"
	// . "github.com/ace2z/HP_COMMON"
	//"encoding/json"
	. "github.com/ace2z/GOGO/Gadgets"
	"github.com/fatih/color"

	"github.com/charmbracelet/lipgloss"
)

var HEX_WHITE = "#ffffff"
var HEX_BLACK = "#000000"

// Color DEFS using Lipgloss
var LG_WHITE_PURPLE lipgloss.Style
var LG_WHITE_ORANGE lipgloss.Style
var LG_YELLOW_MAGENTA lipgloss.Style
var LG_JUST_RED lipgloss.Style
var LG_JUST_BLUE lipgloss.Style
var LG_BLUE_BLUE lipgloss.Style
var LG_RED_RED lipgloss.Style

var LG_WHITE_RED lipgloss.Style
var LG_LIGHT_on_BLUE lipgloss.Style

var HEX_MAGENTA = "#c9088f"
var HEX_YELLOW = "#f6fa0f"
var HEX_PURPLE = "#7D56F4"
var HEX_ORANGE = "#e38007"
var HEX_BLUE = "#04a6cf"
var HEX_RED = "#f28dad"
var HEX_DARK_RED = "#d70000"
var HEX_DARK_BLUE = "#4818cc"
var HEX_LIGHT_BLUE = "#a6f5e9"

func init() {
	fmt.Println("")

	// Populate teh lipgloss colors...
	LG_WHITE_PURPLE = LG_NewColor(HEX_WHITE, HEX_PURPLE)
	LG_WHITE_ORANGE = LG_NewColor(HEX_WHITE, HEX_ORANGE)
	LG_YELLOW_MAGENTA = LG_NewColor(HEX_YELLOW, HEX_MAGENTA)

	LG_JUST_RED = LG_NewColor(HEX_RED)
	LG_JUST_BLUE = LG_NewColor(HEX_BLUE)

	LG_BLUE_BLUE = LG_NewColor(HEX_LIGHT_BLUE, HEX_DARK_BLUE)
	LG_RED_RED = LG_NewColor("#afff00", "#d70087")

	LG_WHITE_RED = LG_NewColor(HEX_WHITE, "#d70000", "-under")
	LG_LIGHT_on_BLUE = LG_NewColor(HEX_DARK_BLUE, "#5fd7ff", "-under")

	// Populate the events (WITH these colors)
	populate_EVENTS_LIST()

	// DEBUG below

	// fmt.Println(LG_RED_RED.Render(RED_FULL))
	// fmt.Println(LG_BLUE_BLUE.Render(BLUE_FULL))

	// fmt.Println(LG_WHITE_ORANGE.Render(RED_by_ONE))
	// fmt.Println(LG_WHITE_PURPLE.Render(BLUE_by_ONE))

	// fmt.Println(LG_WHITE_RED.Render(RED_1))
	// fmt.Println(LG_LIGHT_on_BLUE.Render(BLUE_1))

	// fmt.Println(LG_JUST_RED.Render(RED_INC))
	// fmt.Println(LG_JUST_BLUE.Render(BLUE_INC))

	// DO_EXIT()

}

/*
Lipgloss is pretty fucking Awesome: https://github.com/charmbracelet/lipgloss
** pass a second HEX_COLOR string to set the BACKGROUND color
** accepts -blink and -under
*/
func LG_NewColor(FORE string, ALL_PARAMS ...interface{}) lipgloss.Style {
	var BACK = HEX_BLACK

	//1. Here is the default color object just with FORE and BACK
	var COLOR_STYLE = lipgloss.NewStyle().Foreground(lipgloss.Color(FORE))

	pm := Ingest_FUNC_PARAMS(ALL_PARAMS...)

	//2. First STRING passed is assumed to be the BACKGROUND color
	if pm.IsString() {

		BACK = pm.STRING
		COLOR_STYLE = COLOR_STYLE.Background(lipgloss.Color(BACK))
	}

	//3. The rest are additional styling options
	if pm.HAVE_Param("-under") {
		COLOR_STYLE = COLOR_STYLE.Underline(true)
	}
	if pm.HAVE_Param("-blink") {
		COLOR_STYLE = COLOR_STYLE.Blink(true)
	}

	// Bold is irrelevent because we are using custom Hex colors
	// if pm.HAVE_Param("-bold") {
	// 	COLOR_STYLE = COLOR_STYLE.Bold(true)
	// }

	/* EXAMPLE:
	var style = lipgloss.NewStyle().
	    Bold(true).
	    Foreground(lipgloss.Color("#FAFAFA")).
	    Background(lipgloss.Color("#7D56F4")).
	    PaddingTop(2).
	    PaddingLeft(4).
	    Width(22)


		or

	LG_WHITE_PURPLE = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(HEX_WHITE)).Background(lipgloss.Color(HEX_PURPLE))

	fmt.Println(style.Render("Hello, kitty"))
	*/

	return COLOR_STYLE
}

// Old Colors defs...when using Color Fatih
var red_only = BOLD_MAGENTA
var blue_only = BOLD_CYAN

var red_by_ONE_color = color.New(color.FgHiYellow, color.BgMagenta, color.Bold, color.Underline, color.BlinkSlow)
var blu_by_ONE_color = color.New(color.FgHiBlue, color.BgWhite, color.Bold, color.Underline, color.BlinkSlow)

var just_BLUE_1 = color.New(color.FgBlue, color.BgCyan)
var just_RED_1 = color.New(color.FgYellow, color.BgMagenta, color.Bold) //color.New(color.FgHiRed, color.BgYellow)

var red_hinum = color.New(color.FgYellow, color.BgMagenta, color.Bold)
var blue_hinum = color.New(color.FgBlue, color.BgCyan)

var red_full = color.New(color.FgYellow, color.BgMagenta, color.Bold, color.BlinkSlow)
var blue_full = color.New(color.FgBlue, color.BgCyan, color.BlinkSlow)

var redside_col = color.New(color.FgBlue, color.BgHiMagenta)
var blueside_col = color.New(color.FgYellow, color.BgHiBlue) //color.New(color.FgHiRed, color.BgYellow)
