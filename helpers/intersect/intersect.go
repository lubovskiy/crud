package intersect

func IntersectStr(a, b []string) []string {
	set := make([]string, 0)
	hash := make(map[string]bool)

	for i := 0; i < len(a); i++ {
		hash[a[i]] = true
	}

	for i := 0; i < len(b); i++ {
		if _, found := hash[b[i]]; found {
			set = append(set, b[i])
		}
	}

	return set
}
