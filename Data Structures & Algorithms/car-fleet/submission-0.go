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

	stack := make([]float64, 0, len(position))

	for _, car := range cars {
		arrivalTime := float64(target-car.position)/ float64(car.speed)

		if len(stack) == 0 {
			stack = append(stack, arrivalTime)
		 } else {	
			top := stack[len(stack)-1]
			if arrivalTime > top {
				stack = append(stack, arrivalTime)
			}
		}
	}

	return len(stack)

}
