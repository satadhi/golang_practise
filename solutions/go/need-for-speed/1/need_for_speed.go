package speed

// Car represents a remote controlled car
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery
func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
		distance:     0,
	}
}

// Track represents the race track
type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time
func Drive(car Car) Car {
	if car.battery < car.batteryDrain {
		return car
	}

	car.distance += car.speed
	car.battery -= car.batteryDrain

	return car
}

// CanFinish checks if a car can finish the track
func CanFinish(car Car, track Track) bool {
	maxMoves := car.battery / car.batteryDrain
	maxDistance := maxMoves * car.speed

	return maxDistance >= track.distance
}