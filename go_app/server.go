package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"strings"
)

type ml struct {
	model *BinaryClassifer
}

func responseJSON(w http.ResponseWriter, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func arrayToString(a []float64, delim string) string {
	return strings.Replace(fmt.Sprint(a), " ", delim, -1)
	//return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func (m *ml) testOk(w http.ResponseWriter, r *http.Request) {
	example := [][]float32{
		[]float32{-0.74739256,  1.7759689 ,  0.9164685 ,  1.69462746,  0.15660069,
			-0.25860515,  0.41270785,  0.33380319,  0.93324519,  0.91686549,
			0.37013758,  1.11828495, -0.04145232,  0.28582875, -1.23362585,
			-0.01187642,  0.61897348,  1.26059207, -1.0499095 , -0.55483756,}, //0.8440786283218703
		[]float32{7.16649064e-01, -4.33761121e-01, -9.97765087e-01,  2.40422338e+00,
			1.63855342e-03, -1.52299574e+00,  1.73964321e-01,  2.00667088e-01,
			-9.70451456e-01, -3.68434446e-01, -6.41791845e-01,  2.85841898e-01,
			1.77142870e+00, -4.91911687e-01, -1.18888478e+00, -3.64192469e-01,
			1.91465730e+00, -1.44569499e+00,  1.39303868e+00, -6.81542112e-01,}, //0.24511496453349393
	}
	cat_example := [][]string{
		[]string{"2"},
		[]string{"1"},
	}
	res, _ := m.model.PredictProba(example, 20, cat_example, 1)
	fmt.Println(res)

	responseJSON(w, []byte(arrayToString(res, ",")))
}
//curl -X GET "http://127.0.0.1:8080/test"

func main() {
	model, _ := LoadBinaryClassifierFromFile("model.catboost")
	m := ml{model: model}
	r := mux.NewRouter()
	r.HandleFunc("/test", m.testOk).Methods("GET")
	http.ListenAndServe(":8080", cors.AllowAll().Handler(r))
	log.Println("Listening on port 8080...")
}
//for local runnning:
//go mod init catboost_serving
//go get -u github.com/gorilla/mux
//go get -u github.com/rs/cors

