package animals

import "fmt"

func CatFood(food string) string {
	return fmt.Sprintf("Cat eating {%s}", food)
}
