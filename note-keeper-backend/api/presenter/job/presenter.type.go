package presenter

type JobListPresenter struct {
	Id          int         `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Address     []*Address  `json:"address"`
	Company     *Company    `json:"company"`
	Position    []*Position `json:"position"`
	Skill       []*Skill    `json:"skill"`
}

type Address struct {
	Id       int    `json:"id"`
	No       int    `json:"no"`
	Road     string `json:"road"`
	Ward     string `json:"ward"`
	District string `json:"district"`
	City     string `json:"city"`
	Country  string `json:"country"`
}

type Company struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Position struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Skill struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
