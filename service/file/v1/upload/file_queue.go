package upload

//
// FileQueueEntry contains information of
//
type FileQueueEntry struct {
	FileLogEntry string
	FileName     string
}

//
// FileQueue contains a queue of all files waiting to be processed by the
// server.
//
type FileQueue struct {
	Entries []FileQueueEntry
}

//
// AddFileToQueue adds a single file to the file queue.
//
func (fq *FileQueue) AddFileToQueue() error {
	return nil
}
