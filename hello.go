package hello

func Greeting() string {
	return "Hello from go module"
}

func localOnly() string {
	return "Since local is not capitalized this function cannot be called"
}

type GreetingType struct {
	Hello string
}

func (p *GreetingType) GreetWithReceiver() string {
	return localOnly()
}
