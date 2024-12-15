package entity

type Proxy struct {
	Id        int64  `db:"id" yaml:"id"`
	Address   string `db:"address" yaml:"address"`
	CountryId int64  `db:"country_id" yaml:"country_id"`
}

type ProxyCollection struct {
	Proxies []Proxy `yaml:"proxies"`
}
