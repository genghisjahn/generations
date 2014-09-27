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

	}
	return ""
}

/*
Dexterity
1 (–5): Barely mobile, probably significantly paralyzed
2-3 (–4): Incapable of moving without noticeable effort or pain
4-5 (–3): Visible paralysis or physical difficulty
6-7 (–2): Significant klutz or very slow to react
8-9 (–1): Somewhat slow, occasionally trips over own feet
10-11 (0): Capable of usually catching a small tossed object
12-13 (1): Able to often hit large targets
14-15 (2): Can catch or dodge a medium-speed surprise projectile
16-17 (3): Able to often hit small targets
18-19 (4): Light on feet, able to often hit small moving targets
20-21 (5): Graceful, able to flow from one action into another easily
22-23 (6): Very graceful, capable of dodging a number of thrown objects
24-25 (7): Moves like water, reacting to all situations with almost no effort
Constitution
1 (–5): Minimal immune system, body reacts violently to anything foreign
2-3 (–4): Frail, suffers frequent broken bones
4-5 (–3): Bruises very easily, knocked out by a light punch
6-7 (–2): Unusually prone to disease and infection
8-9 (–1): Easily winded, incapable of a full day’s hard labor
10-11 (0): Occasionally contracts mild sicknesses
12-13 (1): Can take a few hits before being knocked unconscious
14-15 (2): Able to labor for twelve hours most days
16-17 (3): Easily shrugs off most illnesses
18-19 (4): Able to stay awake for days on end
20-21 (5): Very difficult to wear down, almost never feels fatigue
22-23 (6): Never gets sick, even to the most virulent diseases
24-25 (7): Tireless paragon of physical endurance
Intelligence
1 (–5): Animalistic, no longer capable of logic or reason
2-3 (–4): Barely able to function, very limited speech and knowledge
4-5 (–3): Often resorts to charades to express thoughts
6-7 (–2): Often misuses and mispronounces words
8-9 (–1): Has trouble following trains of thought, forgets most unimportant things
10-11 (0): Knows what they need to know to get by
12-13 (1): Knows a bit more than is necessary, fairly logical
14-15 (2): Able to do math or solve logic puzzles mentally with reasonable accuracy
16-17 (3): Fairly intelligent, able to understand new tasks quickly
18-19 (4): Very intelligent, may invent new processes or uses for knowledge
20-21 (5): Highly knowledgeable, probably the smartest person many people know
22-23 (6): Able to make Holmesian leaps of logic
24-25 (7): Famous as a sage and genius
Wisdom
1 (–5): Seemingly incapable of thought, barely aware
2-3 (–4): Rarely notices important or prominent items, people, or occurrences
4-5 (–3): Seemingly incapable of forethought
6-7 (–2): Often fails to exert common sense
8-9 (–1): Forgets or opts not to consider options before taking action
10-11 (0): Makes reasoned decisions most of the time
12-13 (1): Able to tell when a person is upset
14-15 (2): Can get hunches about a situation that doesn’t feel right
16-17 (3): Reads people and situations fairly well
18-19 (4): Often used as a source of wisdom or decider of actions
20-21 (5): Reads people and situations very well, almost unconsciously
22-23 (6): Can tell minute differences among many situations
24-25 (7): Nearly prescient, able to reason far beyond logic
Charisma
1 (–5): Barely conscious, probably acts heavily autistic
2-3 (–4): Minimal independent thought, relies heavily on others to think instead
4-5 (–3): Has trouble thinking of others as people
6-7 (–2): Terribly reticent, uninteresting, or rude
8-9 (–1): Something of a bore or makes people mildly uncomfortable
10-11 (0): Capable of polite conversation
12-13 (1): Mildly interesting, knows what to say to the right people
14-15 (2): Interesting, knows what to say to most people
16-17 (3): Popular, receives greetings and conversations on the street
18-19 (4): Immediately likeable by many people, subject of favorable talk
20-21 (5): Life of the party, able to keep people entertained for hours
22-23 (6): Immediately likeable by almost everybody
24-25 (7): Renowned for wit, personality, and/or looks

case "2_intelligence":
		switch value {
		case 1:
			return ""
		case 2, 3:
			return ""
		case 4, 5:
			return ""
		case 6, 7:
			return ""
		case 8, 9:
			return ""
		case 10, 11:
			return ""
		case 12, 13:
			return ""
		case 14, 15:
			return ""
		case 16, 17:
			return ""
		case 18, 19:
			return ""
		case 20, 21:
			return ""
		case 22, 23:
			return ""
		case 24, 25:
			return ""
		}

*/
