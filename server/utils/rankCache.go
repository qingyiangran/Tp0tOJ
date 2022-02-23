package utils

import (
	"errors"
	"sort"
	"sync"
	"time"
)

type RankCache interface {
	SetCalculator(calculator ScoreCalculator)
	GetRank() []uint64
	AddUser(userId uint64)
	AddChallenge(challengeId uint64, originScore uint64)
	Submit(userId uint64, challengeId uint64, stamp time.Time) error
	WarmUp() error
}

type RAMRankCache struct {
	calculator     ScoreCalculator
	mutex          sync.RWMutex
	rank           []uint64
	challengeScore map[uint64]uint64
	challengeSolve map[uint64][]uint64
	userScore      map[uint64]uint64
	userTime       map[uint64]time.Time
}

var Cache RankCache

func init() {
	//calculator need to be init before usage
	Cache = &RAMRankCache{rank: []uint64{}, challengeScore: map[uint64]uint64{}, challengeSolve: map[uint64][]uint64{}, userScore: map[uint64]uint64{}, userTime: map[uint64]time.Time{}}
}

func (cache *RAMRankCache) SetCalculator(calculator ScoreCalculator) {
	cache.calculator = calculator
}

func (cache *RAMRankCache) GetRank() []uint64 {
	cache.mutex.RLock()
	defer cache.mutex.RUnlock()
	ret := make([]uint64, len(cache.rank))
	copy(ret, cache.rank)
	return ret
}

func (cache *RAMRankCache) AddUser(userId uint64) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	cache.userScore[userId] = 0
	cache.userTime[userId] = time.Time{}
}

func (cache *RAMRankCache) AddChallenge(challengeId uint64, originScore uint64) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	cache.challengeScore[challengeId] = originScore
	cache.challengeSolve[challengeId] = []uint64{}
}

func (cache *RAMRankCache) Submit(userId uint64, challengeId uint64, stamp time.Time) error {
	if err := cache.submitImpl(userId, challengeId, stamp); err != nil {
		return err
	}
	cache.refreshRank()
	return nil
}
func (cache *RAMRankCache) submitImpl(userId uint64, challengeId uint64, stamp time.Time) error {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	oldScore, exist := cache.challengeScore[challengeId]
	if !exist {
		return errors.New("unexist challenge")
	}
	count := uint64(len(cache.challengeSolve[challengeId])) + 1
	newScore := cache.calculator.GetScore(challengeId, count)

	cache.userTime[userId] = stamp
	cache.challengeSolve[challengeId] = append(cache.challengeSolve[challengeId], userId)
	cache.userScore[userId] += cache.calculator.GetIncrementScore(oldScore, count-1)
	cache.challengeScore[challengeId] = newScore

	for index, user := range cache.challengeSolve[challengeId] {
		cache.userScore[user] -= cache.calculator.GetDeltaScoreForUser(oldScore, newScore, uint64(index))
	}
	return nil
}

type ScoreItem struct {
	userId uint64
	score  uint64
	stamp  time.Time
}

type ScoreItems = []*ScoreItem

type RankSorter struct {
	ScoreItems
}

func (items RankSorter) Len() int {
	return len(items.ScoreItems)
}

func (items RankSorter) Swap(i, j int) {
	items.ScoreItems[i], items.ScoreItems[j] = items.ScoreItems[j], items.ScoreItems[i]
}

func (items RankSorter) Less(i, j int) bool {
	return items.ScoreItems[i].score > items.ScoreItems[j].score || (items.ScoreItems[i].score == items.ScoreItems[j].score && items.ScoreItems[i].stamp.Before(items.ScoreItems[j].stamp))
}

func (cache *RAMRankCache) refreshRank() {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	rank := ScoreItems{}
	for user, score := range cache.userScore {
		rank = append(rank, &ScoreItem{
			userId: user,
			score:  score,
			stamp:  cache.userTime[user],
		})
	}
	sort.Sort(RankSorter{rank})
	var rankSaved []uint64
	for _, item := range rank {
		rankSaved = append(rankSaved, item.userId)
	}
	cache.rank = rankSaved
}

func (cache *RAMRankCache) WarmUp() error {
	// TODO:
	// do add & submitImpl then refreshRank

}
