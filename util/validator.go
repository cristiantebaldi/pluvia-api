package util

import (
	"encoding/json"
	"regexp"
	"strings"
	"unicode"

	"github.com/go-playground/locales/br"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/pt_BR"
)

var translation ut.Translator

type RequestError struct {
	Fields []Field `json:"fields"`
}

type Field struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r RequestError) Error() string {
	payload, _ := json.Marshal(r)
	return string(payload)
}

func NewValidator() *validator.Validate {
	br := br.New()
	universalTrasnlator := ut.New(br, br)

	translation, _ = universalTrasnlator.GetTranslator("br")

	validate := validator.New()
	validate.RegisterValidation("ISO8601date", isISO8601Date)

	validate.RegisterTranslation(
		"ISO8601date",
		translation,
		registerTraslator("ISO8601", "{0} deve seguir o formato ISO 8601 (YYYY-MM-DDTHH:mm:ssZ)"),
		translate,
	)

	pt_BR.RegisterDefaultTranslations(validate, translation)

	return validate
}

func isISO8601Date(fl validator.FieldLevel) bool {
	ISO8601DateRegexString :=  "^(?:[1-9]\\d{3}-(?:(?:0[1-9]|1[0-2])-(?:0[1-9]|1\\d|2[0-8])|(?:0[13-9]|1[0-2])-(?:29|30)|(?:0[13578]|1[02])-31)|(?:[1-9]\\d(?:0[48]|[2468][048]|[13579][26])|(?:[2468][048]|[13579][26])00)-02-29)T(?:[01]\\d|2[0-3]):[0-5]\\d:[0-5]\\d(?:\\.\\d{1,9})?(?:Z|[+-][01]\\d:[0-5]\\d)$"
	ISO8601DateRegex := regexp.MustCompile(ISO8601DateRegexString)
	return ISO8601DateRegex.MatchString(fl.Field().String())
} 

func registerTraslator(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(translation ut.Translator) error {
		if err := translation.Add(tag, msg, false); err != nil {
			return err
		}

		return nil
	}
}

func translate(trans ut.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		panic(fe.(error).Error())
	}

	return msg
}

func HandleValidatorFieldError(err error) error {
	errs := err.(validator.ValidationErrors)

	requestError := RequestError{}

	for _, e := range errs {
		field := Field{
			Field: getFormatedField(e.Namespace()),
			Message: e.Translate(translation),
		}

		requestError.Fields = append(requestError.Fields, field)
	}

	return requestError
}

func getFormatedField(field string) string {
	fieldSplited := strings.Split(field, ".")
	fieldConverted := ""

	for i, field := range fieldSplited {
		if i > 0 {
			fieldRune := []rune(field)
			fieldRune[0] = unicode.ToLower(fieldRune[0])
			field = string(fieldRune)

			fieldConverted += field

			if i < len(fieldSplited)-1 {
				fieldConverted += "."
			}
		}
	}
	return fieldConverted
}