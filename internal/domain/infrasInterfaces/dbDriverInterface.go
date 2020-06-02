package infrasInterfaces

type DbDriver interface {
	Query(qString string, fields ...string) error
}
