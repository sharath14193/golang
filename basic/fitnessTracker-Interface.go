package main

import (
    "fmt"
    "time"
)

type Workout interface{
  Duration() time.Duration
  CaloriesBurned() float64
  RecordStats() 
  GetType() string
}
func (c CardioWorkout) CaloriesBurned() float64 {
    calories := c.duration.Minutes() * 10 * (c.avgHeartRate / 100)
    return calories
} 

func (c CardioWorkout) RecordStats() {
  fmt.Printf("Stats recorded\n")
}

func (c CardioWorkout) GetType() string{
  return "Cardio"
}
// CardioWorkout struct
type CardioWorkout struct {
    duration     time.Duration
    distance     float64
    avgHeartRate float64
}

func (c CardioWorkout) Duration() time.Duration {
    return c.duration
}


// StrengthWorkout struct
type StrengthWorkout struct {
    duration time.Duration
    weight   int
    reps     int
    sets     int
    avgHeartRate float64
}

func (s StrengthWorkout) Duration() time.Duration {
    return s.duration
}

func (s StrengthWorkout) CaloriesBurned() float64 {
    calories := s.duration.Minutes() * 5 * (s.avgHeartRate / 100)
    return calories
}

func (s StrengthWorkout) RecordStats() {
  fmt.Printf("Stats recorded\n")
}

func (s StrengthWorkout) GetType() string{
  return "Strength"
}
func summarizeWorkouts(workouts []Workout) {
    for _, workout := range workouts {
        fmt.Println("Type:", workout.GetType())
        fmt.Println("Duration:", workout.Duration())
        fmt.Println("Calories Burned:", workout.CaloriesBurned())
        workout.RecordStats()
    }
}

func main() {
    workout := CardioWorkout{duration: 30 * time.Minute, distance: 5.0, avgHeartRate: 120}

    fmt.Println("Workout Summary:")
    summarizeWorkouts([]Workout{workout})
}
