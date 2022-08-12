package already_exists_test

import (
	"github.com/attestify/go-kernel/error/already_exists"
	"testing"
)

func setup(t *testing.T) {
	t.Parallel()
}

func Test_Instantiate_Error_Successfully(t *testing.T) {
	setup(t)

	var err error = already_exists.New("Some Already Exists Error")

	expectedMessages := "Some Already Exists Error"
	if err.Error() != expectedMessages {
		t.Errorf("the actual error messsage is not the expected error message: \n Actual: %s \n Expected: %s",
			err.Error(), expectedMessages)
	}
}
