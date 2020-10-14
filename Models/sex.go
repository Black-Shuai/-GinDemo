package Models

type Sex struct {
	Id string `gorm:"primary_key"`
	Name string
}

func (Sex)TableName() string  {
	return "db_sex"
}
