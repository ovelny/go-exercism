package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	cows int
}

func (err *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", err.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodderAmount, err := weightFodder.FodderAmount()

	switch {
	case cows == 0:
		err = errors.New("division by zero")
	case (err == ErrScaleMalfunction || err == nil) && fodderAmount < 0:
		err = errors.New("negative fodder")
	case cows < 0:
		err = &SillyNephewError{cows: cows}
	case err == nil && fodderAmount > 0:
		fodderAmount = fodderAmount / float64(cows)
		err = nil
	case err == ErrScaleMalfunction && fodderAmount > 0:
		fodderAmount = (fodderAmount * 2) / float64(cows)
		err = nil
	}

	if err != nil {
		return 0, err
	}

	return fodderAmount, err
}
