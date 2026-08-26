package errors

type userErrors struct {
	MATCH_NOT_FOUND string
}

type backendErrors struct {
}

var UserSide = userErrors{
	MATCH_NOT_FOUND: "Match not found",
}

var BackendSide = backendErrors{}
