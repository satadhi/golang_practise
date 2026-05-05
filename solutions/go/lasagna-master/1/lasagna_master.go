package lasagnamaster

// PreparationTime calculates total prep time.
func PreparationTime(layers []string, prepTime int) int {
	if prepTime == 0 {
		prepTime = 2
	}
	return len(layers) * prepTime
}

// Quantities calculates required noodles and sauce.
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		}
		if layer == "sauce" {
			sauce += 0.2
		}
	}
	return
}

// AddSecretIngredient replaces last item of friend's list into yours.
func AddSecretIngredient(friendList, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

// ScaleRecipe scales quantities for desired portions.
func ScaleRecipe(quantities []float64, portions int) []float64 {
	result := make([]float64, len(quantities))
	scale := float64(portions) / 2.0

	for i, v := range quantities {
		result[i] = v * scale
	}
	return result
}