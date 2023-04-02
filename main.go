package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
)

// templateData provides template parameters.
type templateData struct {
	Service  string
	Revision string
}

// Variables used to generate the HTML page.
var (
	data templateData
	tmpl *template.Template
)

func main() {
	// Initialize template parameters.
	service := os.Getenv("K_SERVICE")
	if service == "" {
		service = "???"
	}

	revision := os.Getenv("K_REVISION")
	if revision == "" {
		revision = "???"
	}

	// Prepare template for execution.
	tmpl = template.Must(template.ParseFiles("index.html"))
	data = templateData{
		Service:  service,
		Revision: revision,
	}

	// Define HTTP server.
	http.HandleFunc("/", helloRunHandler)
	// http.HandleFunc("/results.html", results)

	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// PORT environment variable is provided by Cloud Run.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Print("Hello from Cloud Run! The container started successfully and is listening for HTTP requests on $PORT")
	log.Printf("Listening on port %s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}

}

// helloRunHandler responds to requests by rendering an HTML page.
func helloRunHandler(w http.ResponseWriter, r *http.Request) {

	tmpl.ExecuteTemplate(w, "index.html", "Aminoacid Signature Generator")

	// if err := tmpl.Execute(w, data); err != nil {
	// 	msg := http.StatusText(http.StatusInternalServerError)
	// 	log.Printf("template.Execute: %v", err)
	// 	http.Error(w, msg, http.StatusInternalServerError)
	// }
	corona85 := make(map[string]int)
	corona84 := make(map[string]int)
	corona83 := make(map[string]int)
	corona82 := make(map[string]int)
	corona81 := make(map[string]int)
	corona80 := make(map[string]int)
	corona79 := make(map[string]int)
	corona78 := make(map[string]int)
	corona77 := make(map[string]int)
	corona76 := make(map[string]int)
	corona74 := make(map[string]int)
	corona75 := make(map[string]int)
	corona73 := make(map[string]int)

	if r.Method == http.MethodPost {
		file, _, err := r.FormFile("filename")
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
		fp := csv.NewReader(file)
		data, err := fp.ReadAll()
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}

		records := data

		for z := 8.50; z > 7.30; z -= 0.10 {

			for x := 0; x < len(records); x++ {
				Alanine, _ := strconv.ParseFloat(records[x][3], 32)
				Arginine, _ := strconv.ParseFloat(records[x][4], 32)
				AsparticAcid, _ := strconv.ParseFloat(records[x][5], 32)
				Asparagine, _ := strconv.ParseFloat(records[x][6], 32)
				Cysteine, _ := strconv.ParseFloat(records[x][7], 32)
				GlutamicAcid, _ := strconv.ParseFloat(records[x][8], 32)
				Glutamine, _ := strconv.ParseFloat(records[x][9], 32)
				Glycine, _ := strconv.ParseFloat(records[x][10], 32)
				Histidine, _ := strconv.ParseFloat(records[x][11], 32)
				Isoleucine, _ := strconv.ParseFloat(records[x][12], 32)
				Leucine, _ := strconv.ParseFloat(records[x][13], 32)
				Lysine, _ := strconv.ParseFloat(records[x][14], 32)
				Methionine, _ := strconv.ParseFloat(records[x][15], 32)
				Phenylalanine, _ := strconv.ParseFloat(records[x][16], 32)
				Proline, _ := strconv.ParseFloat(records[x][17], 32)
				Serine, _ := strconv.ParseFloat(records[x][18], 32)
				Threonine, _ := strconv.ParseFloat(records[x][19], 32)
				Tryptophan, _ := strconv.ParseFloat(records[x][20], 32)
				Tyrosine, _ := strconv.ParseFloat(records[x][21], 32)
				Valine, _ := strconv.ParseFloat(records[x][22], 32)

				var ALA, ARG, ASP, ASN, CYS, GLA, GLN, GLY, HIS, ILE, LEU, LYS, MET, PHY, PRO, SER, THR, TRP, TYR, VAL string

				if Alanine >= z {
					ALA = "A"
				}
				if Arginine >= z {
					ARG = "R"
				}
				if AsparticAcid >= z {
					ASP = "D"
				}
				if Asparagine >= z {
					ASN = "N"
				}
				if Cysteine >= z {
					CYS = "C"
				}
				if GlutamicAcid >= z {
					GLA = "E"
				}
				if Glutamine >= z {
					GLN = "Q"
				}
				if Glycine >= z {
					GLY = "G"
				}
				if Histidine >= z {
					HIS = "H"
				}
				if Isoleucine >= z {
					ILE = "I"
				}
				if Leucine >= z {
					LEU = "L"
				}
				if Lysine >= z {
					LYS = "K"
				}
				if Methionine >= z {
					MET = "M"
				}
				if Phenylalanine >= z {
					PHY = "F"
				}
				if Proline >= z {
					PRO = "P"
				}
				if Serine >= z {
					SER = "S"
				}
				if Threonine >= z {
					THR = "T"
				}
				if Tryptophan >= z {
					TRP = "W"
				}
				if Tyrosine >= z {
					TYR = "Y"
				}
				if Valine >= z {
					VAL = "V"
				}

				s := fmt.Sprint(math.Round(z*100) / 100)

				threshold := s
				//id := records[x][0]
				virustype := records[x][1]
				aasignature := ALA + ARG + ASP + ASN + CYS + GLA + GLN + GLY + HIS + ILE + LEU + LYS + MET + PHY + PRO + SER + THR + TRP + TYR + VAL

				switch threshold {
				case "8.5":
					corona85[" "+virustype+"-"+aasignature+" "] = corona85[" "+virustype+"-"+aasignature+" "] + 1

				case "8.4":
					corona84[" "+virustype+"-"+aasignature+" "] = corona84[" "+virustype+"-"+aasignature+" "] + 1

				case "8.3":
					corona83[" "+virustype+"-"+aasignature+" "] = corona83[" "+virustype+"-"+aasignature+" "] + 1

				case "8.2":
					corona82[" "+virustype+"-"+aasignature+" "] = corona82[" "+virustype+"-"+aasignature+" "] + 1

				case "8.1":
					corona81[" "+virustype+"-"+aasignature+" "] = corona81[" "+virustype+"-"+aasignature+" "] + 1

				case "8.0":
					corona80[" "+virustype+"-"+aasignature+" "] = corona80[" "+virustype+"-"+aasignature+" "] + 1

				case "7.9":
					corona79[" "+virustype+"-"+aasignature+" "] = corona79[" "+virustype+"-"+aasignature+" "] + 1

				case "7.8":
					corona78[" "+virustype+"-"+aasignature+" "] = corona78[" "+virustype+"-"+aasignature+" "] + 1

				case "7.7":
					corona77[" "+virustype+"-"+aasignature+" "] = corona77[" "+virustype+"-"+aasignature+" "] + 1

				case "7.6":
					corona76[" "+virustype+"-"+aasignature+" "] = corona76[" "+virustype+"-"+aasignature+" "] + 1

				case "7.5":
					corona75[" "+virustype+"-"+aasignature+" "] = corona75[" "+virustype+"-"+aasignature+" "] + 1

				case "7.4":
					corona74[" "+virustype+"-"+aasignature+" "] = corona74[" "+virustype+"-"+aasignature+" "] + 1

				case "7.3":
					corona73[" "+virustype+"-"+aasignature+" "] = corona73[" "+virustype+"-"+aasignature+" "] + 1

				}
			}
		}

		cor85, err := json.Marshal(corona85)

		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			// } else {
			// 	fmt.Fprint(w, string(cor85))
		}

		cor84, err := json.Marshal(corona84)

		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			// } else {
			// 	fmt.Println(string(cor84))
		}

		cor83, err := json.Marshal(corona83)

		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			// } else {
			//    fmt.Println(string(cor83))
		}

		cor82, err := json.Marshal(corona82)

		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			// } else {
			//    fmt.Println(string(cor82))
		}

		cor81, err := json.Marshal(corona81)

		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			//    } else {
			//        fmt.Println(string(cor81))
		}

		cor80, err := json.Marshal(corona80)

		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			//    } else {
			//        fmt.Println(string(cor80))
		}

		cor79, err := json.Marshal(corona79)

		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			//    } else {
			//        fmt.Println(string(cor79))
		}

		cor78, err := json.Marshal(corona78)

		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			//    } else {
			//        fmt.Println(string(cor78))
		}

		cor77, err := json.Marshal(corona77)

		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			//    } else {
			//        fmt.Println(string(cor77))
		}

		cor76, err := json.Marshal(corona76)

		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			//    } else {
			//        fmt.Println(string(cor76))
		}

		cor75, err := json.Marshal(corona75)

		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			//    } else {
			//        fmt.Println(string(cor75))
		}

		cor74, err := json.Marshal(corona74)

		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			//    } else {
			//        fmt.Println(string(cor74))
		}

		cor73, err := json.Marshal(corona73)

		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			//    } else {
			//        fmt.Println(string(cor73))
		}

		// now write to JSON file

		// jsonFile, err := os.Create("./aasignature.json")

		// if err != nil {
		// 		fmt.Println(err)
		// }
		// defer jsonFile.Close()

		// jsonFile.Write(cor73)
		// jsonFile.Write(cor74)
		// jsonFile.Write(cor75)
		// jsonFile.Write(cor76)
		// jsonFile.Write(cor77)
		// jsonFile.Write(cor78)
		// jsonFile.Write(cor79)
		// jsonFile.Write(cor80)
		// jsonFile.Write(cor81)
		// jsonFile.Write(cor82)
		// jsonFile.Write(cor83)
		// jsonFile.Write(cor84)
		// jsonFile.Write(cor85)

		// jsonFile.Close()

		coronavirus := [][]string{{string(cor85), string(cor84), string(cor83), string(cor82), string(cor81), string(cor80), string(cor79), string(cor78), string(cor77), string(cor76), string(cor75), string(cor74), string(cor73)}}

		filedata := "<table>"
		for _, Col := range coronavirus {

			filedata = filedata + "<th>" + "Aminoacid Signatures at Specific Threshold. (Format: StrainID-AASignature: Frequency)"

			filedata = filedata + "<tr>"
			filedata = filedata + "<th>" + "Threshold 8.5" + "</th>"
			filedata = filedata + "</tr>"

			filedata = filedata + "<tr>"
			filedata = filedata + "<td>" + Col[0] + "</td>"
			// + "<td>" + Col[1] + "</td>" + "<td>" + Col[2] + "</td>" + "<td>" + Col[3] + "</td>" + "<td>" + Col[4] + "</td>"
			// filedata = filedata + "<td>" + Col[1] + "</td>"
			// + "<td>" + Col[1] + "</td>" + "<td>" + Col[2] + "</td>" + "<td>" + Col[3] + "</td>" + "<td>" + Col[4] + "</td>"
			filedata = filedata + "</tr>"

			filedata = filedata + "<tr>"
			filedata = filedata + "<th>" + "Threshold 8.4" + "</th>"
			filedata = filedata + "</tr>"

			filedata = filedata + "<tr>"
			filedata = filedata + "<td>" + Col[1] + "</td>"
			filedata = filedata + "</tr>"

			filedata = filedata + "<tr>"
			filedata = filedata + "<th>" + "Threshold 8.3" + "</th>"
			filedata = filedata + "</tr>"

			filedata = filedata + "<tr>"
			filedata = filedata + "<td>" + Col[2] + "</td>"
			filedata = filedata + "</tr>"

			filedata = filedata + "<tr>"
			filedata = filedata + "<th>" + "Threshold 8.2" + "</th>"
			filedata = filedata + "</tr>"

			filedata = filedata + "<tr>"
			filedata = filedata + "<td>" + Col[3] + "</td>"
			filedata = filedata + "</tr>"

			filedata = filedata + "<tr>"
			filedata = filedata + "<th>" + "Threshold 8.1" + "</th>"
			filedata = filedata + "</tr>"

			filedata = filedata + "<tr>"
			filedata = filedata + "<td>" + Col[4] + "</td>"
			filedata = filedata + "</tr>"

			filedata = filedata + "<tr>"
			filedata = filedata + "<th>" + "Threshold 8.0" + "</th>"
			filedata = filedata + "</tr>"

			filedata = filedata + "<tr>"
			filedata = filedata + "<td>" + Col[5] + "</td>"
			filedata = filedata + "</tr>"

			filedata = filedata + "<tr>"
			filedata = filedata + "<th>" + "Threshold 7.9" + "</th>"
			filedata = filedata + "</tr>"

			filedata = filedata + "<tr>"
			filedata = filedata + "<td>" + Col[6] + "</td>"
			filedata = filedata + "</tr>"

			filedata = filedata + "<tr>"
			filedata = filedata + "<th>" + "Threshold 7.8" + "</th>"
			filedata = filedata + "</tr>"

			filedata = filedata + "<tr>"
			filedata = filedata + "<td>" + Col[7] + "</td>"
			filedata = filedata + "</tr>"

			filedata = filedata + "<tr>"
			filedata = filedata + "<th>" + "Threshold 7.7" + "</th>"
			filedata = filedata + "</tr>"

			filedata = filedata + "<tr>"
			filedata = filedata + "<td>" + Col[8] + "</td>"
			filedata = filedata + "</tr>"

			filedata = filedata + "<tr>"
			filedata = filedata + "<th>" + "Threshold 7.6" + "</th>"
			filedata = filedata + "</tr>"

			filedata = filedata + "<tr>"
			filedata = filedata + "<td>" + Col[9] + "</td>"
			filedata = filedata + "</tr>"

			filedata = filedata + "<tr>"
			filedata = filedata + "<th>" + "Threshold 7.5" + "</th>"
			filedata = filedata + "</tr>"

			filedata = filedata + "<tr>"
			filedata = filedata + "<td>" + Col[10] + "</td>"
			filedata = filedata + "</tr>"

			filedata = filedata + "<tr>"
			filedata = filedata + "<th>" + "Threshold 7.4" + "</th>"
			filedata = filedata + "</tr>"

			filedata = filedata + "<tr>"
			filedata = filedata + "<td>" + Col[11] + "</td>"
			filedata = filedata + "</tr>"

			filedata = filedata + "<tr>"
			filedata = filedata + "<th>" + "Threshold 7.3" + "</th>"
			filedata = filedata + "</tr>"

			filedata = filedata + "<tr>"
			filedata = filedata + "<td>" + Col[12] + "</td>"
			filedata = filedata + "</tr>"

		}

		filedata = filedata + "</table>"

		fmt.Fprintf(w, filedata)

	}
}
