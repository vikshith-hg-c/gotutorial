package creational

type BuildProcess interface {
	SetWeels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehical() VehicalProduct
}

type Director struct {
	builder BuildProcess
}

func (f *Director) SetBuilder(b BuildProcess) {
	f.builder = b
}

func (f *Director) Construct() {
	f.builder.SetSeats().SetWeels().SetStructure()
}

type VehicalProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

type carBuilder struct {
	v VehicalProduct
}

func (c *carBuilder) SetWeels() BuildProcess {
	c.v.Wheels = 4
	return c
}

func (c *carBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

func (c *carBuilder) SetStructure() BuildProcess {
	c.v.Structure = "monocock"
	return c
}

func (c *carBuilder) GetVehical() VehicalProduct {
	return c.v
}

type bikeBuilder struct {
	v VehicalProduct
}

func (b *bikeBuilder) SetWeels() BuildProcess {
	b.v.Wheels = 2
	return b
}

func (b *bikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 2
	return b
}

func (b *bikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "rod"
	return b
}

func (b *bikeBuilder) GetVehical() VehicalProduct {
	return b.v
}
