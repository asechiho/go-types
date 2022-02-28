package time

import (
	"database/sql/driver"
	"encoding/json"
	"encoding/xml"
	"gopkg.in/guregu/null.v4"
	"time"
)

const format = "2006-01-02"

type Date null.Time

func DateFromTime(time2 time.Time) Date {
	return Date(null.TimeFrom(time2))
}

func (d *Date) UnmarshalJSON(data []byte) error {
	var str string
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}

	*d, err = parseDateFromString(str)

	return err
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.format())
}

func (d *Date) UnmarshalXML(decoder *xml.Decoder, element xml.StartElement) error {
	var str string
	err := decoder.DecodeElement(&str, &element)
	if err != nil {
		return err
	}

	*d, err = parseDateFromString(str)

	return err
}

func (d Date) MarshalXML(decoder *xml.Encoder, element xml.StartElement) error {
	return decoder.EncodeElement(d.format(), element)
}

func (d Date) Value() (value driver.Value, err error) {
	return null.Time(d).Value()
}

func (d Date) String() string {
	var str = d.format()
	if str != nil {
		return *str
	}
	return ""
}

func (d Date) format() *string {
	if d.Valid {
		var str = d.Time.Format(format)
		return &str
	}

	return nil
}

func parseDateFromString(value string) (Date, error) {
	if value == "" {
		return Date{}, nil
	}

	var tm, err = time.Parse(format, value)
	if err != nil {
		return Date{}, err
	}

	return Date(null.TimeFrom(tm)), nil
}