package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	personaldata "github.com/Yandex-Practicum/tracker/internal/personaldata"
	spentenergy "github.com/Yandex-Practicum/tracker/internal/spentenergy"
	usererrors "github.com/Yandex-Practicum/tracker/internal/usererrors"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	components := strings.Split(datastring, ",")
	if len(components) != 3 {
		return usererrors.ErrIncorrectNumberOfArgs
	}

	steps, err := strconv.Atoi(components[0])
	if err != nil || steps <= 0 {
		return usererrors.ErrStepsNumber
	}

	t.Steps = steps
	t.TrainingType = components[1]

	duration, err := time.ParseDuration(components[2])
	if err != nil || duration <= 0 {
		return usererrors.ErrDuration
	}

	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance, meanSpeed, calories, err := spentenergy.CalcCal(
		t.TrainingType,
		t.Steps,
		t.Weight,
		t.Height,
		t.Duration,
	)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(`Тип тренировки: %s
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f
`, t.TrainingType, t.Duration.Hours(), distance, meanSpeed, calories), nil
}
