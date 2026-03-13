package main

import (
	"fmt"
)

func fetch() {
	art := `
      @@****@@      
   @%**********#@   
  %*  +*******  *%      OS Void space V 0.1 
@%****************%@
@#****************#@
@%#######%%%%%%%%%%@
@@@@@@@@@@@@@@@@@@@@
@@@@@          @@@@@
  @@@*        +@@@  
   @@@@@@@@@@@@@@   
      @@@@@@@@      
	`
	fmt.Println(art)
}

func math() {
	fmt.Println("Калькулятор")
	var znack string
	var b float64
	var a float64
	var rav float64
	fmt.Println("Введите знак")
	fmt.Scan(&znack)
	fmt.Println("Введите a")
	fmt.Scan(&a)
	fmt.Println("Введите b")
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

func help() {
	fmt.Println("match - калькулятор")
	fmt.Println("fetch - информация о системе")
	fmt.Println("help - это меню")
}

func main() {
	var nad string
	fetch()
	fmt.Println("Привет Pig")

	for {
		fmt.Scan(&nad)

		switch {
		case nad == "math":
			math()
		case nad == "fetch":
			fetch()
		case nad == "help":
			help()
		}
	}

}
