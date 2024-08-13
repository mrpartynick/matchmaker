package models

import "time"

type User struct {
	Name       string    `json:"name"`
	Skill      float64   `json:"skill"`
	Latency    float64   `json:"latency"`
	TimeOfJoin time.Time `json:"time_of_join, omitempty"`
}
