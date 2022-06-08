package coverage

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

var people People
var matrix *Matrix
var matrixFail *Matrix

func MockPeople() People {
	return People{
		Person{
			firstName: "testFirstName1",
			lastName:  "testLastName1",
			birthDay: time.Date(
				2008, 11, 17, 20, 34, 58, 651387237, time.UTC),
		},
		Person{
			firstName: "testFirstName2",
			lastName:  "testLastName2",
			birthDay: time.Date(
				2010, 11, 17, 20, 34, 58, 651387237, time.UTC),
		},
		Person{
			firstName: "testFirstName3",
			lastName:  "testLastName3",
			birthDay: time.Date(
				2009, 11, 17, 20, 34, 58, 651387237, time.UTC),
		},
		Person{
			firstName: "testFirstName3",
			lastName:  "testLastName4",
			birthDay: time.Date(
				2009, 11, 17, 20, 34, 58, 651387237, time.UTC),
		},
		Person{
			firstName: "testFirstName",
			lastName:  "testLastName4",
			birthDay: time.Date(
				2009, 11, 17, 20, 34, 58, 651387237, time.UTC),
		},
	}
}

func TestMain(m *testing.M) {
	// Part 1 preparations
	people = MockPeople()

	// Part 2 preparations
	mockString := "1 2 3 4 5 6 7 8 9 10 \n 11 12 13 14 15 16 17 18 19 20 \n 21 22 23 24 25 26 27 28 29 30"
	var err error
	matrix, err = New(mockString)
	if err != nil {
		panic(err)
	}

	code := m.Run()
	os.Exit(code)
}

func TestLen(t *testing.T) {
	log.Println("TestLen is running")
	len := people.Len()
	require.Equal(t, 5, len)
}

func TestLess(t *testing.T) {
	log.Println("TestLess is running")
	result := people.Less(0, 1)
	assert.Equal(t, false, result)
	result = people.Less(1, 2)
	assert.Equal(t, true, result)
	result = people.Less(2, 3)
	assert.Equal(t, true, result)
	result = people.Less(3, 4)
	assert.Equal(t, false, result)
}

func TestSwap(t *testing.T) {
	log.Println("TestSwap is running")
	person1 := people[0]
	people.Swap(0, 1)
	require.Equal(t, person1, people[1])
}

func TestNew(t *testing.T) {
	log.Println("TestNew is running")
	mockStringData := []int([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30})
	require.Equal(t, mockStringData, matrix.data)
	require.Equal(t, 10, matrix.cols)
	require.Equal(t, 3, matrix.rows)

}

func TestNewRowLength(t *testing.T) {
	log.Println("TestRowLength is running")
	var err error
	mockString := "1 2 3 4 5 6 7 8 9 10 \n 11 12 13 14 15 16 17 18 19 20 \n 21 22 23 24 25 26 27 28 29"
	matrixFail, err = New(mockString)
	if err != nil {
		log.Println(err)
	}
	require.Equal(t, err, fmt.Errorf("Rows need to be the same length"))
}

func TestNewRowAtoi(t *testing.T) {
	log.Println("TestRowAtoi is running")
	var err error
	mockString := "a b c 4 5 6 7 8 9 10 \n 11 12 13 14 15 16 17 18 19 20 \n 21 22 23 24 25 26 27 28 29"
	matrixFail, err = New(mockString)
	if err != nil {
		log.Println(err)
	}
	_, errAtoi := strconv.Atoi("a")
	require.Equal(t, err, errAtoi)
}

func TestSet(t *testing.T) {
	log.Println("TestSet is running")
	row := 0
	col := 0
	value := 99
	matrix.Set(row, col, value)

	log.Println(matrix.data)

	require.Equal(t, matrix.data[row*matrix.cols+col], value)
}

func TestSetWrongNumber(t *testing.T) {
	log.Println("TestSetWrongNumber is running")
	row := -1
	col := -1
	value := 99
	result := matrix.Set(row, col, value)

	log.Println(matrix.data)

	require.Equal(t, false, result)
}

func TestRows(t *testing.T) {
	log.Println("TestRows is running")
	rows := matrix.Rows()
	log.Println(rows)

	require.Equal(t, len(rows), 3)
}

func TestCols(t *testing.T) {
	log.Println("TestCols is running")
	cols := matrix.Cols()
	log.Println(cols)

	require.Equal(t, len(cols), 10)
}
