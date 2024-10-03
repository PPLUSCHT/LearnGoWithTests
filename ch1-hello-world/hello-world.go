package ch1helloworld

type Language struct {
	translationOfWorld string
	helloPrefix        string
}

var English = Language{
	"World",
	"Hello, ",
}

var Spanish = Language{
	"Mundo",
	"Hola, ",
}

func Hello(name string, language Language) string {
	if name == "" {
		name = language.translationOfWorld
	}
	return language.helloPrefix + name
}
