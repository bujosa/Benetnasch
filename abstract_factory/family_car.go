package abstract_factory

type FamilyCar struct{}

// Return door count
func (f *FamilyCar) GetDoors() int {
	return 5
}

// Return wheel count
func (f *FamilyCar) GetWheels() int {
	return 4
}

func (f *FamilyCar) GetSeats() int {
	return 5
}