package convtrad

var convTrad *convtrad

type convtrad struct {
	T2S map[rune]rune
	S2T map[rune]rune
}

func init() {
	convTrad = &convtrad{}

	convTrad.T2S = make(map[rune]rune)
	for k := range trad {
		convTrad.T2S[trad[k]] = simp[k]
	}

	convTrad.S2T = make(map[rune]rune)
	for k := range simp {
		convTrad.S2T[simp[k]] = trad[k]
	}
}

// ToSimp converts string from Traditional Chinese to Simplified Chinese.
func ToSimp(s string) string {
	r := []rune(s)
	for k := range r {
		_, exists := convTrad.T2S[r[k]]
		if exists {
			r[k] = convTrad.T2S[r[k]]
		}
	}
	return string(r)
}

// ToTrad converts string from Simplified Chinese to Traditional Chinese.
func ToTrad(s string) string {
	r := []rune(s)
	for k := range r {
		_, exists := convTrad.S2T[r[k]]
		if exists {
			r[k] = convTrad.S2T[r[k]]
		}
	}
	return string(r)
}
