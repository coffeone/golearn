package animals

import "fmt"

func CatFood(food string) string {
	return fmt.Sprintf("cat eating {%s}", food)
}
