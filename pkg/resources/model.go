package resources

type Authentication struct {
	Mechanism string `json:"mechanism"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Token     string `json:"token"`
}

type ConnectionRequest struct {
	URL        string          `json:"url"`
	ConfigPath string          `json:"configPath"`
	Protocol   string          `json:"protocol"`
	Branch     string          `json:"branch"`
	Auth       *Authentication `json:"auth,omitempty"`
}

type Variable struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
	Default     string `json:"default,omitempty"`
}
type Config struct {
	Versino   string     `json:"version,omitempty"`
	Mode      string     `json:"mode,omitempty"`
	Variables []Variable `json:"variables"`
}

type ProcessTemplateRequest struct {
	ConnectionRequest
	Mode   string                 `json:"mode"`
	Params map[string]interface{} `json:"params"`
}
