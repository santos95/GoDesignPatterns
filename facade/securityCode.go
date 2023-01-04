package main

import "fmt"

type securityCode struct {
	code int
}

func createSecurityXode(code int) *securityCode {
	return &securityCode{code: code}
}

func (s *securityCode) checkCode(code int) error {
	if s.code != code {
		return fmt.Errorf("Security code is incorrect")
	}

	fmt.Println("Security code verified")
	return nil
}
