package utils

import (
	"github.com/alexedwards/scs"
	"polar/global/config"
	"polar/utils/redis"
	"time"
)

const SessionLifeTime = 24 * time.Hour

var SessionManager *scs.Manager

func init() {
	pool := config.GetRedisPool()
	engine := redis.NewStore(pool)

	SessionManager = scs.NewManager(engine)
	SessionManager.Lifetime(SessionLifeTime)
}
