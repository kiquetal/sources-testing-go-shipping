package translation

func Translation(word string, language string) string {
	switch language {
	case "english":
		return "Hello"
	case "spanish":
		return "Hola"
	case "german":
		return "Hallo"
	default:
		return ""
	}
}
