package main

import (
	"fmt"
	"goStuff/CrowdStrikeAPI/api"
	"goStuff/CrowdStrikeAPI/rtr"
	"os"
	"os/signal"
	"syscall"
	"time"
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

func (l *login) LoginUser() {
	fmt.Println("------------------------------------------------------------------------------------------|")
	fmt.Println("You must have explicit, authorized permission to access this application. \nUnauthorized attempts to access or use this app may result in criminal penalties.")
	fmt.Println("------------------------------------------------------------------------------------------|")

	var usr, pswd string
	fmt.Println("Enter Username: ")
	fmt.Scanln(&usr)
	fmt.Println("Enter Password: ")
	fmt.Scanln(&pswd)

        // Change generic user* to username of individuals that will use the application
       	user1 := login{username: usr, password: pswd}
	user2 := login{username: usr, password: pswd}
	
	if user1.username != l.username || user1.password != l.password {
		log.Println("[!] Unauthorized Attempt")
		os.Exit(1)
	}

	if user2.username != l.username || user2.password != l.password {
		log.Println("[!] Unauthorized Attempt")
		return
	}

}


func menu() {
	fmt.Println("1. Total Licenses Used")
	fmt.Println("2. Access Token")
	fmt.Println("3. Get host metadata")
	fmt.Println("4. Start Batch Session")
	fmt.Println("5. Run batch CMD")
	fmt.Println("6. Exit")
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
			var ret string
			fmt.Println("\n\nReturn to main menu? y/n")
			fmt.Scan(&ret)
			if ret == "y" {
				menu()
			} else if ret == "n" {

			}
		case 2:
			var token string
			fmt.Println("Do you need a new access token? Enter yes or no")
			fmt.Scanln(&token)
			if token == "yes" {
				api.AccessToken()
				fmt.Println("An Oauth2 token has been created and saved.")
			} else {
				fmt.Println("Exiting Falcon...")
				os.Exit(0)
			}
		case 3:
			var HIDS string
			fmt.Println("Enter Host: ")
			fmt.Scanln(&HIDS)
			api.FindInfo(api.FindHost(HIDS))

		case 4:
			rtr.StartSession()
		case 5:
			//batch session ID
			var batchID string
			fmt.Println("Please enter the Batch ID: ")
			fmt.Scanln(&batchID)
			rtr.ScriptRun(batchID)
		case 6:
			fmt.Println("Exiting Falcon...")
			os.Exit(0)
		}
	}

}

func main() {
        ulogin := &login{username: "test", password: "test"}
	title()
	programExit()
	menu()

}
