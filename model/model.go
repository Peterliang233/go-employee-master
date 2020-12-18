package model

type Employee struct {
	ID          uint64 `json:"id"`
	RealName    string `json:"real_name"`
	NickName    string `json:"nick_name"`
	EnglishName string `json:"english_name"`
	Sex         string `json:"sex"`
	Age         uint8  `json:"age"`
	Address     string `json:"address"`
	MobilePhone string `json:"mobile_phone"`
	IDCard      string `json:"id_card"`
}

type User struct {
	ID           uint64   `json:"id"`
	Username     string   `json:"username"`
	PasswordHash string   `json:"password_hash"`
	EmployeeID   uint64   `json:"employee_id"`
	Roles        []string `json:"roles"`
	Departments  []string `json:"departments"`
}

type Department struct {
	ID             uint64 `json:"id"`
	DepartmentName string `json:"department_name"`
}

type Role struct {
	ID       uint64 `json:"id"`
	RoleName string `json:"role_name"`
}

type IdentifyAndUsername struct {
	Identify string `json:"identify"`
	Username string `json:"username"`
}
