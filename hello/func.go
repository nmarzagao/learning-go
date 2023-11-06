package main
import "fmt"

func printHello(name string) {
	if name == "Nicolas" {
		fmt.Println("Welcome back!")
	} else {
		fmt.Printf("Hello %s!\n", name)
	}
}

func getName() string {
	var name string

	fmt.Print("Whats your name? ")
	fmt.Scan(&name)
	
	return name
}

func countdown(max int) {
	for i := 0; i < max; i++ {
		fmt.Println(i)
	}
	fmt.Println("Done!")
}