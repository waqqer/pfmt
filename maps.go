package pfmt

func margeMaps(d ...map[string]interface{}) map[string]interface{} {
	if len(d) == 0 {
		return make(map[string]interface{})
	}

	r := make(map[string]interface{}, countUniqueLen(d...))

	for _, v := range d {
		for ind, val := range v {
			r[ind] = val
		}
	}
	
	return r
}

func mergeMapsWithBase(base map[string]interface{}, d ...map[string]interface{}) map[string]interface{} {
	maps := make([]map[string]interface{}, len(d) + 1)
	maps[0] = base
	
	for i := 0; i < len(d); i++ {
		maps[i + 1] = d[i]
	}

	return margeMaps(maps...)
}

func countUniqueLen(d ...map[string]interface{}) int {
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
