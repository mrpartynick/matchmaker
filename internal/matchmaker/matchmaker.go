package matchmaker

import (
	"fmt"
	"log/slog"
	"matchmaker/config"
	"matchmaker/internal/models"
	"math"
	"sync"
)

type MatchMaker struct {
	cfg    *config.Config
	groups []*models.Group

	numberOfGroups int
	mu             sync.Mutex
}

func NewMatchMaker(cfg *config.Config) *MatchMaker {
	return &MatchMaker{
		cfg:    cfg,
		groups: make([]*models.Group, 0),
	}
}

func (mm *MatchMaker) Process(usr *models.User) {
	mm.mu.Lock()
	defer mm.mu.Unlock()

	slog.Debug(fmt.Sprintf("Processing user: %v", usr))

	for i, group := range mm.groups {
		cond1 := math.Abs(usr.Skill-group.Skill) < group.SkillTolerance
		cond2 := math.Abs(usr.Latency-group.Latency) < group.LatencyTolerance

		if cond1 && cond2 {
			group.Users = append(group.Users, *usr)
			group.UsersCounter++

			if group.UsersCounter == group.Cap {
				mm.groups = mm.RemoveGroup(i)
				mm.printGroup(group)
			}
			return
		}
	}
	mm.CreateGroup(usr)
}

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
	fmt.Printf("Group %v: \n", group)
}
