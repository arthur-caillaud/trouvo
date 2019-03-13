package search

import (
	"regexp"
	"strings"
)

func isPrimal(q string) bool {
	re := regexp.MustCompile("\\|\\||&&|!")
	res := re.FindString(q)
	return res == ""
}

func isOr(q string) bool {
	re := regexp.MustCompile("\\|\\|")
	res := re.FindString(q)
	return res != ""
}

func splitWords(q string) []string {
	return strings.Split(q, " ")
}

func parse(q string, op string) (p []string) {
	switch op {
	case "AND":
		return parseAnd(q)
	case "OR":
		return parseOr(q)
	case "NOT":
		return parseNot(q)
	default:
		return []string{q}
	}
}

func parseOr(q string) (p []string) {
	re := regexp.MustCompile("\\|\\|")
	p = re.Split(q, -1)
	return
}

func isAnd(q string) bool {
	re := regexp.MustCompile("&&")
	res := re.FindString(q)
	return res != ""
}

func parseAnd(q string) (p []string) {
	re := regexp.MustCompile("&&")
	p = re.Split(q, -1)
	return
}

func isNot(q string) bool {
	re := regexp.MustCompile("!")
	res := re.FindString(q)
	return res != ""
}

func parseNot(q string) (p []string) {
	re := regexp.MustCompile("!")
	p = re.Split(q, -1)
	return p[1:]
}

func union(slices ...[]int) (union []int) {
	for _, slice := range slices {
		union = append(union, slice...)
	}
	union = filterDuplicates(union)
	return
}

func intersect(slices ...[]int) (inter []int) {
	for _, el := range slices[0] {
		shouldAppend := false
		for _, slice := range slices {
			shouldAppend = false
			for _, _el := range slice {
				if el == _el {
					shouldAppend = true
				}
			}
			if !shouldAppend {
				break
			}
		}
		if shouldAppend {
			inter = append(inter, el)
		}
	}
	return
}

func subtract(a []int, slices ...[]int) (sub []int) {
	for _, el := range a {
		shouldAppend := true
		for _, slice := range slices {
			for _, _el := range slice {
				if el == _el {
					shouldAppend = false
					break
				}
			}
		}
		if shouldAppend {
			sub = append(sub, el)
		}
	}
	return
}

func filterDuplicates(in []int) (out []int) {
	for _, el := range in {
		shouldAppend := true
		for _, _el := range out {
			if _el == el {
				shouldAppend = false
				break
			}
		}
		if shouldAppend {
			out = append(out, el)
		}
	}
	return
}

func makeSortDocClosure(res []*Result) func(i, j int) bool {
	return func(i, j int) bool {
		return res[i].score < res[j].score
	}
}
