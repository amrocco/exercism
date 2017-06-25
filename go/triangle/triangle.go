package triangle

import "math"

const testVersion = 3

//NaT = not a triangle
const NaT = 0

//Equ = equilateral triangle
const Equ = 1

//Iso = isosceles triangle
const Iso = 2

//Sca = scalene triangle
const Sca = 3

//Deg = degenerate triangle
const Deg = 4

//Kind refers to a kind of triangle
type Kind int

//KindFromSides determines what kind of triangle the three given sides make
func KindFromSides(a, b, c float64) Kind {
	if isTriangle(a, b, c) == false {
		return NaT
	}

	if isEquilateral(a, b, c) {
		return Equ
	} else if isScalene(a, b, c) {
		return Sca
	} else if isDegenerate(a, b, c) {
		return Deg
	}

	return Iso
}

func isTriangle(a, b, c float64) bool {
	if !math.IsInf(a+b+c, 0) &&
		!triangleInequality(a, b, c) &&
		allSidesGreaterThanZero(a, b, c) {
		return true
	}

	return false
}

func allSidesGreaterThanZero(a, b, c float64) bool {
	return a > 0 && b > 0 && c > 0
}

func triangleInequality(a, b, c float64) bool {
	return (a+b) < c || (a+c) < b || (b+c) < a
}

func isEquilateral(a, b, c float64) bool {
	return a == b && a == c
}

func isScalene(a, b, c float64) bool {
	//Scalene triangles must also have three unequal angles
	//therefore the degenerate triangle (1, 4, 3) is not also scalene
	return a != b && a != c && b != c && !isDegenerate(a, b, c)
}

func isDegenerate(a, b, c float64) bool {
	return (a+b) == c || (a+c) == b || (b+c) == a
}
