# Deeper Into Go

1. Variable Declarations
    Go is statically typed language.

    Basic Go Types:

| Type     | Example Values                  | Description                  |
|----------|---------------------------------|------------------------------|
| bool     | true, false                     | Boolean values               |
| string   | "Hi", "Hello"                   | A sequence of characters     |
| int      | 0, -28282, 9999                 | Integer numbers              |
| float64  | 10.000001, 0.9929292, -100.003  | Floating point numbers       |

    ```
    var card string = "Ace of Spades"

    card := "Ace of Spades"
    ```
- := defining a new variable into a function

2. Functions and Return Types

    Create a new function to return a card as a string:
    ```
        func nameOfFunc() typeToReturn {
            return dataOfTheType
        }
    ``` 

    Example:
        Informing Go compiler that whenever the new card function is executed, it's going to return a value of type string.
        
        ``` 
            func newCard() string {
                return "Ace of Spade"
            }
        ```

3. Slices and For Loops
    - Array: fixed length list os things
    - Slice: An array that can grow or shrink
        - Every element inside a slice must be of the same type.
    
    - `append(cards, "Six of Spades")` doesn't modify the existing slice. Instead, it returns a new slice. So, we need to atribute the append to a new variable or the existent variable (ex cards) to get the new slice.

    - How to iterate over a slice?
        ```
         for index, element := range sllice {
            // body of this for loop
        } 
        ```

4. OO Approach vs Go Approach
    - Go is not an object-oriented programming language.

    Deck class 
        Deck Instance
            property:
                cards: [string]
            functions:
                print
                shuffle
                saveToFile
    
5. Custom Type Declarations
    Common Pattern:
    - Tell go we want to create an array of strings and add a bunch of functions specifically made to work with it:
            `type deck []string`
    - Functions with 'deck' as a 'receiver'
    
    - create 'cards' folder
        `main.go` - create and manipulate a deck
        `deck.go` - describes what a deck is and how it works
        `deck_test.go` - code to automatically test the deck.

6. Receiver Functions

    ```
    func (d deck) print() {
        for i, card := range d {
            fmt.Println(i, card)
        }
    }
    ```
    - d: the actual copy of the deck we're working with
    - deck: every variable of type 'deck' can call this function on itself.
        `cards.print()` // in this example, cards is of deck type

7. Creating a New Deck

Cards:
- newDeck: create and return a list of playing cards. Essentially an array of strings.
- print: log out the content of a deck of cards.
- shuffle: shuffles all the cards in a deck.
- deal: create a 'hand' of cards.
- saveToFile: save a list of cards to a file on the local machine.
- newDeckFromFile: load a list of cards from the local machine.

8. Slice Range Syntax

- slice[startIndexIncluding:upToNotIncluding]
- Example:

    ```
        fruits := []string{"apple", "banana", "grape", "orange"}
        fmt.Println(fruits[0:3]) // apple, banana, grape

        // give me everything from index of two to the very end of the slice
        fmt.Println(fruits[1:]) // banana, grape, orange
        fmt.Println(fruits[:2]) // apple, banana
    ```

    To remove the specific cards for our hand, we can to deal with the deck:
        cards[:handsize] // get the cards for our hand
        cards[handsize:] // get the rest of the cards in the deck

9. Multiple Return Values

10. Bytes Slices
    - []Byte{"some string"}
    - WriteFile()

11. Deck to String

    - strings.Join()

12. Testing With Go
    - Create a new file ending in _test.go
    - To run all tests in a package, run the command 'go test'
        - The `testing`package runs each test function (whose name starts with `Test...`) and creates a `*testing.T` object for it. This object controls the test state, records failures, logs and allows you to mark the test as failed or skipped.
        - `*testing.T` is the interface between your test function and Go's test system.

13. Assignment: Even or Odd?
    - Create a new package
    - In the main function, create a slice of ints from 0 through 10
    - Iterate through the slice. For each element, check to see if the number is even or odd.
    - If the value is even, print out "even". If it is odd, print out "odd"

:= colon equals // initializing and assigning values
{} curly braces