package defs

//request


type Err struct {
	Error string `json:"error"`
	ErrCode string `json:"err_code"`
}

type ErrRespone struct {
	HttpSC int
	Error Err
}

var(
	ErrRequestBodyParsedFailed=ErrRespone{HttpSC:400,Error:Err{Error:"Request Body is not corret",ErrCode:"001"}}
	ErNotAuthUser=ErrRespone{HttpSC:401,Error:Err{Error:"User authentication failed",ErrCode:"002"}}
)