package hamming

import (
	"errors"
	"fmt"
)

func Distance(a, b string) (int, error) {
  var distance int

  if len(a) != len(b) {
    return 0, errors.New("strands must match lengths")
  }

  fmt.Printf("[%s] <--> [%s]\n", a, b)

  for i := 0 ; i < len(a) ; i++ {
    if a[i] != b[i] {
      distance++
    }
  }

  return distance, nil
}
