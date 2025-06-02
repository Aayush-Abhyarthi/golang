package newpackage

type Car struct {
	Vehicle
	ms string
}

func (c *Car) GetMs() string {
	return c.ms
}
func (c *Car) SetMs(ms string) {
	c.ms = ms
}
