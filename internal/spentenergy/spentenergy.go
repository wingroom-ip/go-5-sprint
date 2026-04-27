package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// validateInputs general function for validating entered data
func validateInputs(steps int, weight, height float64, duration time.Duration) error {
	if steps <= 0 {
		return errors.New("the number of steps must be greater than 0")
	}
	if weight <= 0 {
		return errors.New("weight must be greater than 0")
	}
	if height <= 0 {
		return errors.New("height must be greater than 0")
	}
	if duration <= 0 {
		return errors.New("time must be greater than 0")
	}
	return nil
}

// WalkingSpentCalories calculates the number of calories burned while walking
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// performs checks using validateInputs()
	if err := validateInputs(steps, weight, height, duration); err != nil {
		return 0, err
	}

	mSpeed := MeanSpeed(steps, height, duration)

	durationInMinutes := duration.Minutes()

	calories := ((weight * mSpeed * durationInMinutes) / minInH) * walkingCaloriesCoefficient

	return calories, nil
}

// RunningSpentCalories calculates the number of calories burned while running
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// performs checks using validateInputs()
	if err := validateInputs(steps, weight, height, duration); err != nil {
		return 0, err
	}
	mSpeed := MeanSpeed(steps, height, duration)

	durationInMinutes := duration.Minutes()

	calories := (weight * mSpeed * durationInMinutes) / minInH

	return calories, nil
}

// MeanSpeed calculates average speed
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps == 0 {
		return 0
	}
	if duration <= 0 {
		return 0
	}
	distKm := Distance(steps, height)

	hours := duration.Hours()

	return distKm / hours
}

// Distance calculates the distance in km
func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	return (float64(steps) * stepLength) / mInKm
}
