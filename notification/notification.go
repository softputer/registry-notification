package notification

import (
	"github.com/docker/distribution/notifications"
	"net/http"
	"encoding/json"
	"fmt"
	"strings"
	"os"
	"github.com/Sirupsen/logrus"
)

func PostPull(url string, repo string, tag string) (err error){
	content := "name=" + repo + ":" + tag
	fmt.Println(content)
	resp, err := http.Post(url + "/dockerImages/pull", "application/x-www-form-urlencoded", strings.NewReader(content))
        if err != nil {
        	fmt.Println("POST request err!")
		return err
        }
	defer resp.Body.Close()
	return nil
}

func EventHandler(w http.ResponseWriter, r *http.Request) {
	all_event := notifications.Envelope{}
	json_decoder := json.NewDecoder(r.Body)
	err := json_decoder.Decode(&all_event)
	if err != nil {
		fmt.Println("Decode Registry Event Err")
	}	
	var url string
	if url = os.Getenv("PULL_URL"); len(url) == 0 {
		logrus.Info("PULL_URL is not provided.")	
		return
	}

	for _, event := range all_event.Events {
		//w.Write([]byte(event.Action))
		if event.Action == "pull" {
			if strings.Contains(event.Target.URL, "manifests") {
				var repo, tag string
				repo = event.Target.Repository
				tag = event.Target.Tag
				fmt.Println("pull", repo, tag)
				err := PostPull(url, repo, tag)
				if err != nil {
					fmt.Println("Post Pull Failed")
				} else {
					fmt.Println("Post Success")
				}
			}
		}
		
	}
	w.Write([]byte(r.URL.Path + "\n"))
}
