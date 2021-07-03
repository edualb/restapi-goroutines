package services

import (
	"math"
	"time"

	"github.com/edualb/restapi-goroutines/models"
)

type Create interface {
	Create(w *models.WalkDogService) error
}

type CreateImpl struct {
	db models.WalkDogServiceDB
}

func NewCreate() Create {
	return &CreateImpl{
		db: models.NewWalkDogServiceDB(),
	}
}

func (c *CreateImpl) Create(w *models.WalkDogService) error {
	if len(w.Status) > 0 {
		return WalkDogServiceStatusFieldNotAllowedError
	}
	if w.Price != 0 {
		return WalkDogServicePriceFieldNotAllowedError
	}
	if w.Duration != 0 {
		return WalkDogServiceDurationFieldNotAllowedError
	}
	if w.PetsQty <= 0 {
		return WalkDogServicePetsQtyEmptyError
	}

	c.calculateDuration(w)
	c.calculatePrice(w)

	var s models.WalkDogStatusEnum = models.Scheduled
	w.Status = s.Get()

	err := c.db.CreateWalkDogService(*w)
	return err
}

func (c *CreateImpl) calculatePrice(w *models.WalkDogService) {
	weekDay := w.StartWalkAt.Weekday()
	var price float64

	if weekDay == time.Saturday || weekDay == time.Sunday {
		price = c.calculatePriceByPetsQty(w, 0.70, 0.45)
	} else {
		price = c.calculatePriceByPetsQty(w, 0.55, 0.35)
	}

	w.Price = price
}

func (c *CreateImpl) calculateDuration(w *models.WalkDogService) {
	duration := math.Trunc(w.FinishWalkAt.Sub(w.StartWalkAt).Minutes())
	w.Duration = duration
}

func (c *CreateImpl) calculatePriceByPetsQty(w *models.WalkDogService, onePetPrice float64, moreThanOnePetPrice float64) float64 {
	if w.PetsQty > 1 {
		return (w.Duration * moreThanOnePetPrice) * float64(w.PetsQty)
	}
	return w.Duration * onePetPrice
}
