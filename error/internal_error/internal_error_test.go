package internal_error_test

import (
	"github.com/attestify/go-kernel/error/internal_error"
	"testing"
)

func setup(t *testing.T) {
	t.Parallel()
}

func Test_Instantiate_InternalError_Successfully(t *testing.T) {
	setup(t)

	var err error = internal_error.New("Some Internal Error")

	expectedMessages := "Some Internal Error"
	if err.Error() != expectedMessages {
		t.Errorf("the actual error messsage is not the expected error message: \n Actual: %s \n Expected: %s",
			err.Error(), expectedMessages)
	}
}
