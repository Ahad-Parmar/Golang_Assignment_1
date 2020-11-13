package main

import (
	"fmt"
	"regexp"
)

func CPL(ps string) error {    //Check Password Lever
	if len(ps) < 8 {
		return fmt.Errorf("password length is < 8")
	}
	num := `[0-9]{1}`
	a_z := `[a-z]{1}`
	A_Z := `[A-Z]{1}`
	symbol := `[!@#~$%^&*()+|_]{1}`
	if b, er := regexp.MatchString(num, ps); !b || er != nil {
		return fmt.Errorf("password should be having number")
	}
	if b, er := regexp.MatchString(a_z, ps); !b || er != nil {
		return fmt.Errorf("password should be having Lowercase letter")
	}
	if b, er := regexp.MatchString(A_Z, ps); !b || err != nil {
		return fmt.Errorf("password should be having Uppercase letter")
	}
	if b, er := regexp.MatchString(symbol, ps); !b || er != nil {
		return fmt.Errorf("password sould be having symbol")
	}
	return nil
}

func main() {
	err := CheckPasswordLever("Abcd@18kohli")
	fmt.Println(er)
}