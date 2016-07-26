package notification

import (
	"github.com/docker/distribution/notifications"
	"net/http"
	"encoding/json"
	"fmt"
	"strings"
)

func EventHandler(w http.ResponseWriter, r *http.Request) {
	all_event := notifications.Envelope{}
	json_decoder := json.NewDecoder(r.Body)
	err := json_decoder.Decode(&all_event)
	if err != nil {
		fmt.Println("Decode Registry Event Err")
	}	
	for _, event := range all_event.Events {
		//w.Write([]byte(event.Action))
		if event.Action == "pull" {
			if strings.Contains(event.Target.URL, "manifests") {
				fmt.Println(event.Action)
				fmt.Println(event.Target.URL)
			}
		}
		
	}
	w.Write([]byte(r.URL.Path + "\n"))
}
