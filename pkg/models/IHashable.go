package models

// IHashable provides interface definitions for various models associated with this service
type IHashable interface {
	GetPasswordToHash() string
}

// GetPasswordToHash returns a password from a user orm model
func (u UserORM) GetPasswordToHash() string {
	return u.Password
}

// GetPasswordToHash resturns a password from a user object
func (u User) GetPasswordToHash() string {
	return u.Password
}
