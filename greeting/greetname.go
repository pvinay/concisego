package greeting

import "fmt"

func Greet(name string) string {
    return fmt.Sprintf("Greetings, %s!!\n", name)
}

