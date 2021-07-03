package services

import "errors"

var WalkDogServiceStatusFieldNotAllowedError = errors.New("It is not allowed create a walk dog service with status field filled")
var WalkDogServicePriceFieldNotAllowedError = errors.New("It is not allowed create a walk dog service with price field filled")
var WalkDogServiceDurationFieldNotAllowedError = errors.New("It is not allowed create a walk dog service with duration field filled")
var WalkDogServicePetsQtyEmptyError = errors.New("It is not allowed create a walk dog service without pets")
