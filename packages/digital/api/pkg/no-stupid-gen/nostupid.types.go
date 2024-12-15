package nsg

var Datatypes = struct {
	INT        string
	SMALLINT   string
	BIGINT     string
	DECIMAL    string
	NUMERIC    string
	FLOAT      string
	REAL       string
	DOUBLE     string
	CHAR       string
	VARCHAR    string
	TEXT       string
	BINARY     string
	VARBINARY  string
	BLOB       string
	DATE       string
	TIME       string
	DATETIME   string
	TIMESTAMP  string
	YEAR       string
	BOOLEAN    string
	JSON       string
	UUID       string
	GEOMETRY   string
	POINT      string
	LINESTRING string
	POLYGON    string
	ENUM       string
	SET        string
}{
	INT:        "INT",
	SMALLINT:   "SMALLINT",
	BIGINT:     "BIGINT",
	DECIMAL:    "DECIMAL",
	NUMERIC:    "NUMERIC",
	FLOAT:      "FLOAT",
	REAL:       "REAL",
	DOUBLE:     "DOUBLE PRECISION",
	CHAR:       "CHAR",
	VARCHAR:    "VARCHAR",
	TEXT:       "TEXT",
	BINARY:     "BINARY",
	VARBINARY:  "VARBINARY",
	BLOB:       "BLOB",
	DATE:       "DATE",
	TIME:       "TIME",
	DATETIME:   "DATETIME",
	TIMESTAMP:  "TIMESTAMP",
	YEAR:       "YEAR",
	BOOLEAN:    "BOOLEAN",
	JSON:       "JSON",
	UUID:       "UUID",
	GEOMETRY:   "GEOMETRY",
	POINT:      "POINT",
	LINESTRING: "LINESTRING",
	POLYGON:    "POLYGON",
	ENUM:       "ENUM",
	SET:        "SET",
}

// Constraints groups SQL constraint constants.
var Constraints = struct {
	PRIMARY_KEY    string
	FOREIGN_KEY    string
	UNIQUE         string
	NOT_NULL       string
	CHECK          string
	DEFAULT        string
	INDEX          string
	FULLTEXT       string
	AUTO_INCREMENT string
	ON_DELETE      string
	ON_UPDATE      string
}{
	PRIMARY_KEY:    "PRIMARY KEY",
	FOREIGN_KEY:    "FOREIGN KEY",
	UNIQUE:         "UNIQUE",
	NOT_NULL:       "NOT NULL",
	CHECK:          "CHECK",
	DEFAULT:        "DEFAULT",
	INDEX:          "INDEX",
	FULLTEXT:       "FULLTEXT",
	AUTO_INCREMENT: "AUTO_INCREMENT",
	ON_DELETE:      "ON DELETE",
	ON_UPDATE:      "ON UPDATE",
}

// Actions groups SQL action constants.
var Actions = struct {
	CASCADE     string
	SET_NULL    string
	NO_ACTION   string
	RESTRICT    string
	SET_DEFAULT string
}{
	CASCADE:     "CASCADE",
	SET_NULL:    "SET NULL",
	NO_ACTION:   "NO ACTION",
	RESTRICT:    "RESTRICT",
	SET_DEFAULT: "SET DEFAULT",
}
