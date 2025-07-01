package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	personaldata "github.com/Yandex-Practicum/tracker/internal/personaldata"
	spentenergy "github.com/Yandex-Practicum/tracker/internal/spentenergy"
	usererrors "github.com/Yandex-Practicum/tracker/internal/usererrors"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	comp := strings.Split(datastring, ",")
	if len(comp) != 2 {
		return usererrors.ErrIncorrectNumberOfArgs
	}
	steps, err := strconv.Atoi(comp[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return usererrors.ErrStepsNumber
	}
	ds.Steps = steps
	duration, err := time.ParseDuration(comp[1])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return usererrors.ErrDuration
	}
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance, _, calories, err := spentenergy.CalcCal(
		"Ходьба",
		ds.Steps,
		ds.Weight,
		ds.Height,
		ds.Duration,
	)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(`Количество шагов: %d.
Дистанция составила %.2f км.
Вы сожгли %.2f ккал.
`, ds.Steps, distance, calories), nil
}
