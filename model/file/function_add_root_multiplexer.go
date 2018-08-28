package file

//
// AddRootMultiplexerFromUserID adds a single file to the root folders for the specified
// userID.
//
func AddRootMultiplexerFromUserID(file *ModelAddFile, userID uint64) (uint64, error) {
	switch file.Type {
	case 1:
		return AddRootFolderFromUserID(file, userID)
	case 2:
		return AddRootFileFromUserID(file, userID)
	case 3:
		break
	case 4:
		break
	}
	return 0, ErrUnknownFileType
}
