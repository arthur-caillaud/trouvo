package measure

func PrecisionMeasure(tp, fp int) float64 {
	return float64(tp) / float64(tp+fp)
}

func RecallMeasure(tp, fn int) float64 {
	return float64(tp) / float64(tp+fn)
}

func AccuracyMeasure(tp, tn, fn, fp int) float64 {
	return float64(tp+tn) / float64(tp+fn+fp+tn)
}

func EMeasure(tp, fp, fn int) float64 {
	P := PrecisionMeasure(tp, fp)
	R := RecallMeasure(tp, fn)
	beta := P / R
	return 1 - (beta*beta+1)*P*R/(beta*beta*P+R)
}

func FMeasure(tp, fp, fn int) float64 {
	return 1 - EMeasure(tp, fp, fn)
}

func AVG(scores []float64) (avg float64) {
	for _, score := range scores {
		avg += score
	}
	return 100 * avg / float64(len(scores))
}
