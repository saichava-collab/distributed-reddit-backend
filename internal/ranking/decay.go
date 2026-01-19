
package ranking

import "math"

func HotScore(score int, ageSeconds float64) float64 {
    sign := 1.0
    if score < 0 {
        sign = -1.0
    }
    return sign * math.Log10(math.Max(math.Abs(float64(score)), 1)) +
        ageSeconds/45000.0
}
