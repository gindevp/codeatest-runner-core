package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Result struct {
	Score      int64   `json:"score" bson:"score"`
	Time       float64 `json:"time" bson:"time"`
	MemoryUsed int64   `json:"memory_used" bson:"memory_used"`
}

type Submit struct {
	UserQuiz   primitive.ObjectID `json:"user_quiz" bson:"user_quiz"`
	Language   string             `json:"language" bson:"language"`
	UploadFile string             `json:"upload_file" bson:"upload_file"`
	Status     string             `json:"status" bson:"status"`
	Result     Result             `json:"result" bson:"result"`
	UpdatedAt  primitive.DateTime `json:"updated_at"`
	CreatedAt  primitive.DateTime `json:"created_at"`
}