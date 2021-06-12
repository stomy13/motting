package user

// User システムの利用者。
type User struct {
	userId UserId
}

// UserId ユーザーID。このドメインのID
type UserId string
