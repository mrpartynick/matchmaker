package matchmaker

import (
	"encoding/json"
	"fmt"
	"matchmaker/internal/models"
	"time"
)

func (mm *MatchMaker) CreateGroup(usr *models.User) {
	fmt.Println("Creating group")
	mm.numberOfGroups++
	group := models.Group{
		Number:           mm.numberOfGroups,
		Skill:            usr.Skill,
		Latency:          usr.Latency,
		SkillTolerance:   mm.cfg.SkillTolerance,
		LatencyTolerance: mm.cfg.LatencyTolerance,
		Users:            []models.User{*usr},
		Cap:              mm.cfg.GroupSize,
		UsersCounter:     1,
	}
	mm.groups = append(mm.groups, &group)
}

func (mm *MatchMaker) RemoveGroup(i int) []*models.Group {
	return append(mm.groups[:i], mm.groups[i+1:]...)
}

func (mm *MatchMaker) printGroup(group *models.Group) {
	var stat models.GroupStat
	timeOfPrinting := time.Now()

	for _, usr := range group.Users {
		stat.UserNames = append(stat.UserNames, usr.Name)

		t := timeOfPrinting.Sub(usr.TimeOfJoin).Minutes()

		stat.AvgSkill += usr.Skill
		stat.AvgLatency += usr.Latency
		stat.AvgTime += t

		if usr.Skill < stat.MinSkill {
			stat.MinSkill = usr.Skill
		}
		if usr.Skill > stat.MaxSkill {
			stat.MaxSkill = usr.Skill
		}
		if usr.Latency < stat.MinLatency {
			stat.MinLatency = usr.Latency
		}
		if usr.Latency > stat.MaxLatency {
			stat.MaxLatency = usr.Latency
		}
		if t < stat.MinTime {
			stat.MinTime = t
		}
		if t > stat.MaxTime {
			stat.MaxTime = t
		}
	}
	stat.AvgSkill = stat.AvgSkill / float64(len(group.Users))
	stat.AvgLatency = stat.AvgLatency / float64(len(group.Users))
	stat.AvgTime = stat.AvgTime / float64(len(group.Users))

	m, err := json.Marshal(&stat)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf(string(m))
}
