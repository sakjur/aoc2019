package aoc2019

func Fuel(mass int) int {
	return mass/3-2
}

func FuelForFuel(mass int) int {
	tally := 0
	for weight := Fuel(mass); weight > 0; weight = Fuel(weight) {
		tally += weight
	}
	return tally
}
