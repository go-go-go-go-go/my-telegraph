// Code generated by ent, DO NOT EDIT.

package page

const (
	// Label holds the string label denoting the page type in the database.
	Label = "page"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPath holds the string denoting the path field in the database.
	FieldPath = "path"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldAuthorName holds the string denoting the author_name field in the database.
	FieldAuthorName = "author_name"
	// FieldAuthorURL holds the string denoting the author_url field in the database.
	FieldAuthorURL = "author_url"
	// FieldImageURL holds the string denoting the image_url field in the database.
	FieldImageURL = "image_url"
	// FieldViews holds the string denoting the views field in the database.
	FieldViews = "views"
	// FieldCanEdit holds the string denoting the can_edit field in the database.
	FieldCanEdit = "can_edit"
	// Table holds the table name of the page in the database.
	Table = "pages"
)

// Columns holds all SQL columns for page fields.
var Columns = []string{
	FieldID,
	FieldPath,
	FieldTitle,
	FieldContent,
	FieldURL,
	FieldDescription,
	FieldAuthorName,
	FieldAuthorURL,
	FieldImageURL,
	FieldViews,
	FieldCanEdit,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// PathValidator is a validator for the "path" field. It is called by the builders before save.
	PathValidator func(string) error
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// DefaultURL holds the default value on creation for the "url" field.
	DefaultURL string
	// DefaultDescription holds the default value on creation for the "description" field.
	DefaultDescription string
	// DefaultAuthorName holds the default value on creation for the "author_name" field.
	DefaultAuthorName string
	// DefaultAuthorURL holds the default value on creation for the "author_url" field.
	DefaultAuthorURL string
	// DefaultImageURL holds the default value on creation for the "image_url" field.
	DefaultImageURL string
	// DefaultViews holds the default value on creation for the "views" field.
	DefaultViews int
	// DefaultCanEdit holds the default value on creation for the "can_edit" field.
	DefaultCanEdit bool
)
