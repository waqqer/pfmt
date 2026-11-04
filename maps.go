package pfmt

func margeMaps(d ...map[string]interface{}) map[string]interface{} {
	if len(d) == 0 {
		return make(map[string]interface{})
	}

	r := make(map[string]interface{}, countLen(d...))

	for _, v := range d {
		for ind, val := range v {
			r[ind] = val
		}
	}
	
	return r
}

func mergeMapsWithBase(base map[string]interface{}, d ...map[string]interface{}) map[string]interface{} {
	maps := make([]map[string]interface{}, 0, len(d) + 1)
	maps = append(maps, base)
	maps = append(maps, d...)

	return margeMaps(maps...)
}

func countLen(d ...map[string]interface{}) int {
	l := 0
	for _, v := range d {
		l += len(v)
	}
	return l
}