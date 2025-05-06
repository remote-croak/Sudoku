// Prisoner is created when the specific prisoner number is selected from the holding cells
package sodoku

import (
	//"github.com/yohamta/donburi"
	"github.com/hajimehoshi/ebiten/v2"
	//"log"
	"os"
	"image/jpeg"


)


type Prisoner struct {
	cell *Cell
	image *ebiten.Image
	Processed bool
	onDrag bool
	ID int
	x int
	y int
	size int
	designation int // the number and colour a prison assigns to the unprocessed intake
	destination *Cell
	inHolding bool
	inTransfer bool
	inCellBlock bool
	originCell string
	destinationCell string
	prevX int
	prevY int
	// original location
	
}

func NewPrisoner(designation int)*Prisoner{
	//log.Printf("Prisoner sent to holding cell: %d", num)

	
	p := &Prisoner{
		Processed: false,
		onDrag: false,
		ID: 1,
		x: 0,
		y: 0,
		prevX: 0,
		prevY: 0,
		size: 48,
		designation: designation,
		inHolding: true, // default location being kept
		//intake or block or transfer
		inCellBlock: false,
		inTransfer: false,

	}

	p.image = ebiten.NewImage(p.size,p.size)
	p.setImg()

	return p
}

func (p *Prisoner) setDesignation(num int) {
	p.designation = num
	p.setImg()
}

func (p *Prisoner) getDesignation() int {
	return p.designation
}

// sets the current location the prisoner is being kept. Either in intake, a cellblock, a transfer
func (p *Prisoner) setLoc(cellLoc string){
	switch cellLoc {
		case "intake":
			p.inCellBlock = false
			p.inHolding = true
			p.inTransfer = false

		case "block":
			p.inCellBlock = true
			p.inHolding = false
			p.inTransfer = false

		case "Transferring":
			p.inCellBlock = false
			p.inHolding = false
			p.inTransfer = true

		default:
			p.inCellBlock = false
			p.inHolding = false
			p.inTransfer = false

	}
	
}

func (p *Prisoner) getLoc() string {
	var loc string

		if p.inHolding == true{
			loc = "intake"
		} else if p.inCellBlock == true {
			loc = "block" // TODO: include cell block designation as well.
		} else if p.inTransfer == true {
			loc = "Transferring"
		} else {
			loc = "Panic they are missing!"
		}

	return loc
}

func (p *Prisoner) assignPrisonerCell(hold *Cell){
	p.cell = hold
	p.setImg()

}

// resets prisoner location to last held cell. 
func (p *Prisoner) resetLocation() *Cell{
	x, y := p.cell.getWorldPos()
	p.movePrisoner(x,y)
	return p.cell
}


//TODO: make location path local to project not a system drive
// Needs to be a part of prisoner creation
func (p *Prisoner) prisonerNum() string{
	var img string
	switch p.designation {
	case 1:
		img = "F:/Projects/sodoku/sodoku/resources/num1.jpg"
	case 2:
		img = "F:/Projects/sodoku/sodoku/resources/num2.jpg"
	case 3:
		img = "F:/Projects/sodoku/sodoku/resources/num3.jpg"
	case 4:
		img = "F:/Projects/sodoku/sodoku/resources/num4.jpg"
	case 5:	
		img = "F:/Projects/sodoku/sodoku/resources/num5.jpg"
	case 6:	
		img = "F:/Projects/sodoku/sodoku/resources/num6.jpg"
	case 7:
		img = "F:/Projects/sodoku/sodoku/resources/num7.jpg"
	case 8:
		img = "F:/Projects/sodoku/sodoku/resources/num8.jpg"
	case 9:
		img = "F:/Projects/sodoku/sodoku/resources/num9.jpg"
	default:
		img = "F:/Projects/sodoku/sodoku/resources/emptyCell.jpg"
	}

	return img
}

func (p *Prisoner) setImg(){
	var err error

	imgLoc := p.prisonerNum()
	img, err := os.Open(imgLoc)

	if err != nil {
		return
	}

	defer img.Close()

	loadImage, err := jpeg.Decode(img)

	if err != nil {
		return
	}
	p.image = ebiten.NewImageFromImage(loadImage)

}

func (p *Prisoner) getImg() *ebiten.Image{
	
	return p.image

}

func (p *Prisoner) addPos(x int, y int){
	p.x = x
	p.y = y
	p.prevX, p.prevY = x, y
}

func (p *Prisoner) setPrevPos(){
	p.prevX, p.prevY = p.x, p.y
}

func (p *Prisoner) setStatus() {
	if p.inTransfer == true {
		p.inTransfer = false
	} else {
		p.inTransfer = true
	}
}

func (p *Prisoner) contains(x int, y int) bool{
	
	maxX := p.x + p.size
	maxY := p.y + p.size
	if (x > p.x && x < maxX) && ( y > p.y && y < maxY){
		return true
	} 
	
	return false
}

func (p *Prisoner) movePrisoner(x int, y int){
	p.x, p.y = x,y
}

func (p *Prisoner) Draw(screen *ebiten.Image){
	
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.1, 0.1)
	op.GeoM.Translate(float64(p.x), float64(p.y))

	screen.DrawImage(p.image, op)
	
}