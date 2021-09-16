package unitTest

import (
	repositorium "W/repository"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
	"unicode/utf8"
)

// setup and tear down for tests
func TestMain(m *testing.M) {
	os.Chdir("..")
	initRepo(&repo)
	controller.InitController()
	controller.InitRepo()
	controller.InitServer()
	var repo = repositorium.DbCustomerRepo{}
	ensureTableExists(&repo)
	code := m.Run()
	clearTable(&repo)
	os.Exit(code)
}

func ensureTableExists(repo *repositorium.DbCustomerRepo) {
	repo.Initialize()

}
func clearTable(repo *repositorium.DbCustomerRepo) {
	db, _ := repo.Db.DB()
	db.Exec("DELETE FROM customers")
	db.Exec("ALTER SEQUENCE customers_id_seq RESTART WITH 1")
}

func AssertEquals(expected interface{}, actual interface{}, t *testing.T) {
	if expected != actual {
		t.Fatalf("expected is %v, actual is %v", expected, actual)
	}
}

func AssertContains(expected string, res string, t *testing.T) {
	if strings.Contains(strings.ReplaceAll(stripHtmlTags(res), "\n", ""), expected) {
		print("Pass")
		return
	} else {
		t.Fatalf("Reponse body did not contain %s", expected)
	}
}

func AssertListContainsCustomer(e repositorium.Customer, s []repositorium.Customer, t *testing.T) bool {
	for _, a := range s {
		if a.Email == e.Email {
			return true
		}
	}
	t.Fatalf("List does not contain %v", e)
	return false
}
func InitRandomCustomer() (repositorium.CustomerPogo, string) {
	rand.Seed(time.Now().UnixNano())
	s1 := strconv.FormatInt(rand.Int63(), 10)
	testCustomer := repositorium.CustomerPogo{
		FirstName: "Testing_Name" + s1,
		LastName:  "Testing_LastName" + s1,
		Email:     "Email@unitTest" + s1,
	}
	return testCustomer, s1
}
func InitConcreteCustomer() (repositorium.CustomerPogo, string) {
	rand.Seed(time.Now().UnixNano())
	s1 := strconv.FormatInt(rand.Int63(), 10)
	testCustomer := repositorium.CustomerPogo{
		FirstName: "Testing_Name",
		LastName:  "Testing_LastName",
		Email:     "Email@unitTest" + s1,
	}
	return testCustomer, s1
}

func initRepo(repo *repositorium.DbCustomerRepo) {
	repo.Initialize()
}

const (
	htmlTagStart = 60 // Unicode `<`
	htmlTagEnd   = 62 // Unicode `>`
)

func stripHtmlTags(s string) string {
	// Setup a string builder and allocate enough memory for the new string.
	var builder strings.Builder
	builder.Grow(len(s) + utf8.UTFMax)

	in := false // True if we are inside an HTML tag.
	start := 0  // The index of the previous start tag character `<`
	end := 0    // The index of the previous end tag character `>`

	for i, c := range s {
		// If this is the last character and we are not in an HTML tag, save it.
		if (i+1) == len(s) && end >= start {
			builder.WriteString(s[end:])
		}

		// Keep going if the character is not `<` or `>`
		if c != htmlTagStart && c != htmlTagEnd {
			continue
		}

		if c == htmlTagStart {
			// Only update the start if we are not in a tag.
			// This make sure we strip out `<<br>` not just `<br>`
			if !in {
				start = i
			}
			in = true

			// Write the valid string between the close and start of the two tags.
			builder.WriteString(s[end:start])
			continue
		}
		// else c == htmlTagEnd
		in = false
		end = i + 1
	}
	s = builder.String()
	return s
}
