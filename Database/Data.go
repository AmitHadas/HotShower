package Database

type Boiler struct {
	ID 				int
	Desired_Temp	int
	Max_Temp    	int
}

type User struct {
	User_Name 	string
	Pass    	string
	Boiler_ID   int
}

var UserMap = make(map[string]User)
var BoilerMap = make(map[int]Boiler)

func AddNewBoiler(boiler_id int, desired_temp int, max_temp int){
	newBoiler := Boiler{boiler_id, desired_temp, max_temp}
	BoilerMap[boiler_id] = newBoiler
}

func AddNewUser(user User, desired_temp int, max_temp int) {
	UserMap[user.User_Name] = user
	boiler, ok := BoilerMap[user.Boiler_ID]
	if !ok {
		AddNewBoiler(user.Boiler_ID, desired_temp, max_temp)
	} else {
		boiler.Desired_Temp = desired_temp
		boiler.Max_Temp = max_temp
	}
}

func GetBoilerByUserName(userName string, pass string) *Boiler {

	user, valid_user := UserMap[userName]
	if valid_user {
		if user.Pass == pass {
			boiler, valid_boiler := BoilerMap[user.Boiler_ID]

			if valid_boiler {
				return &boiler
			}
		}
	}
	return nil
}

func UserExists(userName string) bool {
	_, valid_user := UserMap[userName]
	if valid_user{
		return true
	}
	return false
}
