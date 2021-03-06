package handler

import (
	"sync"
	"time"

	"github.com/checkr/flagr/pkg/config"
	"github.com/checkr/flagr/pkg/entity"

	"github.com/sirupsen/logrus"
)

var (
	singletonEvalCache     *EvalCache
	singletonEvalCacheOnce sync.Once
)

// EvalCache is the in-memory cache just for evaluation
type EvalCache struct {
	mapCache     map[uint]*entity.Flag
	mapCacheLock sync.RWMutex

	refreshTimeout  time.Duration
	refreshInterval time.Duration

	redisCache *redisCache
}

// GetEvalCache gets the EvalCache
var GetEvalCache = func() *EvalCache {
	singletonEvalCacheOnce.Do(func() {
		ec := &EvalCache{
			mapCache:        make(map[uint]*entity.Flag),
			refreshTimeout:  config.Config.EvalCacheRefreshTimeout,
			refreshInterval: config.Config.EvalCacheRefreshInterval,
			redisCache:      &redisCache{},
		}
		singletonEvalCache = ec
	})
	return singletonEvalCache
}

// Start starts the polling of EvalCache
func (ec *EvalCache) Start() {
	err := ec.reloadMapCache()
	if err != nil {
		panic(err)
	}
	go func() {
		for range time.Tick(ec.refreshInterval) {
			err := ec.reloadMapCache()
			if err != nil {
				logrus.WithField("err", err).Error("reload evaluation cache error")
			}
		}
	}()
}

// GetByFlagIDs gets the flags by flagIDs from the EvalCache
func (ec *EvalCache) GetByFlagIDs(flagIDs []uint) map[uint]*entity.Flag {
	m := make(map[uint]*entity.Flag)

	ec.mapCacheLock.RLock()
	for _, flagID := range flagIDs {
		f, ok := ec.mapCache[flagID]
		if ok {
			m[flagID] = f
		}
	}
	ec.mapCacheLock.RUnlock()
	return m
}

// GetByFlagID gets the flag by flagID
func (ec *EvalCache) GetByFlagID(flagID uint) *entity.Flag {
	ec.mapCacheLock.RLock()
	f := ec.mapCache[flagID]
	ec.mapCacheLock.RUnlock()
	return f
}

func (ec *EvalCache) reloadMapCache() error {
	if config.Config.NewRelicEnabled {
		defer config.Global.NewrelicApp.StartTransaction("eval_cache_reload", nil, nil).End()
	}

	fs := []entity.Flag{}
	q := entity.NewFlagQuerySet(getDB())
	if err := q.All(&fs); err != nil {
		return err
	}
	m := make(map[uint]*entity.Flag)
	for i := range fs {
		ptr := &fs[i]
		err := ptr.Preload(getDB())
		if err != nil {
			return err
		}
		m[ptr.ID] = ptr
	}

	for _, f := range m {
		err := f.PrepareEvaluation()
		if err != nil {
			return err
		}
	}

	ec.mapCacheLock.Lock()
	ec.mapCache = m
	ec.mapCacheLock.Unlock()
	return nil
}

// TODO implement the redis cache before hitting DB
type redisCache struct {
}

func (rc *redisCache) GetFlags() ([]entity.Flag, error) {
	return nil, nil
}

func (rc *redisCache) SetFlags(fs []entity.Flag) error {
	return nil
}

func (rc *redisCache) Lock() error {
	return nil
}

func (rc *redisCache) Unlock() error {
	return nil
}
