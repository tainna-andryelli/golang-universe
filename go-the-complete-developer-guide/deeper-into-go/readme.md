# Deeper Into Go

This document covers key concepts and practical examples for working with Go, focusing on variable declarations, functions, slices, custom types, and more.

---

## 1. Variable Declarations

Go is a statically typed language.

### Basic Go Types

| Type     | Example Values                  | Description                  |
|----------|---------------------------------|------------------------------|
| bool     | true, false                     | Boolean values               |
| string   | "Hi", "Hello"                   | A sequence of characters     |
| int      | 0, -28282, 9999                 | Integer numbers              |
| float64  | 10.000001, 0.9929292, -100.003  | Floating point numbers       |

```go
var card string = "Ace of Spades"
card := "Ace of Spades" // := defines and assigns a new variable inside a function
```

---

## 2. Functions and Return Types

Create a function that returns a card as a string:

```go
func nameOfFunc() typeToReturn {
    return dataOfTheType
}
```

**Example:**

```go
func newCard() string {
    return "Ace of Spade"
}
```

---

## 3. Slices and For Loops

- **Array:** Fixed-length list of items.
- **Slice:** Dynamic array that can grow or shrink. All elements must be of the same type.

Appending to a slice returns a new slice:

```go
cards := []string{"Ace", "King"}
cards = append(cards, "Six of Spades")
```

Iterating over a slice:

```go
for index, element := range slice {
    // loop body
}
```

---

## 4. OO Approach vs Go Approach

Go is not an object-oriented language, but you can use structs and methods.

**Deck Example:**

- **Properties:** cards ([]string)
- **Functions:** print, shuffle, saveToFile

---

## 5. Custom Type Declarations

Create a custom type and add methods:

```go
type deck []string

func (d deck) print() {
    for i, card := range d {
        fmt.Println(i, card)
    }
}
```

---

## 6. Receiver Functions

Receiver functions allow you to associate methods with types:

```go
func (d deck) print() { ... }
cards.print() // cards is of type deck
```

---

## 7. Creating a New Deck

**Key Functions:**
- `newDeck`: Create and return a list of playing cards.
- `print`: Display the deck.
- `shuffle`: Shuffle the deck.
- `deal`: Create a hand of cards.
- `saveToFile`: Save the deck to a file.
- `newDeckFromFile`: Load a deck from a file.

---

## 8. Slice Range Syntax

```go
fruits := []string{"apple", "banana", "grape", "orange"}
fmt.Println(fruits[0:3]) // apple, banana, grape
fmt.Println(fruits[1:])  // banana, grape, orange
fmt.Println(fruits[:2])  // apple, banana
```

To deal cards:

```go
hand := cards[:handSize]    // hand of cards
remaining := cards[handSize:] // rest of the deck
```

---

## 9. Multiple Return Values

Go functions can return multiple values.

---

## 10. Byte Slices

- `[]byte("some string")`
- Use `WriteFile()` to write bytes to a file.

---

## 11. Deck to String

- Use `strings.Join()` to convert a slice to a string.

---

## 12. Testing With Go

- Create test files ending with `_test.go`.
- Run all tests in a package with `go test`.
- The `testing` package provides a `*testing.T` object for each test.

---

## 13. Assignment: Even or Odd?

- Create a new package.
- In `main`, create a slice of ints from 0 to 10.
- Iterate through the slice and print "even" or "odd" for each value.

---

**Go Syntax Reference:**
- `:=` (colon equals): initialize and assign values
- `{}` (curly braces): code blocks
