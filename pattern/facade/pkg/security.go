package pkg

import "fmt"

type securityCode struct {
	code int
}

func newSecurityCode(New int) *securityCode {
	return &securityCode{code: New}
}

func (s *securityCode) checkCode(incomingCode int) error {
	if s.code != incomingCode {
		fmt.Errorf("Security Code is incorrect")
	}
	fmt.Println("SecurityCode Verified")
	return nil
}
