package Bridage

import "testing"

func TestRectangle_Draw(t *testing.T) {
	redRectangle := Rectangle{}
	redRectangle.Constructor(1111, 222, &RedRectangle{})
	redRectangle.Do()

	blueRectangle := Rectangle{}
	blueRectangle.Constructor(3333, 3444, &BlueRectangle{})

	blueRectangle.Do()
}
