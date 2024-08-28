package mgmt

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gen2brain/malgo"
)

var gCtx *malgo.AllocatedContext

func GetDevices() string {
	devs, err := gCtx.Devices(malgo.Capture)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	var capDevs []CaptureDevice
	for _, dev := range devs {
		var capDev = CaptureDevice{Name: dev.Name(), ID: string(dev.ID.String())}
		capDevs = append(capDevs, capDev)
	}
    b, err := json.Marshal(capDevs)
    if err != nil {
        fmt.Println(err)
        return ""
    }
	return string(b)
}
func RunServer(ctx *malgo.AllocatedContext) {

	gCtx = ctx


	fs := http.FileServer(http.Dir("./html")) 
	http.Handle("/", fs)
	http.HandleFunc("/devices", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			fmt.Fprint(w, GetDevices())
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not allowed")
		}
	})

	fmt.Println("Server listening on port 8093...")
	http.ListenAndServe(":8093", nil)
}