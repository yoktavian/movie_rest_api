package sql

type FakeSql interface {
	Get() map[string]interface{}
}

type fakeSql struct{}

func (f *fakeSql) Get() map[string]interface{} {
	return map[string]interface{}{
		"id":     "1",
		"name":   "spiderman",
		"link":   "http://spiderman.com",
		"rating": 4,
	}
}

func NewFakeSql() FakeSql {
	return &fakeSql{}
}
