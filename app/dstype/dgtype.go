package dstype

type Scanneable interface {
	Scan(string, *int) error
}
