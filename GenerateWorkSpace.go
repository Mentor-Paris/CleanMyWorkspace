package CleanMyWorkspace

import "math/rand"

// GenerateWorkSpace generate the souk with random waste
func GenererateWorkSpace() *[][]*string {
	w, h := rand.Intn(25)+25, rand.Intn(25)+25
	result := make([][]*string, h)
	for i := range result {
		result[i] = make([]*string, w)
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if rand.Intn(100) < 20 {
				result[i][j] = randomWaste()
			} else {
				result[i][j] = nil
			}
		}
	}
	return &result
}
