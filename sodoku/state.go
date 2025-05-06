package sodoku

type State interface {
	imprisonPrisoner(*Prisoner) error
	releasePrisoner() error
	emptyCell() error
	//transferPrisoner() error
}