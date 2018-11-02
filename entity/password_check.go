package entity
func PasswordCheck(passwd string) string {

indNum := [4]int{0, 0, 0, 0}
spCode := []byte{'!', '@', '#', '$', '%', '^', '&', '*', '_', '-'}

	if len(passwd) < 6 {
	return"password too short"
}

	passwdByte := []byte(passwd)

	for _, i := range passwdByte {
		if i >= 'A' && i <= 'Z' {
			indNum[0] = 1
			continue
	}

	if i >= 'a' && i <= 'z' {
		indNum[1] = 1
		continue
	}

	if i >= '0' && i <= '9' {
		indNum[2] = 1
		continue
	}

	notEnd := 0
	for _, s := range spCode {
	if i == s {
		indNum[3] = 1
		notEnd = 1
		break
	}
}

	if notEnd != 1 {
		return "Unsupport code"
	}
	}

	codeCount := 0

	for _, i := range indNum {
		codeCount += i
	}

	if codeCount < 3 {
		return "Too simple password"
	}

	return "right"
}