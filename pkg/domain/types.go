package domain

const (
	Unasigned int = 0
	Asigned       = 1
)

type Journey struct {
	ID     int64 `json:"id"`
	People int   `json:"people"`
}

type JourneyDB struct {
	ID     int64
	People int
	Status int
}

type Car struct {
	ID    int64 `json:"id"`
	Seats int   `json:"seats"`
}

type CarDB struct {
	ID         int64
	Seats      int
	ID_Journey int64
}
