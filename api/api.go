package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Defined variables for communication between api endpoints
var (
	BaseURL   = ""  // Base API URL
	FindAID   = ""  // Get agent ID
	HostInfo  = "" // Get host info
	AuthToken = ""// Request or Revoke Token
	Token     = os.Getenv("TOKEN")
  ManageHost = ""       // Del or restore host
)

type Resources struct {
	Resources []HostMeta `json:"resources"`
}

type HostMeta struct {
	DeviceID  string `json:"device_id"`
	CID       string `json:"cid"`
	AgentVER  string `json:"agent_version"`
	BiosDEV   string `json:"bios_version"`
	ExtIP     string `json:"external_ip"`
	MAC       string `json:"mac_address"`
	HostName  string `json:"hostname"`
	FirstSeen string `json:"first_seen"`
	LastSeen  string `json:"last_seen"`
	IntIP     string `json:"local_ip"`
	Domain    string `json:"machine_domain"`
	OSVersion string `json:"os_version"`
	Platform  string `json:"platform_name"`
}

type HostMaker struct {
	ProductDes        string `json:"product_type_desc"`
	SiteName          string `json:"site_name"`
	Status            string `json:"status"`
	SystemManf        string `json:"system_manufacturer"`
	SystemProduct     string `json:"system_product_name"`
	ModifiedTimeStamp string `json:"modified_timestamp"`
}

type HostPolicy struct {
}

//TODO handle logging for error
// AccessToken func generates to new token. This token expires every 30 min
func AccessToken() string {

	// Creates log file for any errors that occur. Will log to log.txt
	f, err := os.Create(`location for file`)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	// Object that formats body with client/secret that will be sent
 //  to token endpoint
	body := url.Values{}
	body.Set("client_id", os.Getenv("CS_CLIENT_ID"))  // Client ID
	body.Set("client_secret", os.Getenv("CS_SECRET")) // Secret
	if err != nil {
		log.Println("Issue with data: ", err)
	}

	// First request for the required access token. This token is only
	// active for 30 minutes.
	req, err := http.NewRequest("POST", BaseURL+AuthToken, strings.NewReader(body.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	if err != nil {
		log.Println("Attn: ", err)
	}
	client := &http.Client{}

	defer req.Body.Close()

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error: ", err)
	}

	// Data contains the body of the response, in this case the auth token
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Could not read response. Pleace check error log.")
	}

	return string(data)

}

