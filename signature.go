// package main

// import (
// 	"fmt"
// 	"github.com/go-gota/gota/dataframe"
// 	"log"
// 	"os"
// )

// func main() {

// 	signaturefile, err := os.Open("assets/CoronaSignature.csv")

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	df := dataframe.ReadCSV(signaturefile)

// 	fil := df.Filter(
// 		dataframe.F{0, "Threshold", "==", 8.5},
// 	)

// 	fil2 := fil.Filter(
// 		dataframe.F{2, "STRAIN", "==", "SARS 2"},
// 		dataframe.F{2, "STRAIN", "==", "SARS"},
// 		dataframe.F{2, "STRAIN", "==", "MERS"},
// 		dataframe.F{2, "STRAIN", "==", "OC43"},
// 		dataframe.F{2, "STRAIN", "==", "HKU1"},
// 		dataframe.F{2, "STRAIN", "==", "229E"},
// 		dataframe.F{2, "STRAIN", "==", "NL63"},
// 	)

// 	corona := fil2.Arrange(
// 		dataframe.Sort("STRAIN"),
// 	)
// 	fmt.Println(corona)

// 	outputFile, err := os.Create("output85.csv")

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = corona.WriteCSV(outputFile)

// 	if err != nil {
// 		log.Fatalln("There was an error writing the dataframe contents to the file")
// 	}

// }