package models

type User struct {
	Nickname   string `json:"nickname"`
	Age        string `json:"age"`
	Gender     string `json:"gender"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Uuid       string `json:"uid"`
	Image      string `json:"image"`
}

type UserInfo struct {
	Nickname   string `json:"nickname"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Image      string `json:"image"`
}


var UserErrors struct {
	InvalidEmail       string
	InvalideAge        string
	InvalideGender     string
	InvalideFirst_Name string
	InvalideLast_Name  string
	Invalidnickname    string
	InvalidPassword    string
	UserAlreadyExist   string
	UserNotExist       string
} = struct {
	InvalidEmail       string
	InvalideAge        string
	InvalideGender     string
	InvalideFirst_Name string
	InvalideLast_Name  string
	Invalidnickname    string
	InvalidPassword    string
	UserAlreadyExist   string
	UserNotExist       string
}{
	InvalideAge:        "invalide Age",
	InvalideGender:     "invalide Gender",
	InvalideFirst_Name: "invalide First_Nam",
	InvalideLast_Name:  "invalide Last_Name",
	InvalidEmail:       "invalid email",
	Invalidnickname:    "invalid nickname",
	InvalidPassword:    "invalid password",
	UserAlreadyExist:   "user already exist",
	UserNotExist:       "user doesn't exist",
}

var Errors struct {
	InvalidCredentials   string
	InvalidEmail         string
	LongEmail            string
	InvalidUsername      string
	InvalidPassword      string
	UserAlreadyExist     string
	EmailAlreadyExist    string
	UsernameAlreadyExist string
	// other
	ErrorHashingPass string
} = struct {
	InvalidCredentials   string
	InvalidEmail         string
	LongEmail            string
	InvalidUsername      string
	InvalidPassword      string
	UserAlreadyExist     string
	EmailAlreadyExist    string
	UsernameAlreadyExist string
	// other
	ErrorHashingPass string
}{
	InvalidCredentials:   "Invalid Credentials",
	InvalidEmail:         "Invalid Email ex: exmaple@mail.com",
	LongEmail:            "Email must be between 5 and 50 characters.",
	InvalidUsername:      "Username must be between 3 and 15 characters.",
	InvalidPassword:      "Password must be between 6 and 30 characters.",
	UserAlreadyExist:     "Email or Username Already Exist",
	EmailAlreadyExist:    "Email Already Exist",
	UsernameAlreadyExist: "Username Already Exist",
	// other
	ErrorHashingPass: "Error Hashing Password",
}
