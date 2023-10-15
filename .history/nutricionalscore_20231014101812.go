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
