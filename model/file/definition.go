package file

const (
	//
	// TypeUnknown the type id of an unknown file type.
	//
	TypeUnknown = uint64(0)

	//
	// TypeFolder the type id of a folder.
	//
	TypeFolder = uint64(1)

	//
	// TypeFile the type id of a file.
	//
	TypeFile = uint64(2)

	//
	// TypeWeakLink the type id of a weak link.
	//
	TypeWeakLink = uint64(3)

	//
	// TypeHardLink the type id of a hard link.
	//
	TypeHardLink = uint64(4)
)
