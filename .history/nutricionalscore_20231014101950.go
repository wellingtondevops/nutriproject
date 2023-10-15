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
