package domain

type Park struct {
	ParkId    uint `postgres:"primaryKey"`
	ParkName  string
	CourtType string
	IsFree    bool
}
