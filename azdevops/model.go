package azdevops

type buildDefinition struct {
	Name       string `json:"name"`
	ID         int    `json:"id,omitempty"`
	Repository struct {
		Properties struct {
			LabelSources       string `json:"labelSources"`
			ReportBuildStatus  string `json:"reportBuildStatus"`
			ConnectedServiceID string `json:"connectedServiceId"`
		} `json:"properties"`
		Type               string `json:"type"`
		Name               string `json:"name,omitempty"`
		URL                string `json:"url,omitempty"`
		DefaultBranch      string `json:"defaultBranch"`
		Clean              string `json:"clean"`
		CheckoutSubmodules bool   `json:"checkoutSubmodules"`
	} `json:"repository"`
	Process struct {
		YamlFilename string `json:"yamlFilename"`
		Type         int    `json:"type"`
	} `json:"process"`
	Queue struct {
		ID int `json:"id"`
	} `json:"queue"`
	Variables []variable `json:"variables"`
}

type variable struct {
	Value         string `json:"value"`
	IsSecret      bool   `json:"isSecret"`
	AllowOverride bool   `json:"allowOverride"`
}

type azureEndpoint struct {
	Data struct {
		SubscriptionID   string `json:"SubscriptionId"`
		SubscriptionName string `json:"SubscriptionName"`
		CreationMode     string `json:"creationMode"`
	} `json:"data"`
	ID            string `json:"id,omitempty"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	URL           string `json:"url"`
	Authorization struct {
		Parameters struct {
			ServicePrincipalID  string `json:"serviceprincipalid"`
			TenantID            string `json:"tenantid"`
			ServicePrincipalKey string `json:"serviceprincipalkey"`
		} `json:"parameters"`
		Scheme string `json:"scheme"`
	} `json:"authorization"`
	IsReady bool `json:"isReady"`
}

type githubEndpoint struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	URL           string `json:"url"`
	Authorization struct {
		Parameters struct {
			Token string `json:"AccessToken"`
		} `json:"parameters"`
		Scheme string `json:"scheme"`
	} `json:"authorization"`
	IsReady bool `json:"isReady"`
}
