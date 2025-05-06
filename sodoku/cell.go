package sodoku

import (
	"github.com/hajimehoshi/ebiten/v2"
	//"image/jpeg"
	//"os"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
//	"fmt"
	"log"
	//"strconv"
)


//TODO: Refactor so that all functions have the same local call of cell.

type Cell struct {
	// State refactor
	// hasPrisoner State
	// vacant state
	// occupied state
	// transfer state
	// current state

	vacant State
	occupied State
	current State



	prisoner *Prisoner

	image *ebiten.Image
	size int
	margin int
	test State // Can be in a filled or empty state
	// location Figure out the necessary variables
	// x and y coordinates for drawing purposes
	x int
	y int

	// relative positon to block
	relX int
	relY int

	// a world position that represents the cell's place in the game screen.
	worldPosX int
	worldPosY int

	label string
	hasPrisoner bool // remove and replace by appropriate state
	location string

	singleDigitValue int
	existsInRow int
	existsInCol int
	prisonerID int
//	cellID

}


func NewCell(loc string, hasPrisoner bool) *Cell{

	//log.Printf("Create Cell")

	cell := &Cell{
		size: 48,
		margin: 4,
		relX: 0,
		relY: 0, 
		worldPosX: 0,
		worldPosY: 0,
		hasPrisoner: hasPrisoner,
		location: loc,
		prisonerID: 0,
	}

	vacantCellState := &VacantCell{
		cell: cell,
	}
	occupiedCellState := &OccupiedCell{
		cell: cell,
	}

	cell.vacant = vacantCellState
	cell.occupied = occupiedCellState
	cell.setState(cell.vacant)

	cell.image = ebiten.NewImage(cell.size, cell.size)
	
	return cell
}

func (cell *Cell) info(sdv int, row int, col int){
	cell.singleDigitValue = sdv
	cell.existsInRow = row
	cell.existsInCol = col

}

func (cell *Cell) imprisonPrisoner(criminal *Prisoner) error {
	cell.hasPrisoner = true
	return cell.current.imprisonPrisoner(criminal)
}

func (cell *Cell) emptyCell() error {
	cell.hasPrisoner = false
	return cell.current.emptyCell()
}

func (cell *Cell) releasePrisoner() error {
	cell.hasPrisoner = false
	return cell.current.releasePrisoner()
}

func (cell *Cell) setState(state State){
	cell.current = state
}

func (cell *Cell) getLocation() string{
	return cell.location
}

func (cell *Cell) setLocation(loc string){
	cell.location = loc
}


func (cell *Cell) setSize(size int){

	cell.size = size
}

func (cell *Cell) getSize() int{

	return cell.size
}
func (cell *Cell) setDrawPos(x int, y int){
	
	cell.x = x
	cell.y = y

}
func (cell *Cell) setRelPos(x int, y int){
	
	cell.relX = x
	cell.relY = y
}

func (cell *Cell) getRelPos() (int, int){
	
	return cell.relX, cell.relY
}

func (cell *Cell) setWorldPos(x int, y int){
	
	cell.worldPosX = cell.relX + x
	cell.worldPosY = cell.relY + y
}

func (cell *Cell) getWorldPos() (int, int){
	
	return cell.worldPosX, cell.worldPosY
}
func (cell *Cell) getOriginPoints() (int, int){

	xOP, yOP := cell.x, cell.y 

	return xOP, yOP
}

func (cell *Cell) getEndPoints() (int, int){

	xEP, yEP := cell.x + cell.size, cell.y + cell.size

	return xEP, yEP 
}

func (cell *Cell) setLabel(label string){
	
	cell.label = label
}

func (cell *Cell) placePrisoner(prisoner *Prisoner){
	cell.prisoner = prisoner
}

func (cell *Cell) addPrisoner(){
	
	cell.hasPrisoner = true
}


func(cell *Cell) removePrisoner(){
	cell.hasPrisoner = false
	//c.prisoner.setDesignation(0)
}

func (cell *Cell) emptyHold(){
	cell.hasPrisoner = false

}

func (cell *Cell) getPrisoner() *Prisoner{
	
	// c.hasPrisoner = false
	// if prisoner is transfered then the prisoner data needs to be come empty.
	return cell.prisoner

}

func (cell *Cell) checkPrisoner() bool{
	
	return cell.hasPrisoner
}

func (cell *Cell) getImg() *ebiten.Image{

	return cell.image
}

func (cell *Cell) contains(x int, y int) bool {
	
	xOP, yOP := cell.worldPosX, cell.worldPosY
	xEP, yEP := cell.worldPosX + cell.size, cell.worldPosY + cell.size


	if (xOP < x && xEP > x) && (yOP < y && yEP > y) {
		log.Printf("cell origin: %d, %d\n mouse origin: %d, %d\n cell end: %d, %d", xOP, yOP, x, y, xEP, yEP)
		if cell.hasPrisoner == true{
			log.Printf("There be a prisoner in here")
			return true
		} else if !cell.hasPrisoner {
				log.Printf("Your prisoner is in another cell")
				return true
			}
		}
	
	return false
}

func (cell *Cell) Draw(cellImage *ebiten.Image) {
	
	op := &ebiten.DrawImageOptions{}

	
	cell.image = cell.prisoner.getImg()
	op.GeoM.Scale(0.1, 0.1)

	op.GeoM.Translate(float64(cell.x), float64(cell.y))
	cellImage.DrawImage(cell.image, op)
}

