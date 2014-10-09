package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"

	"math"
	"math/rand"
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
	Name      string
	Abilities map[string]float64
	Alleles   map[string]Allele
	Gender    string
	EyeColor  string
	HairColor string
	Vision    string
	Height    Height
	AvgScore  float64
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

type Person struct {
	Trait
	Father Trait
	Mother Trait
}

func Procreate(name string, father Person, mother Person) (Person, error) {
	fatherTraits := father.GetContributionTraits()
	motherTraits := mother.GetContributionTraits()
	return GeneratePerson(name, fatherTraits, motherTraits)
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

func (p *Person) setHeight() {
	result := Height{}
	totalinches := (48. + p.Abilities[STRENGTH] + p.Abilities[CONSTITUTION])
	if p.Gender == "X" {
		totalinches = totalinches * 0.9
	}
	feet, _ := math.Modf(totalinches / 12.0)
	inches := math.Mod(totalinches, 12.0)
	result.Feet = int(feet)
	result.Inches = int(inches)
	p.Height = result
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
