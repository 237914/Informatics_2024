package Lab6

import (
 "fmt"
)

type Pig struct {
 Name string
 Age int
 Weight float64
}

func NewPig(name string, age int, weight float64) *Pig {
 return &Pig{
  Name: name,
  Age: age,
  Weight: weight,
 }
}

func (p *Pig) Speak() {
 fmt.Printf("%s говорит: Хрю-хрю!\n", p.Name)
}

func (p *Pig) Eat(foodWeight float64) {
 p.Weight += foodWeight
 fmt.Printf("%s покушала и теперь весит %.2f кг\n", p.Name, p.Weight)
}

func (p *Pig) GrowOlder() {
 p.Age++
 fmt.Printf("%s теперь на год старше, ей %d года\n", p.Name, p.Age)
}

func RunLab6() {
 myPig := NewPig("Пеппа", 2, 15.0)
 myPig.Speak()
 myPig.Eat(2.5)
 myPig.GrowOlder()
}