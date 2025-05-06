package sodoku

import(
	
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"fmt"
	"log"
	
)

type Prison struct {
	screen *ebiten.Image
	blocks []*Block
	intake *Intake
	input *Input
	prisoner *Prisoner

	cell *Cell
	cellImage *ebiten.Image
	cellSize int
	blockMult int

	blockSize int
	xEP int
	yEP int

	margin int
	xPos int
	yPos int
	departure *Cell
	destination *Cell
	destinationBlock *Block
	cellDATA string
	blockDATA string
	wardenLoc string
	id string



}

func NewPrison(input *Input) (*Prison, error) {
	//blocks := []*Block{}
	
	p := &Prison{
		cellSize: 48,
		blockMult: 3,
		margin: 4,
		xPos: 0,
		yPos: 0,
		input: input,
		cellDATA: "cell",
		blockDATA: "block",
		wardenLoc: "wandering",
	}



	p.intake = NewIntake()

	p.blockSize = p.cellSize*p.blockMult + (p.blockMult+1) * p.margin
	blockRow := 0
	for col := 0; col < p.blockMult; col++{
		for row := 0; row < p.blockMult; row++ {

			//nBlock := 3 * col + row + 1 //debug label Block Num
		//	p.blockLabel(calcSDV(row, col)) //block label should be in the block constructor
			singleDigitValue := calcSDV(row, col)
			p.blockLabel(singleDigitValue)

			if singleDigitValue == 1{
				blockRow = singleDigitValue
			} else if singleDigitValue == 4{
				blockRow = singleDigitValue
			} else if singleDigitValue == 7{
				blockRow = singleDigitValue
			}

			x := row * p.blockSize + (row + 1) * p.margin
			y := col * p.blockSize + (col + 1) * p.margin
			b := NewBlock(x, y, p.id)
			b.info(singleDigitValue, blockRow, row + 1)
			p.blocks = append(p.blocks, b)
			//TODO b.blockLabel(row,col) // this is where the block's label is generated
			//log.Printf("Block Num: %d, x: %d, y: %d, blockSize: %d", nBlock, x, y, p.blockSize)
		}
	}


	for test := 0; test < len(p.blocks); test++{
		log.Printf("Display Block ID\n", p.blocks[test].blockID)
		for cellTest := 0; cellTest < len(p.blocks[test].cells); cellTest++{
			log.Printf("Display cell label\n", p.blocks[test].cells[cellTest].label)
		}
	}
	
	return p, nil
}

func (p *Prison) Update(){

	if p.prisoner == nil{
		p.prisonerSearch()
		if p.prisoner !=nil{

			p.prisoner.setPrevPos()
		}
	}

	if p.prisoner != nil{
		// if left mouse button released then revert prisoner to it's original location
		if !p.input.mousePressed{
			
			
			// Retrieve a cell without a prisoner by searcing prison block 
			if(p.prisoner.getLoc() == "intake" && p.blockContains(p.input.mouseX, p.input.mouseY)){
				if (p.blockSearch() == nil){

					// Checks if prison is valid but doesn't stop placement.
					// Must give a slight visible error warning when a conflict occurs
					// but shouldn't make it impossible to place a prisoner in a wrong spot.
					// let the player figure it out when needed.
					log.Printf("Before Validation Value: %d", p.destinationBlock.singleDigitValue)
					p.validatePrison(p.destinationBlock, p.destination, p.prisoner.getDesignation())
					
					p.destination.imprisonPrisoner(p.prisoner)
					p.prisoner.setLoc("block")
					//p.destination.addPrisoner()

					//log.Printf("current prisoner number: %d", p.destination.prisoner.getDesignation())
					p.prisoner.assignPrisonerCell(p.destination)
					p.prisoner.movePrisoner(p.input.mouseX, p.input.mouseY)
					
					p.prisoner = nil
					p.input.clearPrisoner()
					log.Printf("Transfer Complete")

				} 
			
			} else if (p.prisoner.getLoc() == "block" && p.intake.contains(p.input.mouseX, p.input.mouseY)){
					if (p.intake.cellSearch(p.input.mouseX, p.input.mouseY) == nil){
						p.intake.destination.imprisonPrisoner(p.prisoner)
						p.prisoner.setLoc("intake")
						p.prisoner.assignPrisonerCell(p.intake.destination)
						p.prisoner.movePrisoner(p.input.mouseX, p.input.mouseY)
						p.prisoner = nil
						p.input.clearPrisoner()
						log.Printf("Transfer Complete")
					}

			} else if p.prisoner.getLoc() == "Transferring"{
		
				prevCell := p.prisoner.resetLocation()
				prevCell.imprisonPrisoner(p.prisoner)
				prevCell.addPrisoner()

				p.prisoner = nil
			}

		} else {
			log.Printf("Transferring Prisoner")
			
			x, y := p.input.prisonerPosition()
			p.prisoner.movePrisoner(x,y)
		}	
	}
}


