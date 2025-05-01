package spentenergy

import (
	"fmt"
	"time"
)

const (
	mInKm                 = 1000.0
	minInH                = 60
	stepLengthCoefficient = 0.45
	walkingCaloriesCoeff  = 0.5
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("no steps")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("no duration")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("weight must be positive")
	}
	if height <= 0 {
		return 0, fmt.Errorf("height must be positive")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := ((weight * meanSpeed * durationInMinutes) / minInH) * walkingCaloriesCoeff

	// Округляем до 2 знаков после запятой
	return float64(int(calories*100)) / 100, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("no steps")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("no duration")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("weight must be positive")
	}
	if height <= 0 {
		return 0, fmt.Errorf("height must be positive")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * meanSpeed * durationInMinutes) / minInH

	// Округляем до 2 знаков после запятой
	return float64(int(calories*100)) / 100, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 || steps <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	return distance / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	if steps <= 0 {
		return 0
	}
	lengthStep := height * stepLengthCoefficient
	return (float64(steps) * lengthStep) / mInKm
}
