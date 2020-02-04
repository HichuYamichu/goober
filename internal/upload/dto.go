package upload

type uploadResult struct {
	URL     string `json:"url"`
	Name    string `json:"name"`
	Success bool   `json:"success"`
}
