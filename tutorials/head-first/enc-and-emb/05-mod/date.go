package mypackage

import "errors"

// Structure: ----------------------------------------------------------------
type Date struct {
  year int
  month int
  day int
}

// Set-methods (setters): ----------------------------------------------------
func (d *Date) SetYear(year int) error {
  if year < 1900 {
    return errors.New("invalid year")
  }
  d.year = year
  return nil
}

func (d *Date) SetMonth(month int) error {
  if month < 1 || month > 12 {
    return errors.New("invalid month")
  }
  d.month = month
  return nil 
}

func (d *Date) SetDay(day int) error {
  if day < 1 || day > 31 {
    return errors.New("invalid day")
  }
  d.day = day
  return nil
}

// Get methods (getters) -----------------------------------------------------
func (d *Date) Year() int {
  return d.year
}

func (d *Date) Month() int {
  return d.month
}

func (d *Date) Day() int {
  return d.day
}
