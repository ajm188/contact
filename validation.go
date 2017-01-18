package main

import (
	"net/mail"
)

func ensureValidEmail(possibleAddr string) (string, error) {
	addrStruct, err := mail.ParseAddress(possibleAddr)
	if err != nil {
		return "", err
	}

	addrString := addrStruct.String()
	/*
	 * if possibleAddr is foo@bar.com, then addrStruct.String() will be
	 * <foo@bar.com>, so we need to chop off the <,>
	 */
	return addrString[1:(len(addrString) - 1)], nil
}
