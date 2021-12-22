package main

import "fmt"

type Music interface {
	playMusic()
}
type Video interface {
	playVideo()
}

type Mobile struct {
}

func (m Mobile) playMusic() {

	fmt.Println("play music")
}

func (m Mobile) playVideo() {
	fmt.Println("play video")
}

type Computer struct {
}

type Pet interface {
	eat()
}

type Dog struct {
}

type Cat struct {
}

func (dog Dog) eat() {
	fmt.Println("dog eat>>>")
}
func (cat Cat) eat() {
	fmt.Println("cat eat>>>")
}

func main() {
	m := Mobile{}
	m.playMusic()
	m.playVideo()

	var pet Pet
	pet = Dog{}
	pet.eat()
	pet = Cat{}
	pet.eat()
}
