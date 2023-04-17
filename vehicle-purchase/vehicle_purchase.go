package purchase

import "strings"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	const recommandation string = "is clearly the better choice."
	if option1 < option2 {
		return strings.Join([]string{option1, recommandation}, " ")
	}
	return strings.Join([]string{option2, recommandation}, " ")
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	switch {
	case age < 3:
		return originalPrice * 0.80
	case age >= 10:
		return originalPrice / 2
	default:
		return originalPrice * 0.70
	}
}
