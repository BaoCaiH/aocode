package day19

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/BaoCaiH/aocode/2023/utils"
)

type ruleSet struct {
	c, o string
	n    int
	t    string
}

func Airflow() {
	sc, file := utils.Read("./day19/input.txt")
	// sc, file := utils.Read("./day19/input_1_example.txt")
	if sc == nil {
		return
	}
	defer file.Close()

	isParts := false
	workflow := map[string][]ruleSet{}
	parts := []map[string]int{}
	line := ""
	for sc.Scan() {
		line = sc.Text()
		if len(line) == 0 {
			isParts = true
			continue
		}
		if isParts {
			xmasTmp := strings.Split(strings.Trim(line, "{}"), ",")
			xmas := map[string]int{}
			for _, tmp := range xmasTmp {
				c := strings.Split(tmp, "=")[0]
				n, _ := strconv.Atoi(strings.Split(tmp, "=")[1])
				xmas[c] = n
			}
			parts = append(parts, xmas)
			continue
		}
		c := strings.Split(line, "{")[0]
		rules := strings.Split(strings.Trim(strings.Split(line, "{")[1], "}"), ",")
		ruleSets := []ruleSet{}
		for _, r := range rules {
			rs := strings.Split(r, ":")
			if len(rs) == 1 {
				ruleSets = append(ruleSets, ruleSet{"", "", -1, r})
				continue
			}
			n, _ := strconv.Atoi(rs[0][2:])
			ruleSets = append(ruleSets, ruleSet{rs[0][0:1], rs[0][1:2], n, rs[1]})
		}
		workflow[c] = ruleSets
	}

	total := 0
	done := false
	curr := "in"
	for _, p := range parts {
		done = false
		curr = "in"
		for !done {
			ruleSets := workflow[curr]
			for _, rule := range ruleSets {
				c, o, n, t := rule.c, rule.o, rule.n, rule.t
				if c == "" {
					if t == "A" {
						done = true
						total += p["x"] + p["m"] + p["a"] + p["s"]
						break
					}
					if t == "R" {
						done = true
						break
					}
					curr = t
					break
				}
				if o == ">" {
					if p[c] > n {
						if t == "A" {
							done = true
							total += p["x"] + p["m"] + p["a"] + p["s"]
							break
						}
						if t == "R" {
							done = true
							break
						}
						curr = t
						break
					}
				} else {
					if p[c] < n {
						if t == "A" {
							done = true
							total += p["x"] + p["m"] + p["a"] + p["s"]
							break
						}
						if t == "R" {
							done = true
							break
						}
						curr = t
						break
					}
				}
			}
		}
	}
	fmt.Printf("\tPart 1: %d\n", total)

	total = 0
	base := []struct {
		t string
		b map[string]struct{ l, h int }
	}{{"in", map[string]struct{ l, h int }{"x": {1, 4000}, "m": {1, 4000}, "a": {1, 4000}, "s": {1, 4000}}}}
	for len(base) > 0 {
		sets := base[0]
		r, ms := sets.t, sets.b
		base = base[1:]
		if r == "A" {
			total += (ms["x"].h - ms["x"].l + 1) * (ms["m"].h - ms["m"].l + 1) * (ms["a"].h - ms["a"].l + 1) * (ms["s"].h - ms["s"].l + 1)
			continue
		}
		if r == "R" {
			continue
		}
		for _, rule := range workflow[r] {
			c, o, n, t := rule.c, rule.o, rule.n, rule.t
			if c == "" {
				base = append(base, struct {
					t string
					b map[string]struct {
						l int
						h int
					}
				}{t, ms})
				break
			}
			if o == ">" {
				if n >= ms[c].l && n <= ms[c].h {
					tmp := map[string]struct{ l, h int }{}
					for k, v := range ms {
						if k == c {
							tmp[k] = struct {
								l int
								h int
							}{n + 1, v.h}
						} else {
							tmp[k] = v
						}
					}
					base = append(base, struct {
						t string
						b map[string]struct {
							l int
							h int
						}
					}{t, tmp})
					ms[c] = struct {
						l int
						h int
					}{ms[c].l, n}
				} else if n < ms[c].l {
					tmp := map[string]struct{ l, h int }{}
					for k, v := range ms {
						if k == c {
							tmp[k] = struct {
								l int
								h int
							}{n + 1, v.h}
						} else {
							tmp[k] = v
						}
					}
					base = append(base, struct {
						t string
						b map[string]struct {
							l int
							h int
						}
					}{t, tmp})
				} else {
					ms[c] = struct {
						l int
						h int
					}{ms[c].l, n}
				}
			} else {
				if n >= ms[c].l && n <= ms[c].h {
					tmp := map[string]struct{ l, h int }{}
					for k, v := range ms {
						if k == c {
							tmp[k] = struct {
								l int
								h int
							}{v.l, n - 1}
						} else {
							tmp[k] = v
						}
					}
					base = append(base, struct {
						t string
						b map[string]struct {
							l int
							h int
						}
					}{t, tmp})
					ms[c] = struct {
						l int
						h int
					}{n, ms[c].h}
				} else if n < ms[c].l {
					ms[c] = struct {
						l int
						h int
					}{n, ms[c].h}
				} else {
					tmp := map[string]struct{ l, h int }{}
					for k, v := range ms {
						if k == c {
							tmp[k] = struct {
								l int
								h int
							}{v.l, n - 1}
						} else {
							tmp[k] = v
						}
					}
					base = append(base, struct {
						t string
						b map[string]struct {
							l int
							h int
						}
					}{t, tmp})
				}
			}
		}
	}
	fmt.Printf("\tPart 2: %d\n", total)
}
