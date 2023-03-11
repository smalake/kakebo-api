package model

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// func GetUser() string {
// 	db := ConnectDB()
// 	rows, err := db.Query("SELECT * FROM users")
// 	if err != nil {
// 		logging.WriteErrorLog(err.Error(), true)
// 		return "error"
// 	}
// 	defer db.Close()
// 	defer rows.Close()

// 	var users []User
// 	for rows.Next() {
// 		var user User
// 		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
// 			logging.WriteErrorLog(err.Error(), true)
// 			return "error"
// 		}
// 		users = append(users, user)
// 	}

// 	// JSONに変換
// 	jsonBytes, err := json.Marshal(users)
// 	if err != nil {
// 		logging.WriteErrorLog(err.Error(), true)
// 		return "error"
// 	}
// 	userJson := string(jsonBytes)
// 	return userJson
// }

// func GetUserData(id int) User {
// 	db := ConnectDB()
// 	query := "SELECT name, email FROM users WHERE id = ?"
// 	var user User
// 	rows := db.QueryRow(query, id)

// 	defer db.Close()

// 	if err := rows.Scan(&user.Name, &user.Email); err != nil {
// 		logging.WriteErrorLog(err.Error(), true)
// 	}

// 	return user
// }
