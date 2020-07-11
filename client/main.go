package main

import (
	"fmt"
	"goStuff/CrowdStrikeAPI/api"
)

func main() {
	
	//1. Total license count
	//2. Search for host
	//3. Contain host
	//4. refresh token

	//fmt.Println(api.AccessToken())
	fmt.Println(api.FindHost("Mac"))

}
