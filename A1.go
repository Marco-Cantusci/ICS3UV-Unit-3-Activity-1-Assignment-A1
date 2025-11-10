/*
 * @author Marco Cantusci
 * @verion 1.0.0
 * @date 2025-11-10
 * @fileoverview Display and calculate the area of a circle with a radius of 5.6 cm.
 */

package main

import "fmt"

func main() {

	// Stores radius as 5.6
	var radius float64 = 5.6

	// Stores PI as 3.14
	const PI float64 = 3.14

	// formula/calculation of area
	var formula float64 = PI * radius * radius

	// Print radius
	fmt.Println("Radius:", radius, "cm")

	// Print area
	fmt.Println("The area of a circle with a radius of", radius, "cm is", formula, "cm²")

	fmt.Println("\nDone.")

}
