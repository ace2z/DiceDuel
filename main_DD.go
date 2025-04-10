package main

import (
	// = = = = = Native Libraries

	. "local/CORE"
	. "local/DICEY"

	"time"

	//"reflect"
	//"bytes"
	//"encoding/binary"
	//"encoding/gob"
	//"reflect"
	//"github.com/qdm12/reprint"
	//. "local/INGEST_ENGINE"

	. "github.com/ace2z/GOGO/Gadgets"
	//"github.com/jinzhu/copier"
	//"github.com/rivo/tview" // https://github.com/rivo/tview/tree/master
	// "strings"
	// "time"
	//"github.com/Delta456/box-cli-maker/v2"
	//"github.com/fatih/color"
	//"os"
)

var VERSION_NUM = ""

//  *************************** MAIN AREA ****************************

func SUPER_FREAK_Test(ALL_PARAMS ...interface{}) {

	pm := Ingest_FUNC_PARAMS(ALL_PARAMS...)

	// Will search the MAP and see if we have a match for -sheep
	// if pm.HAVE_Param("-sheep") {
	// 	W.Println(" Sheep are SPECTACULAR!!!!")
	// }

	// if pm.HAVE_Param("--maptest") {
	// 	SHOW_STRUCT(pm.MAP)
	// }
	// if pm.HAVE_Param("--int") {
	// 	W.Println("Int: ", pm.INT)
	// }
	// if pm.HAVE_Param("-float") {
	// 	W.Println("Float: ", pm.FLOAT)
	// }

	// if pm.HAVE_Param("-i64") {
	// 	W.Println("Int64: ", pm.INT64)
	// }

	// if pm.HAVE_Param("-newcolor") {
	// 	pm.COLOR.Println("You asked to print in THIS COLOR")
	// }
	// if pm.IsColor() {
	// 	pm.COLOR.Println("You asked to print in THIS COLOR")
	// }

	// if pm.HAVE_Param("-testint") {
	// 	W.Println("Test INT you provided: ", pm.INT)
	// }

	// // if pm.IsString() {
	// // 	W.Println("String: ", pm.STRING)
	// // }
	// if pm.IsMAP(1) {
	// 	M.Print("MAP is: ")
	// 	SHOW_STRUCT(pm.MAP)

	// 	mval_name := pm.MAP["name"].(string)
	// 	Y.Print("Mval NAME is: ")
	// 	C.Println(mval_name)
	// }

	// if pm.IsARR(1) {
	// 	W.Println("")
	// 	W.Println(" []STRING Arr:")
	// 	Y.Println(pm.ARR)
	// }

	// // If there is a GENERIC type structure we need to get, use this fucntion
	// // If you want have multiples of the same GENERIC, pass an INT (like 2) to get the 2nd one
	// if Param_GENERIC[TOLSTOY](pm, 2) {
	// 	W.Println("")
	// 	W.Println("")
	// 	W.Println("YO, Here is your GENERIC: ")
	// 	SHOW_STRUCT(pm.GENERIC)

	// 	//To access the actual VALUE in a particular generic, use assertion like this:
	// 	Y.Print("ACTUAL - Generic NAME is: ")
	// 	C.Println(pm.GENERIC.(TOLSTOY).unexport_Field)
	// 	W.Println("")
	// }

	// if pm.HAVE_Param("-myfloat") {
	// 	W.Println("NAMED param float provided is: ")
	// 	Y.Println(pm.FLOAT)
	// }

	// if pm.IsFloat(2) {
	// 	W.Print("Float provided is: ")
	// 	Y.Println(pm.FLOAT)
	// }

	if pm.HAVE_Param("-bees") {
		W.Println("Bees are: ")
		Y.Println(pm.PIPE)
		C.Println("2ndVAL in pipe: ", pm.PIPE[2])
	}

	if pm.HAVE_Param("-testint") {
		W.Println("Test INT you provided: ", pm.INT)
	}

	if pm.IsINT() {
		W.Println("\nFound INT: ", pm.INT)
	}

	if pm.IsFloat(2) {
		W.Println("\nFound FLOAT: ", pm.FLOAT)
	}

	if pm.IsString() {
		W.Println("\nSTRING asked for: ", pm.STRING)
	}
	if pm.IsString(2) {
		C.Println("\nSECOND string: ", pm.STRING)
	}

	if Param_GENERIC[TOLSTOY](pm) {
		W.Println("")
		W.Println("")
		W.Println("YO, Here is your GENERIC: ")
		SHOW_STRUCT(pm.GENERIC)

		//To access the actual VALUE in a particular generic, use assertion like this:
		Y.Print("ACTUAL - Generic NAME is: ")
		C.Println(pm.GENERIC.(TOLSTOY).unexport_Field)
		W.Println("")
	}

	if pm.HAVE_Param("-darcy") {
		W.Println("Darcy is: ")
		Y.Println(pm.GENERIC)

	}

} //end of

