package internal

type FileInfo struct {
	Count     int64
	Name      string
	Size      string
	Kind      string
	DateAdded string
}

type FileColumnLengths struct {
	Count    int
	Filename int
	Size     int
	Kind     int
	Date     int
}
