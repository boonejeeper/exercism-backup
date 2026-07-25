package thefarm

import (
    "errors"
    "fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, numCows int) (float64, error) {
    fodderAmount, fodderAmountErr := fc.FodderAmount(numCows)
    if nil != fodderAmountErr {
        return 0.0, fodderAmountErr
    }
    fatteningFactor, fatteningFactorErr := fc.FatteningFactor()
    if nil != fatteningFactorErr {
        return 0.0, fatteningFactorErr
    }
    
    return (fodderAmount * fatteningFactor)/float64(numCows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, numCows int) (float64, error) {
    if numCows <= 0 {
        return 0.0, errors.New("invalid number of cows")
    }
    return DivideFood(fc, numCows)
}


// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(numCows int) error {
    switch {
    case numCows < 0:
        return &InvalidCowsError{
            numberOfCows: numCows,
            message: "there are no negative cows",
        }
    case numCows == 0:
        return &InvalidCowsError{
            numberOfCows: numCows,
            message: "no cows don't need food",
        }
    default: 
        return nil
    }
}

type InvalidCowsError struct {
    numberOfCows int
    message string
}

func (e *InvalidCowsError) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s", e.numberOfCows, e.message)
}


// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
