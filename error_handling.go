package art

import (
	"fmt"
)

func HandleError(errorCode string) error {
	var errMsg string
	switch errorCode {
	case "un_brackets":
		errMsg = "\033[31mUnbalanced square brackets.\033[0m"
	case "no_space":
		errMsg = "\033[31mThe arguments are not separated by a space.\033[0m"
	case "missing_argument":
		errMsg = "\033[31mThere is no second argument, i.e., empty strings are invalid.\033[0m"
	case "not_a_number":
		errMsg = "\033[31mThe first argument is not a number.\033[0m"
	case "missing_decode_flag":
		errMsg = "\033[31mMultiline art detected. Please use the \033[0;30m-m\033[0m \033[31mflag for multiline decoding.\033[0m"
	case "missing_encode_flag":
		errMsg = "\033[31mMultiline art detected. Please use the \033[0;30m-me\033[0m \033[31mflag for multiline encoding.\033[0m"
	case "missing_flag":
		errMsg = "\033[31mMultiline art detected without decoding or encoding flag.\033[0m"
	case "missing_-e_flag":
		errMsg = "\033[31mThe first argument is not a number.\033[0m"
	default:
		errMsg = "\033[31mUnknown error.\033[0m"
	}

	return fmt.Errorf("%s", errMsg)
}
