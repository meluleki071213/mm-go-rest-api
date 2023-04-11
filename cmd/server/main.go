package main

import "fmt"

func Run() error {
	fmt.Println("Starting up the application...")
	return nil
}

func main()  {
	fmt.Println("EntryPoint: Rest API")

	if err:= Run(); err != nil {
		fmt.Println(err)
	}
}
