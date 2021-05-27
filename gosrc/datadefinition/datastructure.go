package datadefinition

type Event struct {
	EventSource     string `json:"event_source" csv:"event_source"`
	EventRef        string `json:"event_ref" csv:"event_ref"`
	EventType       int    `json:"event_type" csv:"event_type"`
	EventDate       string `json:"event_date" csv:"event_date"`
	CallingNumber   int    `json:"calling_number" csv:"calling_number"`
	CalledNumber    int    `json:"called_number" csv:"called_number"`
	Location        string `json:"location" csv:"location"`
	DurationSeconds int    `json:"duration_seconds" csv:"duration_seconds"`
	Attr1           string `json:"attr_1" csv:"attr_1"`
	Attr2           string `json:"attr_2" csv:"attr_2"`
	Attr3           string `json:"attr_3" csv:"attr_3"`
	Attr4           string `json:"attr_4" csv:"attr_4"`
	Attr5           string `json:"attr_5" csv:"attr_5"`
	Attr6           string `json:"attr_6" csv:"attr_6"`
	Attr7           string `json:"attr_7" csv:"attr_7"`
	Attr8           string `json:"attr_8" csv:"attr_8"`
}
