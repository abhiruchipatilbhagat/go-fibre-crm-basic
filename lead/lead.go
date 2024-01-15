package lead

import(

)

type Lead struct{
	gorm.Model
	Name 		string
	Company 	string
	Email 		string
	Phone 		int
}