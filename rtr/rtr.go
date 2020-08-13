package RTR

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var (
	BaseURL = ""
	
	iniTSession = ""    // step 1 start session with host(s)
	cmdSend     = ""   // step 2 endpoint to send commands to
	refresh     = "" 
	
	uploadScript = ""
	runScript    = ""

	
	uploadFile = "" 

	response = ""
)

type BatchSess struct {
	Existing_Batch string   `json:"existing_batch_id"`
	HostIDs        []string `json:"host_ids"`
	QueueOffline   bool     `json:"queue_offline"`
}

type BatchCmd struct {
	BaseCMD       string   `json:"base_command"`
	BatchID       string   `json:"batch_id"`
	CmdString     string   `json:"command_string"`
	OptionalHosts []string `json:"optional_hosts:"`
	PersistAll    string   `json:"persist_all"`
}

func StartSession() {
	URLValue := url.Values{}

	URLValue.Set("timeout", "30")
	URLValue.Set("timeout_duration", "60s")

	group := BatchSess{
		//Existing_Batch: "",
		HostIDs: []string{"", ""},
		//QueueOffline:   true,
	}
	jsonData, err := json.Marshal(group)

	req, err := http.NewRequest("POST", BaseURL+iniTSession, bytes.NewBuffer(jsonData))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", " Bearer <token>")
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Println(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode, resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func ScriptRun() {

	payload := &BatchCmd{
		BaseCMD:   "put",
		BatchID:   "",
		CmdString: "put testfile.txt",
		//OptionalHosts: []string{}
		PersistAll: "true",
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
	}

	req, err := http.NewRequest("POST", BaseURL+cmdSend, bytes.NewBuffer(jsonData))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", " Bearer <token>")
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Println(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
}
