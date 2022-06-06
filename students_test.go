package coverage

import (
	"os"
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
	}
}

func TestLen(t *testing.T) {
	people := MockPeople()
	len := people.Len()
	require.Equal(t, 3, len)
}

func TestLess(t *testing.T) {
	people := MockPeople()
	result1 := people.Less(0, 1)
	assert.Equal(t, false, result1)
	result2 := people.Less(1, 2)
	assert.Equal(t, true, result2)
}

func TestSwap(t *testing.T) {
	people := MockPeople()
	person1 := people[0]
	people.Swap(0, 1)
	require.Equal(t, person1, people[1])
}
