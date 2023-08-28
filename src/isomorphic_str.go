package main

func isIsomorphic(s string, t string) bool {
	if len(s) != len(s) {
		return false
	}

	sMap := make(map[byte]byte)
	tMap := make(map[byte]byte)

	for i, _ := range s {
		sChar := s[i]
		tChar := t[i]

		if value, ok := sMap[sChar]; ok {
			if value != tChar {
				return false
			}
		} else {
			if value, ok := tMap[tChar]; ok {
				if value != sChar {
					return false
				}
			}
			sMap[sChar] = tChar
			tMap[tChar] = sChar
		}
	}

	return true
}
