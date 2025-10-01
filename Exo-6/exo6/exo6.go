package main

import (
	"fmt"
	"time"
)

type Workout interface {
	Duration() time.Duration
	CaloriesBurned() float64
	RecordStats()
	GetType() string
}

type CardioWorkout struct {
	duration     time.Duration
	distance     float64
	avgHeartRate float64
}

func (c CardioWorkout) Duration() time.Duration {
	return c.duration
}

func (c CardioWorkout) CaloriesBurned() float64 {
	return c.duration.Minutes() * 10.0 * (c.avgHeartRate / 100.0)
}

func (c CardioWorkout) RecordStats() {
	fmt.Printf("   [Record] Cardio  | Duration: %v | Distance: %.2f km | Avg HR: %.0f bpm | Calories: %.1f\n",
		c.duration, c.distance, c.avgHeartRate, c.CaloriesBurned())
}

func (c CardioWorkout) GetType() string {
	return "Cardio"
}

type StrengthWorkout struct {
	duration time.Duration
	weight   int
	reps     int
	sets     int
}

func (s StrengthWorkout) Duration() time.Duration {
	return s.duration
}

func (s StrengthWorkout) CaloriesBurned() float64 {
	return s.duration.Minutes() * 5.0 * (float64(s.weight) / 10.0)
}

func (s StrengthWorkout) RecordStats() {
	fmt.Printf("   [Record] Strength| Duration: %v | Weight: %d kg | Sets: %d | Reps: %d | Calories: %.1f\n",
		s.duration, s.weight, s.sets, s.reps, s.CaloriesBurned())
}

func (s StrengthWorkout) GetType() string {
	return "Strength"
}

func summarizeWorkouts(workouts []Workout) {
	for i, w := range workouts {
		fmt.Printf("%d) %s — Duration: %v — Calories: %.1f\n", i+1, w.GetType(), w.Duration(), w.CaloriesBurned())
		w.RecordStats()
	}
}

func main() {

	var workouts []Workout

	workouts = append(workouts, CardioWorkout{
		duration:     30 * time.Minute,
		distance:     5.2,
		avgHeartRate: 150,
	})

	workouts = append(workouts, StrengthWorkout{
		duration: 45 * time.Minute,
		weight:   80,
		sets:     4,
		reps:     10,
	})

	fmt.Println("Workout Summary:")
	summarizeWorkouts(workouts)

}
