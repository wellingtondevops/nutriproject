package main

type ScoreType int

const (
	Food ScoreType = iota
	Beverage
	Water
	Cheese
)

type NutricionalScore struct {
	Value     int
	Positive  int
	Negative  int
	ScoreType ScoreType
}
type EnergyKJ float64

type SugerGram float64

type SaturadedFattyAcids float64

type SodiumMilligram float64

type FruitsPercent float64

type FibreGram float64

type ProteinGram float64

type NutriciolData struct {
	Energy              EnergyKJ
	Sugars              SugerGram
	SaturadedFattyAcids SaturadedFattyAcids
	Sodium              SodiumMilligram
	Fruits              FruitsPercent
	Fribe               FibreGram
	Protein             ProteinGram
	IsWater             bool
}

func GetNutritionalScore(n NutriciolData, st ScoreType) NutricionalScore {

	value := 0
	positive := 0
	negative := 0

	if st != water {
		fruitPoints := n.Fruits.GetPoints()
		fibrePoints := n.Fribe.GetPoints()

		negative = n.Energy.GetPoints() + n.Sugars.GetPoints + n.SaturadedFattyAcids.GetPoints() + n.Sodium.GetPoints()
		positive = fruitPoints + fibrePoints + n.Protein.GetPoints()
	}

	return NutricionalScore{
		Value:     value,
		Positve:   positive,
		Negative:  negative,
		ScoreType: st,
	}

}
