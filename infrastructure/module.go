package infrastructure

import (
	"FSchedule/infrastructure/repository"
	"github.com/farseer-go/data"
	"github.com/farseer-go/eventBus"
	"github.com/farseer-go/fs/modules"
	"github.com/farseer-go/queue"
	"github.com/farseer-go/redis"
)

type Module struct {
}

func (module Module) DependsModule() []modules.FarseerModule {
	return []modules.FarseerModule{data.Module{}, redis.Module{}, eventBus.Module{}, queue.Module{}}
}

func (module Module) PreInitialize() {
}

func (module Module) Initialize() {
}

func (module Module) PostInitialize() {
	repository.InitRepository()
}

func (module Module) Shutdown() {
}
