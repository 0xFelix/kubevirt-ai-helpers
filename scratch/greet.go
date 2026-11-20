package scratch

import "fmt"

// Greet returns a greeting.
func Greet(name string) string {
	if name == "" {
		return "hello"
	}
	return fmt.Sprintf("hello %s", name)
}
