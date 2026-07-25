package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
    LanguageName() string
    Greet(visitorName string) string
}

func SayHello(visitorName string, greeter Greeter) string {
    var greetingString string
	greetingString = fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(visitorName))
    return greetingString
}

type Italian struct {}

func (_ Italian) LanguageName() string {
    return "Italian"
}

func (_ Italian) Greet(visitorName string) string {
    return fmt.Sprintf("Ciao %s!", visitorName)
}

type Portuguese struct {}

func (_ Portuguese) LanguageName() string {
    return "Portuguese"
}

func (_ Portuguese) Greet(visitorName string) string {
    return fmt.Sprintf("Olá %s!", visitorName)
}