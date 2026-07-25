package collatzconjecture

import (
  "errors"
)

func doStep(n int) (int) {
	if n % 2 == 1 || n < 1 {
    return (3 * n) + 1
  } else {
    return n / 2
  }
}

func CollatzConjecture(n int) (int, error) {

  if n < 1 {
    return 0, errors.New("invalid starting integer, must be positive integer")
  }
  iterationCount := 0 
  for {
    if n == 1 {
      break
    }
      
    iterationCount = iterationCount + 1
    n = doStep(n)
  }
  return iterationCount, nil
}
