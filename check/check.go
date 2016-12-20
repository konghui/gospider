package check

type CheckUrl interface {
	SetString(string)
	GetString(string) bool
}
