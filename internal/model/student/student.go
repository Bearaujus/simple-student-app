package student

type Student struct {
	SID   string `json:"sid,omitempty"`
	Name  string `json:"name,omitempty"`
	Age   int    `json:"age,omitempty"`
	Grade int    `json:"grade,omitempty"`
}
