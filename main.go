package main

import (
	"fmt"
	"sort"
)

func main() {

	trait1Father := GenerateTraits("Y")
	var keys []string
	for k := range trait1Father.Abilities {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	trait2Father := GenerateTraits("Y")

	trait1Mother := GenerateTraits("X")
	trait2Mother := GenerateTraits("X")

	person1, _ := GeneratePerson(trait1Father, trait1Mother)
	person2, _ := GeneratePerson(trait2Father, trait2Mother)
	PrintPerson("#1", person1, keys)
	PrintPerson("#2", person2, keys)
	if person1.Gender != person2.Gender {
		for i := 0; i < 10; i++ {
			child := Person{}
			if person1.Gender == "X" {
				child, _ = Procreate(person1, person2)
			} else {
				child, _ = Procreate(person2, person1)
			}
			PrintPerson(fmt.Sprintf("Child %v", i+1), child, keys)
		}

	} else {
		fmt.Println("Try again.")
	}
}

func PrintPerson(name string, p Person, keys []string) {
	fmt.Printf("----------\n")
	fmt.Println(name)
	for _, k := range keys {
		fmt.Printf("%v: %v F:%v M:%v\n", k, p.Abilities[k], p.Father.Abilities[k], p.Mother.Abilities[k])
	}
	fmt.Println("Gender:", p.Gender)
}
