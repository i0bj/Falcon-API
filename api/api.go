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
const (
	BaseURL   = ""  // Base API URL
	FindAID   = ""  // Get agent ID
	HostInfo  = "" // Get host info
	AuthToken = ""// Request or Revoke Token
)

type Token struct {
	MainToken string `json:"access_token"`
	ExpiresIn int    `json:"expires_in"`
}

type InLicense struct {
	Total int16 `json:"total,omitempty"`
}

type OutLicense struct {
	Pagination InLicense `json:"pagination,omitempty"`
}

type MetaLicense struct {
	Meta OutLicense `json:"meta,omitempty"`
}

type HostDetails struct {
	Resources []HostMeta `json:"resources"`
}

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

// AccessToken func generates to new token. This token expires every 30 min
func AccessToken() string {
	/* Object that formats body with client/secret that will be sent
	   to token endpoint*/
	body := url.Values{}
	body.Set("client_id", os.Getenv("CS_CLIENT_ID"))  // Client ID
	body.Set("client_secret", os.Getenv("CS_SECRET")) // Secret

	/*First request for the required access token. This token is only
	  active for 30 minutes.*/
	req, err := http.NewRequest("POST", BaseURL+AuthToken, strings.NewReader(body.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	if err != nil {
		log.Println("Attn: ", err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error: ", err)
	}
	defer resp.Body.Close()

	// Data contains the body of the response, in this case the auth token
	tok := &Token{}
	err = json.NewDecoder(resp.Body).Decode(&tok)
	if err != nil {
		log.Println(err)
	}

	return tok.MainToken

}



//LicenseTotal will fetch the total number of licenses in use.
func LicenseTotal(q string)  {
	params := url.Values{}
	params.Add("limit", fmt.Sprintf("%s", q)) //5000 should give the total number of hosts
	req, err := http.NewRequest("GET", BaseURL+FindAID+params.Encode(), nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", " Bearer <token>") // find out why it only works when token is hardcoded
	//refresh token
	if err != nil {
		log.Println(err, "Cannot find total licenses used.")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error: ", err)
	}
	defer resp.Body.Close()

	val := &MetaLicense{} //variable that will contain response JSON and place it in MetaLicense struct

	err = json.NewDecoder(resp.Body).Decode(&val)
	if err != nil {
		log.Fatal(err)
	}

	if val.Meta.Pagination.Total >= 1300 {
		fmt.Println("[!] You are over the allocated limit:", val.Meta.Pagination.Total)
		var answer string
		fmt.Println("Would you like to delete duplicate hosts? yes/no")
		fmt.Scanln(&answer)
		if answer == "yes" {
			fmt.Println("fun to delete duplictes by last check in") //TODO function to remove duplicates
		}
	} else if val.Meta.Pagination.Total <= 1300 {
		fmt.Printf("[+] Total: %d", val.Meta.Pagination.Total)
	}

//FindHost will fetch the HID/AID using the provided query.
func FindHost() (*HostSearch, error) {
	var i, p, h string

	fmt.Println("Select from the options below:")
	var choice int
	fmt.Println("1. Platform")
	fmt.Println("2. Hostname")
	fmt.Println("3. IP")
	fmt.Scanln(&choice)
	switch choice {
	case 1:

		fmt.Println("Enter Platform name: ex Windows")
		fmt.Scanln(i)
	case 2:
		var h string
		fmt.Println("Enter hostname: ")
		fmt.Scanln(h)

	case 3:
		var i string
		fmt.Println("Enter IP")
		fmt.Scanln(i)

	}

	params := url.Values{}
	params.Add("filter", fmt.Sprintf("hostname: '%s'", h))
	params.Add("filter", fmt.Sprintf("platform_name: '%s'", p))
	req, err := http.NewRequest("GET", BaseURL+FindAID+params.Encode(), nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", " Bearer <token>") 
	//refresh token

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error: ", err)
	}
	defer resp.Body.Close()

	// variable holding the struct to which the json response will be placed in.
	//var ret string
	var ret HostSearch
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
