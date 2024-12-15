package entity

type Country struct {
	Id   int64  `db:"id" yaml:"id"`
	Name string `db:"name" yaml:"name"`
}

type CountryCollection struct {
	Countries []Country `yaml:"countries"`
}
