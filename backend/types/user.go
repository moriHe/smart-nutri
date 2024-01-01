package types

type User struct {
	Id             int  `json:"id"`
	ActiveFamilyId *int `json:"activeFamilyId"`
}

type PostUser struct {
	FireUid string `json:"fireUid"`
}
