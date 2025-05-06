package sodoku

import (
	"fmt"
)

type OccupiedCell struct {
	cell *Cell
}

func (occupied *OccupiedCell) imprisonPrisoner(prisoner *Prisoner) error{

	return fmt.Errorf("A prisoner already occupies the cell. Find another!")
}

func (occupied *OccupiedCell) releasePrisoner() error{
	// if there are more prisoners in the cell then don't empty
	// if prisoner != nil {
		
	// 	occupied.cell.prisoner = prisoner
	
	// } else {
		
	// 	occupied.emptyCell()
	// }

	occupied.cell.prisoner = nil
	occupied.cell.prisoner = NewPrisoner(0)
	occupied.cell.setState(occupied.cell.vacant)

	return nil
}

func (occupied *OccupiedCell) emptyCell() error{

	occupied.cell.prisoner = nil
	occupied.cell.prisoner = NewPrisoner(0)
	occupied.cell.setState(occupied.cell.vacant)

	return nil
}