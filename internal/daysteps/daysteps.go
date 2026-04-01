package daysteps

import (
	"fmt"
	"log"
	personaldata "sprint-5/internal/personaldata"
	spentenergy "sprint-5/internal/spentenergy"
	"strconv"
	"strings"
	"time"
)

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	strSplit := strings.Split(datastring, ",")
	if len(strSplit) != 2 {
		log.Println("длина слайса была равна 2, так как в строке данных у нас количество шагов и продолжительность.")
		return fmt.Errorf("длина слайса была равна 2, так как в строке данных у нас количество шагов и продолжительность.")
	}

	countSteps, err := strconv.Atoi(strSplit[0])
	if err != nil {
		return err
	}

	if countSteps <= 0 {
		log.Println("кол-во шагов должно быть больше нуля.")
		return fmt.Errorf("кол-во шагов должно быть больше нуля.")
	}

	ds.Steps = countSteps

	t, err := time.ParseDuration(strSplit[1])
	if err != nil {
		return err
	}

	if t <= 0 {
		log.Println("продолжительность равна нулю или отрицательная.")
		return fmt.Errorf("продолжительность равна нулю или отрицательная.")
	}

	ds.Duration = t

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию

	if ds.Steps <= 0 || ds.Duration <= 0 || ds.Weight <= 0 || ds.Height <= 0 {
		fmt.Println("кол-во шагов или продолжительность или вес или высота равны или меньше нуля")
		return "", fmt.Errorf("кол-во шагов или продолжительность или вес или высота равны или меньше нуля")
	}

	distance := spentenergy.Distance(ds.Steps, float64(ds.Height))
	if distance <= 0 {
		return "", fmt.Errorf("отрицательная или нулевая дистанция")
	}
	walkingSpentCalories, err := spentenergy.WalkingSpentCalories(ds.Steps, float64(ds.Weight), float64(ds.Height), ds.Duration)
	if err != nil {
		return "", err
	}

	if walkingSpentCalories <= 0 {
		return "", fmt.Errorf("потраченные калл. равны или меньше нуля")
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, walkingSpentCalories), nil
}
