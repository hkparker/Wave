package helpers

func Setup() {
	setEnvironment()
	setHostname()
	setDatabase()
}

func StringIncludedIn(set []string, member string) bool {
	for _, str := range set {
		if str == member {
			return true
		}
	}
	return false
}

func StringsExcept(set []string, member string) []string {
	new_set := make([]string, 0)
	for _, str := range set {
		if str != member {
			new_set = append(new_set, str)
		}
	}
	return new_set
}
