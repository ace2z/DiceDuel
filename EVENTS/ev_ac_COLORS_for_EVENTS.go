package EVENTS

import (
	// "flag"
	"fmt"
	//. "local/CORE"
	// . "github.com/ace2z/HP_COMMON"
	//"encoding/json"
	//. "github.com/ace2z/GOGO/Gadgets"
	//"github.com/fatih/color"
	//"github.com/charmbracelet/lipgloss"
)

func init() {
	fmt.Println("")

	// Populate the events list with the LipGloss colors
	populate_EVENTS_LIST()

	// LG_Print("Terry LG Test", LG_WHITE_ORANGE)
	// Y.Println(" Thats all folks!")
	// DO_EXIT()
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

// // Old Colors defs...when using Color Fatih
// var red_only = BOLD_MAGENTA
// var blue_only = BOLD_CYAN

// var red_by_ONE_color = color.New(color.FgHiYellow, color.BgMagenta, color.Bold, color.Underline, color.BlinkSlow)
// var blu_by_ONE_color = color.New(color.FgHiBlue, color.BgWhite, color.Bold, color.Underline, color.BlinkSlow)

// var just_BLUE_1 = color.New(color.FgBlue, color.BgCyan)
// var just_RED_1 = color.New(color.FgYellow, color.BgMagenta, color.Bold) //color.New(color.FgHiRed, color.BgYellow)

// var red_hinum = color.New(color.FgYellow, color.BgMagenta, color.Bold)
// var blue_hinum = color.New(color.FgBlue, color.BgCyan)

// var red_full = color.New(color.FgYellow, color.BgMagenta, color.Bold, color.BlinkSlow)
// var blue_full = color.New(color.FgBlue, color.BgCyan, color.BlinkSlow)

// var redside_col = color.New(color.FgBlue, color.BgHiMagenta)
// var blueside_col = color.New(color.FgYellow, color.BgHiBlue) //color.New(color.FgHiRed, color.BgYellow)
