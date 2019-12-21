package models

// Point is a simple type with a location for geospatial
// queries.
type Point struct {
	UserId      string   `json:"user_id"`
	Description string   `json:"description"`
	MemberName  string   `json:"member_name"`
	Location    Location `json:"location"`
}
