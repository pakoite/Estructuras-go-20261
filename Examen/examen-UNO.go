package main

import (
    "fmt"
    "math/rand"
)

type Card struct {
    Color string
    Value string
}

type Player struct {
    Name string
    Hand []Card
}

type Game struct {
    Deck        []Card
    Players     []Player
    DiscardPile []Card
    Turn        int
}

func (g *Game) InitDeck() {
    colors := []string{"Rojo", "Verde", "Azul", "Amarillo"}
    values := []string{"0","1","2","3","4","5","6","7","8","9"}

    for _, c := range colors {
        for _, v := range values {
            g.Deck = append(g.Deck, Card{Color: c, Value: v})
        }
    }

    rand.Shuffle(len(g.Deck), func(i, j int) {
        g.Deck[i], g.Deck[j] = g.Deck[j], g.Deck[i]
    })
}

func (g *Game) DealCards() {
    // TODO: repartir 5 cartas a cada jugador
}

func (g *Game) DrawCard() Card {
    // TODO: tomar carta del deck
    return Card{}
}

func (g *Game) PlayTurn() {
    player := &g.Players[g.Turn]

    fmt.Println("Turno de:", player.Name)

    // TODO:
    // - mostrar cartas
    // - validar carta jugable (mismo color o número)
    // - permitir robar si no puede jugar
}

func (g *Game) NextTurn() {
    g.Turn = (g.Turn + 1) % len(g.Players)
}

func main() {
    game := Game{
        Players: []Player{
            {Name: "Jugador 1"},
            {Name: "Jugador 2"},
        },
    }

    game.InitDeck()
    game.DealCards()

    for {
        game.PlayTurn()
        game.NextTurn()
    }
}