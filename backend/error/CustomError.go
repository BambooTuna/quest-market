package error

import (
	"errors"
	"fmt"
	"strings"
)

type CustomError string

const (
	NoUpdatesWereFound    CustomError = "一件も更新できませんでした"
	DuplicateRegistration CustomError = "そのメールアドレスはすでに使われています"
	AccountNotFound       CustomError = "アカウントが見つからないか、パスワードが間違っています"
	AuthenticationFailed  CustomError = "アカウントが見つからないか、パスワードが間違っています"
	SqlRequestFailed      CustomError = "Internal Server Error"
	RequestFieldEmpty     CustomError = "Internal Server Error"
	BindJSONFailed        CustomError = "不正なJson形式です"

	ItemSoldError     CustomError = "売り切れ"
	ItemNotFoundError CustomError = "アイテムが見つかりません"
	PurchaseYourself  CustomError = "自分自身の出品を買おうとしている"
	CannotBuy         CustomError = "購入できませんでした"
	LackOfMoney       CustomError = "金不足"
)

func ValidateError(fieldName string, errType string) CustomError {
	errorMessage := "を正しく入力してください"
	switch errType {
	case "email":
		errorMessage = "がメールアドレスの形式になっていません"
	case "min":
		errorMessage = "は１以上で入力してください"
	case "max":
		errorMessage = "が大きすぎます"
	case "gte":
		errorMessage = "が短すぎます"
	case "lte":
		errorMessage = "が長すぎます"
	case "required":
		errorMessage = "を入力してください"
	}
	return CustomError(fmt.Sprintf("%s%s", fieldName, errorMessage))
}

func Error(message CustomError) error {
	return errors.New(string(message))
}

func Errors(messages []CustomError) error {
	r := make([]string, len(messages))
	for i, e := range messages {
		r[i] = string(e)
	}
	return errors.New(strings.Join(r, ","))
}
