package main

import(
  "fmt"
  "draftalyzer/draft"
)

var heroIds = []uint{16, 20, 10, 96, 83} // simulated input
var heroes []draft.Hero

func main(){
  for i := range heroIds{
    hero := draft.Hero{heroIds[i], draft.GetHeroNameFromId(heroIds[i])}
    heroes = append(heroes, hero)
  }
  fmt.Println(heroes)
}
