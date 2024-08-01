package author

type Author struct {
	Name        string `json:"name" xml:"name" form:"name" query:"name"`
	YearOfBirth int    `json:"year-of-birth" xml:"year-of-birth" form:"year-of-birth" query:"year-of-birth"`
}
