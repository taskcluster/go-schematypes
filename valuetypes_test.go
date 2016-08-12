package schematypes

import "testing"

var (
	aInt     int
	aInt8    int8
	aInt16   int16
	aInt32   int32
	aInt64   int64
	aUint    uint
	aUint8   uint8
	aUint16  uint16
	aUint32  uint32
	aUint64  uint64
	aString  string
	aFloat32 float32
	aFloat64 float64
	aBool    bool
	pInt     = &aInt
	pInt8    = &aInt8
	pInt16   = &aInt16
	pInt32   = &aInt32
	pInt64   = &aInt64
	pUint    = &aUint
	pUint8   = &aUint8
	pUint16  = &aUint16
	pUint32  = &aUint32
	pUint64  = &aUint64
	pString  = &aString
	pFloat32 = &aFloat32
	pFloat64 = &aFloat64
	pBool    = &aBool
)

func TestInteger(t *testing.T) {
	testCase{
		Schema: Integer{
			MetaData: MetaData{
				Title:       "my-title",
				Description: "my-description",
			},
			Minimum: -240,
			Maximum: 240,
		},
		Match: `{
      "type": "integer",
      "title": "my-title",
      "description": "my-description",
      "minimum": -240,
      "maximum": 240
    }`,
		Valid: []string{
			"32", "-32", "0",
		},
		Invalid: []string{
			"0.4", "1.3", "242", "-254",
		},
		TypeMatch: []interface{}{
			pInt16,
			pInt32,
			pInt64,
			pInt,
		},
		TypeMismatch: []interface{}{
			aInt8,
			aInt16,
			aInt32,
			aInt64,
			aUint8,
			aUint16,
			aUint32,
			aUint64,
			aString,
			aFloat32,
			aFloat64,
			aBool,
			pInt8,
			pUint8,
			pUint16,
			pUint32,
			pUint64,
			pString,
			pFloat32,
			pFloat64,
			pBool,
		},
	}.Test(t)
}

func TestNumber(t *testing.T) {
	var a float64
	var b float32
	var c int16
	var d string
	var e []string
	var f struct{ float64 }
	var g uint64
	var h int8
	testCase{
		Schema: Number{
			MetaData: MetaData{
				Title:       "my-title",
				Description: "my-description",
			},
			Minimum: -240.5,
			Maximum: 240,
		},
		Match: `{
      "type": "number",
      "title": "my-title",
      "description": "my-description",
      "minimum": -240.5,
      "maximum": 240
    }`,
		Valid: []string{
			"32", "-32", "0", "32.3", "-55.3",
		},
		Invalid: []string{
			"242", "-254",
		},
		TypeMatch: []interface{}{
			&a, &b,
		},
		TypeMismatch: []interface{}{
			&c, &d, &e, &f, &g, &h,
		},
	}.Test(t)
}

func TestBoolean(t *testing.T) {
	testCase{
		Schema: Boolean{
			MetaData: MetaData{
				Title:       "my-title",
				Description: "my-description",
			},
		},
		Match: `{
      "type": "boolean",
      "title": "my-title",
      "description": "my-description"
    }`,
		Valid: []string{
			"true", "false",
		},
		Invalid: []string{
			"242", "-254", "{}", "[]", "null",
		},
		TypeMatch: []interface{}{
			pBool,
		},
		TypeMismatch: []interface{}{
			aInt8,
			aInt16,
			aInt32,
			aInt64,
			aUint8,
			aUint16,
			aUint32,
			aUint64,
			aString,
			aFloat32,
			aFloat64,
			aBool,
			pInt8,
			pInt16,
			pInt32,
			pInt64,
			pUint8,
			pUint16,
			pUint32,
			pUint64,
			pString,
			pFloat32,
			pFloat64,
		},
	}.Test(t)
}

func TestString(t *testing.T) {
	testCase{
		Schema: String{
			MetaData: MetaData{
				Title:       "my-title",
				Description: "my-description",
			},
		},
		Match: `{
      "type": "string",
      "title": "my-title",
      "description": "my-description"
    }`,
		Valid: []string{
			"\"Some slightly longer string\"", "\"some string\"",
		},
		Invalid: []string{
			"242", "-254", "{}", "[]", "null", "true", "false",
		},
		TypeMatch: []interface{}{
			pString,
		},
		TypeMismatch: []interface{}{
			aInt8,
			aInt16,
			aInt32,
			aInt64,
			aUint8,
			aUint16,
			aUint32,
			aUint64,
			aString,
			aFloat32,
			aFloat64,
			aBool,
			pInt8,
			pInt16,
			pInt32,
			pInt64,
			pUint8,
			pUint16,
			pUint32,
			pUint64,
			pFloat32,
			pFloat64,
			pBool,
		},
	}.Test(t)
}

func TestStringLength(t *testing.T) {
	testCase{
		Schema: String{
			MetaData: MetaData{
				Title:       "my-title",
				Description: "my-description",
			},
			MinimumLength: 5,
			MaximumLength: 10,
		},
		Match: `{
      "type": "string",
      "title": "my-title",
      "description": "my-description",
      "minLength": 5,
      "maxLength": 10
    }`,
		Valid: []string{
			"\"12345678\"", "\"123456789\"",
		},
		Invalid: []string{
			"242", "-254", "{}", "[]", "null", "true", "false", "\"ddd\"",
			"\"12344567890123456789\"",
		},
		TypeMatch: []interface{}{
			pString,
		},
		TypeMismatch: []interface{}{
			aInt8,
			aInt16,
			aInt32,
			aInt64,
			aUint8,
			aUint16,
			aUint32,
			aUint64,
			aString,
			aFloat32,
			aFloat64,
			aBool,
			pInt8,
			pInt16,
			pInt32,
			pInt64,
			pUint8,
			pUint16,
			pUint32,
			pUint64,
			pFloat32,
			pFloat64,
			pBool,
		},
	}.Test(t)
}

func TestStringPattern(t *testing.T) {
	testCase{
		Schema: String{
			MetaData: MetaData{
				Title:       "my-title",
				Description: "my-description",
			},
			Pattern: "^[a-z]+$",
		},
		Match: `{
      "type": "string",
      "title": "my-title",
      "description": "my-description",
      "pattern": "^[a-z]+$"
    }`,
		Valid: []string{
			"\"dfasfdsafdsafsa\"", "\"a\"", "\"sds\"",
		},
		Invalid: []string{
			"242", "-254", "{}", "[]", "null", "true", "false", "\"\"", "\"-\"",
			"\"asda dsfsf\"", "\"asd4\"", "\"A\"", "\"azS\"",
		},
		TypeMatch: []interface{}{
			pString,
		},
		TypeMismatch: []interface{}{
			aInt8,
			aInt16,
			aInt32,
			aInt64,
			aUint8,
			aUint16,
			aUint32,
			aUint64,
			aString,
			aFloat32,
			aFloat64,
			aBool,
			pInt8,
			pInt16,
			pInt32,
			pInt64,
			pUint8,
			pUint16,
			pUint32,
			pUint64,
			pFloat32,
			pFloat64,
			pBool,
		},
	}.Test(t)
}
