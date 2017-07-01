package models

// AlertmanagerInstance describes the Alertmanager instance alert was collected
// from
type AlertmanagerInstance struct {
	Name string `json:"name"`
	URI  string `json:"uri"`
	// Source links to alert source for given alertmanager instance
	Source string `json:"source"`
	// all silences matching current alert in this upstream
	Silences map[string]Silence `json:"silences"`
	// TODO also export InhibitedBy here if it ever becomes needed
}

// AlertmanagerAPIStatus describes the Alertmanager instance overall health
type AlertmanagerAPIStatus struct {
	Name  string `json:"name"`
	URI   string `json:"uri"`
	Error string `json:"error"`
}

// AlertmanagerAPICounters returns number of Alertmanager instances in each
// state
type AlertmanagerAPICounters struct {
	Total   int `json:"total"`
	Healthy int `json:"healthy"`
	Failed  int `json:"failed"`
}

// AlertmanagerAPISummary describes the Alertmanager instance overall health
type AlertmanagerAPISummary struct {
	Counters  AlertmanagerAPICounters `json:"counters"`
	Instances []AlertmanagerAPIStatus `json:"instances"`
}