type META_TMP struct {
	Rname     string
	unexpName string
}
type TOLSTOY struct {
	Name           string
	Age            int
	Sex            string
	unexport_Field string

	hiddenTIME time.Time

	Meta         META_TMP
	lowStuffMeta META_TMP
}

func main() {

	var terryINT = 8675309
	var test64 int64 = 1234567890123457899
	var testSTR = "New Terry Test String"
	var tbool = true

	var asl = TOLSTOY{
		Name:           "Terry",
		Age:            55,
		Sex:            "Non-Binary",
		unexport_Field: "Morgan Freeman",
		Meta: META_TMP{
			Rname:     "Bonnie",
			unexpName: "HellHole",
		},
		lowStuffMeta: META_TMP{
			Rname:     "Frank",
			unexpName: "Akimba",
		},
	}

	visit := COPY_DeepCopy(asl).(TOLSTOY)
	visit.Name = "CharlieCox"
	visit.Age = 30
	visit.Sex = "Male"
	visit.unexport_Field = "La Fay"

	// W.Println(" MAPS ")
	myMap := NEW_Map()
	myMap["name"] = "John Doe"
	myMap["bodycount"] = 30
	myMap["city"] = "New York"

	var MAP2 = COPY_DeepCopy(myMap, SILENT).(MAP_TYPE)
	MAP2["Updated destMAP"] = "SellTheKid_MarkoMark"
	MAP2["name"] = "Bonnie Kriel"
	MAP2["bodycount"] = 145
	// SHOW_STRUCT(myMap)
	// SHOW_STRUCT(MAP2, GREEN)

	// // Slices
	// W.Println("")
	// G.Println(" SLICES")

	var SLICE = []string{"Spiderman", "IronMan", "Captain America"}
	var slicey2 = COPY_DeepCopy(SLICE, SILENT).([]string)
	slicey2[0] = "HULK"
	slicey2[1] = "Thor"
	slicey2[2] = "Black Widow"

	// SHOW_STRUCT(SLICE)
	// W.Println("")
	// SHOW_STRUCT(slicey2, GREEN)

	// Y.Println("--")
	// DO_EXIT("-silent")

	//SUPER_FREAK_Test("-sheep", "-maptest", myMap, "--int", 4, "-float", 5.56, "Your String TERRY", "-i64", test64, "-newcolor", BW, "SECOND String FOUND")
	//SUPER_FREAK_Test(myMap, "--testint", 4, "-flotest", 1.23, MAP2, SLICE, 5.56, asl, "FirstSTRING passed", 8.67, slicey2, "SECOND string data", visit, MAGENTA_WHITE)
	//SUPER_FREAK_Test("Your String TERRY", myMap1, "Second Simple STRING", asl, mapTWO, SEC)

	SUPER_FREAK_Test(myMap, "--testint", 4, "-myfloat", 1.23, 5.56, MAP2, SLICE, "-darcy", asl, 43, visit, "FirstSTRING passed", 8.67, "-bees|3|50", 69, "SECOND StringStuff", slicey2)

	// Error handling
	if terryINT != 0 {
	}
	if test64 != 0 {
	}
	if testSTR != "" {
	}
	if tbool {
	}
	if asl != (TOLSTOY{}) {
	}
	if myMap != nil {
	}
	if len(SLICE) > 0 {
	}

	DO_EXIT()

	//1. Init the program (and the command line parameters)
	CLI_PARAM_INIT()
	MASTER_INIT("DiceDuel", VERSION_NUM)

	//2. Main Dice Engine
	if INPUT_RED_DICE != "" && INPUT_BLUE_DICE != "" {
		Dice_Engine_INIT(INPUT_RED_DICE, INPUT_BLUE_DICE)

		DO_EXIT("-silent")
	}

	// otherwise go into a loop
	for {

		//1. Main Dice Engine
		Dice_Engine_INIT("", "")

		//2. Show the Dice History  (last 10 hands)
		ShowDice_HISTORY(false)
	}

} // end of main
