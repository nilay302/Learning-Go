package main

import "fmt"

func main(){
	cards:= []string{newCard(),newCard()}
	cards = append(cards, "my heart")
	for i, card := range cards{
	      fmt.Println(i, card)
	}
}

func newCard() string{
	return "ace of spades"
}