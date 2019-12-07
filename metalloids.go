package main

import (
	"fmt"
)

const avogrado float64 = 6.0221413e+23	// Declares a typed constant of float64 type
const grams	= 100.0		// Declares an untyped constant. However unspecified types are float64 by default

type amu float64	// Creates a new type named amu based off of float64

// Creates a struct type named metalloid. Struct types sort encompasses other primitive and derived types
type metalloid struct {
	name	string
	number	int32
	weight	amu
}

// create a variable of a slice of the metalloid struct type
var metalloids = []metalloid{
	metalloid{ "Boron", 5, 10.81},		// remember to add the comma after each struct initialization
	metalloid{ "Silicon", 14, 28.085},
	metalloid{ "Germanium", 32, 74.63},
	metalloid{ "Arsenic", 33, 74.921},
	metalloid{ "Antimony", 51, 121.760},
	metalloid{ "Tellerium", 52, 127.60},
	metalloid{ "Polonium", 84, 209.0},
}

// create a method named float of the amu type whose logic/operation returns a float64
func (mass amu) float() float64 {
	return float64(mass)
}

// create a function calcuates the number of moles. Notice the subtle differences in defining a method and function
func moles(mass amu) float64 {
	return float64(mass)/grams
}

// a function to return the number of atoms moles
func atoms(moles float64)	float64 {
	return moles*avogrado
}

// a function to print column headers that aligns with the values of the output. This particular function takes no argument but returns a value of string type.
func headers() string {
	return fmt.Sprintf("%-15s %-15s %-15s Atoms in %.2f Grams\n", "Element", "Number", "AMU", grams)
}

// now we make sense of all the variables, functions, methods, types etc to run our program. Remember the main function is the entry point of our program.
func main() {
	fmt.Print(headers())
	for _,m	:=	range metalloids {
		fmt.Printf("%-10s	%-10d	%-10.3f	%e\n", m.name, m.number, m.weight.float(), atoms(moles(m.weight)))
	}	
}