package main

import (
	"fmt"
)

// DI container
type Godee struct {
	services map[string]any
}

func NewGodee() *Godee {
	return &Godee{
		services: make(map[string]any),
	}
}

func (d *Godee) Register(name string, service any) {
	d.services[name] = service
}

func Get[T any](d *Godee, name string) T {
	s, ok := d.services[name]
	if !ok {
		panic("service not found: " + name)
	}
	return s.(T)
}

// Example service
type UserService struct {
	Name string
}

func main() {
	di := NewGodee()

	di.Register("users", &UserService{Name: "John"})

	userService := Get[*UserService](di, "users")
	fmt.Println(userService.Name)
}
