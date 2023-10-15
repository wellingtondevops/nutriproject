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

func (e EnergyKJ) GetPoints(st ScoreType) int {

}

func (s SugerGram) GetPoints(st ScoreType) int {

}

func (sfa SaturadedFattyAcids) GetPoints(st ScoreType) int {

}

func (s SodiumMilligram) GetPoints(st ScoreType) int {

}

func (f FruitsPercent) GetPoints(st ScoreType) int {

}

func (f FibreGram) GetPoints(st ScoreType) int {

}

func (p ProteinGram) GetPoints(st ScoreType) int {

}
func GetNutritionalScore(n NutriciolData, st ScoreType) NutricionalScore {

	value := 0
	positive := 0
	negative := 0

	if st != water {
		fruitPoints := n.Fruits.GetPoints(st)
		fibrePoints := n.Fribe.GetPoints(st)

		negative = n.Energy.GetPoints(st) + n.Sugars.GetPoints(st) + n.SaturadedFattyAcids.GetPoints(st) + n.Sodium.GetPoints(st)
		positive = fruitPoints + fibrePoints + n.Protein.GetPoints(st)
	}

	return NutricionalScore{
		Value:     value,
		Positive:  positive,
		Negative:  negative,
		ScoreType: st,
	}

}
