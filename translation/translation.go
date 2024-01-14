package translation

func Translation(word string, language string) string {

	switch language {
	case "english":
		return "hello"
	case "spanish":
		return "hola"
	case "german":
		return "hallo"
	default:
		return ""
	}
}
