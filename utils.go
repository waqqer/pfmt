package pfmt

func margeRules(d ...Rule) Rule {
	if len(d) == 0 {
		return make(Rule)
	}

	r := make(Rule, countUniqueRulesLen(d...))

	for _, v := range d {
		for ind, val := range v {
			r[ind] = val
		}
	}

	return r
}

func mergeRulesWithBase(base Rule, d ...Rule) Rule {
	maps := make([]Rule, len(d)+1)
	maps[0] = base

	for i := 0; i < len(d); i++ {
		maps[i+1] = d[i]
	}

	return margeRules(maps...)
}

func countUniqueRulesLen(d ...Rule) int {
	count := 0
	for i, current := range d {
		for key := range current {
			seen := false
			for j := 0; j < i; j++ {
				if _, exists := d[j][key]; exists {
					seen = true
					break
				}
			}
			if !seen {
				count++
			}
		}
	}
	return count
}
