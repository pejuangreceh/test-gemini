package requests

import (
	"errors"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"gopkg.in/thedevsaddam/govalidator.v1"
)

func BlacklistValidation(field string) validation.RuleFunc {
	return func(value interface{}) error {
		val, ok := value.(string)
		if !ok {
			return errors.New("The " + field + " is not a string")
		}

		if val == "" {
			return nil
		}

		match, _ := regexp.MatchString(`^[^'"\[\]<>\{\}]+$`, val)
		if !match {
			return errors.New("The " + field + " contains unsafe characters")
		}

		return nil
	}
}
func (r TextRequest) Validate(validationType string) interface{} {
	if validationType == "thedevsaddam" {
		return govalidator.MapData{
			"question": []string{"required", "blacklist"},
		}
	} else {
		return validation.ValidateStruct(&r,
			validation.Field(&r.Question, validation.Required, validation.By(BlacklistValidation("question"))),
		)
	}

	return nil
}
