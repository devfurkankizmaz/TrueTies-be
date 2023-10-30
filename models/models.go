package models

type Survey struct {
	ID        uint       `json:"id"`
	Title     string     `json:"title"`
	OwnerID   uint       `json:"owner_id"`
	Questions []Question `json:"questions"`
}

type Question struct {
	ID      uint     `json:"id"`
	Text    string   `json:"text"`
	Options []string `json:"options"`
}

type Answer struct {
	ID        uint       `json:"id"`
	SurveyID  uint       `json:"survey_id"`
	UserID    uint       `json:"user_id"`
	Responses []Response `json:"responses"`
}

type Response struct {
	QuestionID uint   `json:"question_id"`
	Answer     string `json:"answer"`
}
