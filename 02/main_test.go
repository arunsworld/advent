package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func Test_A_Password_Validator(t *testing.T) {
	t.Run("on given password with policies in raw format", func(t *testing.T) {
		pwdData := `1-3 a: abcde
					1-3 b: cdefg
					2-9 c: ccccccccc`
		validator, err := NewPwdValidator(strings.NewReader(pwdData))
		if err != nil {
			t.Fatal(err)
		}
		t.Run("separates policy from password", func(t *testing.T) {
			expectedValidator := PwdValidator{
				&PolicyAndData{
					Policy: Policy{letter: 'a', min: 1, max: 3},
					Data:   "abcde",
				},
				&PolicyAndData{
					Policy: Policy{letter: 'b', min: 1, max: 3},
					Data:   "cdefg",
				},
				&PolicyAndData{
					Policy: Policy{letter: 'c', min: 2, max: 9},
					Data:   "ccccccccc",
				},
			}
			if !reflect.DeepEqual(validator, expectedValidator) {
				fmt.Printf("%#v", validator)
				t.Fatal("validator not as expected")
			}
			t.Run("validates password against old policy", func(t *testing.T) {
				validator.Validate(OldPolicyStandard)

				expectedValidator[0].IsValid = true
				expectedValidator[2].IsValid = true
				if !reflect.DeepEqual(validator, expectedValidator) {
					fmt.Printf("%#v", validator)
					t.Fatal("validator after validation not as expected")
				}
				t.Run("and counts valid passwords", func(t *testing.T) {
					if validator.Count() != 2 {
						t.Fatalf("valid counts didn't match. Expected 2, got: %d", validator.Count())
					}
				})
			})
			t.Run("validates password against new policy", func(t *testing.T) {
				validator.Validate(NewPolicyStandard)

				expectedValidator[0].IsValid = true
				expectedValidator[2].IsValid = false
				if !reflect.DeepEqual(validator, expectedValidator) {
					fmt.Printf("%#v", validator)
					t.Fatal("validator after validation not as expected")
				}
				t.Run("and counts valid passwords", func(t *testing.T) {
					if validator.Count() != 1 {
						t.Fatalf("valid counts didn't match. Expected 1, got: %d", validator.Count())
					}
				})
			})
		})
	})
}
