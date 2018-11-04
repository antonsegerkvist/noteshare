package permission

const (

	//
	// ACCOUNT SETTINGS.
	//

	// PermissionEditAccount specifies if a group can edit the account.
	PermissionEditAccount = 0x00000000

	// PermissionEditAccountLayout specifies if a group can edit the account layout.
	PermissionEditAccountLayout = 0x00000001

	//
	// USER SETTINGS.
	//

	// PermissionEditOwnProfile specifies if a group can edit its own profile.
	PermissionEditOwnProfile = 0x01000000

	// PermissionEditAllProfile specifies if a group can edit all user profiles.
	PermissionEditAllProfile = 0x01000001

	//
	// FILE SETTINGS.
	//

	// PermissionAddFolder specifies if a group can add folders to the folder tree.
	PermissionAddFolder = 0x02000000
)
