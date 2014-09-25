package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Trait struct {
	Abilities map[string]float64
	Gender    string
}

type Person struct {
	Trait
	Father Trait
	Mother Trait
}

func Procreate(father Person, mother Person) (Person, error) {
	fatherTraits := father.GetContributionTraits()
	motherTraits := mother.GetContributionTraits()
	return GeneratePerson(fatherTraits, motherTraits)
}

func (p *Person) GetContributionTraits() Trait {
	result := Trait{}
	result.Abilities = make(map[string]float64)
	for key, _ := range p.Father.Abilities {
		pick := rand.Float64()
		if pick <= .5 {
			result.Abilities[key] = p.Father.Abilities[key]
		} else {
			result.Abilities[key] = p.Mother.Abilities[key]
		}
	}
	if p.Gender == "X" {
		result.Gender = "X"
	} else {
		result.Gender = GetGender()
	}
	return result
}

func GeneratePerson(fatherTrait Trait, motherTrait Trait) (Person, error) {
	person := Person{}
	person.Abilities = make(map[string]float64)
	person.Father = fatherTrait
	person.Mother = motherTrait
	for key, _ := range fatherTrait.Abilities {
		err := errors.New("")
		if fatherTrait.Abilities[key] <= 2.0 {
			fmt.Println("Father multi: ", key, fatherTrait.Abilities[key])
		}
		if motherTrait.Abilities[key] <= 2.0 {
			fmt.Println("Mother multi: ", key, motherTrait.Abilities[key])
		}
		if person.Abilities[key], err = selectValue(fatherTrait.Abilities[key], motherTrait.Abilities[key]); err != nil {
			return person, err
		}

	}
	person.Gender = GetGender()
	return person, nil
}

func selectValue(fatherValue float64, motherValue float64) (float64, error) {
	if fatherValue > 2 && motherValue > 2 {
		if fatherValue >= motherValue {
			return fatherValue, nil
		} else {
			return motherValue, nil
		}
	}
	if (fatherValue > 2.0 && motherValue <= 2.0) || (motherValue > 2.0 && fatherValue <= 2.0) {
		if multiplierValue := float64(int(fatherValue * motherValue)); multiplierValue > 1.0 {
			return float64(int(fatherValue * motherValue)), nil
		} else {
			return 0.0, errors.New("Cannot thrive.  Value is less than 1.")
		}

	}
	return 3.0, nil
}

func GenerateTraits(gender string) Trait {
	trait := Trait{}
	trait.Abilities = make(map[string]float64)
	trait.Abilities["1_strength"] = GetValue()
	trait.Abilities["2_intelligence"] = GetValue()
	trait.Abilities["3_wisdom"] = GetValue()
	trait.Abilities["4_dexterity"] = GetValue()
	trait.Abilities["5_constitution"] = GetValue()
	trait.Abilities["6_charisma"] = GetValue()
	trait.Gender = gender
	return trait
}

func GetValue() float64 {
	rand.Seed(time.Now().UnixNano())
	scoreType := float64(rand.Intn(10)) + 1.0
	if scoreType == 1 {
		return float64(rand.Intn(2000)) / 1000.0
	} else {
		return float64(rand.Intn(16)) + 3.0
	}
}

func GetGender() string {
	rand.Seed(time.Now().UnixNano())
	if gender := rand.Intn(2) == 1; gender {
		return "Y"
	} else {
		return "X"
	}
}
