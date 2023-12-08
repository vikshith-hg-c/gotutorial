package creational

import "testing"

func TestBuilderPattern(t *testing.T) {
	manifacture := Director{}

	carBuilder := &carBuilder{}
	bikeBuilder := &bikeBuilder{}
	manifacture.SetBuilder(bikeBuilder)
	manifacture.Construct()
	manifacture.SetBuilder(carBuilder)
	manifacture.Construct()

	car := bikeBuilder.GetVehical()
	bike := carBuilder.GetVehical()

	if car.Wheels != 4 {
		t.Errorf("Wheels of car = 4 but git %d\n", car.Wheels)
	}

	if car.Seats != 5 {
		t.Errorf("Seats of car = 5 but git %d\n", car.Seats)
	}

	if car.Structure != "monocock" {
		t.Errorf("Structure of car = monocock but git %s\n", car.Structure)
	}

	if bike.Wheels != 2 {
		t.Errorf("Wheels of car = 4 but git %d\n", bike.Wheels)
	}

	if bike.Seats != 2 {
		t.Errorf("Seats of car = 5 but git %d\n", bike.Seats)
	}

	if bike.Structure != "rod" {
		t.Errorf("Structure of car = monocock but git %s\n", bike.Structure)
	}
}
