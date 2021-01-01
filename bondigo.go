package main
import (
	"fmt"
	"net/http"
	"flag"
	"encoding/json"
)

func main(){
	var parada string 
	var bondi string
	flag.StringVar(&parada, "parada", "2664", "Numero de parada a consultar.")
	flag.StringVar(&bondi, "bondi", "0", "Numero de bondi para filtrar.")
	flag.Parse()


	resp, err := http.Get("http://api.montevideo.gub.uy/transporteRest/siguientesParada/" + parada)
	if err != nil {
        	print(err)
    	}
	defer resp.Body.Close()
	var data map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&data)

	if err != nil {
        	print(err)
    	}

	for _, lineas := range data{
		fmt.Println(lineas)
	}
}
