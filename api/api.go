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
// HostSearch struct will hold the host ID which can be used for additional queries.
type HostSearch struct {
	Resources []string `json:"resources"`
}
type Meta struct {
	QueryTime uint32 `json:"query_time"`
	PoweredBy string `json:"powered_by"`
	TraceID   string `json:"trace_id"`
}

type HostMeta struct {
	DeviceID  string `json:"device_id"`
	CID       string `json:"cid"`
	AgentLoad string `json:"agent_load_flags"`
	AgentTime string `json:"agent_local_time"`
	AgentVER  string `json:"agent_version"`
	BiosDEV   string `json:"bios_manufacturer"`
	BiosVER   string `json:"bios_version"`
	ConfBase  string `json:"config_id_base"`
	ConfBuild string `json:"config_id_build"`
	ConfPlat  string `json:"config_id_platform"`
	ExtIP     string `json:"external_ip"`
	MAC       string `json:"mac_address"`
	HostName  string `json:"hostname"`
	FirstSeen string `json:"first_seen"`
	LastSeen  string `json:"last_seen"`
	IntIP     string `json:"local_ip"`
	Domain    string `json:"machine_domain"`
	OSVersion string `json:"os_version"`
	PlatID    string `json:"platform_id"`
	Platform  string `json:"platform_name"`
}

type HostPolicy struct {
	PolicyType   string `json:"policy_type"`
	PolicyID     string `json:"policy_id"`
	Applied      bool   `json:"applied"`
	SettingsHash string `json:"settings_hash"`
	AssignedDate string `json:"assigned_date"`
	AppliedDate  string `json:"applied_date"`
}

type HostMaker struct {
	ProductType       string `json:"product_type"`
	ProductDes        string `json:"product_type_desc"`
	SiteName          string `json:"site_name"`
	Status            string `json:"status"`
	SystemManf        string `json:"system_manufacturer"`
	SystemProduct     string `json:"system_product_name"`
	ModifiedTimeStamp string `json:"modified_timestamp"`
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

func HostData() *HostSearch {
	params := url.Values{}
	params.Add("filter", "platform_name: 'Mac'")

	req, err := http.NewRequest("GET", BaseURL+FindAID+params.Encode(), nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", " Bearer <token>") 
	//refresh token
	if err != nil {
		f, _ := os.Create(`C:\temp\CrowdstrikeLogs\log.txt`)
		log.Println("Error: ", err)
		log.SetOutput(f)
	}
	fmt.Println(req)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error: ", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	// variable holding the struct to which the json response will be placed in.
	//var ret string
	var ret HostSearch
	if err := json.Unmarshal(respBody, &ret); err != nil {
		fmt.Println(err)
	}

	//fmt.Printf("Host: %s, MAC: %s", ret.HostName, ret.Domain)
	//return string(respBody)
	return &ret
}

