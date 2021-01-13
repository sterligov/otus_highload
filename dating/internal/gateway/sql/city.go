package sql

type City struct {
	ID        int64
	Name      string
	CountryID *City
}
