package main

import "sort"

func (p *Person) Describe() []string {
	result := make([]string, 0)
	keys := sortedKeys(p.Abilities)

	for _, v := range keys {
		result = append(result, getAbilityDescription(v, p.Abilities[v]))
	}

	return result
}

type sortedMap struct {
	m map[string]float64
	s []string
}

func (sm *sortedMap) Len() int {
	return len(sm.m)
}

func (sm *sortedMap) Less(i, j int) bool {
	return sm.m[sm.s[i]] > sm.m[sm.s[j]]
}

func (sm *sortedMap) Swap(i, j int) {
	sm.s[i], sm.s[j] = sm.s[j], sm.s[i]
}

func sortedKeys(m map[string]float64) []string {
	sm := new(sortedMap)
	sm.m = m
	sm.s = make([]string, len(m))
	i := 0
	for key, _ := range m {
		sm.s[i] = key
		i++
	}
	sort.Sort(sm)
	return sm.s
}

func getAbilityDescription(ability string, value float64) string {
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
