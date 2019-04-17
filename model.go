package autorestapi

type Model interface {
	New() (Model, error)
	Check() error
	Save() error
	Find(query interface{}, pageSize, pageIndex int, sorted string) ([]Model, error)
	FindOne(query interface{}) (Model, error)
	Count(query interface{}) (int, error)
	Delete(query interface{}) error
	DeleteAll(query interface{}) error
	UpdateOne(query interface{}, update interface{}) error
	UpdateAll(query interface{}, update interface{}) error
}
