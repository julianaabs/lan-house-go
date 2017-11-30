package main

import(
	"fmt"
	"time"
	"math/rand"
)

var (
	ch = make(chan *Adolescente)
)

type Adolescente struct{
	name string
}

func randomName() int{
	s1:= rand.NewSource(time.Now().UnixNano())
	r1:= rand.New(s1)

	result:= r1.Intn(25) + 65

	return result
}

func randomTime() int{
	s1:= rand.NewSource(time.Now().UnixNano())
	r1:= rand.New(s1)

	result:= r1.Intn(53) + 7 

	return result
}

func (this *Adolescente) intervalOnline(){
	fmt.Println(this.name, "está online.")
	intervalo:= randomTime()
	time.Sleep(time.Duration(intervalo/2) * time.Second)
	fmt.Println("liberou a máquina após passar", intervalo, "minutos on-line.")
	ch <- this
}

func NewAdolescente() *Adolescente{
	return &Adolescente{name: string(randomName())}
}

func runTeen(){

	for i:= 0; i<26; i++ {
		NewAdolescente()
	}

	for{
		select {
			case adolescente:= <- ch:
				NewAdolescente()
				fmt.Println(adolescente.name, "está aguardando.")
		}
	}

}

func main(){
	go runTeen()
	for{

	}
}