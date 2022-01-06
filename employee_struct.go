package main

type name struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type department struct {
	DepartmentName string `json:"department_name"`
}

type level struct {
	HighLevel bool  `json:"high_level"`
	LowLevel  bool  `json:"low_level"`
	Amount    int64 `json:"amount"`
}

type salary struct {
	Level level `json:"level"`
}

type streetName struct {
	HouseNo  string `json:"house_no"`
	Locality string `json:"locality"`
	Landmark string `json:"landmark"`
}

type address struct {
	StreetName streetName `json:"street_name"`
	CityName   string     `json:"city_name"`
}

type Employee struct {
	Name          name       `json:"name"`
	Department    department `json:"department"`
	Salary        salary     `json:"salary"`
	Address       address    `json:"address"`
	FamilyMembers []string   `json:"family_members"`
}
