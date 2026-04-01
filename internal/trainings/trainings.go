package trainings

import (
	"fmt"
	personaldata "sprint-5/internal/personaldata"
	spentenergy "sprint-5/internal/spentenergy"
	"strconv"
	"strings"
	"time"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	strSplit := strings.Split(datastring, ",")
	if len(strSplit) != 3 {
		return fmt.Errorf(`длина слайса была равна 3, так как в строке данных у нас количество шагов, вид активности и продолжительность.`)
	}
	countSteps, err := strconv.Atoi(strSplit[0])
	if err != nil {
		return err
	}

	if countSteps <= 0 {
		return fmt.Errorf("кол-во шагов меньше или равно нулю.")
	}

	t.Steps = countSteps

	category := strSplit[1]
	t.TrainingType = category

	tm, err := time.ParseDuration(strSplit[2])
	if err != nil {
		return err
	}

	if tm <= 0 {
		return fmt.Errorf("продолжительность равна нулю или отрицательное значение.")
	}
	t.Duration = tm

	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	if t.Steps <= 0 {
		return "", fmt.Errorf("кол-во шагов меньше или равно нулю")
	}

	if t.Height <= 0 {
		return "", fmt.Errorf("рост меньше или равен нулю")
	}

	if t.Weight <= 0 {
		return "", fmt.Errorf("вес меньше или равен нулю")
	}

	if t.Duration <= 0 {
		return "", fmt.Errorf("продолжительность меньше или равна нулю")
	}

	distance := spentenergy.Distance(t.Steps, float64(t.Height))
	if distance <= 0 {
		return "", fmt.Errorf("отрицательная или нулевая дистанция")
	}

	averageSpeed := spentenergy.MeanSpeed(t.Steps, float64(t.Height), t.Duration)
	if averageSpeed <= 0 {
		return "", fmt.Errorf("отрицательная или нулевая средняя скорость")
	}

	category := t.TrainingType
	walkingSpentCalories, err := spentenergy.WalkingSpentCalories(t.Steps, float64(t.Weight), float64(t.Height), t.Duration)
	if err != nil {
		return "", err
	}

	runningSpentCalories, err := spentenergy.RunningSpentCalories(t.Steps, float64(t.Weight), float64(t.Height), t.Duration)
	if err != nil {
		return "", err
	}

	switch category {
	case "Ходьба":
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", category, t.Duration.Hours(), distance, averageSpeed, walkingSpentCalories), nil
	case "Бег":
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", category, t.Duration.Hours(), distance, averageSpeed, runningSpentCalories), nil
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
}
