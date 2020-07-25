package main

import (
	"fmt"
	"goStuff/CrowdStrikeAPI/api"
)

func title() {

       fmt.Println(" ██████╗  ██████╗       ███████╗ █████╗ ██╗      ██████╗ ██████╗ ███╗   ██╗")
       fmt.Println("██╔════╝ ██╔═══██╗      ██╔════╝██╔══██╗██║     ██╔════╝██╔═══██╗████╗  ██║")
       fmt.Println("██║  ███╗██║   ██║█████╗█████╗  ███████║██║     ██║     ██║   ██║██╔██╗ ██║")
       fmt.Println("██║   ██║██║   ██║╚════╝██╔══╝  ██╔══██║██║     ██║     ██║   ██║██║╚██╗██║")
       fmt.Println("╚██████╔╝╚██████╔╝      ██║     ██║  ██║███████╗╚██████╗╚██████╔╝██║ ╚████║")
       fmt.Println("╚═════╝  ╚═════╝       ╚═╝     ╚═╝  ╚═╝╚══════╝ ╚═════╝ ╚═════╝ ╚═╝  ╚═══╝ ")

	
	fmt.Println("         	  ____                      ")
	fmt.Println("                 /   \\                     ")
	fmt.Println("                 /    \\                    ")
	fmt.Println("      ///////////^^^^^^\\\\\\\\\\\\         ")
	fmt.Println("     ///////////        \\\\\\\\\\\\        ")
	fmt.Println("    ////// \\\\\\      //////  \\\\\\       ")
	fmt.Println("   /////       \\      //     \\\\\\      ")
	fmt.Println("  ////          //    \\        \\\\     ")
	fmt.Println(" ///           ///||\\\\          \\\\   ")
	fmt.Println("//            | //||\\ |          \\\\   ")
	fmt.Println("/               ^^^^^^             \\ ")

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Made with <3 by i_obj")
	return
}



func main() {
	
	title()
	var choice int
	for ok := true; ok; ok = (choice != 3) {
		n, err := fmt.Scanln(&choice)
		if n > 1 || err != nil {
			fmt.Println("[!] Invalid input")
			fmt.Println("[!] Entry not found, try again.")
			continue
		}
		switch choice {
		case 1:
			fmt.Println("test1")
		case 2:
			fmt.Println(api.AccessToken())
		case 3:
			fmt.Println("Exiting Falcon...")
			os.Exit(2)
		}
	}
	
	//1. Total license count
	//2. Search for host
	//3. Contain host
	//4. refresh token

	//fmt.Println(api.AccessToken())
	fmt.Println(api.FindHost("Mac"))

}
