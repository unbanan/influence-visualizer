package context

import (
	"contest-influence/server/internal/config"
	"contest-influence/server/internal/repos"
)

type Context struct {
	InfluenceDBRepo *repos.InfluenceDBRepo
}

func NewContext(config config.ServiceConfig) *Context {
	return &Context{
		InfluenceDBRepo: repos.NewInfluenceDBRepo(config.InfluenceDB),
	}
}
