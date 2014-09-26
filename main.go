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
	person1.Gender = "Y"
	person2, _ := GeneratePerson(trait2Father, trait2Mother)
	person2.Gender = "X"
	PrintPerson("Father", person1, keys)
	PrintPerson("Mother", person2, keys)
	for i := 0; i < 10; i++ {
		if child, chldErr := Procreate(person1, person2); chldErr == nil {
			PrintPerson(fmt.Sprintf("Child %v", i+1), child, keys)
		} else {
			fmt.Printf("----------\n")
			fmt.Printf("Child %v, something didn't go right.\n%v\n", i+1, chldErr)
		}

	}
}

func PrintPerson(name string, p Person, keys []string) {
	fmt.Printf("----------\n")
	fmt.Println(name)
	fmt.Printf("Eyes: %v\n 1:%v 2:%v \n", p.EyeColor, p.Alleles["ec1"], p.Alleles["ec2"])
	for _, k := range keys {
		//fmt.Printf("%v: %v F:%v M:%v\n", k, p.Abilities[k], p.Father.Abilities[k], p.Mother.Abilities[k])
		_ = k
		//fmt.Printf("%v: %v \n", k, p.Abilities[k])
	}
	fmt.Println("Gender:", p.Gender)
}
