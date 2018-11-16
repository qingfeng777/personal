package moudle

type Resume struct {

}

type Personal struct {
	Image        string
	Name         string
	Position     string
	Addr         string
	Introduction string
	Email        string
	Phone        string
	//home page
}

type Experience struct {
	Position    string
	CompanyName string
	CompanyAddr string
	Title       string
	Desc        string
	Skill       []string
	StartAt     int64   //  Accuracy months
	EndAt       int64
}


type Education struct {
	Index uint8 //show index ,begin from 0
	CollageName string
	StartAt int64
	EndAt int64
	Profession string
	Level string //高中、大专、本科、硕士、博士
	Desc string
}

type Skill struct {
	Name string
	Expert uint8 //year
	Degree uint8 //1~100
	Desc string

}

