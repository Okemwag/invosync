package model

// defines a contract that other structs can implement
type Tabler interface {
	TableName() string
}

//This struct is specifically designed for interaction with a database
type DBBook struct {
	Isbn int `json:"isbn"`
	Name string `json:"name"`
	Publisher string `json:"publisher"` 
}

func (DBBook) TableName() string {
	return "books"
}