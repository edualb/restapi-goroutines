package models

import "time"

type WalkDogService struct {
	Status       WalkDogStatus `json:"status,omitempty"`
	Price        float64       `json:"price,omitempty"`
	Duration     float64       `json:"duration,omitempty"`
	PetsQty      int8          `json:"pets_qty"`
	StartWalkAt  time.Time     `json:"start_walk_at"`
	FinishWalkAt time.Time     `json:"finish_walk_at,omitempty"`
}
