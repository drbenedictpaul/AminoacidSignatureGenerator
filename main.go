package main

import (
	"fmt"
	"encoding/csv"
	"math"
	"strconv"
	"html/template"
	"log"
	"net/http"
	"os"

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

		s := fmt.Sprint(math.Round(z*100)/100)

		// signature := [][]string{{s, records[x][0],records[x][1], records[x][2], ALA+ARG+ASP+ASN+CYS+GLA+GLN+GLY+HIS+ILE+LEU+LYS+MET+PHY+PRO+SER+THR+TRP+TYR+VAL}}

		// //  fmt.Println(records[x][0])
		// filedata := "<table>"
		// for _, Col := range signature {
		// 	filedata = filedata + "<tr>"
		// 	filedata = filedata + "<td>" + Col[0] + "</td>" + "<td>" + Col[1] + "</td>" + "<td>" + Col[2] + "</td>" + "<td>" + Col[3] + "</td>" + "<td>" + Col[4] + "</td>" 
		// 	filedata = filedata + "</tr>"
		// }
	
		// filedata = filedata + "</table>"
		
		// fmt.Fprintf(w, filedata)
		// // fmt.Println(signature)
		
		threshold := s
		//id := records[x][0]
		virustype := records[x][1]
		aasignature := ALA+ARG+ASP+ASN+CYS+GLA+GLN+GLY+HIS+ILE+LEU+LYS+MET+PHY+PRO+SER+THR+TRP+TYR+VAL

		
	// sars2SignCount := make(map[string]int)
	// sarsSignCount := make(map[string]int)
	// mersSignCount := make(map[string]int)
	// oc43SignCount := make(map[string]int)
	// e229SignCount := make(map[string]int)
	// nl63SignCount := make(map[string]int)
	// hku1SignCount := make(map[string]int)

	
		
		
		switch threshold {
		case "8.5":
			// sars2SignCount[" " + threshold + " " + aasignature + " "] = sars2SignCount[" " + threshold + " " + aasignature] + 1
			corona85[" " + virustype + " " + aasignature + " "] = corona85[" " + virustype + " " + aasignature + " "] + 1

		case "8.4":
			// sarsSignCount[" " + threshold + " " + aasignature + " "] = sarsSignCount[" " + threshold + " " + aasignature] + 1
			corona84[" " + virustype + " " + aasignature + " "] = corona84[" " + virustype + " " + aasignature + " "] + 1

		case "8.3":
			// mersSignCount[" " + threshold + " " + aasignature + " "] = mersSignCount[" " + threshold + " " + aasignature] + 1
			corona83[" " + virustype + " " + aasignature + " "] = corona83[" " + virustype + " " + aasignature + " "] + 1
			
		case "8.2":
			// oc43SignCount[" " + threshold + " " + aasignature + " "] = oc43SignCount[" " + threshold + " " + aasignature] + 1
			corona82[" " + virustype + " " + aasignature + " "] = corona82[" " + virustype + " " + aasignature + " "] + 1

		case "8.1":
			// e229SignCount[" " + threshold + " " + aasignature + " "] = e229SignCount[" " + threshold + " " + aasignature] + 1
			corona81[" " + virustype + " " + aasignature + " "] = corona81[" " + virustype + " " + aasignature + " "] + 1

		case "8.0":
			// nl63SignCount[" " + threshold + " " + aasignature + " "] = nl63SignCount[" " + threshold + " " + aasignature] + 1
			corona80[" " + virustype + " " + aasignature + " "] = corona80[" " + virustype + " " + aasignature + " "] + 1

		case "7.9":
			// hku1SignCount[" " + threshold + " " + aasignature + " "] = hku1SignCount[" " + threshold + " " + aasignature] + 1
			corona79[" " + virustype + " " + aasignature + " "] = corona79[" " + virustype + " " + aasignature + " "] + 1

		case "7.8":
			// hku1SignCount[" " + threshold + " " + aasignature + " "] = hku1SignCount[" " + threshold + " " + aasignature] + 1
			corona78[" " + virustype + " " + aasignature + " "] = corona78[" " + virustype + " " + aasignature + " "] + 1

		case "7.7":
			corona77[" " + virustype + " " + aasignature + " "] = corona77[" " + virustype + " " + aasignature + " "] + 1

		case "7.6":
			corona76[" " + virustype + " " + aasignature + " "] = corona76[" " + virustype + " " + aasignature + " "] + 1

		case "7.5":
			corona75[" " + virustype + " " + aasignature + " "] = corona75[" " + virustype + " " + aasignature + " "] + 1

		case "7.4":
			corona74[" " + virustype + " " + aasignature + " "] = corona74[" " + virustype + " " + aasignature + " "] + 1

		case "7.3":
			corona73[" " + virustype + " " + aasignature + " "] = corona73[" " + virustype + " " + aasignature + " "] + 1
			
		}

		
	
		
	}
	
}

	}
	fmt.Fprint(w, corona85)
	fmt.Fprint(w, corona84)
	fmt.Fprint(w, corona83)
	fmt.Fprint(w, corona82)
	fmt.Fprint(w, corona81)
	fmt.Fprint(w, corona80)
	fmt.Fprint(w, corona79)
	fmt.Fprint(w, corona78)
	fmt.Fprint(w, corona77)
	fmt.Fprint(w, corona75)
	fmt.Fprint(w, corona74)
	fmt.Fprint(w, corona73)

	}