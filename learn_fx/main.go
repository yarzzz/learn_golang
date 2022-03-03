package main

import (
	"context"
	"fmt"

	"go.uber.org/fx"
)

type A struct {
	Name string
}

func NewA() A {
	return A{Name: "a"}
}

type B struct {
	Name string
}

func NewB(a A) B {
	return B{Name: a.Name + "b"}
}

type C struct {
	fx.In

	XA A
	YB B
}

type D struct {
	X string
	Y string
	Z string
}

func NewNewD(x string) func(C) D {
	return func(c C) D {
		return D{
			X: c.XA.Name,
			Y: c.YB.Name,
			Z: x,
		}
	}
}

func main() {
	var d D
	app := fx.New(
		fx.Provide(NewA),
		fx.Provide(NewB),
		fx.Provide(NewNewD("123")),
		fx.Populate(&d),
	)
	err := app.Start(context.Background())
	if err != nil {
		panic(err)
	}
	err = app.Stop(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(d.X)
	fmt.Println(d.Y)
	fmt.Println(d.Z)
}
