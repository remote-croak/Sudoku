package sodoku

import (
	"github.com/hajimehoshi/ebiten/v2"
	//"log"
	"strconv"
)

type Block struct {
	cells []*Cell
	prisoners []int
	image *ebiten.Image
	size int
	margin int
	mult int
	x int
	y int
	label string
	pX int
	pY int
	blockSize int
	blockID string
	singleDigitValue int
	existsInRow int
	existsInCol int
}
//*************************************
// Parameters:	size - int value dictating the size of an individual cell that can be extrapolated for greater sizes in regards to blocks and prisons
//		margin - int value representing space between cells
//		mult - int value giving the number of cells in a particular block
//		x - int value given for initial x positioning of a cell
//		y - int value given for initial y positioning of a cell
//		blockLabel - string value given for identifying a particular block in the prison
//*************************************
func NewBlock(x int, y int, blockID string) *Block {

	//prisoners := []int{1,2,3,4,5,6,7,8,9}

	b := &Block{
		
		size: 48,
		margin: 4,
		mult: 3,
		x: x, // local init x pos
		y: y, // local init y pos
		pX: 0,
		pY: 0,
		blockSize: 0,
		blockID: blockID,
	}
	// to write a label outside of the for loop I need to include the calculation that
	// requires the for loop to exist and to place the calculation in the correct
	// position in the slice.
	// this is relatively annoying and complex when label creation can be done within
	// the file and structure. All I need to do is pass the calculation to the object

	b.blockSize = b.size*b.mult + (b.mult + 1) * b.margin
	b.image = ebiten.NewImage(b.blockSize, b.blockSize)
	for col := 0; col < b.mult; col++{
		for row := 0; row < b.mult; row++{

			sdv := 3 * col + row + 1
			cellLabel := blockID + strconv.Itoa(sdv)
			cellX := row * b.size + (row + 1) * b.margin
			cellY := col * b.size + (col + 1) * b.margin
			c := NewCell("block", false)
			c.prisoner = NewPrisoner(0)
			c.setRelPos(cellX, cellY)
			c.setLabel(cellLabel)
			c.info(sdv, row, col)
			b.cells = append(b.cells, c)
		}
	}

	return b
}

func (block *Block) info(sdv int, row int, col int){
	block.singleDigitValue = sdv
	block.existsInRow = row
	block.existsInCol = col
}

func (b *Block) LocInPrison(x int, y int){
	b.pX = b.x + x
	b.pY = b.y + y
}

// mouse in block

// block contains given coordinates
func (b *Block) contains(x int, y int) bool {
	
	maxX := b.pX + b.blockSize
	maxY := b.pY + b.blockSize
	bY := b.pY
	bX := b.pX

	if (bX < x && maxX > x) && (bY < y && maxY > y) {
		return true
	}

	return false
}

func (b *Block) getLabel() string {
	return b.label
}

func (b *Block) setPos(x int, y int) {
	b.x = x
	b.y = y
}

func (b *Block) getPos() (int, int){
	return b.x, b.y
}

func (b *Block) getOriginPoints() (int, int){
	xOP, yOP := b.cells[0].getOriginPoints()
	return xOP, yOP
}

func (b *Block) getEndPoints() (int, int){
	xEP, yEP := b.cells[len(b.cells)-1].getEndPoints()
	return xEP, yEP
}

func (b *Block) unplacedPrisoners() []int{
	return b.prisoners
}

func (b *Block) Draw(blockImage *ebiten.Image) {

	b.image.Fill(frameColour)
	//b.cell.Draw(blockImage)
	for c := 0; c < len(b.cells); c++{
		cell := b.cells[c]
		cell.setWorldPos(b.pX, b.pY)
		cell.setDrawPos(cell.relX, cell.relY)
		cell.Draw(b.image)

		// TODO: Test this part outside of the for loop for block cells
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(b.x), float64(b.y))
		blockImage.DrawImage(b.image, op)
	}
}

