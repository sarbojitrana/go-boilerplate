package errs



import(
	"strings"
)


type FieldError struct{					// form based errors
	Field 	string 	 `json:"field"`
	Error 	string	 `json:"error"`
}


type ActionType string 					// for cases when the sessions expire


const (
	ActionTypeRedirect ActionType = "redirect"
)


type Action struct{
	Type 	ActionType  `json:"type"`		// what to do
	Message string  	`json:"message"`	
	Value	string		`json:"value"`		//where to be redirected
}

type HTTPError struct{
	Code		string			`json:"code"`			//snake-code uppercase simple message
	Message		string			`json:"message"`		// more detailed msg
	Status		int				`json:"status"`
	Override	bool			`json:"override"`		// in places where the client is not supposed to intervene
	Errors		[]FieldError	`json:"errors"`
	Action		*Action			`json:"action"`			// if there are any actions to be taken 

}


func (e *HTTPError) Error() string{
	return e.Message
}

func (e *HTTPError) Is(target error) bool{
	_,ok := target.(*HTTPError)

	return ok
}




func (e *HTTPError) WithMessage(message string) *HTTPError{
	return &HTTPError{
		Code:	e.Code,
		Message: message,
		Status: e.Status,
		Override: e.Override,
		Errors: e.Errors,
		Action: e.Action,
	}
}


func MakeUpperCaseWithUnderscores(str string) string{
	return strings.ToUpper(strings.ReplaceAll(str, " ", "_"))
}