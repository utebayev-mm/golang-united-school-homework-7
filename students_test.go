package coverage

import (
	"log"
	"os"
	"testing"

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
	require.Equal(t, 3, len)
}

func TestLess(t *testing.T) {
	log.Println("TestLess is running")
	result1 := people.Less(0, 1)
	assert.Equal(t, false, result1)
	result2 := people.Less(1, 2)
	assert.Equal(t, true, result2)
}

func TestSwap(t *testing.T) {
	log.Println("TestSwap is running")
	person1 := people[0]
	people.Swap(0, 1)
	require.Equal(t, person1, people[1])
}

func TestNew(t *testing.T) {
	mockStringData := []int([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30})

	require.Equal(t, mockStringData, matrix.data)
	require.Equal(t, 10, matrix.cols)
	require.Equal(t, 3, matrix.rows)
}

func TestSet(t *testing.T) {
	row := 0
	col := 0
	value := 99
	matrix.Set(row, col, value)

	log.Println(matrix.data)

	require.Equal(t, matrix.data[row*matrix.cols+col], value)
}

func TestRows(t *testing.T) {
	rows := matrix.Rows()
	log.Println(rows)

	require.Equal(t, len(rows), 3)
}

func TestCols(t *testing.T) {
	cols := matrix.Cols()
	log.Println(cols)

	require.Equal(t, len(cols), 10)
}
