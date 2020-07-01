package main

import (
	"fmt"
	"math"
)

type Position struct {
	x float64
	y float64
	z float64
}
func (vec Position) display() {
	fmt.Print(vec.x , vec.y, vec.z)
}

// origin to point a
type PositionVector struct {
	a Position
}

//TODO: find magnitude of a vector
// Magnitude : Sqrt( x*x + y*y + z*z) where x,y and z are position cordinate
func (vec1 PositionVector) magnitude() float64 {
	return  math.Sqrt( math.Pow(vec1.a.x,2) + math.Pow(vec1.a.y,2)+ math.Pow(vec1.a.z,2) )
}

//TODO:
// Equal
func (vec1 PositionVector) Equal(vec2 PositionVector) bool {
	return 	vec1.a.x == vec2.a.x && vec1.a.y == vec2.a.y && vec1.a.z == vec2.a.z
}

//TODO:
// Addition of two vectors
func (vec1 PositionVector) Add(vec2 PositionVector) PositionVector {
	return  PositionVector{
		a: Position{
			vec1.a.x + vec2.a.x,
			 vec1.a.y + vec2.a.y,
			vec1.a.z + vec2.a.z,
		},
	}
}


func main() {
	vec1 := PositionVector{
		Position{
			4.0,3.0,0.0,
		},
	}

	vec2 := PositionVector{
		Position{
			4.0,3.0,0.20,
		},
	}

	fmt.Println(vec2.Equal(vec1))
	fmt.Println(vec2.Add(vec1))
	magnitudePosvec := vec1.magnitude()
	fmt.Println(magnitudePosvec)
}
