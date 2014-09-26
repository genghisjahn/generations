package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Trait struct {
	Abilities map[string]float64
	Alleles   map[string]bool
	Gender    string
	EyeColor  string
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
	result.Alleles = make(map[string]bool)
	for key, _ := range p.Father.Abilities {
		pick := rand.Float64()
		if pick <= .5 {
			result.Abilities[key] = p.Father.Abilities[key]
		} else {
			result.Abilities[key] = p.Mother.Abilities[key]
		}
	}
	for key, _ := range p.Father.Alleles {
		pick := rand.Float64()
		if pick < .5 {
			result.Alleles[key] = p.Father.Alleles[key]
		} else {
			result.Alleles[key] = p.Mother.Alleles[key]
		}
		//log.Printf("%v %v %v %v", key, result.Alleles[key], p.Father.Alleles[key], p.Mother.Alleles[key])
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
	person.Alleles = make(map[string]bool)
	person.Father = fatherTrait
	person.Mother = motherTrait
	errText := ""
	for key, _ := range fatherTrait.Abilities {
		err := errors.New("")
		if person.Abilities[key], err = selectValue(key, fatherTrait.Abilities[key], motherTrait.Abilities[key]); err != nil {
			errText += fmt.Sprintf("%v\n", err)
		}

	}
	for key, _ := range fatherTrait.Alleles {

		person.Alleles[key] = selectAllele(fatherTrait.Alleles[key], motherTrait.Alleles[key])
		if fatherTrait.Alleles[key] == motherTrait.Alleles[key] && person.Alleles[key] != fatherTrait.Alleles[key] {
			log.Println("WTF??!?!?!?!")
		}
	}
	person.Gender = GetGender()
	person.setEyeColor()
	if errText == "" {
		return person, nil
	} else {
		return person, errors.New(errText)
	}

}

func selectAllele(fatherAllele bool, motherAllele bool) bool {
	pickP := rand.Float64()
	if pickP <= .5 {
		return fatherAllele
	} else {
		return motherAllele
	}
}

func selectValue(ability string, fatherValue float64, motherValue float64) (float64, error) {
	if fatherValue > 2 && motherValue > 2 {
		if fatherValue >= motherValue {
			return fatherValue, nil
		} else {
			return motherValue, nil
		}
	}
	if (fatherValue > 2.0 && motherValue <= 2.0) || (motherValue > 2.0 && fatherValue <= 2.0) {
		if multiplierValue := float64(int(fatherValue * motherValue)); multiplierValue > 3.0 {
			return float64(int(fatherValue * motherValue)), nil
		} else {
			return 3.0, nil
		}

	}
	return 0.0, errors.New(fmt.Sprintf("Cannot thrive because of extremely low %v.", ability))
}

func GenerateTraits(gender string) Trait {
	trait := Trait{}
	trait.Abilities = make(map[string]float64)
	trait.Alleles = make(map[string]bool)
	trait.Abilities["1_strength"] = GetValue()
	trait.Abilities["2_intelligence"] = GetValue()
	trait.Abilities["3_wisdom"] = GetValue()
	trait.Abilities["4_dexterity"] = GetValue()
	trait.Abilities["5_constitution"] = GetValue()
	trait.Abilities["6_charisma"] = GetValue()

	trait.Alleles["ec1p"] = GetAllele()
	trait.Alleles["ec1m"] = GetAllele()
	trait.Alleles["ec2p"] = GetAllele()
	trait.Alleles["ec2m"] = GetAllele()

	trait.Gender = gender
	return trait
}

func GetValue() float64 {
	rand.Seed(time.Now().UnixNano())
	scoreType := float64(rand.Intn(10)) + 1.0
	if scoreType == 1 {
		key := rand.Intn(8)
		multivals := make(map[int]float64)
		multivals[0] = 0.25
		multivals[1] = 0.5
		multivals[2] = 0.75
		multivals[3] = 1.0
		multivals[4] = 1.25
		multivals[5] = 1.5
		multivals[6] = 1.75
		multivals[7] = 2.0
		return multivals[key]
	} else {
		return float64(rand.Intn(16)) + 3.0
	}
}

func GetAllele() bool {
	rand.Seed(time.Now().UnixNano())
	if result := rand.Intn(2) == 1; result {
		return true
	} else {
		return false
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

func (p *Person) setEyeColor() {
	ec1p := p.Alleles["ec1p"]
	ec1m := p.Alleles["ec1m"]
	ec2p := p.Alleles["ec2p"]
	ec2m := p.Alleles["ec2m"]

	eyecolor := ""
	if ec1p || ec1m {
		eyecolor = "brown"
	} else {
		if ec2p || ec2m {
			eyecolor = "green"
		} else {
			eyecolor = "blue"
		}

	}
	p.EyeColor = eyecolor
}
