package util

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"

	"github.com/lib/pq"
)

const (
	CodeNumericValueOutOfRange    = "22003"
	CodeInvalidTextRepresentation = "22P02"
	CodeNotNullViolation          = "23502"
	CodeForeignKeyViolation       = "23503"
	CodeUniqueViolation           = "23505"
	CodeCheckViolation            = "23514"
	CodeLockNotAvailable          = "55P03"
)

// Error is a human-readable database error. Message should always be a
// non-empty, readable string, and is returned when you call err.Error(). The
// other fields may or may not be empty.
type PgxError struct {
	Message    string `json:"message"`
	Code       string `json:"code"`
	Constraint string `json:"constraint"`
	Severity   string `json:"severity"`
	Routine    string `json:"routine"`
	Table      string `json:"table"`
	Detail     string `json:"detail"`
	Column     string `json:"column"`
}

func (dbe *PgxError) Error() string {
	return dbe.Message
}

// Constraint is a custom database check constraint you've defined, like "CHECK
// balance > 0". Postgres doesn't define a very useful message for constraint
// failures (new row for relation "accounts" violates check constraint), so you
// can define your own. The Name should be the name of the constraint in the
// database. Define GetError to provide your own custom error handler for this
// constraint failure, with a custom message.
type Constraint struct {
	Name     string
	GetError func(*pq.Error) *PgxError
}

var constraintMap = map[string]*Constraint{}
var constraintMu sync.RWMutex

// RegisterConstraint tells dberror about your custom constraint and its error
// handling. RegisterConstraint panics if you attempt to register two
// constraints with the same name.
func RegisterConstraint(c *Constraint) {
	constraintMu.Lock()
	defer constraintMu.Unlock()
	if _, dup := constraintMap[c.Name]; dup {
		panic("dberror: RegisterConstraint called twice for name " + c.Name)
	}
	constraintMap[c.Name] = c
}

// capitalize the first letter in the string
func capitalize(s string) string {
	r, size := utf8.DecodeRuneInString(s)
	return fmt.Sprintf("%c", unicode.ToTitle(r)) + s[size:]
}

var columnFinder = regexp.MustCompile(`Key \((.+)\)=`)
var valueFinder = regexp.MustCompile(`Key \(.+\)=\((.+)\)`)

// findColumn finds the column in the given pq Detail error string. If the
// column does not exist, the empty string is returned.
//
// detail can look like this:
//    Key (id)=(3c7d2b4a-3fc8-4782-a518-4ce9efef51e7) already exists.
func findColumn(detail string) string {
	results := columnFinder.FindStringSubmatch(detail)
	if len(results) < 2 {
		return ""
	} else {
		return results[1]
	}
}

// findColumn finds the column in the given pq Detail error string. If the
// column does not exist, the empty string is returned.
//
// detail can look like this:
//    Key (id)=(3c7d2b4a-3fc8-4782-a518-4ce9efef51e7) already exists.
func findValue(detail string) string {
	results := valueFinder.FindStringSubmatch(detail)
	if len(results) < 2 {
		return ""
	} else {
		return results[1]
	}
}

var foreignKeyFinder = regexp.MustCompile(`not present in table "(.+)"`)

// findForeignKeyTable finds the referenced table in the given pq Detail error
// string. If we can't find the table, we return the empty string.
//
// detail can look like this:
//    Key (account_id)=(91f47e99-d616-4d8c-9c02-cbd13bceac60) is not present in table "accounts"
func findForeignKeyTable(detail string) string {
	results := foreignKeyFinder.FindStringSubmatch(detail)
	if len(results) < 2 {
		return ""
	}
	return results[1]
}

var parentTableFinder = regexp.MustCompile(`update or delete on table "([^"]+)"`)

func findParentTable(message string) string {
	match := parentTableFinder.FindStringSubmatch(message)
	if len(match) < 2 {
		return ""
	}
	return match[1]
}

func FilterColumns(columns *string, defaultColumns string, table string) string {
    if columns == nil {
        return defaultColumns
    }

	var builder strings.Builder
	for i, col := range strings.Split(*columns, ",") {
		if i > 0 {
			builder.WriteString(",")
		}

		builder.WriteString(table + "." + col)
	}

    return builder.String()
}

