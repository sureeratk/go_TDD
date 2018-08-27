package iteration

import "strings"

func Repeat(character string, itr int) string {
	//var repeated string
	// for i := 0; i < itr; i++ {
	// 	repeated += character
	//return repeated

	return strings.Repeat(character, itr)

}
