package main

import (
	"fmt"
	"goStuff/CrowdStrikeAPI/api"
	"goStuff/CrowdStrikeAPI/rtr"
	"os"
	"os/signal"
	"syscall"
	"time"
	
	"golang.org/x/crypto/bcrypt"
)

func title() {

       fmt.Println(" ██████╗  ██████╗       ███████╗ █████╗ ██╗      ██████╗ ██████╗ ███╗   ██╗")
       fmt.Println("██╔════╝ ██╔═══██╗      ██╔════╝██╔══██╗██║     ██╔════╝██╔═══██╗████╗  ██║")
       fmt.Println("██║  ███╗██║   ██║█████╗█████╗  ███████║██║     ██║     ██║   ██║██╔██╗ ██║")
       fmt.Println("██║   ██║██║   ██║╚════╝██╔══╝  ██╔══██║██║     ██║     ██║   ██║██║╚██╗██║")
       fmt.Println("╚██████╔╝╚██████╔╝      ██║     ██║  ██║███████╗╚██████╗╚██████╔╝██║ ╚████║")
       fmt.Println("╚═════╝  ╚═════╝       ╚═╝     ╚═╝  ╚═╝╚══════╝ ╚═════╝ ╚═════╝ ╚═╝  ╚═══╝ ")

	
	fmt.Println("         	        ____                      ")
	fmt.Println("                 /   \\                      ")
	fmt.Println("                 /    \\                     ")
	fmt.Println("      ///////////^^^^^^\\\\\\\\\\\\          ")
	fmt.Println("     ///////////        \\\\\\\\\\\\         ")
	fmt.Println("    ////// \\\\\\      //////  \\\\\\        ")
	fmt.Println("   /////       \\      //       \\\\\\       ")
	fmt.Println("  ////          //    \\           \\\\      ")
	fmt.Println(" ///           ///||\\\\             \\\     ")
	fmt.Println("//            | //||\\ |               \\    ")
	fmt.Println("/               ^^^^^^                   \   ")

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Made with <3 by i_obj")
}

type login struct {
	username string
	password string
}


