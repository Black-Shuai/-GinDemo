package Models

type GeneralSort struct {
	Id int `json:"id"`
	GeneralSort string `json:"general_sort"`
}

func (GeneralSort)TableName() string  {
	return "db_general_sort"
}