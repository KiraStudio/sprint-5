package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("кол-во шагов меньше или равно нулю.")
	}

	if weight <= 0 {
		return 0, fmt.Errorf("вес не может быть нулевым или отрицательным.")
	}

	if height <= 0 {
		return 0, fmt.Errorf("рост не может быть нулевым или отрицательным.")
	}

	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность не может быть нулевая или отрицательная.")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := int(duration / time.Minute)

	walkingSpentCalories := ((weight * meanSpeed * float64(durationInMinutes)) / minInH) * walkingCaloriesCoefficient
	if walkingSpentCalories < 0 {
		return 0, fmt.Errorf("Ошибка при вычислении каллорий")
	}

	return walkingSpentCalories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("кол-во шагов меньше или равно нулю.")
	}

	if weight <= 0 {
		return 0, fmt.Errorf("вес не может быть нулевым или отрицательным.")
	}

	if height <= 0 {
		return 0, fmt.Errorf("рост не может быть нулевым или отрицательным.")
	}

	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность не может быть нулевая или отрицательная.")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := int(duration / time.Minute)

	return (weight * meanSpeed * float64(durationInMinutes)) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || height <= 0 || duration <= 0 {
		return 0
	}
	var hours float64
	destination := Distance(steps, height)
	hours = duration.Hours()
	averageSpeed := destination / hours

	return averageSpeed
}

func Distance(steps int, height float64) float64 {
	if steps <= 0 || height <= 0 {
		return 0
	}

	stepLength := stepLengthCoefficient * height
	destination := stepLength * float64(steps)
	destination /= mInKm

	return destination
}
