package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Contender interface {
	ThrowAttack() int
	RecieveAttack(intensity int)
	IsAlive() bool
	GetName() string
	GetLife() int
}

type Police struct {
	Life   int // 0..100
	Armour int // 0..50
}

func (p Police) ThrowAttack() int {
	return rand.Intn(4)
}

func (p *Police) RecieveAttack(intensity int) {
	if p.Armour > 0 {
		diff := (p.Armour - intensity)
		hasArmour := diff > 0
		if hasArmour {
			p.Armour -= intensity
			intensity = 0
		} else {
			p.Armour = 0
			intensity = -diff // intensity -= p.Armour
		}
	}
	p.Life -= intensity
}

func (p Police) GetName() string {
	return "Policia"
}

func (p Police) IsAlive() bool {
	return p.Life > 0
}

func (p Police) GetLife() int {
	return p.Life
}

type Criminal struct {
	Life int // 0..100
}

func (c *Criminal) ThrowAttack() int {
	return rand.Intn(8)
}

func (c *Criminal) RecieveAttack(intensity int) {
	c.Life -= intensity
}

func (c *Criminal) GetName() string {
	return "Criminal"
}

func (c Criminal) IsAlive() bool {
	return c.Life > 0
}

func (c Criminal) GetLife() int {
	return c.Life
}

func main() {
	var police Police = Police{
		Life:   10,
		Armour: 5,
	}
	var criminal Criminal = Criminal{
		Life: 15,
	}

	randomValueBetweenOneAndZero := rand.Intn(2)
	policeHitFirst := randomValueBetweenOneAndZero == 1

	var areBothAlive = police.IsAlive() && criminal.IsAlive()
	turns := 0

	for areBothAlive {
		if policeHitFirst {
			intesity := police.ThrowAttack()
			fmt.Println("Policia tira golpe con intensidad =", intesity)
			criminal.RecieveAttack(intesity)

			if criminal.IsAlive() {
				intesity := criminal.ThrowAttack()
				fmt.Println("Criminal tira golpe con intensidad =", intesity)
				police.RecieveAttack(intesity)
			}
		} else {
			intesity := criminal.ThrowAttack()
			fmt.Println("Criminal tira golpe con intensidad =", intesity)
			police.RecieveAttack(intesity)

			if police.IsAlive() {
				intesity := police.ThrowAttack()
				fmt.Println("Policia tira golpe con intensidad =", intesity)
				criminal.RecieveAttack(intesity)
			}
		}

		turns++
		fmt.Printf("PoliceLife = %d, CriminalLife = %d\nFin de turno %d\n", police.Life, criminal.Life, turns)
		areBothAlive = police.IsAlive() && criminal.IsAlive()
		time.Sleep(3 * time.Second)
	}
}
