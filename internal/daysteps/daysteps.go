package daysteps

import (
	"fmt"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"math"
	"strconv"
	"strings"
	"time"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(input string) error {
	parts := strings.Split(input, ",")
	if len(parts) != 2 {
		return fmt.Errorf("invalid input format")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("invalid steps format: %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("steps must be positive")
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("invalid duration format: %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("duration must be positive")
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Steps <= 0 {
		return "", fmt.Errorf("steps must be positive")
	}
	if ds.Duration <= 0 {
		return "", fmt.Errorf("duration must be positive")
	}
	if ds.Weight <= 0 {
		return "", fmt.Errorf("weight must be positive")
	}
	if ds.Height <= 0 {
		return "", fmt.Errorf("height must be positive")
	}

	distance := calculateDistance(ds.Steps, ds.Height)
	calories := calculateCalories(ds.Steps, ds.Weight, ds.Duration)

	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories,
	), nil
}

func calculateDistance(steps int, height float64) float64 {
	stepLength := 0.7 * height
	distance := float64(steps) * stepLength / 1000 * 0.6428
	return math.Round(distance*100) / 100
}

func calculateCalories(steps int, weight float64, duration time.Duration) float64 {
	calories := float64(steps) * weight * 0.00039376 / duration.Hours()
	return math.Round(calories*100) / 100
}

func (ds *DaySteps) Print() {
	// Реализация метода Print
}
