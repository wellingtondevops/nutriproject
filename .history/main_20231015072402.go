package main

import "fmt"

func main() {
	ns := GetNutritionalScore(NutriciolData{
		Energy:              EnergyFromKcal(0),
		Sugars:              SugarGram(10),
		SaturadedFattyAcids: SaturadedFattyAcids(2),
		Sodium:              SodiumMilligram(500),
		Fruits:              FruitsPercent(60),
		Fibre:               FibreGram(4),
		Protein:             ProteinGram(2),
	}, Food)

	fmt.Printf("Nutricional Score: %d\n", ns.Value)
	fmt.Printf("NutriScore: %s\n", ns.GetNutriScore())
}
