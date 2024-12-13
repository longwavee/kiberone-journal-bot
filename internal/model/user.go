package model

type Worker struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"user_name"`
	TutorWork int    `json:"tutor_work"`
	AssisWork int    `json:"assis_work"`
	Outwork   int    `json:"outwork"`
}
