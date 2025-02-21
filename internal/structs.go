package internal

type FileInfo struct {
	Count     int64
	Name      string
	Size      int64
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
