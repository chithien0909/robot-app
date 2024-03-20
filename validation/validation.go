package validation

import (
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

type StructValidation struct {
	Validator *validator.Validate
	Trans     ut.Translator
}

func New() *StructValidation {
	translator := en.New()
	uni := ut.New(translator, translator)
	trans, _ := uni.GetTranslator("en")
	return &StructValidation{
		Validator: validator.New(),
		Trans:     trans,
	}
}

func (v *StructValidation) RegisterValidate() {
	v.registerDateRange()
	v.registerEnums()

}

func (v *StructValidation) registerDateRange() {
	_ = v.Validator.RegisterValidation("dateRange", func(fl validator.FieldLevel) bool {
		if fl.Field().IsZero() {
			return true
		}
		if fl.Field().String() == "" {
			return true
		}
		values := strings.Split(fl.Field().String(), ",")
		return len(values) == 2
	})
	_ = v.Validator.RegisterTranslation("dateRange", v.Trans, func(ut ut.Translator) error {
		return ut.Add("dateRange", "{0} must be date range format: start,end, example: 2020-01-01,2020-01-02", true)
	}, func(ut ut.Translator, err validator.FieldError) string {
		t, _ := ut.T("dateRange", err.Field())
		return t
	})
}

func (v *StructValidation) registerEnums() {

	_ = v.Validator.RegisterValidation("enum", func(fl validator.FieldLevel) bool {
		enumString := fl.Param()
		value := strings.ReplaceAll(fl.Field().String(), "-", "")
		enumSlice := strings.Split(enumString, ";")
		for _, val := range enumSlice {
			if value == val {
				return true
			}
		}
		return false
	})

	_ = v.Validator.RegisterTranslation("enum", v.Trans, func(ut ut.Translator) error {
		return ut.Add("enum", "{0}", true)
	}, func(ut ut.Translator, err validator.FieldError) string {
		replacer := *strings.NewReplacer(";", ",")
		resultErrors := err.Field() + " must be one of " + replacer.Replace(err.Param())

		t, _ := ut.T("enum", resultErrors)
		return t
	})

}

func (v *StructValidation) Validate(i interface{}) error {
	err := v.Validator.Struct(i)
	if err == nil {
		return nil
	}

	transErrors := make([]string, 0)
	for _, e := range err.(validator.ValidationErrors) {
		transErrors = append(transErrors, e.Translate(v.Trans))
	}
	return fmt.Errorf("%s", strings.Join(transErrors, " \n "))
}
