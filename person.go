package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Trait struct {
	Abilities map[string]float64
	Alleles   map[string]Allele
	Gender    string
	EyeColor  string
	HairColor string
}

type Allele struct {
	Pos1 bool
	Pos2 bool
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
	person.EyeColor = getEyeColor(person.Alleles["ec1"], person.Alleles["ec2"])
	person.HairColor = getHairColor(person.Alleles["ec1"], person.Alleles["ec2"])
	if errText == "" {
		return person, nil
	} else {
		return person, errors.New(errText)
	}

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
	trait.Abilities["1_strength"] = GetValue()
	trait.Abilities["2_intelligence"] = GetValue()
	trait.Abilities["3_wisdom"] = GetValue()
	trait.Abilities["4_dexterity"] = GetValue()
	trait.Abilities["5_constitution"] = GetValue()
	trait.Abilities["6_charisma"] = GetValue()

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

func GetGender() string {
	rand.Seed(time.Now().UnixNano())
	if gender := rand.Intn(2) == 1; gender {
		return "Y"
	} else {
		return "X"
	}
}

func (p *Person) Describe() []string {
	result := make([]string, len(p.Abilities))

	return result
}

func abilityDescription(ability string, value float64) string {
	switch ability {
	case "1_strength":
		switch value {
		case 1:
			return "Morbidly weak, has significant trouble lifting own limbs."
		case 2, 3:
			return "Needs help to stand, can be knocked over by strong breezes."
		case 4, 5:
			return "Knocked off balance by swinging something dense."
		case 6, 7:
			return "Difficulty pushing an object of their weight."
		case 8, 9:
			return "Has trouble even lifting heavy objects."
		case 10, 11:
			return "Can literally pull their own weight."
		case 12, 13:
			return "Carries heavy objects for short distances."
		case 14, 15:
			return "Visibly toned, throws small objects for long distances."
		case 16, 17:
			return "Carries heavy objects with one arm."
		case 18, 19:
			return "Can break objects like wood with bare hands."
		case 20, 21:
			return "Able to out-wrestle a work animal or catch a falling person."
		case 22, 23:
			return "Can pull very heavy objects at appreciable speeds."
		case 24, 25:
			return "Pinnacle of brawn, able to out-lift several people."
		}

	case "2_intelligence":
		switch value {
		case 1:
			return "Animalistic, no longer capable of logic or reason."
		case 2, 3:
			return "Barely able to function, very limited speech and knowledge."
		case 4, 5:
			return "Often resorts to charades to express thoughts."
		case 6, 7:
			return "Often misuses and mispronounces words."
		case 8, 9:
			return "Has trouble following trains of thought, forgets most unimportant things."
		case 10, 11:
			return "Knows what they need to know to get by."
		case 12, 13:
			return "Knows a bit more than is necessary, fairly logical."
		case 14, 15:
			return "Able to do math or solve logic puzzles mentally with reasonable accuracy."
		case 16, 17:
			return "Fairly intelligent, able to understand new tasks quickly."
		case 18, 19:
			return "Very intelligent, may invent new processes or uses for knowledge."
		case 20, 21:
			return "Highly knowledgeable, probably the smartest person many people know."
		case 22, 23:
			return "Able to make Holmesian leaps of logic."
		case 24, 25:
			return "Famous as a sage and genius."
		}
	case "3_wisdom":
		switch value {
		case 1:
			return "Seemingly incapable of thought, barely aware."
		case 2, 3:
			return "Rarely notices important or prominent items, people, or occurrences."
		case 4, 5:
			return "Seemingly incapable of forethought."
		case 6, 7:
			return "Often fails to exert common sense."
		case 8, 9:
			return "Forgets or opts not to consider options before taking action."
		case 10, 11:
			return "Makes reasoned decisions most of the time."
		case 12, 13:
			return "Able to tell when a person is upset."
		case 14, 15:
			return "Can get hunches about a situation that doesn’t feel right."
		case 16, 17:
			return "Reads people and situations fairly well."
		case 18, 19:
			return "Often used as a source of wisdom or decider of actions."
		case 20, 21:
			return "Reads people and situations very well, almost unconsciously."
		case 22, 23:
			return "Can tell minute differences among many situations."
		case 24, 25:
			return "Nearly prescient, able to reason far beyond logic."
		}
	case "4_dexterity":
		switch value {
		case 1:
			return "Barely mobile, probably significantly paralyzed."
		case 2, 3:
			return "Incapable of moving without noticeable effort or pain."
		case 4, 5:
			return "Visible paralysis or physical difficulty."
		case 6, 7:
			return "Significant klutz or very slow to react."
		case 8, 9:
			return "Somewhat slow, occasionally trips over own feet."
		case 10, 11:
			return "Capable of usually catching a small tossed object."
		case 12, 13:
			return "Able to often hit large targets."
		case 14, 15:
			return "Can catch or dodge a medium-speed surprise projectile."
		case 16, 17:
			return "Able to often hit small targets."
		case 18, 19:
			return "Light on feet, able to often hit small moving targets."
		case 20, 21:
			return "Graceful, able to flow from one action into another easily."
		case 22, 23:
			return "Very graceful, capable of dodging a number of thrown objects."
		case 24, 25:
			return "Moves like water, reacting to all situations with almost no effort."
		}
	case "5_constitution":
		switch value {
		case 1:
			return "Minimal immune system, body reacts violently to anything foreign."
		case 2, 3:
			return "Frail, suffers frequent broken bones."
		case 4, 5:
			return "Bruises very easily, knocked out by a light punch."
		case 6, 7:
			return "Unusually prone to disease and infection."
		case 8, 9:
			return "Easily winded, incapable of a full day’s hard labor."
		case 10, 11:
			return "Occasionally contracts mild sicknesses."
		case 12, 13:
			return "Can take a few hits before being knocked unconscious."
		case 14, 15:
			return "Able to labor for twelve hours most days."
		case 16, 17:
			return "Easily shrugs off most illnesses."
		case 18, 19:
			return "Able to stay awake for days on end."
		case 20, 21:
			return "Very difficult to wear down, almost never feels fatigue."
		case 22, 23:
			return "Never gets sick, even to the most virulent diseases."
		case 24, 25:
			return "Tireless paragon of physical endurance."
		}
	case "6_charisma":
		switch value {
		case 1:
			return "Barely conscious, probably acts heavily autistic."
		case 2, 3:
			return "Minimal independent thought, relies heavily on others to think instead."
		case 4, 5:
			return "Has trouble thinking of others as people."
		case 6, 7:
			return "Terribly reticent, uninteresting, or rude."
		case 8, 9:
			return "Something of a bore or makes people mildly uncomfortable."
		case 10, 11:
			return "Capable of polite conversation."
		case 12, 13:
			return "Mildly interesting, knows what to say to the right people."
		case 14, 15:
			return "Interesting, knows what to say to most people."
		case 16, 17:
			return "Popular, receives greetings and conversations on the street."
		case 18, 19:
			return "Immediately likeable by many people, subject of favorable talk."
		case 20, 21:
			return "Life of the party, able to keep people entertained for hours."
		case 22, 23:
			return "Immediately likeable by almost everybody."
		case 24, 25:
			return "Renowned for wit, personality, and/or looks."
		}

	}
	return ""
}
