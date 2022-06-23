package common

import "go.mongodb.org/mongo-driver/bson/primitive"

// Fetch projects collection struct
type FetchProjectStruct struct{
	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID string `json:"userId"`
	ProjectName string `json:"projectName"`
	ProjectDescription string `json:"projectDescription"`
	RepoLink string `json:"repoLink"`
	Languages []string `json:"languages"`
	OtherLanguages []string `json:"otherLanguages"`
	Allied []string `json:"allied"`
	ProjectType []string `json:"projectType"`
	ContributorCount string `json:"contributorCount"`
	Documentation string `json:"documentation"`
	Public string `json:"public"`
	Company string `json:"company"`
	CompanyName string `json:"companyName"`
	Funded string `json:"funded"`
	CreatedDate string `json:"createdDate"`
	PublishedDate string `json:"publishedDate"`
	ClosedDate string `json:"closedDate"`
	IsActive int `json:"isActive"`
}

// Save projects collection struct
type SaveProjectStruct struct{
	UserID string `json:"userId"`
	ProjectName string `json:"projectName"`
	ProjectDescription string `json:"projectDescription"`
	RepoLink string `json:"repoLink"`
	Languages []string `json:"languages"`
	OtherLanguages []string `json:"otherLanguages"`
	Allied []string `json:"allied"`
	ProjectType []string `json:"projectType"`
	ContributorCount string `json:"contributorCount"`
	Documentation string `json:"documentation"`
	Public string `json:"public"`
	Company string `json:"company"`
	CompanyName string `json:"companyName"`
	Funded string `json:"funded"`
	CreatedDate string `json:"createdDate"`
	PublishedDate string `json:"publishedDate"`
	ClosedDate string `json:"closedDate"`
	IsActive int `json:"isActive"`
}

// Fetch users struct
type FetchUserStruct struct {
	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email 		string `json:"email"`
	Name 		string `json:"name"`
	Location 	string `json:"location"`
	ImageLink 	string `json:"imageLink"`
	RepoUrl 	string `json:"repoUrl"`
	Source 		string `json:"source"`
	CreatedDate string `json:"createdDate"`
}

// Save users struct
type SaveUserStruct struct {
	Email 		string `json:"email"`
	Name 		string `json:"name"`
	Location 	string `json:"location"`
	ImageLink 	string `json:"imageLink"`
	RepoUrl 	string `json:"repoUrl"`
	Source 		string `json:"source"`
	CreatedDate string `json:"createdDate"`
}

// Fetch user sessions Struct
type FetchUserSessionStruct struct{
	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID string `json:"userId"`
	SessionID string `json:"sessionId"`
}

// Save user sessions Struct
type SaveUserSessionStruct struct{
	UserID string `json:"userId"`
	SessionID string `json:"sessionId"`
}

// Fetch contributor preferences Struct
type FetchContributorPreferencesStruct struct{
	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID string `json:"userId"`
	Languages []string `json:"languages"`
	OtherLanguages []string `json:"otherLanguages"`
	Allied []string `json:"allied"`
	ProjectType []string `json:"projectType"`
	NotificationFrequency string `json:"notificationFrequency"`
	ContributorCount string `json:"contributorCount"`
	PaidJob string `json:"paidJob"`
	Relocation string `json:"relocation"`
	Qualification string `json:"qualification"`
}

// Save contributor preferences Struct
type SaveContributorPreferencesStruct struct{
	UserID string `json:"userId"`
	Languages []string `json:"languages"`
	OtherLanguages []string `json:"otherLanguages"`
	Allied []string `json:"allied"`
	ProjectType []string `json:"projectType"`
	NotificationFrequency string `json:"notificationFrequency"`
	ContributorCount string `json:"contributorCount"`
	PaidJob string `json:"paidJob"`
	Relocation string `json:"relocation"`
	Qualification string `json:"qualification"`
}