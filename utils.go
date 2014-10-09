package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var AbilityKeys []string

func GetProtoParents() (Person, Person) {
	trait1Father := GenerateTraits("Y")
	trait1Mother := GenerateTraits("X")
	trait2Father := GenerateTraits("Y")
	trait2Mother := GenerateTraits("X")

	protoFather, _ := GeneratePerson("Father", trait1Father, trait1Mother)
	protoFather.Gender = "Y"
	protoMother, _ := GeneratePerson("Mother", trait2Father, trait2Mother)
	protoMother.Gender = "X"
	return protoFather, protoMother
}

func getAbilityKeys() []string {
	if len(AbilityKeys) == 0 {
		keys := make([]string, 0, 0)
		keys = append(keys, STRENGTH)
		keys = append(keys, INTELLIGENCE)
		keys = append(keys, WISDOM)
		keys = append(keys, DEXTERITY)
		keys = append(keys, CONSTITUTION)
		keys = append(keys, CHARISMA)

		AbilityKeys = keys
		return keys
	}
	return AbilityKeys

}

func GeneratePerson(name string, fatherTrait Trait, motherTrait Trait) (Person, error) {
	person := Person{}
	person.Name = name
	person.Abilities = make(map[string]float64)
	person.Alleles = make(map[string]Allele)
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
		fatherAllele := fatherTrait.Alleles[key]
		motherAllel := motherTrait.Alleles[key]
		newAllele := Allele{Pos1: fatherAllele.Select(), Pos2: motherAllel.Select()}
		person.Alleles[key] = newAllele
	}
	person.Gender = GetGender()
	person.AvgScore = person.setAbilityAverage()
	person.EyeColor = getEyeColor(person.Alleles["ec1"], person.Alleles["ec2"])
	person.HairColor = getHairColor(person.Alleles["ec1"], person.Alleles["ec2"])
	person.Vision = getColorBlind(person.Alleles["ec1"], person.Alleles["ec2"], person.Gender)
	person.Handedness = setDdAllele(person.Alleles["h1"], "Right handed", "Left handed")
	person.setHeight()
	if errText == "" {
		return person, nil
	} else {
		return person, errors.New(errText)
	}

}

func GenerateTraits(gender string) Trait {
	trait := Trait{}
	trait.Abilities = make(map[string]float64)
	trait.Abilities[STRENGTH] = GetValue()
	trait.Abilities[INTELLIGENCE] = GetValue()
	trait.Abilities[WISDOM] = GetValue()
	trait.Abilities[DEXTERITY] = GetValue()
	trait.Abilities[CONSTITUTION] = GetValue()
	trait.Abilities[CHARISMA] = GetValue()

	trait.Alleles = make(map[string]Allele)
	trait.Alleles["ec1"] = GenerateAllele()
	trait.Alleles["ec2"] = GenerateAllele()
	trait.Alleles["hc1"] = GenerateAllele()
	trait.Alleles["hc2"] = GenerateAllele()

	trait.Alleles["cb1"] = GenerateAllele()
	trait.Alleles["cb2"] = GenerateAllele()

	trait.Alleles["h1"] = GenerateAllele()

	trait.Gender = gender
	return trait
}

func GetValue() float64 {
	rand.Seed(time.Now().UnixNano())
	scoreType := float64(rand.Intn(10)) + 1.0
	if scoreType == 1 {
		key := rand.Intn(12)
		multivals := make(map[int]float64)
		multivals[0] = 0.3
		multivals[1] = 0.4
		multivals[2] = 0.5
		multivals[3] = 0.6
		multivals[4] = 0.7
		multivals[5] = 0.8
		multivals[6] = 0.9
		multivals[7] = 1.0
		multivals[8] = 1.1
		multivals[9] = 1.2
		multivals[10] = 1.3
		multivals[11] = 1.4
		return multivals[key]
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