func menu() {
	fmt.Println("1. Total Licenses Used")
	fmt.Println("2. Access Token")
	fmt.Println("3. Get Host Metadata")
	fmt.Println("4. Start Batch Session")
	fmt.Println("5. Run Batch CMD")
	fmt.Println("6. Bulk Delete Hosts")
	fmt.Println("7. Exit")
	// choice variable will hold the selection you make
	// and will be used in the switch statement below
	var choice int
	for ok := true; ok; ok = (choice != 6) {
		n, err := fmt.Scanln(&choice)
		if n > 1 || err != nil {

			continue
		}
		switch choice {
		case 1:
			for i := 0; i < 30; i++ {
				time.Sleep(50 * time.Millisecond)
				rtr.ProgressBar(i+1, 30, "Calculating", 10, "=")
			}

			api.LicenseTotal("5000")
			var answer string
			fmt.Println("\nReturn To Menu? ")
			fmt.Scan(&answer)
			if answer == "yes" {
				menu()
			} else {
				os.Exit(2)
			}
		case 2:
			var token string
			fmt.Println("Do you need a new access token? Enter yes or no")
			fmt.Scanln(&token)
			if token == "yes" {
				api.AccessToken()
				fmt.Println("An Oauth2 token has been created and saved.")
			} else {
				for i := 0; i < 30; i++ {
					time.Sleep(50 * time.Millisecond)
					rtr.ProgressBar(i+1, 30, "Exiting Falcon..", 10, "=")
				}
				os.Exit(0)
			}
		case 3:
			var HIDS string
			fmt.Println("Enter Host: ")
			fmt.Scanln(&HIDS)
			api.FindInfo(api.FindHost(HIDS))
			var ret string
			fmt.Println("\n\n[!] Perform another lookup? yes or no?")
			fmt.Scan(&ret)
			if ret == "yes" {
				var HIDS string
				fmt.Println("Enter Host: ")
				fmt.Scanln(&HIDS)
				api.FindInfo(api.FindHost(HIDS)) //TODO clean up, have the function return to menu possibly as opposed to asking for another hostname.
			} else {
				menu()

			}

		case 4:
			rtr.StartSession()
			var answer string
			fmt.Println("\nReturn To Menu? ")
			fmt.Scan(&answer)
			if answer == "yes" {
				menu()
			} else {
				for i := 0; i < 30; i++ {
					time.Sleep(50 * time.Millisecond)
					rtr.ProgressBar(i+1, 30, "Exiting Falcon..", 10, "=")
				}
				os.Exit(2)
			}
		case 5:
			//batch session ID
			var batchID string
			fmt.Println("Please enter the Batch ID: ")
			fmt.Scanln(&batchID)
			rtr.ScriptRun(batchID)
			var answer string
			fmt.Println("\nReturn To Menu? ")
			fmt.Scan(&answer)
			if answer == "yes" {
				menu()
			} else {
				for i := 0; i < 30; i++ {
					time.Sleep(50 * time.Millisecond)
					rtr.ProgressBar(i+1, 30, "Exiting Falcon..", 10, "=")
				}
				os.Exit(2)
			}
		
		case 6:
			fmt.Println("[!] After Entering Hosts Press CTRL + Z, Then Enter.")
			api.DeleteHosts()
			var answer string
			fmt.Println("Delete Additional Hosts?")
			fmt.Scan(&answer)
			if answer == "yes" {
				api.DeleteHosts()
			} else {
				for i := 0; i < 30; i++ {
					time.Sleep(50 * time.Millisecond)
					rtr.ProgressBar(i+1, 30, "Exiting Falcon..", 10, "=")
				}
				os.Exit(2)
		case 7:
			for i := 0; i < 30; i++ {
				time.Sleep(50 * time.Millisecond)
				rtr.ProgressBar(i+1, 30, "Exiting Falcon..", 10, "=")
			}
			os.Exit(0)
		}
	}

}
func main() {
	
	// Change username and add bcrypt hashed passwords to login struct
	ulogin1 := &login{username: "user", password: "$2y$10$Rfz1xP4NzZukBYGuhZ10meItmPAquov2xbKtsIwXbZLgqtUO/YcRm"} 
	ulogin2 := &login{username: "user", password: "$2y$10$Rfz1xP4NzZukBYGuhZ10meItmPAquov2xbKtsIwXbZLgqtUO/YcRm"}
	ulogin3 := &login{username: "user", password: "$2y$10$Rfz1xP4NzZukBYGuhZ10meItmPAquov2xbKtsIwXbZLgqtUO/YcRm"}
	ulogin4 := &login{username: "user", password: "$2y$10$Rfz1xP4NzZukBYGuhZ10meItmPAquov2xbKtsIwXbZLgqtUO/YcRm"}
	fmt.Println("------------------------------------------------------------------------------------------|")
	fmt.Println("You must have explicit, authorized permission to access this application. \nUnauthorized attempts to access or use this app may result in criminal penalties.")
	fmt.Println("------------------------------------------------------------------------------------------|")

	var usr, pswd string
	fmt.Println("Enter Username: ")
	fmt.Scanln(&usr)
	fmt.Println("Enter Password: ")
	fmt.Scanln(&pswd)
	
        // When using Bcrypt you cannot compare hash values because bcrypt generates a 128 bit salt that is part of the generated hash.
	// You will need to compare the bcrypt hash with the input from the user.
	if ulogin1.username != usr || bcrypt.CompareHashAndPassword([]byte(ulogin1.password), []byte(pswd)) != nil {
		if ulogin2.username != usr || bcrypt.CompareHashAndPassword([]byte(ulogin2.password), []byte(pswd)) != nil {
			if ulogin3.username != usr || bcrypt.CompareHashAndPassword([]byte(ulogin3.password), []byte(pswd)) != nil {
				if ulogin4.username != usr || bcrypt.CompareHashAndPassword([]byte(ulogin4.password), []byte(pswd)) != nil {
					log.Println("[!] Unauthorized.")
					os.Exit(1)

				} else {
					title()
					programExit()
					menu()

				}
			} else {
				title()
				programExit()
				menu()
			}
		} else {

			title()
			programExit()
			menu()

		}
	} else {
		title()
		programExit()
		menu()
	}
}
