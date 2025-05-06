// This should be the game manager! It contains all relevent global data, variable, action for use in other files.
// checking if an end state is reached
// transfering game states
// loading generation functions.
// loading levels or scenes.

package sodoku

import (
	// Github Repo's
	// Import Game Engine Ebiten
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	//"github.com/hajimehoshi/ebiten/v2/inpututil"

	// Import ECS Donburi for Ebiten Engine
	//"github.com/yohamta/donburi"
	"fmt"
	//"log"
)

const (
prisonSize = 3
)

// TODO: Restructure variable positioning
type Game struct {

	prison *Prison
	prisonImage *ebiten.Image

	cell *Cell

	prisoner *Prisoner
	prisonerImage *ebiten.Image

	width int
	height int
	cellSize int
	input *Input
	
}



func NewGame(screenWidth int, screenHeight int) (*Game, error) {
	
	g := &Game{
		
		width: screenWidth,
		height: screenHeight,
		cellSize: 48,
		//intake: prisonerIntake(),
		input: NewInput(),
	}

	//g.intake = NewIntake()

	var err error
	//world := donburi.NewWorld()
	//pris := world.Create(Position)
	

	g.prison, err = NewPrison(g.input)
	//entry := world.Entry(pris)

	if err != nil {
		return nil, err
	}

	return g, nil
}

func (g *Game) Update() error {

	g.input.Update() // this updates input to be used in other files.
	g.prison.Update()
	// how to avoid the double prisonerSearch input and move the prisoner?

	 
// how to make it so I can search both the intake and the the prison for a prisoner transfer?
	//intake search


	

	return nil
}
// REFACTOR: This doesn't need to be here
// if mouse cursor is within the bounds of a cell and a block. display data
// func (g *Game) cellExists(x int, y int){
	
// 	exists := g.prison.SearchBlock(x,y)

// 	if exists == nil{
// 		log.Printf("Doesn't Exist")
// 	} else {
// 		label := exists.label
// 		log.Printf("Does it Exist? %s\n", label)
// 	}
// }

func (g *Game) Layout(outsideWidth, outsideHeight int) (int,int){
	return 1920, 1080
	}


func (g *Game) Draw(screen *ebiten.Image) {
	
	screen.Fill(backgroundColour)

	prisoner := g.input.releasePrisoner()

	ebitenutil.DebugPrintAt(screen, "Sodoku", 0,1000)
	if g.prisonImage == nil {
		//TODO: make dimensions reflect the size of the Blocks x 3
		g.prisonImage = ebiten.NewImage(512,512)
	}

	

	op := &ebiten.DrawImageOptions{}
	sw, sh := screen.Bounds().Dx(), screen.Bounds().Dy()
	bw, bh := g.prisonImage.Bounds().Dx(), g.prisonImage.Bounds().Dy()
	x := (sw - bw)/2
	y := (sh - bh)/2


	// Load Unprocessed "Prisoners"
	// [1][2][3][4][5][6][7][8][9] 
	// this is just a series of images that the warden can click and drag from
	// onClick a prisoner is generated that is then dragged to where the warden chooses.
	// need to figure out how to store newly generated prisoners so that when they return
	// to the hold they are kept instead of destroyed.
	// This may be an issue object pooling can solve. Along with an identification
	// variable stored in the Prisoner struct.
	// up = unprocess prisoner from stack
	// all 9 in each hold do not have to be drawn! Only  the first image in the array.
	// can also limit the compute needed if only the top and the next prisoners need to be created 
	// for each hold until the hold is "empty"
	// p.unprocessedHold1.draw(g.prisoner)
	

	g.prison.AddPos(x, y)
	g.prison.Draw(g.prisonImage, screen)
	
	
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Prison Origin from Prison: %d,%d", g.prison.ShowPosX(), g.prison.ShowPosY()), 0, 400)
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(g.prisonImage, op)

	if g.input.mousePressed && prisoner != nil{
		prisoner.Draw(screen)
	}
	//ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Prison Origin from World: %d,%d", sw, sh), 0, 500)	
}
