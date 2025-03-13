package geo

import "errors"

type Landmark struct {
  name string
  Coordinates
}

// Get method ----------------------------------------------------------------
func (l *Landmark) Name() string {
  return l.name
}

// Set method ----------------------------------------------------------------
func (l *Landmark) SetName(name string) error {
  if name == "" {
    return errors.New("invalid name")
  }
  l.name = name
  return nil 
}
