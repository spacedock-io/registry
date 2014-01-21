package models

type Error struct {
  Message string
}

var (
  TagCreateErr = &Error{ "Unable to create tag." }
)

func (err *Error) Error() string {
  return err.Message
}
