package sodoku

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"log"
)

type Input struct{
	//mouse inputs
	mousePressed bool
	mouseReleased bool
	mouseX int
	mouseY int
	offsetX int
	offsetY int
	prisonerHeld *Prisoner

}

func NewInput() *Input{
	return &Input{}
}

func (input *Input) grabPrisoner(prisoner *Prisoner){
	input.prisonerHeld = prisoner

}

func (input *Input) releasePrisoner() *Prisoner{

	return input.prisonerHeld
}

func (input *Input) clearPrisoner(){
	input.prisonerHeld = nil
}

// offsets image of the held prisoner's position to have the mouse cursor
// in the middle of the image
func(input *Input) prisonerPosition() (int, int){
	// based on the image from the prisoner.
	// doesn't need to be a hardcoded value
	// just saves computation
	x := input.mouseX - 24
	y := input.mouseY - 24

	return x,y 
}



func (input *Input) Update(){

	input.mouseX, input.mouseY = ebiten.CursorPosition()
	

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft){

		log.Printf("Left Mouse Button Pressed")
		input.mousePressed = true
		input.cellExists(input.mouseX, input.mouseY)
	} else if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft){
		log.Printf("Left Mouse Button Released")
		input.mousePressed = false
	}
}

// input file needs to deal with the moving of prisoners with the mouse
// undo and redo
// and information display.
// but how to do this?
func (input *Input) cellExists(x, y int) {

	// if position is over any part of the prison it should check the prisonf
	// if position is over any part of the intake cells it should check the intake
	// how best to do this when dealing with an individual input file.
	//exists := g.prison.SearchBlock(x,y)

	// if exists == nil{
	// 	log.Printf("Doesn't Exist")
	// } else {
	// 	label := exists.label
	// 	log.Printf("Does it Exist? %s\n", label)
	// }

}
