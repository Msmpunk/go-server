package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"encoding/base64"
	"fmt"
)

type Response struct {
    Content string `json:"content"`
}

func GetDoc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Abre el archivo
	file, err := os.Open("archivos/archivo.txt")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	// Lee los contenidos del archivo
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}

	// Codifica los contenidos del archivo a base64
	encoded := base64.StdEncoding.EncodeToString(contents)

	defer file.Close()

	 // Crea la respuesta JSON
	 response := Response{Content: encoded}

	 // Codifica la respuesta JSON y escribe en la respuesta HTTP
	 jsonResponse, err := json.Marshal(response)
	 if err != nil {
		 fmt.Println("Error al codificar la respuesta JSON:", err)
		 return
	 }
	 w.Header().Set("Content-Type", "application/json")

	     // Decodifica la cadena en Base64
		 decoded, err := base64.StdEncoding.DecodeString(encoded)
		 if err != nil {
			 fmt.Println("Error al decodificar la cadena en Base64:", err)
			 return
		 }
	 
		 // Convierte el resultado a una cadena
		 str := string(decoded)
	 
		 fmt.Println("Cadena decodificada:", str)

	 json.NewEncoder(w).Encode(jsonResponse)
	// json.NewEncoder(w).Encode(results)
}
