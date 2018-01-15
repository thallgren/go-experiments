package testing


import (
	. "testing"
	. "strconv"
)

func TestBug5225(t *T) {
	for a := 0; a < 10; a++ {
		t.Run(`a` + Itoa(a), func(t *T) {
			for b := 0; b < 10; b++ {
				t.Run(`b` + Itoa(b), func(t *T) {
					for c := 0; c < 10; c++ {
						t.Run(`c` + Itoa(c), func(t *T) {
							if a == 9 && b % 3 == 0 && c % 5 == 0 {
								t.Errorf("Error %d%d", b, c)
							}
						})
					}
				})
			}
		})
	}
}

