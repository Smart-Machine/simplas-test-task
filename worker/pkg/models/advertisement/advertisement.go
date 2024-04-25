package advertisement

type Advertisement struct {
	ID         string            `json:"_id"`
	Categories map[string]string `json:"categories"`
	Title      map[string]string `json:"title"`
	Type       string            `json:"type"`
	Posted     float64           `json:"posted"`
}
