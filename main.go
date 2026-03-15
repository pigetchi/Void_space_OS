package main

import (
	"fmt"
	"os"

	"math/rand/v2"

	"github.com/fatih/color"
	"github.com/klauspost/cpuid/v2"
)

var level int
var xp int
var name string

func fetch() { //analog fastfetch
	fmt.Println("----------------------------") //system information
	fmt.Println("OS: Void Space V0.7")
	fmt.Println("Cpu", cpuid.CPU.BrandName)
	fmt.Println("----------------------------")
	art := color.BlueString(`        
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

/*func translator() {
	var gaster map[string]string = map[string]string{
		"A": "✌",
		"B": "👌",
		"C": "👍",
		"D": "👎",
		"E": "☜",
		"F": "☞",
		"G": "☝",
		"H": "☟",
		"I": "✋",
		"J": "☺",
		"K": "😐",
		"L": "☹",
		"M": "💣",
		"N": "☠",
		"O": "⚐",
		"P": "🏱",
		"Q": "✈",
		"R": "☼",
		"S": "💧",
		"T": "❄",
		"U": "🕆",
		"V": "✞",
		"W": "🕈",
		"X": "✠",
		"Y": "✡",
		"Z": "☪",
		"1": "📂",
		"2": "📄",
		"3": "🗏",
		"4": "🗐",
		"5": "🗄",
		"6": "⌛",
		"7": "🖮",
		"8": "🖰",
		"9": "🖲",
		"0": "📁",
	}
	text := "HELLO"

} */

func info() {
	fmt.Println("-------------------------------")
	fmt.Println("OS: Void Space")
	fmt.Println("Version: V0.7")
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

func help() {
	fmt.Println("math - calculator")
	fmt.Println("fetch - system information")
	fmt.Println("help - available commands")
	fmt.Println("sample - create a template Go")
	fmt.Println("info - information about the program")
	fmt.Println("random - guess the number")
	fmt.Println("exit - exit")
}

func main() {
	krasivo := color.GreenString("V.S.│~ ")
	var nad string
	fetch()

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
		default:
			fmt.Printf("command %s does not exist, type help for a list of commands", nad)
		}
	}
}
