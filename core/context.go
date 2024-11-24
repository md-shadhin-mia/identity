package core

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type Context struct {
	*gin.Context
}

func NewContext(c *gin.Context) *Context {
	return &Context{Context: c}
}

type HandlerFunc func(*Context)

func (c *Context) BindValidateJson(obj any) error {
	validate, trans := setupValidator()

	if err := c.ShouldBindJSON(&obj); err != nil {
		return err
	}

	if err := validate.Struct(obj); err != nil {
		errors := formatValidationErrors(err, trans)
		c.JSON(http.StatusBadRequest, gin.H{"validation_errors": errors})
		return err
	}
	return nil
}

func setupValidator() (*validator.Validate, ut.Translator) {
	// Initialize validator
	validate := validator.New()

	// Register function to get json tag as field name
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return fld.Name
		}
		return name
	})

	// Initialize translator
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(validate, trans)

	// Register custom validation messages
	registerCustomTranslations(validate, trans)

	return validate, trans
}

// 3. Custom error message registration
func registerCustomTranslations(validate *validator.Validate, trans ut.Translator) {
	// Register custom translation for required tag
	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is a required field", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	// Register custom translation for email tag
	validate.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "{0} must be a valid email address", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())
		return t
	})

	// Register custom translation for min tag
	validate.RegisterTranslation("min", trans, func(ut ut.Translator) error {
		return ut.Add("min", "{0} must be at least {1} characters long", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("min", fe.Field(), fe.Param())
		return t
	})
}

// 4. Error handling function
func formatValidationErrors(err error, trans ut.Translator) map[string]string {
	errors := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			errors[e.Field()] = e.Translate(trans)
		}
	}
	return errors
}
