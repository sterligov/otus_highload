package v1

type Inserted struct {
	InsertedID int64 `json:"inserted_id"`
}

type Affected struct {
	Affected int64 `json:"affected"`
}
