package forms

import (
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	form := New(url.Values{})

	if !form.Valid() {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	form := New(url.Values{})

	form.Required("missing-field1", "missing-field2", "missing-field3")
	if form.Valid() {
		t.Error("form shows valid when required fields are missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	form = New(postedData)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("form shows invalid when required fields are not missing")
	}

}

func TestForm_Has(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("email", "test@example.org")

	form := New(postedData)
	if !form.Has("email") {
		t.Error("form shows invalid when expected field is present")
	}

	form = New(postedData)
	if form.Has("invalid-field") {
		t.Error("form shows valid when expected field is not present")
	}
}

func TestForm_MinLength(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("email", "test@example.org")

	form := New(postedData)
	form.MinLength("email", 3)
	if !form.Valid() {
		t.Error("form shows invalid when field value complies with minimum length of 3")
	}

	isError := form.Errors.Get("email")
	if isError != "" {
		t.Error("error not expected but got one")
	}

	postedData = url.Values{}
	postedData.Add("email", "e@")

	form = New(postedData)
	form.MinLength("email", 100)
	if form.Valid() {
		t.Error("form shows valid when field value does not comply with minimum length of 100")
	}

	isError = form.Errors.Get("email")
	if isError == "" {
		t.Error("error expected but got none")
	}

	form = New(postedData)

	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("form shows valid with min length check for non-existent field 'x'")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("email", "test@example.org")

	form := New(postedData)
	form.IsEmail("email")
	if !form.Valid() {
		t.Error("form shows invalid when form field for email is valid")
	}

	postedData = url.Values{}
	postedData.Add("email", "e@@example")

	form = New(postedData)
	form.IsEmail("email")
	if form.Valid() {
		t.Error("form shows valid when form field for email is not valid")
	}

	form = New(url.Values{})

	form.IsEmail("x")
	if form.Valid() {
		t.Error("form passes email validation for non-existent field 'x'")
	}

}
