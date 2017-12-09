package main

func scan(s string) *group {
	root := newGroup(nil)

	current := root
	inGarbage := false
	for i := 1; i < len(s)-1; i++ {
		if s[i] == '>' {
			// end of garbage
			inGarbage = false
			continue
		}

		if inGarbage {
			if s[i] == '!' {
				// escape: skip next character
				i++
			}

			continue
		}

		switch s[i] {
		case '{':
			// start of group
			current = newGroup(current)
			current.parent.children = append(current.parent.children, current)
		case '}':
			// end of group
			current = current.parent
		case ',':
			// next group
		case '<':
			// start of garbage
			inGarbage = true
		}
	}

	return root
}
