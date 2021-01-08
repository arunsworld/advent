package main

import (
	"log"
	"reflect"
	"testing"
)

func Test_That_Expense_Calculator(t *testing.T) {
	t.Run("when given the contents of an expense report", func(t *testing.T) {
		expenseReport := `1721
					979
					366
					299
					675
					1456`
		t.Run("is able to convert it to an array of numbers", func(t *testing.T) {
			expensesList, err := expenseReportContentsToList(expenseReport)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(expensesList, []int{1721, 979, 366, 299, 675, 1456}) {
				log.Println(expensesList)
				t.Fatal("expenses list not as expected")
			}
			t.Run("then multiply the first two numbers that when added give 2020", func(t *testing.T) {
				result, err := productOfTwoSum(expensesList, 2020)
				if err != nil {
					log.Fatal(err)
				}
				if result != 514579 {
					log.Println(result)
					t.Fatal("final result of two sum not as expected")
				}
			})
			t.Run("then multiply the first three numbers that when added give 2020", func(t *testing.T) {
				result, err := productOfThreeSum(expensesList, 2020)
				if err != nil {
					log.Fatal(err)
				}
				if result != 241861950 {
					log.Println(result)
					t.Fatal("final result of three sum not as expected")
				}
			})
		})
	})
}
