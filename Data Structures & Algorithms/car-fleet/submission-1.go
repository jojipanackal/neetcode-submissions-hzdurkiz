import "slices"

func carFleet(target int, position []int, speed []int) int {

	type Car struct {
		speed int
		position int
	}

	cars := make([]Car, len(position))

	for i := range position {
		cars[i] = Car {
			speed: speed[i],
			position: position[i],
		}
	}

	slices.SortFunc(cars, func(a, b Car) int {
		if a.position > b.position {
			return -1
		}
		if a.position < b.position {
			return 1
		}
		return 0
	})

	fleets := 0
	lastArrivalTime := float64(0)

	for _, car := range cars {
		arrivalTime := float64(target-car.position)/ float64(car.speed)
		
		if arrivalTime > lastArrivalTime {
			fleets++
			lastArrivalTime = arrivalTime
		}
		
	}

	return fleets

}
