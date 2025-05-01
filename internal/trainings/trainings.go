package trainings

import (
	"fmt"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
	"strconv"
	"strings"
	"time"
)

type Training struct {
	Steps                 int           //количество шагов
	TrainingType          string        //тип тренировки
	Duration              time.Duration //длительность тренировки
	personaldata.Personal               //структура Personal из пакета personaldata
}

func (t *Training) Parse(datastring string) (err error) {
	dataSlise := strings.Split(datastring, ",")
	if len(dataSlise) != 3 {
		return fmt.Errorf("invalid data")
	}

	// Проверка шагов
	steps, err := strconv.Atoi(dataSlise[0])
	if err != nil {
		return fmt.Errorf("conversion steps error: %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("steps must be positive")
	}
	t.Steps = steps

	t.TrainingType = dataSlise[1]

	// Проверка продолжительности
	duration, err := time.ParseDuration(strings.Replace(dataSlise[2], "h", "h", 1))
	if err != nil {
		return fmt.Errorf("conversion time error: %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("duration must be positive")
	}
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	var distance, meanSpeed, calories float64
	var err error
	switch t.TrainingType {
	case "Ходьба":
		distance = spentenergy.Distance(t.Steps, t.Height)
		meanSpeed = spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Бег":
		distance = spentenergy.Distance(t.Steps, t.Height)
		meanSpeed = spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", t.TrainingType)
	}
	if err != nil {
		fmt.Println("Error:", err)
	}
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, meanSpeed, calories), nil
}

func (t *Training) Print() {

}
