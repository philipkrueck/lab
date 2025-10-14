package main

import "fmt"

type Coordinate struct {
	Lat, Long float64
}

func main() {
	// `make` assignment
	cities := make(map[string]Coordinate)

	cities["london"] = Coordinate{51.5098, 0.1180}
	cities["new york"] = Coordinate{51.5098, 0.1180}
	fmt.Println("cities:", cities)

	// literal assignment
	countries := map[string]string{
		"germany":  "berlin",
		"portugal": "lisbon",
		"denmark":  "copenhagen",
		"norway":   "oslo",
	}
	fmt.Println("countries:", countries)

	// mutating map
	countries["argentina"] = "buenos aires"
	delete(countries, "denmark")
	fmt.Println("countries:", countries)

	// testing if a value exists
	country := "portugal"
	capital, countryExists := countries[country]
	if countryExists {
		fmt.Println("The capital of", country, "is", capital)
	} else {
		fmt.Println(country, " doesn't exist in map, so the retrieved value will be the zero value for string:", capital)
	}
}
