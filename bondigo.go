package main
import (
	"fmt"
	"net/http"
	"flag"
	"io/ioutil"
	"encoding/json"
//	"reflect"
)

func main(){
	parada := flag.String("parada", "2664", "Numero de parada a consultar.")
//	bondi := flag.String("bondi", "0", "Numero de bondi para filtrar.")
	flag.Parse()


	resp, err := http.Get("http://api.montevideo.gub.uy/transporteRest/siguientesParada/" +  *parada)
	if err != nil {
        	print(err)
    	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
        	print(err)
    	}
	sb := string(body)
  var results []map[string]interface{}
	json.Unmarshal([]byte(sb),&results)

	for key, result := range results {
		fmt.Println("Reading Value for Key :", key)
		fmt.Println(
			"- Linea: ",result["linea"],
			"- Minutos: ", result["minutos"],
			"- Real: ", result["real"])
	}

		}
