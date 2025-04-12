package main

import (
	// = = = = = Native Libraries

	. "local/CORE"
	. "local/DICEY"

	//	"time"

	//"fmt"
	//"strings"
	//"text/template"

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

	"github.com/montanaflynn/stats"
)

var VERSION_NUM = ""

//  *************************** MAIN AREA ****************************

func main() {

	//1. Init the program (and the command line parameters)
	CLI_PARAM_INIT()
	MASTER_INIT("DiceDuel", VERSION_NUM)

	DRED_ab := []float64{5, 6}
	DBLUE_ab := []float64{2, 3}

	/*
	In order of IMPORTANCE:

		// Measure RED A vs RED B   and BLUE A vs BLUE B
		1 == Covariance in statistics measures the degree to which two variables change together. It indicates whether the variables tend to move in the same direction (positive covariance) or opposite directions (negative covariance). In the context of a population, covariance is calculated using the entire set of data points, providing a precise measure of the relationship between the two variables.

		// Measure RED A vs RED B   and BLUE A vs BLUE B
		2 == Entropy in a set of numbers measures the amount of uncertainty or randomness within the data. Higher entropy indicates a more diverse and unpredictable set, while lower entropy suggests a more homogeneous and predictable set. In essence, it quantifies the "surprise" or "information" content of each number within the set, considering the distribution of all possible values.

										Euclidean distance tells you the straight-line distance between two sets of numbers, treated as points in a multidimensional space. It essentially measures how "far apart" those points are. A smaller Euclidean distance indicates that the two points (or sets of numbers) are more similar or closer together, while a larger distance suggests they are more dissimilar or farther apart.
		calculates the straight-line distance between two points in a multi-dimensional space. It's a common and well-understood metric, making it a good starting point for many tasks.


								Manhattan distance is a metric used to determine the distance between two points in a grid-like path. Unlike Euclidean distance, which measures the shortest possible line between two points, Manhattan distance measures the sum of the absolute differences between the coordinates of the points
								also known as L1 distance, measures the distance between two points by summing the absolute differences along each dimension. It's particularly useful when dealing with grid-based data or when certain dimensions have more weight than others.

				MID HING
				In essence, the midhinge provides a measure of the central tendency of a dataset that is resistant to the influence of extreme values, similar to how the median is.

	*/

	red_MID, _ := stats.Midhinge(DRED_ab)
	blu_MID, _ := stats.Midhinge(DBLUE_ab)

	preci := 2
	// manDIST, _ := stats.ManhattanDistance(DRED_ab, DBLUE_ab)
	// round_MAN, _ := stats.Round(manDIST, 4)

	eucDIST, _ := stats.EuclideanDistance(DRED_ab, DBLUE_ab)
	round_EUC, _ := stats.Round(eucDIST, preci)

	copop, _ := stats.CovariancePopulation(DRED_ab, DBLUE_ab)
	round_copop, _ := stats.Round(copop, preci)

	e_red, _ := stats.Entropy(DRED_ab)
	red_ENTROP, _ := stats.Round(e_red, preci)
	e_blue, _ := stats.Entropy(DBLUE_ab)
	blue_ENTROP, _ := stats.Round(e_blue, preci)

	// C.Print("MANHATT_Dist:")
	// Y.Println(round_MAN)

	C.Print("EUCL_Dist:")
	Y.Println(round_EUC)

	C.Print("Co POPUlation:")
	Y.Println(round_copop)

	Y.Print("Entropy RED: ")
	M.Println(red_ENTROP)

	W.Print("Entropy BLU: ")
	C.Println(blue_ENTROP)

	Y.Print("RED_Mid: ")
	M.Println(red_MID)

	W.Print("BLU_Mid: ")
	C.Println(blu_MID)

	//SHOW_Type(rounded)

	DO_EXIT()

	var allowed_DICE = []string{"1", "2", "3", "4", "5", "6"}
	rres := beta_Read_User_Input("RED Dice: ", allowed_DICE, "-min|2", "-max|2", MAGENTA, NOEOL)
	r2 := beta_Read_User_Input("   BLUE Dice: ", allowed_DICE, "-min|2", "-max|2", CYAN)

	Y.Print("\nRed Dice: ")
	M.Println(rres)
	Y.Print("\nBLUE Dice: ")
	C.Println(r2)

	DO_EXIT(SILENT)

	INIT_DiceDuel()

} // end of main
