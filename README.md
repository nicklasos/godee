# Godee

Prototype of DI container if Go

```
type UserService struct {
	Name string
}

func main() {
	di := NewGodee()

	di.Register("users", &UserService{Name: "John"})

	userService := Get[*UserService](di, "users")
	fmt.Println(userService.Name)
}
```