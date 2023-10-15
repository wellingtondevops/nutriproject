package main

import "fmt"

func main() {
	ns := GetNutritionalScore(NutriciolData{
		Energy:              EnergyFromKcal(),
		Sugars:              SugarGram(),
		SaturatedFattyAcids: SaturadedFattyAcids(),
		Sodium:              SodiumMilligram(),
		Fruits:              FruitsPercent(),
		Fibre:               FibreGram(),
		Protein:             ProteinGram(),
	}, Food)

	fmt.Printf("Nutricional Score: %d\n", ns.Value)
}
