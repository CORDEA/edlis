package models

type Comment struct {
	Id       string `json:"id" bson:"_id"`
	SlideId  string `json:"slide_id" bson:"slide_id"`
	UserId   string `json:"user_id" bson:"user_id"`
	Lang     string `json:"lang" bson:"lang"`
	Comment  string `json:"comment" bson:"comment"`
	Type     string `json:"type" bson:"type"`
	PostedAt string `json:"posted_at" bson:"posted_at"`
}
