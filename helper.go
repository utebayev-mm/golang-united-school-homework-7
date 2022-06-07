package coverage

import "time"

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
