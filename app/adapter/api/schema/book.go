package schema

type Book struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	OfficeID string `json:"office_id"`
	Isbn     string `json:"isbn"`
}
