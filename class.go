package main

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
