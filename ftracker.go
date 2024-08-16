package ftracker

import (
	"fmt"
	"math"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep   = 0.65  // средняя длина шага в метрах.
	mInKm     = 1000  // количество метров в километре.
	minInH    = 60    // количество минут в часе.
	kmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с.
	cmInM     = 100   // количество сантиметров в метре.
)

// distance возвращает дистанцию (в километрах), которую преодолел пользователь за время тренировки.
//
// Параметры:
//
// action int — количество совершенных действий (число шагов при ходьбе и беге, либо гребков при плавании).
func distance(action int) float64 {
	return float64(action) * lenStep / mInKm
}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
//
// Параметры:
//
// action int — количество совершенных действий (число шагов при ходьбе и беге, либо гребков при плавании).
// duration float64 — длительность тренировки в часах.
func meanSpeed(action int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	dist := distance(action)
	return dist / duration
}

// swimmingMeanSpeed возвращает среднюю скорость при плавании.
//
// Параметры:
//
// lengthPool int — длина бассейна в метрах.
// countPool int — сколько раз пользователь переплыл бассейн.
// duration float64 — длительность тренировки в часах.
func swimmingMeanSpeed(lengthPool, countPool int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	return float64(lengthPool) * float64(countPool) / mInKm / duration
}

// ShowTrainingInfo возвращает строку с информацией о тренировке.
//
// Параметры:
//
// action int — количество совершенных действий (число шагов при ходьбе и беге, либо гребков при плавании).
// trainingType string — вид тренировки (Бег, Ходьба, Плавание).
// duration float64 — длительность тренировки в часах.
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
// lengthPool int — длина бассейна в метрах.
// countPool int — сколько раз пользователь переплыл бассейн.
func ShowTrainingInfo(action int, trainingType string, duration, weight, height float64, lengthPool, countPool int) string {
	var dist, speed, calories float64

	switch trainingType {
	case "Бег":
		dist = distance(action) * mInKm                 // Переводим расстояние в метры
		speed = meanSpeed(action, duration) * kmhInMsec // Переводим среднюю скорость в м/с
		calories = RunningSpentCalories(action, weight, duration)
	case "Ходьба":
		dist = distance(action) * mInKm                 // Переводим расстояние в метры
		speed = meanSpeed(action, duration) * kmhInMsec // Переводим среднюю скорость в м/с
		calories = WalkingSpentCalories(action, duration, weight, height)
	case "Плавание":
		dist = swimmingMeanSpeed(lengthPool, countPool, duration) * duration * mInKm // Переводим расстояние в метры
		speed = swimmingMeanSpeed(lengthPool, countPool, duration) * kmhInMsec       // Переводим среднюю скорость в м/с
		calories = SwimmingSpentCalories(lengthPool, countPool, duration, weight)
	default:
		return "неизвестный тип тренировки"
	}

	return fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f м.\nСкорость: %.2f м/с\nСожгли калорий: %.2f\n",
		trainingType, duration, dist, speed, calories,
	)
}

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18   // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 1.79 // среднее количество сжигаемых калорий при беге.
)

// RunningSpentCalories возвращает количество потраченных калорий при беге.
//
// Параметры:
//
// action int — количество совершенных действий (число шагов при ходьбе и беге, либо гребков при плавании).
// weight float64 — вес пользователя.
// duration float64 — длительность тренировки в часах.
func RunningSpentCalories(action int, weight, duration float64) float64 {
	speed := meanSpeed(action, duration)
	return ((runningCaloriesMeanSpeedMultiplier * speed * runningCaloriesMeanSpeedShift) * weight / mInKm) * duration * minInH
}

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
// Параметры:
//
// action int — количество совершенных действий (число шагов при ходьбе и беге, либо гребков при плавании).
// duration float64 — длительность тренировки в часах.
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
func WalkingSpentCalories(action int, duration, weight, height float64) float64 {
	speedMps := meanSpeed(action, duration) * kmhInMsec // Средняя скорость в м/с
	return ((walkingCaloriesWeightMultiplier * weight) + (math.Pow(speedMps, 2)/height)*walkingSpeedHeightMultiplier*weight) * duration * minInH
}

// Константы для расчета калорий, расходуемых при плавании.
const (
	swimmingCaloriesMeanSpeedShift   = 1.1 // среднее количество сжигаемых калорий при плавании относительно скорости.
	swimmingCaloriesWeightMultiplier = 2   // множитель веса при плавании.
)

// SwimmingSpentCalories возвращает количество потраченных калорий при плавании.
//
// Параметры:
//
// lengthPool int — длина бассейна в метрах.
// countPool int — сколько раз пользователь переплыл бассейн.
// duration float64 — длительность тренировки в часах.
// weight float64 — вес пользователя.
func SwimmingSpentCalories(lengthPool, countPool int, duration, weight float64) float64 {
	speed := swimmingMeanSpeed(lengthPool, countPool, duration)
	return (speed + swimmingCaloriesMeanSpeedShift) * swimmingCaloriesWeightMultiplier * weight * duration
}
