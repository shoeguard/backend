package forms

type UserRegistrationForm struct {
	PhoneNumber        string `json:"phone_number"         example:"01043214321"   valid:"required, numeric, stringlength(11|11)"` // username field
	Password           string `json:"password"             example:"the!@#$pas123" valid:"required, minstringlength(8)"`           // password field
	IsStudent          bool   `json:"is_student"                                   valid:"required"`
	PartnerPhoneNumber string `json:"partner_phone_number" example:"01012341234"   valid:"numeric, stringlength(11|11)"`
	Nickname           string `json:"nickname"             example:"mengmota"      valid:"required, minstringlength(4)"`
}

type UserReadOnlyInfo struct {
	PhoneNumber        string `json:"phone_number"         example:"01043214321" valid:"required, numeric, stringlength(11|11)"` // username field
	IsStudent          bool   `json:"is_student"                                 valid:"required"`
	PartnerPhoneNumber string `json:"partner_phone_number" example:"01012341234" valid:"numeric, stringlength(11|11)"`
	Nickname           string `json:"nickname"             example:"mengmota"    valid:"required, minstringlength(4)"`
}

type UserModifiableInfo struct {
	Password           string `json:"password"             example:"the!@#$pas123" valid:"minstringlength(8)"` // password field
	IsStudent          *bool  `json:"is_student"`
	PartnerPhoneNumber string `json:"partner_phone_number" example:"01012341234"   valid:"numeric, stringlength(11|11)"`
	Nickname           string `json:"nickname"             example:"mengmota"      valid:"minstringlength(4)"`
}
