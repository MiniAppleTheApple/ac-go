{{ .Content }}

func (xs {{ .Type }}) Compare(ys {{ .Type }}) bool {
	if len(xs) != len(ys) {
		return false
	}

	for i := range xs {
		if xs[i] != ys[i] {
			return false
		}
	}

	return true
}
