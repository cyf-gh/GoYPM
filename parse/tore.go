package ypm_parse

import (
	"stgogo/comn/str"
	"strings"
	"errors"
)

func IsOperationPackage( msg string ) bool {
	return strings.Contains( msg, "@")
}

func SplitHeadBody( raw string ) ( string, string, error ){
	if !IsOperationPackage(raw) {
		return "", "", errors.New("[Invalid Message] no @" )
	}

	markAt := "@"
	markAtCount := 1
	for  {
		markAtCountCur := strings.Count(raw, st_comn_str.WrapWith( markAt, "[", "]" ));
		if markAtCountCur == 0 {
			break;
		}
		markAtCount = markAtCountCur
		markAt = st_comn_str.WrapWith( markAt, "[", "]" )
	}

	if markAtCount != 1 {
		return "", "", errors.New("[Invalid Message] Too many " + markAt + "detected." )
	}

	head___body := strings.Split( raw, markAt)

	if len(head___body) > 2 {
		return "", "", errors.New("[Invalid Message] split parsing error.")
	} else if len(head___body) == 1 {
		return head___body[0], "", nil
	}
	head := head___body[0]
	body := head___body[1]
	return head, body, nil
}

func SplitParaments( body string ) ( []string ) {
	commaFlag := "$^ccaammoomm^$"
	slashFlag := "\\$^ssllaasshh^$"
	body = strings.ReplaceAll( body, "\\\\", slashFlag)
	body = strings.ReplaceAll( body, "\\,", commaFlag)
	body = strings.ReplaceAll( body, slashFlag, "\\")

	params := strings.Split( body, "," )
	for i, p := range params {
		params[i] = strings.ReplaceAll( p, commaFlag, "\\,")
	}
	return params
}