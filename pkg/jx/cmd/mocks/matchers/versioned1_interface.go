// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	versioned1 "github.com/banzaicloud/bank-vaults/operator/pkg/client/clientset/versioned"
	"github.com/petergtz/pegomock"
	"reflect"
)

func AnyVersioned1Interface() versioned1.Interface {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(versioned1.Interface))(nil)).Elem()))
	var nullValue versioned1.Interface
	return nullValue
}

func EqVersioned1Interface(value versioned1.Interface) versioned1.Interface {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue versioned1.Interface
	return nullValue
}
