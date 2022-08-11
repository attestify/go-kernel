package validation_error_test

import (
	"github.com/attestify/go-kernel/error/validation_error"
	"testing"
)

func setup(t *testing.T) {
	t.Parallel()
}

func Test_Instantiate_ValidationError_Successfully(t *testing.T) {
	setup(t)

	var err error = validation_error.New("Some Random Error")

	expectedMessages := "Some Random Error"
	if err.Error() != expectedMessages {
		t.Errorf("the actual error messsage is not the expected error message: \n Actual: %s \n Expected: %s",
			err.Error(), expectedMessages)
	}
}
