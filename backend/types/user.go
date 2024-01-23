package types

type User struct {
	Id             int  `json:"id"`
	ActiveFamilyId *int `json:"activeFamilyId"`
}

type UserFamily struct {
	Id         int    `json:"id"`
	FamilyId   int    `json:"familyId"`
	FamilyName string `json:"familyName"`
	Role       string `json:"role"`
}
