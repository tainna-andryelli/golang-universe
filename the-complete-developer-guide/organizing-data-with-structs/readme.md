# Organizing Data With Structs

## Structs in Go

A `struct` in Go is a composite data type that groups together related properties, often of different types, into a single unit.

Instead of representing a card as just a string like `"Ace of Spades"`, you can define a `card` struct with separate fields for each property:

- `suit` (string): the suit of the card (e.g., "Spades")
- `value` (string): the value of the card (e.g., "Ace")

An instance of this struct for the Ace of Spades would look like:

- `suit`: `"Spades"`
- `value`: `"Ace"`

Structs make your code more organized, readable, and type-safe by clearly defining the structure of your data.

---
## Defining Structs

Create a new struct for a `person` who has the following fields:

- firstName <string>
- lastName <string>

Create a definition of what a person is:
```go
type person struct {
	firstName string
	lastName string
}
```

Create a values that mactch up with this struct:
```go
tainna := person{
	firstName: "Tainna", 
	lastName: "Ribeiro",
}
```

---
## Updating Struct Values
 
When you create a variable in go and do not assign any value to it, go assigns to a "zero value".

```go
var alex person
```

What is the zero value? For go means:

type | zero value |
----- | ---------- |
string | "" |
int | 0 |
float | 0 |
bool | false |

```go
fmt.Printf("%+v", alex) // {firstName: lastName:}
```
## Embedding Structs

```go
type person struct {
	firstName string
	lastName string
	contact contactInfo // we also can use just contactInfo, the same thing of contactInfo contactInfo the name of the field and the type
}

type contactInfo struct {
	email string
	zipCode int
}
```

## Structs With Receiver Functions

```go

func (p person) print() {
    fmt.Printf("%+v", p)
}
```