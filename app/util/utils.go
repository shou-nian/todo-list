package util

func QueryAll(all map[string]string) []map[string]string {
	var result []map[string]string

	for k, v := range all {
		result = append(result, map[string]string{k: v})
	}

	return result
}

func ReverseStatus(status string) string {
	if status == "open" {
		return "close"
	}
	return "open"
}
