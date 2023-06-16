package main 

var commands = map[string]func() string {
    "hello": func() string {
        return "hello, world!" 
    },
    "foo": func() string {
        return "bar"
    },
    "inspire": GetInspirationQuote,
}

func ExecuteCommand(c string) string {
    callable, ok := commands[c]
    if !ok {
        return ""
    }

    return callable()
}
