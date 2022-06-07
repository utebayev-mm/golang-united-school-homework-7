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

func TestMain(m *testing.M) {
	people = MockPeople()
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
