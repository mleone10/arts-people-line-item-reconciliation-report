package artspeople_test

import (
	"testing"

	"github.com/mleone10/arts-people-reconciler/artspeople"
)

func TestNewCurrencyFromString_ValidCurrencyString(t *testing.T) {
	actual, _ := artspeople.NewCurrencyFromString("1234.56")
	expected := artspeople.Currency(123456)

	if actual != expected {
		t.Errorf("Expected currency %v, got %v", expected, actual)
	}
}

func TestNewCurrencyFromString_TooManyDecimels(t *testing.T) {
	_, err := artspeople.NewCurrencyFromString("1234.56.78")
	if err == nil {
		t.Errorf("Expected an error due to too many decimels, but received none")
	}
}

func TestNewCurrencyFromString_EmptyString(t *testing.T) {
	actual, err := artspeople.NewCurrencyFromString("")
	expected := artspeople.Currency(0)

	if err != nil {
		t.Errorf("Unexpected error when parsing empty string")
	}

	if actual != expected {
		t.Errorf("Expected currency %v, got %v", expected, actual)
	}
}

func TestNewCurrencyFromString_NegativeAmount(t *testing.T) {
	actual, _ := artspeople.NewCurrencyFromString("-1234.56")
	expected := artspeople.Currency(-123456)

	if actual != expected {
		t.Errorf("Expected currency %v, got %v", expected, actual)
	}
}
