package models

import (
	"fmt"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

// Date represents a date without time information
type Date struct {
	Year  int
	Month time.Month
	Day   int
}

// NewDate creates a new Date from a time.Time
func NewDate(t time.Time) Date {
	return Date{
		Year:  t.Year(),
		Month: t.Month(),
		Day:   t.Day(),
	}
}

// ToTime converts a Date to time.Time
func (d Date) ToTime() time.Time {
	return time.Date(d.Year, d.Month, d.Day, 0, 0, 0, 0, time.UTC)
}

// MarshalJSON implements json.Marshaler
func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.ToTime().Format("2006-01-02") + `"`), nil
}

// UnmarshalJSON implements json.Unmarshaler
func (d *Date) UnmarshalJSON(data []byte) error {
	t, err := time.Parse(`"2006-01-02"`, string(data))
	if err != nil {
		return err
	}
	*d = NewDate(t)
	return nil
} 

func (d *Date) UnmarshalBSONValue(t bsontype.Type, data []byte) error {

	switch t {
	case bsontype.DateTime:
		var dt time.Time
		err := bson.UnmarshalValue(t, data, &dt)
		if err != nil {
			return err
		}
		*d = NewDate(dt)
		return nil

	case bsontype.EmbeddedDocument:
		var doc struct {
			Year  int        `bson:"Year"`
			Month int        `bson:"Month"`
			Day   int        `bson:"Day"`
		}
		err := bson.Unmarshal(data, &doc)
		if err != nil {
			return err
		}
		*d = Date{
			Year:  doc.Year,
			Month: time.Month(doc.Month),
			Day:   doc.Day,
		}
		return nil

	default:
		return fmt.Errorf("unsupported BSON type for Date: %v", t)
	}
}

func (d Date) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(d.ToTime())
}
