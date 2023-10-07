package main

import "testing"

func TestTrue(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			t.Errorf("Unexpected panic occured: %v", err)
		}
	}()
	main()
}
