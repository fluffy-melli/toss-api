package types

type Owner string

const (
	OwnerPersonal  Owner = "개인"
	OwnerCorporate Owner = "법인"
	OwnerUnknown   Owner = "미확인"
)
