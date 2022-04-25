package main

import (
	"fmt"
)

type SubmissionAttrAction int64

const (
	ADD SubmissionAttrAction = iota
	UPDATE
	DELETE
	APPEND
)

func (a SubmissionAttrAction) String() string {
	return []string{"add", "update", "delete", "append"}[a]
}

func main() {

	var a SubmissionAttrAction = ADD

	fmt.Println(a.String())

}
