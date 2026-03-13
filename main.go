package main

import (
	"fmt"
	"os"

	"github.com/klauspost/cpuid/v2"
)

func fetch() { //analog neofetch
	fmt.Println("----------------------------") //system information
	fmt.Println("OS: Void Space V0.2")
	fmt.Println("Cpu", cpuid.CPU.BrandName)
	fmt.Println("----------------------------")
	art := `        //art
      @@****@@      
   @%**********#@   
  %*  +*******  *%    
@%****************%@ 
@#****************#@
@%################%@
@@@@@@@@@@@@@@@@@@@@
@@@@@          @@@@@
  @@@*        +@@@  
   @@@@@@@@@@@@@@   
      @@@@@@@@      
	`
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

}
`
	file.WriteString(data)

}

func help() {
	fmt.Println("match - calculator")
	fmt.Println("fetch - system information")
	fmt.Println("help - available commands")
	fmt.Println("sample - create a template Go")
}

func main() {
	var nad string
	fetch()

	for {
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
		}
	}

}
