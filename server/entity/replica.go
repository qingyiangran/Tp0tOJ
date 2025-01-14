package entity

import "time"

type Replica struct {
	ReplicaId   uint64    `gorm:"primaryKey"`
	ChallengeId uint64    `gorm:"not null"`
	Challenge   Challenge `gorm:"foreignKey:ChallengeId;references:ChallengeId;"`
	Singleton   bool
	Status      string    `gorm:"check: status <> ''"`
	Flag        string    `gorm:"check: flag <> ''"`
	CreatedAt   time.Time `gorm:"autoCreateTime;not null;"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime;not null;"`
}
