package Models

type Sex struct {
	Id string
	Name string
}

func (Sex)TableName() string  {
	return "db_sex"
}
