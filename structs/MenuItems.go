package structs

type MenuItem struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	ItemType    string  `json:"itemtype"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	ImageID     string  `json:"imageid"`
}
