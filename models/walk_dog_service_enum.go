package models

type WalkDogStatusEnum int

const (
	Scheduled = iota
	Canceled
	Done
)

type WalkDogStatus string

func (w WalkDogStatusEnum) Get() WalkDogStatus {
	return [...]WalkDogStatus{
		"SCHEDULED",
		"CANCELED",
		"DONE",
	}[w]
}
