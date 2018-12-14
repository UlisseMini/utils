package test_test

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
	"utils/test"
)

func init() {
	err := os.Chdir("./testdata")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd := exec.Command("go", "build", "one.go")
	err = cmd.Run()
	if err != nil {
		fmt.Println("failed to build ./testdata/one.go")
		os.Exit(1)
	}

	cmd = exec.Command("go", "build", "two.go")
	err = cmd.Run()
	if err != nil {
		fmt.Println("failed to build ./testdata/two.go")
		os.Exit(1)
	}
}

func Test_test(t *testing.T) {
	tc := test.Case{
		Path:     "./one",
		Expected: "one",
	}

	err := tc.Run()
	if err != nil {
		t.Fatal(err)
	}

	tc = test.Case{
		Path:     "./two",
		Expected: "one",
	}

	err = tc.Run()
	if err == nil {
		t.Fatal("expected an error got nil")
	}

	epErr := `Expected "one" got "two"`
	if err.Error() != epErr {
		t.Fatalf("expected %q got %q", epErr, err.Error())
	}
}
