package models

type Group struct {
	Number  int
	Skill   float64
	Latency float64

	SkillTolerance   float64
	LatencyTolerance float64

	Users []User

	Cap          int
	UsersCounter int
}
