package spentenergy

import (
	"time"

	usererrors "github.com/Yandex-Practicum/tracker/internal/usererrors"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	cal, err := spentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}
	return cal * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	cal, err := spentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}
	return cal, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if steps <= 0 || duration <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	return distance / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	return float64(steps) * height * stepLengthCoefficient / mInKm
}

func spentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, usererrors.ErrStepsNumber
	}
	if weight <= 0 {
		return 0, usererrors.ErrWeight
	}
	if height <= 0 {
		return 0, usererrors.ErrHeight
	}
	if duration <= 0 {
		return 0, usererrors.ErrDuration
	}
	speed := MeanSpeed(steps, height, duration)
	return (weight * speed * duration.Minutes()) / minInH, nil
}

func CalcCal(activity string, steps int, weight, height float64, duration time.Duration) (distance, speed, calories float64, err error) {
	distance = Distance(steps, height)
	speed = MeanSpeed(steps, height, duration)

	switch activity {
	case "Бег":
		calories, err = RunningSpentCalories(steps, weight, height, duration)
	case "Ходьба":
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
	default:
		err = usererrors.ErrActivity
	}

	return
}
