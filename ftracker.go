package ftracker

import (
	"fmt"
)

// distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
func distance(action int) float64 {
	const lenStep = 0.65 // средняя длина шага.
	const mInKm = 1000   // количество метров в километре.
	return float64(action) * lenStep / mInKm
}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
func meanSpeed(action int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	distance := distance(action)
	return distance / duration
}

// RunningSpentCalories возвращает количество потраченных калорий при беге.
func RunningSpentCalories(action int, weight, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	speed := meanSpeed(action, duration)
	return (18*speed + 1.79) * weight * duration
}

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
func WalkingSpentCalories(action int, duration, weight, height float64) float64 {
	if duration == 0 {
		return 0
	}
	speed := meanSpeed(action, duration)
	return (0.035*weight + (speed*speed)/(height*100)*0.029) * weight * duration
}

// swimmingMeanSpeed возвращает среднюю скорость при плавании.
func swimmingMeanSpeed(lengthPool, countPool int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	const mInKm = 1000 // количество метров в километре.
	return float64(lengthPool*countPool) / float64(mInKm) / duration
}

// SwimmingSpentCalories возвращает количество потраченных калорий при плавании.
func SwimmingSpentCalories(lengthPool, countPool int, duration, weight float64) float64 {
	if duration == 0 {
		return 0
	}
	speed := swimmingMeanSpeed(lengthPool, countPool, duration)
	return speed*2*weight + 1.1*weight*duration
}

// ShowTrainingInfo возвращает строку с информацией о тренировке.
func ShowTrainingInfo(action int, trainingType string, duration, weight, height float64, lengthPool, countPool int) string {
	switch trainingType {
	case "Бег":
		distance := distance(action)
		speed := meanSpeed(action, duration)
		calories := RunningSpentCalories(action, weight, duration)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, distance, speed, calories)
	case "Ходьба":
		distance := distance(action)
		speed := meanSpeed(action, duration)
		calories := WalkingSpentCalories(action, duration, weight, height)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, distance, speed, calories)
	case "Плавание":
		distance := swimmingMeanSpeed(lengthPool, countPool, duration)
		speed := distance / duration
		calories := SwimmingSpentCalories(lengthPool, countPool, duration, weight)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, distance, speed, calories)
	default:
		return "неизвестный тип тренировки"
	}
}
