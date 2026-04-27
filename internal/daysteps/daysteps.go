package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse the method parses the data string and
// writes the data into the corresponding fields
// of the DaySteps structure
func (ds *DaySteps) Parse(dataString string) (err error) {

	parseData := strings.Split(dataString, ",")
	if len(parseData) != 2 {
		return errors.New("insufficient data: 2 values ​​expected")
	}

	// getting the number of steps int
	steps, err := strconv.Atoi(strings.TrimSpace(parseData[0]))
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("The number of steps must be greater than 0")
	}
	ds.Steps = steps

	// getting duration time.Duration
	duration, err := time.ParseDuration(strings.TrimSpace(parseData[1]))
	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("duration must be greater than 0")
	}
	ds.Duration = duration

	return nil
}

// ActionInfo the method calculates and displays
// detailed information about day walking
func (ds DaySteps) ActionInfo() (string, error) {

	dist := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, dist, calories), nil
}
