package sodoku

import (
	"fmt"
)

type VacantCell struct {
	cell *Cell
}

func (vacant *VacantCell) imprisonPrisoner(criminal *Prisoner) error{
	vacant.cell.prisoner = criminal
	vacant.cell.prisonerID = criminal.getDesignation()
	vacant.cell.setState(vacant.cell.occupied)
	return nil
}

func (vacant *VacantCell) releasePrisoner() error{
	return fmt.Errorf("Thank you Warden! But our prisoner is in another cell!")
}

func (vacant *VacantCell) emptyCell() error{
	return fmt.Errorf("Warden the cell is empty!")
}