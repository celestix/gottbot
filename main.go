package main

type API struct {
	Components struct {
		Schemas map[string]struct {
			Properties struct {
				Description string `json:"description"`
				Type        string `json:"type"`
				Format      string `json:"format,omitempty"`
			} `json:"properties,omitempty"`
			Required []string `json:"required,omitempty"`
		} `json:"schemas"`
	} `json:"components"`
	Paths any `json:"paths"`
}

func main() {

}
