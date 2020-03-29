package main

import "fmt"

type army interface {
	name()	string
	power()	int
}

type nationalArmy struct {
	armyName	string
	unitPower	int
	soldiers	int
}

type heroArmy struct {
	heroName	string
	heroPower	int
	sideKicksPower	int
}

type chuckArmy struct {

}

func (c chuckArmy) name() string {
	return "The greatest Chuck Norris"
}

func (c chuckArmy) power() int {
	return 9999999
}

func (n nationalArmy) name() string{
	return n.armyName
}

func (n nationalArmy) power() int {
	return n.unitPower * n.soldiers
}

func (h heroArmy) name() string {
	return fmt.Sprintf("Army with %s", h.heroName)
}

func (h heroArmy) power() int {
	return h.heroPower + h.sideKicksPower
}

func fightAgainstThanos(armies []army) {
	thanosLife := 9999999

	for _, army := range armies	{
		fmt.Printf("%s fight against Thanos with a force of %d\n", army.name(), army.power())
		thanosLife -= army.power()
	}
	if thanosLife <= 0 {
		fmt.Println("The terrible Thanos has been defeated")
	} else {
		fmt.Printf("Thanos destroy entire universe, and his life is %d yet\n", thanosLife)
	}
}

func main() {
	army1 := nationalArmy{
		armyName:  "Air Forces",
		unitPower: 10,
		soldiers:  1000,
	}
	army2 := nationalArmy{
		armyName:  "Marines",
		unitPower: 20,
		soldiers:  1000,
	}
	army3 := nationalArmy{
		armyName:  "Space Forces",
		unitPower: 40,
		soldiers:  1000,
	}
	army4 := heroArmy{
		heroName:       "Iron Man",
		heroPower:      100000,
		sideKicksPower: 50000,
	}
	army5 := heroArmy{
		heroName:       "Black Panther",
		heroPower:      50000,
		sideKicksPower: 100000,
	}

	chuck := chuckArmy{}

	armies := []army{army1, army2, army3, army4, army5, chuck}

	fightAgainstThanos(armies)
}
