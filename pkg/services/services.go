package services

func AccessByUser() (string, error) {
	return "User Access Granted", nil
}

func AccessByAdmin() (string, error) {
	return "Admin Access Granted", nil
}
