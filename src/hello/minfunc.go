package

import "math"

func MinInt(x, y int) int {

  return int(math.Min(float64(x), float64(y))) // Need to convert number to float64 (function Min rules)

}
