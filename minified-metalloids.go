package main

import (
	"fmt"
)

const avogrado, grams float64 = 6.0221413e+23, 100.0	// Multiple same type one liner declaration

// Creates a struct type named metalloid. Struct types sort encompasses other primitive and derived types
type metalloid struct {
	name	string
	number	int32
	mass	float64
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

func atoms(mass float64) (weight float64) {
	weight = (mass/grams)*avogrado
	return 
}

// a function to print column headers that aligns with the values of the output. This particular function takes no argument but returns a value of string type.
func headers() string {
	return fmt.Sprintf("%-15s %-15s %-10s Atoms in %.2f Grams\n", "Element", "Number", "AMU", grams)
}

/*
// This block uses a function to produce the items output. Notice the subtle difference with its method version
func mString (m metalloid) string {
	return fmt.Sprintf("%-10s	%-10d	%-10.3f %e", m.name, m.number, m.mass, atoms(m.mass))
}

// now we make sense of all the variables, functions, methods, types etc to run our program. Remember the main function is the entry point of our program.
func main() {
	fmt.Print(headers())
	for _,m	:=	range metalloids {
		fmt.Println(mString(m))
	}	
}
*/

func (m metalloid) String() string {
	return fmt.Sprintf("%-10s	%-10d	%-10.3f %e", m.name, m.number, m.mass, atoms(m.mass))
}

// now we make sense of all the variables, functions, methods, types etc to run our program. Remember the main function is the entry point of our program.
func main() {
	fmt.Print(headers())
	for _,m	:=	range metalloids {
		fmt.Println(m)
	}	
}