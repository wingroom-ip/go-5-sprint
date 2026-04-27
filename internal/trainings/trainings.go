package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse the method parses the data string and
// writes the data into the corresponding fields
// of the Training structure
func (t *Training) Parse(dataString string) (err error) {

	parseData := strings.Split(dataString, ",")
	if len(parseData) != 3 {
		return errors.New("insufficient data: 3 values ​​expected")
	}

	// getting the number of steps int
	steps, err := strconv.Atoi(parseData[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("The number of steps must be greater than 0")
	}
	t.Steps = steps

	// getting the type of train string
	trainType := parseData[1]
	if trainType == "" {
		return errors.New("workout type not specified")
	}
	t.TrainingType = trainType

	// getting duration of train time.Duration
	duration, err := time.ParseDuration(parseData[2])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("duration must be greater than 0")
	}
	t.Duration = duration

	return nil
}

// ActionInfo the method calculates and displays
// detailed information about training
func (t Training) ActionInfo() (string, error) {

	dist := spentenergy.Distance(t.Steps, t.Height)

	mSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	switch t.TrainingType {
	case "Бег":
		calories, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			t.TrainingType, t.Duration.Hours(), dist, mSpeed, calories), nil
	case "Ходьба":
		calories, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			t.TrainingType, t.Duration.Hours(), dist, mSpeed, calories), nil
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
}
