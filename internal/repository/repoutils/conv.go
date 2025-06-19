package repoutils

func ConvertMeiliHitsToIntSlice(hits []any) []int {
	if len(hits) == 0 {
		return []int{}
	}

	mapSlice := convertRawHitsToMapSlice(hits)

	ints := make([]int, 0, len(mapSlice))

	for _, hit := range mapSlice {
		idAny, ok := hit["id"]
		if !ok {
			continue
		}

		id, ok := idAny.(int)
		if !ok {
			continue
		}

		ints = append(ints, id)
	}

	return ints
}

func convertRawHitsToMapSlice(hits []any) []map[string]any {
	mapSlice := make([]map[string]any, 0, len(hits))

	for _, rawHit := range hits {
		hit, ok := rawHit.(map[string]any)
		if !ok {
			continue
		}
		mapSlice = append(mapSlice, hit)
	}

	return mapSlice
}