func GetError(err error) error {
	if err == nil {
		return nil
	}

	switch sqlxerr := err.(type) {
	case *pq.Error:
		fmt.Println(sqlxerr.Code)
		switch sqlxerr.Code {
		case CodeUniqueViolation:
			columnName := findColumn(sqlxerr.Detail)
			if columnName == "" {
				columnName = "value"
			}
			valueName := findValue(sqlxerr.Detail)
			var msg string
			if valueName == "" {
				msg = fmt.Sprintf("A %s already exists with that value", columnName)
			} else {
				msg = fmt.Sprintf("A %s already exists with this value (%s)", columnName, valueName)
			}
			dbe := &PgxError{
				Message:    msg,
				Code:       string(sqlxerr.Code),
				Severity:   sqlxerr.Severity,
				Constraint: sqlxerr.Constraint,
				Table:      sqlxerr.Table,
				Detail:     sqlxerr.Detail,
			}
			if columnName != "value" {
				dbe.Column = columnName
			}
			return dbe
		case CodeForeignKeyViolation:
			columnName := findColumn(sqlxerr.Detail)
			if columnName == "" {
				columnName = "value"
			}
			foreignKeyTable := findForeignKeyTable(sqlxerr.Detail)
			var tablePart string
			if foreignKeyTable == "" {
				tablePart = "in the parent table"
			} else {
				tablePart = fmt.Sprintf("in the %s table", foreignKeyTable)
			}
			valueName := findValue(sqlxerr.Detail)
			var msg string
			switch {
			case strings.Contains(sqlxerr.Message, "update or delete"):
				parentTable := findParentTable(sqlxerr.Message)
				// in this case pqerr.Table contains the child table. there's
				// probably more work we could do here.
				msg = fmt.Sprintf("Can't update or delete %[1]s records because the %[1]s %s (%s) is still referenced by the %s table", parentTable, columnName, valueName, sqlxerr.Table)
			case valueName == "":
				msg = fmt.Sprintf("Can't save to %s because the %s isn't present %s", sqlxerr.Table, columnName, tablePart)
			default:
				msg = fmt.Sprintf("Can't save to %s because the %s (%s) isn't present %s", sqlxerr.Table, columnName, valueName, tablePart)
			}
			return &PgxError{
				Message:    msg,
				Code:       string(sqlxerr.Code),
				Column:     sqlxerr.Column,
				Constraint: sqlxerr.Constraint,
				Table:      sqlxerr.Table,
				Routine:    sqlxerr.Routine,
				Severity:   sqlxerr.Severity,
			}
		case CodeNumericValueOutOfRange:
			msg := strings.Replace(sqlxerr.Message, "out of range", "too large or too small", 1)
			return &PgxError{
				Message:  capitalize(msg),
				Code:     string(sqlxerr.Code),
				Severity: sqlxerr.Severity,
			}
		case CodeInvalidTextRepresentation:
			msg := sqlxerr.Message
			// Postgres tweaks with the message, play whack-a-mole until we
			// figure out a better method of dealing with these.
			if !strings.Contains(sqlxerr.Message, "invalid input syntax for type") {
				msg = strings.Replace(sqlxerr.Message, "input syntax for", "input syntax for type", 1)
			}
			msg = strings.Replace(msg, "input value for enum ", "", 1)
			msg = strings.Replace(msg, "invalid", "Invalid", 1)
			return &PgxError{
				Message:  msg,
				Code:     string(sqlxerr.Code),
				Severity: sqlxerr.Severity,
			}
		case CodeNotNullViolation:
			msg := fmt.Sprintf("No %[1]s was provided. Please provide a %[1]s", sqlxerr.Column)
			return &PgxError{
				Message:  msg,
				Code:     string(sqlxerr.Code),
				Column:   sqlxerr.Column,
				Table:    sqlxerr.Table,
				Severity: sqlxerr.Severity,
			}
		case CodeCheckViolation:
			constraintMu.RLock()
			c, ok := constraintMap[sqlxerr.Constraint]
			constraintMu.RUnlock()
			if ok {
				return c.GetError(sqlxerr)
			} else {
				return &PgxError{
					Message:    sqlxerr.Message,
					Code:       string(sqlxerr.Code),
					Column:     sqlxerr.Column,
					Table:      sqlxerr.Table,
					Severity:   sqlxerr.Severity,
					Constraint: sqlxerr.Constraint,
				}
			}
		default:
			return &PgxError{
				Message:    sqlxerr.Message,
				Code:       string(sqlxerr.Code),
				Column:     sqlxerr.Column,
				Constraint: sqlxerr.Constraint,
				Table:      sqlxerr.Table,
				Routine:    sqlxerr.Routine,
				Severity:   sqlxerr.Severity,
			}
		}
	default:
		fmt.Println(err)
		return err
	}
}
