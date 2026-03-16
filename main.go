package main

import (
	"fmt"
	"os"

	"math/rand/v2"

	"github.com/fatih/color"
	"github.com/klauspost/cpuid/v2"
)

func fetch() { //analog fastfetch
	fmt.Println("----------------------------") //system information
	fmt.Println("OS: Void Tree V0.8")
	fmt.Println("Cpu", cpuid.CPU.BrandName)
	fmt.Println("----------------------------")
	art := color.GreenString(`        
              @@              
            @@  @@            
          @@@@@@@ @@          
         @@@@    @@@@         
        @@@@      @@@@        
       @@    @@@@    @@       
     @@@   @@    @@   @@@     
     @@  @@@      @@@  @@     
    @@@@@@          @@@@@@    
 @@@@@@@     @@@@     @@@@@@@ 
 @ @@@     @@    @@     @@@ @ 
 @@@     @@@      @@@     @@@ 
 @@    @@@    @@    @@@    @@ 
  @@ @@@    @@  @@    @@@ @@  
   @@@    @@@     @@    @@@@  
    @@@@@@@@@@@@@@@@@@@@@@      
	`)
	fmt.Println(art)
}

func math() { //calculator
	fmt.Println("calculator")
	var znack string
	var b float64
	var a float64
	var rav float64
	fmt.Println("Enter sign")
	fmt.Scan(&znack)
	fmt.Println("Enter a")
	fmt.Scan(&a)
	fmt.Println("Enter b")
	fmt.Scan(&b)
	switch {
	case znack == "+":
		rav = a + b
	case znack == "-":
		rav = a - b
	case znack == "*":
		rav = a * b
	case znack == "/":
		rav = a / b
	default:
		fmt.Println("Error")
	}
	fmt.Println(rav)
}

func sample() {
	file, err := os.Create("sample.go")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	data := `
	package main

import (
	"fmt"

	"os"
)

func main() {
	fmt.Println("Hello, World!")
}
`
	file.WriteString(data)

}

func info() {
	fmt.Println("-------------------------------")
	fmt.Println("OS: Void Tree")
	fmt.Println("Version: V0.8")
	fmt.Println("Created by: pigetchi")
	fmt.Println("Vritten in: GO")
	fmt.Println("License: MIT")
	fmt.Println("-------------------------------")
}

func random() {
	var nad int
	var diop int
	fmt.Println("enter a range from 1 to")
	fmt.Scan(&diop)
	num := rand.IntN(diop)
	fmt.Println("guess")
	fmt.Scan(nad)
	switch {
	case nad == num:
		fmt.Println("you guessed it")
	default:
		fmt.Println("try again")
	}
}

func user() {
	var name string
	var nad string
	fmt.Println("Do you want to create/change a name?")
	fmt.Scan(&nad)
	if nad == "yes" {
		fmt.Scan(&name)
	}
	os.WriteFile("player.txt", []byte(name), 0644)
}

func help() {
	fmt.Println("math - calculator")
	fmt.Println("fetch - system information")
	fmt.Println("help - available commands")
	fmt.Println("sample - create a template Go")
	fmt.Println("info - information about the program")
	fmt.Println("random - guess the number")
	fmt.Println("name - see your name")
	fmt.Println("exit - exit")
}

func main() {
	username, _ := os.ReadFile("player.txt")
	krasivo := color.GreenString("🌿│ ~ ")
	var nad string
	fetch()
	user()
	fmt.Println("Hello,", string(username))

	for {
		fmt.Println(" ")
		fmt.Print(krasivo)
		fmt.Scan(&nad)

		switch {
		case nad == "math":
			math()
		case nad == "fetch":
			fetch()
		case nad == "help":
			help()
		case nad == "sample":
			sample()
		case nad == "exit":
			os.Exit(0)
		case nad == "info":
			info()
		case nad == "random":
			random()
		case nad == "name":
			fmt.Println(string(username))
		default:
			fmt.Printf("command %s does not exist, type help for a list of commands", nad)
		}
	}
}
