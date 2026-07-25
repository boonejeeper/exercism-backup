package raindrops

import (
    "strings"
    "fmt"
)

func Convert(number int) string {
    var out strings.Builder
    if number % 3 == 0 {
        out.WriteString("Pling")
    }
    if number % 5 == 0 {
        out.WriteString("Plang")
    }
    if number % 7 == 0 {
        out.WriteString("Plong")
    }
    if out.Len() == 0 {
    	out.WriteString(fmt.Sprintf("%d", number))
    }
    return out.String()
}
