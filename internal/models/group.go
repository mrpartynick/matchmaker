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

type GroupStat struct {
	Number int

	MinSkill float64
	AvgSkill float64
	MaxSkill float64

	MinLatency float64
	AvgLatency float64
	MaxLatency float64

	MinTime float64
	AvgTime float64
	MaxTime float64

	UserNames []string
}
