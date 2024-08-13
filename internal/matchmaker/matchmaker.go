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
			slog.Debug(fmt.Sprintf("Adding user: %v to group: %d", usr, group.Number))
			group.Users = append(group.Users, *usr)
			group.UsersCounter++

			if group.UsersCounter == group.Cap {
				slog.Debug(fmt.Sprintf("Group %d is full. Printing...", group.Number))
				mm.groups = mm.RemoveGroup(i)
				mm.printGroup(group)
			}
			return
		}
	}
	slog.Debug(fmt.Sprintf("There is no group for user %v . Creating new...", usr))
	mm.CreateGroup(usr)
}
