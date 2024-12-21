package models

type User struct {
	Uuid                 string `gorm:"primarykey"`
	UserName             string
	PasswordDigest       string
	NickName             string
	Status               string
	Avatar               string `gorm:"size:1000"`
	UserFileStoreID      string
	UserMainFileFolderID string
}

const (
	// PasswordCount password encryption difficulty
	PasswordCount = 12
	// super admin
	StatusSuperAdmin = "super_admin"
	// admin User
	StatusAdmin = "common_admin"
	// active User
	StatusActiveUser = "active"
	// inactive User
	StatusInactiveUser = "inactive"
	// Suspend User
	StatusSuspendUser = "suspend"
)
