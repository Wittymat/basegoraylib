package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	//"math/rand"
	//"math"
	"fmt"
	//"time"
)


// ICI TU FOUS LES DONNEES DE TON JEU
type Game struct {
	Title string
	ScreenWidth int32
	ScreenHeight int32
	gameOver bool

}

// LA J'AI MIS DES VARIABLES POUR LES COULEURS
var (

	BGC = rl.Color{30,30,30,255}

)

// FONCTION MAIN TU TOUCHES A RIEN NORMALEMENT
func main() {
	fmt.Println("init")
	game := Game{}
	game.Init()

	rl.InitWindow(game.ScreenWidth, game.ScreenHeight, game.Title)
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		game.Update()
		game.Draw()
	}

	rl.CloseWindow()
}

func (g *Game) Init(){

	g.Title = "Base"
	g.ScreenWidth = 800
	g.ScreenHeight = 500

	g.gameOver = false

}

// FONCTION UPDATE POUR LA LOGIQUE DE TON JEU
func (g *Game) Update(){
	// update logic here
	g.inputs()
	// inputs
	if !g.gameOver {
	}
}

// FONCTION DRAW POUR AFFICHER TES TRUCS
func (g *Game) Draw(){
	rl.ClearBackground(BGC)
	rl.BeginDrawing()
	// draw here

	rl.EndDrawing()
}

// FONCTION D'INPUT FAITE MAISON
func (g *Game) inputs() {
	/*
	  KeySpace        = 32
    KeyRight        = 262
    KeyLeft         = 263
    KeyDown         = 264
    KeyUp           = 265 */


	if rl.IsKeyPressed(32){
		g.restart()
	}
}


// OSEF
func (g *Game) restart() {
	g.gameOver = false

}
