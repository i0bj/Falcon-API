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

func programExit() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r[!] Ctrl+C pressed. Program exiting..")
		os.Exit(0)
	}()
}




func main() {
        programExit()
	title()

        fmt.Println("1. Total Licenses Used")
	fmt.Println("2. Get host metadata")
	fmt.Println("3. Access Token")
	fmt.Println("4. Find a Host")
	fmt.Println("5. Start Batch Session")
	fmt.Println("6. Run CMD")
	fmt.Println("7. Exit")
	// choice variable will hold the selection you make
	// and will be used in the switch statement below
	var choice int
	for ok := true; ok; ok = (choice != 3) {
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
			api.LicenseTotal("5000")
		case 2:
			api.FindInfo(api.FindHost(""))
		case 3:
			fmt.Println(api.AccessToken())
		case 4:
			fmt.Println(api.FindHost(""))
		case 5:
			rtr.StartSession()
		case 6:
			rtr.ScriptRun("Batch session ID")
		case 7:
			fmt.Println("Exiting Falcon...")
			os.Exit(0)
		}
	}
