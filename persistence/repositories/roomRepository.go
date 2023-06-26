package repositories

type RoomRepository struct {
	Rooms []Room
}

func (rr *RoomRepository) Save(room Room) {
	rr.Rooms = append(rr.Rooms, room)
}

func (rr *RoomRepository) FindAll() []Room {
	return rr.Rooms
}
