package main

import (
	"fmt"
	"os"

	"github.com/klauspost/cpuid/v2"
)

func fetch() { //analog neofetch
	fmt.Println("----------------------------") //system information
	fmt.Println("OS: Void Space V0.5")
	fmt.Println("Cpu", cpuid.CPU.BrandName)
	fmt.Println("----------------------------")
	art := `        
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

func help() {
	fmt.Println("math - calculator")
	fmt.Println("fetch - system information")
	fmt.Println("help - available commands")
	fmt.Println("sample - create a template Go")
	fmt.Println("exit - exit")
}

func main() {
	var nad string
	fetch()

	for {
		fmt.Println(" ")
		fmt.Print("V.S.│~ ")
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
		default:
			fmt.Printf("command %s does not exist, type help for a list of commands", nad)
		}
	}

}