// TODO find parameter name that is not c but relates to prison cells.
func (p *Prison) prisonerSearch(){
	inputMX, inputMY := p.input.mouseX, p.input.mouseY
	if p.intake.contains(inputMX, inputMY){
		p.wardenLoc = "Warden is searching prison intake"
		// check if prisoner exists at location on button click
		if (p.input.mousePressed){

			p.prisoner = p.intake.cellSearch(inputMX, inputMY)
			//p.destination = p.prisoner.getCellData()
			p.input.grabPrisoner(p.prisoner)	
		}

		//p.prisoner = cells.prisoner
	} else if p.blockContains(inputMX, inputMY){
		p.wardenLoc = "Warden is searching prison blocks"

		if (p.input.mousePressed){
			p.prisoner = p.blockSearch()
			p.input.grabPrisoner(p.prisoner)

			if p.prisoner != nil{
				log.Printf("Prisoner: %d", p.prisoner.designation)
			}
			
		}



	} else {
		p.wardenLoc = "Warden is wandering"
	}

}

func (p *Prison) blockSearch() *Prisoner{
	inputMX, inputMY := p.input.mouseX, p.input.mouseY
 
	for block := 0; block < len(p.blocks); block++{
		curBlock := p.blocks[block]

		if curBlock.contains(inputMX, inputMY){
			for cell := 0; cell < len(curBlock.cells); cell++{
				curCell := curBlock.cells[cell]

				//log.Printf("mouse position: %d, %d\n cell position: %d, %d", inputMX, inputMY, curCell.worldPosX, curCell.worldPosY)

				log.Printf("searching block for prisoner")
				if curCell.contains(inputMX, inputMY){
					// if cell doesn't have a prisoner then
					if !curCell.hasPrisoner{
						//assign current cell as the prisoner's destination
						p.destination = curCell
						log.Printf("curBlock: %d", curBlock.singleDigitValue)
						p.destinationBlock = curBlock
					// 	//check if prisoner exists and grab them.
						log.Printf("prisoner")
					} else if curCell.hasPrisoner{
						log.Printf("Prisoner! I'm here to free you!")
						prisoner := curCell.getPrisoner()
						curCell.releasePrisoner()
						return prisoner
					}
				}
			} 
		}
	}

	return nil
}

// area the prison blocks cover in the prison
func (p *Prison) blockContains(x int, y int) bool{

	if (p.xPos < x && p.xEP > x) && (p.yPos < y && p.yEP > y){
		return true
	}

	return false
}

func (prison *Prison) validatePrison(target *Block, targetCell *Cell, prisonerID int) bool{

	// A target block provides three pieces of information.
	// It's current position in the array represented as a singleDigitValue
	// the row it exists in
	// the column it exists in.
	targetBlock := target.singleDigitValue // this represents the location that all other blocks are evaluated around
	//pivotBlock := target.
	targetRow := target.existsInRow // this represents the value of the first position in a row (1, 4, 7)
	targetCol := target.existsInCol // this represents the value of the top position in a column (1, 2, 3)
	
	// Check target block if placement is valid
	log.Printf("Start Validation")
	log.Printf("Single Digit Value: %d", prison.blocks[targetBlock - 1].singleDigitValue)
	//count := 0

	for cellCheck := 0; cellCheck < len(target.cells); cellCheck++{
		if target.cells[cellCheck].prisonerID == prisonerID{
			//count++
			//log.Printf("matching designations: %d", count)
			log.Printf("There's a matching prisoner")
			return false
		}
	}

	log.Printf("Target Row: %d", targetRow)
	log.Printf("Target Col: %d", targetCol)
	for count := 0; count < 3; count++{
		log.Printf("Target Row: %d", (targetRow - 1  + count))
		nextTarget := prison.blocks[targetRow - 1 + count]
		//prison.conflictInBlock(nextTarget)

		for cellCheck := 0; cellCheck < len(nextTarget.cells); cellCheck++{
			if nextTarget.cells[cellCheck].prisonerID == prisonerID{
				prison.horizontalConflict(nextTarget,targetCell.existsInRow, prisonerID)
				//count++
				//log.Printf("matching designations: %d", count)
				log.Printf("There's a matching prisoner in the row")
				return false
			}
		}

		nextTarget = prison.blocks[targetCol - 1 + count]
		for cellCheck := 0; cellCheck < len(nextTarget.cells); cellCheck++{
			if nextTarget.cells[cellCheck].prisonerID == prisonerID{
				prison.verticalConflict(nextTarget, targetCell.existsInCol, prisonerID)
				//count++
				//log.Printf("matching designations: %d", count)
				log.Printf("There's a matching prisoner in the column")
				return false
			}
		}

		
		
	}
	

	log.Printf("End Validation")
	return true
}

