package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

var (
	STRENGTH     = "1_strength"
	INTELLIGENCE = "2_intelligence"
	WISDOM       = "3_wisdom"
	DEXTERITY    = "4_dexterity"
	CONSTITUTION = "5_constitution"
	CHARISMA     = "6_charisma"
)

type Trait struct {
	Abilities map[string]float64
	Alleles   map[string]Allele
	Gender    string
	EyeColor  string
	HairColor string
	Height    Height
	AvgScore  float64
}

type Allele struct {
	Pos1 bool
	Pos2 bool
}

type Height struct {
	Feet   int
	Inches int
}

func (p *Person) GetHash() string {
	jsonbody, _ := json.Marshal(p)
	h := hmac.New(sha512.New, nil)
	h.Write([]byte(jsonbody))
	hash := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return hash
}

func (p *Person) setAbilityAverage() float64 {
	total := 0.0
	for k, _ := range p.Abilities {
		total += p.Abilities[k]
	}
	return total / float64(len(p.Abilities))
}

func (a *Allele) Select() bool {
	rand.Seed(time.Now().UnixNano())
	pick := rand.Float64()
	if pick < .5 {
		return a.Pos1
	}
	return a.Pos2
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
	result.Alleles = make(map[string]Allele)
	for key, _ := range p.Father.Abilities {
		pick := rand.Float64()
		if pick < .5 {
			result.Abilities[key] = p.Father.Abilities[key]
		} else {
			result.Abilities[key] = p.Mother.Abilities[key]
		}
	}
	result.Alleles = p.Alleles
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
	person.Height = getHeight(person.Gender, person.Abilities[STRENGTH], person.Abilities[CONSTITUTION])
	if errText == "" {
		return person, nil
	} else {
		return person, errors.New(errText)
	}

}

func getHeight(gender string, strength float64, constitution float64) Height {
	result := Height{}
	totalinches := (48.0 + strength + constitution)
	if gender == "X" {
		totalinches = totalinches * 0.9
	}
	feet, _ := math.Modf(totalinches / 12.0)
	inches := math.Mod(totalinches, 12.0)
	result.Feet = int(feet)
	result.Inches = int(inches)
	return result
}

func getHairColor(hc1 Allele, hc2 Allele) string {
	result := ""
	if hc1.Pos1 && hc1.Pos2 {
		result = "Black"
	} else {
		if hc1.Pos1 || hc1.Pos1 {
			result = "Brown"
		} else {
			if hc2.Pos1 || hc2.Pos2 {
				result = "Blonde"
			} else {
				result = "Red"
			}
		}
	}
	return result
}

func getEyeColor(ec1 Allele, ec2 Allele) string {
	result := ""
	if ec1.Pos1 || ec1.Pos2 {
		result = "Brown"
	} else {
		if ec2.Pos1 || ec2.Pos2 {
			result = "Green"
		} else {
			result = "Blue"
		}
	}
	return result
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

	trait.Gender = gender
	return trait
}

func GenerateAllele() Allele {
	rand.Seed(time.Now().UnixNano())
	allele := Allele{Pos1: false, Pos2: false}
	pick := rand.Float64()
	if pick > .5 {
		allele.Pos1 = true
	}
	pick = rand.Float64()
	if pick > .5 {
		allele.Pos2 = true
	}
	return allele
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

type ClassRequirements struct {
	AbilityName  string
	MininumScore float64
}
type Class struct {
	Name         string
	Requirements []ClassRequirements
}

func buildClass(name string, requirements map[string]float64) Class {
	result := Class{Name: name}
	for k, v := range requirements {
		requirement := ClassRequirements{AbilityName: k, MininumScore: v}
		result.Requirements = append(result.Requirements, requirement)
	}
	return result
}

func buildClassRequirements() []Class {
	classes := make([]Class, 0)

	requirements := make(map[string]float64) // := because this is the first one.
	requirements[WISDOM] = 9.0
	classes = append(classes, buildClass("Cleric", requirements))

	requirements = make(map[string]float64)
	requirements[WISDOM] = 12.0
	requirements[CHARISMA] = 12.0
	classes = append(classes, buildClass("Druid", requirements))

	requirements = make(map[string]float64)
	requirements[STRENGTH] = 9.0
	classes = append(classes, buildClass("Fighter", requirements))

	requirements = make(map[string]float64)
	requirements[STRENGTH] = 12.0
	requirements[INTELLIGENCE] = 9
	requirements[WISDOM] = 13.0
	requirements[CONSTITUTION] = 9.0
	requirements[CHARISMA] = 17.0
	classes = append(classes, buildClass("Paladin", requirements))

	requirements = make(map[string]float64)
	requirements[STRENGTH] = 13.0
	requirements[WISDOM] = 14.0
	requirements[CONSTITUTION] = 14.0
	classes = append(classes, buildClass("Ranger", requirements))

	requirements = make(map[string]float64)
	requirements[INTELLIGENCE] = 9.0
	requirements[DEXTERITY] = 6.0
	classes = append(classes, buildClass("Magic-User", requirements))

	requirements = make(map[string]float64)
	requirements[INTELLIGENCE] = 15.0
	requirements[DEXTERITY] = 16.0
	classes = append(classes, buildClass("Illusionist", requirements))

	requirements = make(map[string]float64)
	requirements[DEXTERITY] = 9.0
	classes = append(classes, buildClass("Thief", requirements))

	requirements = make(map[string]float64)
	requirements[STRENGTH] = 12.0
	requirements[INTELLIGENCE] = 11.0
	requirements[DEXTERITY] = 12.0
	classes = append(classes, buildClass("Assasin", requirements))

	requirements = make(map[string]float64)
	requirements[STRENGTH] = 15.0
	requirements[WISDOM] = 15.0
	requirements[DEXTERITY] = 15.0
	requirements[CONSTITUTION] = 11.0
	classes = append(classes, buildClass("Monk", requirements))

	requirements = make(map[string]float64)
	requirements[STRENGTH] = 15.0
	requirements[INTELLIGENCE] = 11.0
	requirements[DEXTERITY] = 15.0
	requirements[CONSTITUTION] = 10.0
	requirements[CHARISMA] = 15.0
	classes = append(classes, buildClass("Bard", requirements))

	return classes
}

func (p *Person) selectClasses() []string {
	results := make([]string, 0)
	classes := buildClassRequirements()

	for _, value1 := range classes {
		isClass := true
		for _, value2 := range value1.Requirements {
			if p.Abilities[value2.AbilityName] < value2.MininumScore {
				isClass = false
			}
		}
		if isClass {
			results = append(results, value1.Name)
		}
	}

	return results
}
