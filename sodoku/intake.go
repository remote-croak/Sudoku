package sodoku

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"strconv"
)

type Intake struct{

	holds [][]*Cell // collection of holding cells.

	destination *Cell

	x int
	y int

	intakeSize int

	screenX int
	screenY int

	designation int
	label string
}

func NewIntake() *Intake{

	intake := &Intake{
		x: 0,
		y: 0,
		intakeSize: 0,
		screenX: 0,
		screenY: 0,
		label: "",

	}

	intake.holds = intake.newCriminal()

	return intake
}

func (intake *Intake) newCriminal() [][]*Cell{

	holds := make([][]*Cell, 9)

	for cell := 0; cell < 9; cell++{

		holds[cell] = make([]*Cell, 10)

		// assign prisoners to a holding cell. Each holding cell, 9 in total, should have 9 prisoners.
		// Each prisoner should be selected and removed individually leaving the rest to remain in the the holding cell
		for prisonerNum := 0; prisonerNum <= 9; prisonerNum++{

			hold := NewCell("intake", true)
			prisoner := NewPrisoner(cell+1)
			hold.imprisonPrisoner(prisoner)
			prisoner.assignPrisonerCell(hold)
			label := "cell: " + strconv.Itoa(cell+1) + " prisoner: " + strconv.Itoa(prisonerNum + 1)
			hold.setLabel(label)
			holds[cell][prisonerNum] = hold
			// check if the corrent number of prisoners are being held in each hold.
			// create a check to see if the appropriate prisoners are being generated
			// prisoners should have a label like prisoner 1 cell 1
		}
	}

	return holds
}

func (intake *Intake) setLabel(label string) {

	intake.label = label
}

func (intake *Intake) getLabel() string {

	return intake.label
}

func (intake *Intake) AddPos(x int, y int){

	intake.screenX = x
	intake.screenY = y
}

func (intake *Intake) getPos() (int, int){

	return intake.x, intake.y
}

func (intake *Intake) getOriginPoints() (int, int){
	
	xOP, yOP := intake.holds[0][0].getOriginPoints()

	return xOP, yOP
}

func (intake *Intake) getEndPoints() (int, int){

	xEP, yEP := intake.holds[len(intake.holds)-1][0].getEndPoints()

	return xEP, yEP
}

func (intake *Intake) contains(x int, y int) bool{

	xOP, yOP := intake.getOriginPoints()
	xEP, yEP := intake.getEndPoints()

	if (xOP < x && xEP > x) && (yOP < y && yEP > y){

		return true
	}

	return false
}

// TODO: change x and y parameters into something more descriptive of mouseInput
func (intake *Intake) cellSearch(x int, y int) *Prisoner {
	log.Printf("Searching Intake")

	for cell := 0; cell < len(intake.holds); cell++ {

		log.Printf("Cell pos: %d,%d", x, y)
	
		if intake.holds[cell][0].contains(x,y) {
			if intake.holds[cell][0].hasPrisoner {
				log.Printf("has prisoner")
				log.Printf(intake.holds[cell][0].label )
				
				prisonerCell, nextPrisoner := intake.holds[cell][0], intake.holds[cell][1:]
				intake.holds[cell] = nextPrisoner
				prisoner := prisonerCell.getPrisoner()

				log.Printf("how many cells left: %d", len(intake.holds[cell]))

				if len(intake.holds[cell]) == 1 {
					intake.holds[cell][0].emptyCell()
				}
				
				return prisoner

			} else {
				
				log.Printf("doesn't have prisoner")
				log.Printf("Cell: %s", intake.holds[cell][0].label)
				intake.destination = intake.holds[cell][0]
			}
		}
	}

	return nil
}

func (intake *Intake) Draw(screen *ebiten.Image){

	for cell := 0; cell < len(intake.holds); cell++{
		prisoner := intake.holds[cell][0]
		cellX := (float64(intake.screenX) - float64(intake.screenY)*0.1) + float64(cell * (prisoner.getSize() + 24))
		cellY := float64(intake.screenY + prisoner.getSize())*0.5
		prisoner.setWorldPos(int(cellX), int(cellY))
		prisoner.setDrawPos(int(cellX), int(cellY))
		prisoner.Draw(screen)
		
	}
}