// func (prison *Prison) conflictInBlock(targetBlock *Block) bool{

// 	for count := 0; count < 3; count++{
// 		for cellCheck := 0; cellCheck < len(targetBlock.cells); cellCheck++{
// 			if targetBlock.cells[cellCheck].prisonerID == prisonerID{
// 				//count++
// 				//log.Printf("matching designations: %d", count)
// 				log.Printf("There's a matching prisoner")
// 				return false
// 			}
// 		}
// 			return true
// 		//col;
// 			return true
// 	}
// 	return false
// }

// Two seperate functions so an extra parameter doesn't have to be added to control
// the direction of the search. 
// Also lets it remain grouped to the direction searched for blocks

// checking each row within a single column for a prisoner number that conflicts
func (prison *Prison) verticalConflict(targetBlock *Block, cellInCol int, prisonerID int) bool{
	for count := 0; count < 3; count++{
		checkCell := 3 * count + cellInCol
		if targetBlock.cells[checkCell].prisonerID == prisonerID {
			log.Printf("This isn't the column cell you're looking for")
			return true
		}
	}

	return false
}

// checking each column along a single row for a prisoner number that conflicts
func (prison *Prison) horizontalConflict(targetBlock *Block, cellInRow int, prisonerID int) bool{
	for count := 0; count < 3; count++{
		checkCell := cellInRow + count
		if targetBlock.cells[checkCell].prisonerID == prisonerID {
			log.Printf("This isn't the row cell you're looking for")
			return true
		}
	}
	
	 return false
}

// calculates Single digit Value(1-9) from two given numbers (row and col)
func calcSDV(row int, col int) int {
	value := 3*col + row + 1
	return value
}
// determines the label(A to I) for a specific block based on the single digit value parameter
// parameters: sdv = single digit value taken in
func (p *Prison) blockLabel (sdv int) {
	switch sdv {
	case 1:
		p.id = "A"
	case 2:
		p.id = "B"
	case 3:
		p.id = "C"
	case 4:
		p.id = "D"
	case 5:
		p.id = "E"
	case 6:
		p.id = "F"
	case 7:
		p.id = "G"
	case 8:
		p.id = "H"
	case 9:
		p.id = "I"
	default:
		p.id = "Z"
	}
}

func (p *Prison) AddPos(x int, y int){

	p.xPos = x
	p.yPos = y
	p.intake.AddPos(x, y)
	p.setEndPoints()
	
}

func (p *Prison) setEndPoints(){

	p.xEP = p.xPos + ((p.blockSize + 4) * 3)
	p.yEP = p.yPos + ((p.blockSize + 4) * 3)

}

func (p *Prison) ShowPosX() int{

	return p.xPos
}
func (p *Prison) ShowPosY() int{

	return p.yPos
}

//TODO: Calculate size of Block Image using Cell Image dimensions
// Figure out how to seperate Cell Image Draw into it's own struct
func (p *Prison) Draw(prisonImage *ebiten.Image, screen *ebiten.Image ){
	 p.screen = screen


	//blockSize := p.cellSize*p.blockMult + (p.blockMult+1) * p.margin	
	 for b := 0; b < len(p.blocks); b++{
		p.blocks[b].LocInPrison(p.xPos, p.yPos)
		p.blocks[b].Draw(prisonImage)
		
	 }

	//  if p.input.mousePressed{
	// 	p.prisoner.Draw(screen)
	//  }

	 p.intake.Draw(screen)
	 ebitenutil.DebugPrintAt(screen, fmt.Sprintf("cursor at: %d, %d", p.input.mouseX, p.input.mouseY), 0, 500)
	 ebitenutil.DebugPrintAt(screen, fmt.Sprintf("cursor at: %s, %s", p.cellDATA, p.blockDATA), 0, 700)
	 inXOP, inYOP := p.intake.getOriginPoints()
	 inXEP, inYEP := p.intake.getEndPoints()
	 ebitenutil.DebugPrintAt(screen, fmt.Sprintf("dimensions of intake\n origin: %d, %d\n end: %d, %d", inXOP, inYOP, inXEP, inYEP), 0, 800)
	 ebitenutil.DebugPrintAt(screen, fmt.Sprintf("warden is in: %s", p.wardenLoc), 0, 900)
	 if (p.prisoner != nil){
	 	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("prisoner is grabbed: %s", p.prisoner.inTransfer), 0, 950)
	} else {
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf(""), 0, 950)
	}

}
