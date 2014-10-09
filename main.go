package main

import (
	"fmt"
)

func main() {

	pFather, pMother := GetProtoParents()
	PrintPerson(pFather)
	PrintPerson(pMother)

	for i := 0; i < 10; i++ {
		if child, chldErr := Procreate(fmt.Sprintf("Child %v", i+1), pFather, pMother); chldErr == nil {
			PrintPerson(child)
		} else {
			fmt.Printf("----------\n")
			fmt.Printf("Child %v, something didn't go right.\n%v\n", i+1, chldErr)
		}
	}
}

func PrintPerson(p Person) {
	fmt.Printf("----------\n")
	fmt.Println(p.Name)
	fmt.Println("Hash:", p.GetHash())
	fmt.Println("AvgScore:", p.AvgScore)
	fmt.Println("Eye Color: ", p.EyeColor)
	fmt.Println("Vision:", p.Vision)
	fmt.Println("Hair Color: ", p.HairColor)
	fmt.Printf("Height: %v ft.   %v in.\n", p.Height.Feet, p.Height.Inches)

	fmt.Println("Description: ")
	for _, v := range p.Describe() {
		fmt.Println(v)
	}

	fmt.Println("Gender:", p.Gender)
	fmt.Println("Classes:", p.selectClasses())
}
