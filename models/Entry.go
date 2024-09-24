package models

type Entry struct {
	UuId                             string
	AppId                            string `json:"app_id"`
	AnswerTime                       string `json:"Answer time"`
	Created                          string `json:"Created"`
	Elaborate                        string `json:"Elaborate"`
	ElaborateTimestampEnd            string `json:"Elaborate timestamp end"`
	ElaborateTimestampStart          string `json:"Elaborate timestamp start"`
	Greeting                         string `json:"Greeting"`
	GreetingTimestampEnd             string `json:"Greeting timestamp end"`
	GreetingTimestampStart           string `json:"Greeting timestamp start"`
	LanguageMedication               string `json:"Language medication"`
	LanguageMedicationTimestampEnd   string `json:"Language medication timestamp end"`
	LanguageMedicationTimestampStart string `json:"Language medication timestamp start"`
	Language1                        string `json:"Language1"`
	Language1TimestampEnd            string `json:"Language1 timestamp end"`
	Language1TimestampStart          string `json:"Language1 timestamp start"`
	Language2                        string `json:"Language2"`
	Language2TimestampEnd            string `json:"Language2 timestamp end"`
	Language2TimestampStart          string `json:"Language2 timestamp start"`
	Name                             string `json:"Name"`
	NameInk                          string `json:"Name ink"`
	NameInkTimestampEnd              string `json:"Name ink timestamp end"`
	NameInkTimestampStart            string `json:"Name ink timestamp start"`
	NameTimestampEnd                 string `json:"Name timestamp end"`
	NameTimestampStart               string `json:"Name timestamp start"`
	Picture                          string `json:"Picture"`
	PictureTimestampEnd              string `json:"Picture timestamp end"`
	PictureTimestampStart            string `json:"Picture timestamp start"`
	Radio1                           string `json:"Radio1"`
	Radio1TimestampStart             string `json:"Radio1 timestamp start"`
	Radio2                           string `json:"Radio2"`
	Radio2TimestampStart             string `json:"Radio2 timestamp end"`
	Retell                           string `json:"Retell"`
	RetellTimestampEnd               string `json:"Retell timestamp end"`
	RetellTimestampStart             string `json:"Retell timestamp start"`
	Slider1                          string `json:"Slider1"`
	Slider1TimestampStart            string `json:"Slider1 timestamp end"`
	Slider2                          string `json:"Slider2"`
	Slider2TimestampStart            string `json:"Slider2 timestamp end"`
	Slider3                          string `json:"Slider3"`
	Slider3TimestampStart            string `json:"Slider3 timestamp start"`
	SubmissionID                     string `json:"Submission ID"`
	Suggestion                       string `json:"Suggestion"`
	SuggestionTimestampEnd           string `json:"Suggestion timestamp end"`
	SuggestionTimestampStart         string `json:"Suggestion timestamp start"`
	TapRate                          string `json:"TapRate"`
	TapRateTimestampEnd              string `json:"TapRate timestamp end"`
	TapRateTimestampStart            string `json:"TapRate timestamp start"`
	DataVersion                      string `json:"data_version"`
	NotificationDatetime             string `json:"notification_datetime"`
	SessionId                        string `json:"session_id"`
	StudyId                          string `json:"study_id"`
	Annotations                      string
}
