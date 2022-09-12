package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Student struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name,omitempty" json:"name,omitempty"`
	RollNumber  string             `bson:"roll_number,omitempty" json:"roll_number,omitempty"`
	CourseName string             `bson:"course_name,omitempty" json:"course_name,omitempty"`
	EmailId     string             `bson:"email_id, omitempty" json:"email_id,omitempty"`
	DOB         primitive.DateTime `bson:"dob, omitempty" json:"dob,omitempty"`
	Subjects      []string           `bson:"subjects, omitempty" json:"subjects,omitempty"`

}