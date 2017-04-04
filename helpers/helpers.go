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
