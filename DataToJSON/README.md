JSON - JavaScript Object Notation
It's just key-value pairs in JS and it's easier to transmit than XML. 
It's human readable, machine transferable and generally the preferred
way to send and receive data via REST APIs. It's not the most efficient
way but it's the web-developer preferred way.

Marshalling - is the process of transforming the meory representation of
an object to a sata format suitable for storage or transmission, and it
is typically used when data must be moved between different parts of a
computer program or from one program to another.

The inverse of marshalling is called unmarshalling or demarshaling.

With Respect to GoLang,

Marshalling/Unmarshalling is Golang trying to convert struct into JSON 
object and JSON objects into GoLang structs.
We can transfer to JSON and from JSON(string literal v=byte slice) back 
into GoLang struct.
GoLang is a web backend language.

In GoLang when we put a string in backtick '', it is treated as raw string literal.
Meaning that it is read exactly as is a UTF-8 stringwith no escape characters
only runes(ut-8 characters) versus the "" double quote which allows
for escaoed characters.
In "", we need to escape new lines, tabs and other characters that do
not need to be escaped in backticks ''.
That is to say if we put a line break in a backtick string, it is 
interpreted as a '\n'.

Marshal Function- converts a GoLang struct into JSON object using json.Marshal
Its GoLang's way of saying "encode/convert to JSON Object" because GoLang is a strictly
typed language and JSON is a dynamically typed language.

Exposed v/s Not-Exposed Fields
As with all structs in Go, only fields with capital first letter are
visible to external programs/packages like JSON Marshaller.



func Marshal(v interface{}) ([]byte, error)
It takes a v interface{}, which can be 'any' go type.

It returns two things -
1. a slice if byte []byte, containing the literal string that is the JSON
    object
2. error- for error handling


Go Types represented in JSON by JSON Marshaller 
bool for JSON booleans,
float64 for JSON numbers,
string for JSON strings,
nil for JSON null.

Only data that can be represented as JSON will be encoded/converted by the
json.marshal() function


Unmarshal Function in the JSON Package

Converting JSON Object into a GoLang struct, we use the json.Unmarshal.
This is GoLangs way of parsing this JSON object into a valid Golang Struct

func Unmarshal(data []byte, v interface{}) error

unmarshal takes two parameters
1. a slice of bytes - this is raw string, the JSON object that is to be parsed
2. a pointer to struct to parse the JSON into

Unmarshal returns-
a error if anything went wring with parsing


How does Unmarshal decide which fields to try and parse -

For any key found in the JSON, Unmarshall will try to match it to key 
found in the struct wiht the following logic

1. It will first look for an exported(field with a capital letter) with
a tag json:"FieldName" //FieldName represents any member of a struct
2. A exported(field member with a capital letter) with the name 
FieldName
3. Any exported field name, that matches the fieldname if casesensitivity
is not an issue, e.g. fieLDName, FIELDNAME, fieldname

ONLY FIELDS FOUNDS IN THE destination type(struct) will be decoded.


For more reading - 
https://medium.com/go-walkthrough/go-walkthrough-encoding-json-package-9681d1d37a8f

Official GoLang Documentation for JSON package
https://golang.org/pkg/encoding/json/#Marshal

