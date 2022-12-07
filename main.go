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

		 signature := [][]string{{s, records[x][0],records[x][1], records[x][2], ALA+ARG+ASP+ASN+CYS+GLA+GLN+GLY+HIS+ILE+LEU+LYS+MET+PHY+PRO+SER+THR+TRP+TYR+VAL}}

		 filedata := "<table>"
		for _, Col := range signature {
			filedata = filedata + "<tr>"
			filedata = filedata + "<td>" + Col[0] + "</td>" + "<td>" + Col[1] + "</td>" + "<td>" + Col[2] + "</td>" + "<td>" + Col[3] + "</td>" + "<td>" + Col[4] + "</td>" 
			filedata = filedata + "</tr>"
		}
	
		filedata = filedata + "</table>"
		
		fmt.Fprintf(w, filedata)
}
	}
}
}