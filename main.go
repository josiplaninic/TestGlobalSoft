package main

import (
	"fmt"
)

type Employee struct {
	Name     string
	Position string
	Tasks    []string
}

func (e *Employee) AddTask(task string) {
	e.Tasks = append(e.Tasks, task)
}

func (e *Employee) CompleteTasks() {
	
	for i := len(e.Tasks) - 1; i >= 0; i-- {
		fmt.Printf("Employee <%s> completed task <%s>\n", e.Name, e.Tasks[i])
		if len(e.Tasks) > 0 {
			e.Tasks = e.Tasks[:len(e.Tasks)-1]
		}
	}
}

func main() {

	marko := Employee{"Marko Markic", "Manager", []string{"inicijaliziraniZadatakMarko"}}
	ivan := Employee{"Ivan Ivankovic", "PR", []string{"inicijaliziraniZadatakIvan"}}

	var proizvodi = map[string]int {
		"meso" : 180,
		"mlijeko" : 20,
		"sir" : 300,
		"sok" : 10,
		"majica" : 460,
	}

	printResults(proizvodi)


	marko.AddTask("task1marko")
	marko.AddTask("task2marko")
	marko.AddTask("task3marko")
	
	ivan.AddTask("task1ivan")
	ivan.AddTask("task2ivan")
	ivan.AddTask("task3ivan")


	marko.CompleteTasks()
	ivan.CompleteTasks()

}

func printResults(mapa map[string]int) {
	var mostExpensive string
	biggestPrice := 0

	for name, price := range mapa {
		if price > 150 {
			fmt.Printf("\nProizvod <%s> ima cijenu <%d>\n", name, price)
		}

		if price > biggestPrice {
			biggestPrice = price
			mostExpensive = name
		}
	}

	fmt.Printf("\nNajskuplji proizvod je <%s> s cijenom od <%d>\n", mostExpensive, biggestPrice)
}
