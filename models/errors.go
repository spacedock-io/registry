package models

type Error struct {
  Message string
}

var (
  NotFoundErr = &Error{ "Not found." }
  TagCreateErr = &Error{ "Unable to create tag." }
  TagSaveErr = &Error{ "Unable to save tag." }
)

func (err *Error) Error() string {
  return err.Message
}
