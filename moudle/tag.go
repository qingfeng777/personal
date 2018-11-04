package moudle

type Tag struct {
	Name     string   `json:"name"`
	Slug     string   `json:"slug"`
	Desc     string   `json:"desc"`
	Extends  []Extend `json:"extends"`
	CreateAt int64    `json:"create_at"`
	id       int      `json:"id"`
}
