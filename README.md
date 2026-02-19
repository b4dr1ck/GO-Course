# GO Lang - Basics

## Main Module

Jedes Programm benötigt einen Einstiegspunkt: das **Main Package**.
Dies ist wie folgt aufgebaut und ruft beim Start sofoft die deklarierte main-Function auf:

```go
package main

func main() {
  // ... code here ...
}
```

## Module und Packages

Ein Module wird mittels `go mod init <module url>` erstellt.

Jedes Module muss mindestens ein main package und eine `main() function als Einstiegspunkt haben.
Weitere Packages müssen als Folder erstellt werden.

Aufbau eines simplen Modules:

```
├── bank.go           # main module (wird beim Start aufgerufen)
├── communication.go  # communication-module (innerhalb main-package)
├── fileops           # fileops-package
│   └── fileops.go    # fileops module
└── go.mod            # module example.com/bank
```

Wenn man ein Package erstellen möchte, legt man einen entsprechenden Order mit dem Package-Namen an und erstellt darin die jew. Scripts welche dann mittels `package name` dem Package zugeteilt werden.
Alle Functions innerhalb eines Packages dessen Name mit Upper-Case beginnen, werden exportiert und können dann außerhalb verwendet werden.

Der gesamte Pfad des Module muss angegeben werden wenn man ein internes Package importieren möchte.

```go
// fileops package
package fileops

func GetFile(filename string) {
  //... code here ...
}
```

```go
// main package
package main

import "example.com/bank/fileops" // gesamter Pfad zum Package

func main() {
  // Zugriff auf Functions via Packagename und Punkt-Operator
  fileops.GetFile("myFile.txt")
}

```

## Externe Packages

Installation (Beschreibung meist auf Seite zu finden)
https://pkg.go.dev/

```
go get <url>
```

Beispiel

```
go get github.com/Pallinder/go-randomdata
```

Danach wird das go.mod File updated und das Package kann verwendet werden.

```go
// Eintrag in go.mod
require github.com/Pallinder/go-randomdata v1.2.0 // indirect
```

```go
// Package verwenden
import "github.com/Pallinder/go-randomdata"

fmt.Println("Phone: ", randomdata.PhoneNumber())
```

## Kontrollstrukturen

**IF / ELESE IF / ELSE**

```go
if a < 10 {
  fmt.Println("kleiner als 10")
} else if a == 10 {
  fmt.Println("ist genau 10")
} else {
  fmt.Println("ist groesser als 10")
}
```

**SWITCH**

```go
switch a {
case 1:
  fmt.Println("Ist genau 1")
case 2:
  fmt.Println("Ist genau 2")
case 3:
  fmt.Println("Ist genau 3")
default:
  fmt.Printf("Ist %v\n", a)
}
```

**FOR**

```go
// Endlos-Schleife
for {
  fmt.Println("waiting...")
}

// klassische for-schleife
for i:=0; i<10;i++ {
  fmt.Printf("%v",i)
}

// while-schleife
i:=1
for i < 10 {
  i++
  fmt.Printf("%v",i)
}

// iteriere über die Map (oder Array)
for key, value := range urls {
  fmt.Printf("[%v] => %v\n", key, value)
}
```

## Datentypen / Variablen

primitve Datentypen

```go
var text string = "Hallo"
var float float64 = 100.5
var integer int = 5
var boolean bool = true
```

**shorthand** Variablendeklaration

```go
value := 23.5
a,b,c := 1, 2, 3
```

Konstanten

```go
const myFile = "test.txt"
```

## Custom Type / Type Aliases

Eigene Types definieren

```go
type number int

var myNumber number
myNumber = 5

```

Eigene Methoden für eigene Typen definieren

```go
import "fmt"

// custom type "str"
type str string

// declare method
func (s str) output() {
	fmt.Printf("This is my String: %s", s)
}
func main() {
	var text str
	text = "Hello World"
	text.output()
}
```

## Type Assertions

Datentyp bestimmen mittels `value.(T)

```go
var a any = 10
typeInt, ok := a.(int) // check if a is of type int
if ok {
  fmt.Printf("a is of type int: %v\n", typeInt)
}
```

Datentype bestimmen mittels Type-Switch

```go
// type assertions with type-switch
switch a.(type) {
case int:
  fmt.Printf("Integer: %v\n", a)
case float64:
  fmt.Printf("Float: %v\n", a)
case string:
  fmt.Printf("String: %v\n", a)
case bool:
  fmt.Printf("bool: %v\n", a)
default:
  fmt.Print("Unknown Type\n")
}
```

## Functions

```go
func (receiver-argument) name (param1 type, ...) returntype {
  // ... code ...//
  return
}

// Beispiel
func plus(a int, b int) int {
  return a + b
}
```

Functions als Werte

```go
func doubleNumber(n int) int {
  return n*n
}

func transformNumbers(numbers []int, transform func(int) int) []int {
  transformedNumbers := []int{}
  for _,value := range numbers {
    transformedNumbers = append(transformedNumbers,transform(value))
  }
  return transformedNumbers
}

numbers := []int{1, 2, 3, 4, 5}
double := transformNumbers(numbers, doubleNumber)
```

auch Type Aliases (custom Types) funktionieren hier

```go
type transformFn func(int) int // declare type

// use custon function type
func transformNumbers(numbers []int, transform transformFn) []int
// ....

```

Anonyme Functions sind auch möglich

```go
numbers := []int{1, 2, 3, 4, 5}
double := transformNumbers(numbers, func(number int) int { return number * number})
fmt.Println(double)
```

## Structs

Structs werden verwendet um Daten zu gruppieren (vergleichbar mit einem Dictonary oder JSON-Object).

**Deklaration**

```go
type user struct {
	firstName    string
	lastName     string
	birthdate    string
	creationTime time.Time
}
```

**Zuweisen der Werte**

```go
var newPerson user
newPerson = user{
  firstName: "Patrick",
  lastName: "Reiter",
  birthdate: "25.12.1988"
}
```

**Zuweisen über Constructor-Funktion (inkl. Error-Handling)**

```go
func newUser(firstName, lastName, birthdate string) (user, error) {
  if firstName == "" || lastName == "" || birthdate == "" {
    user{},errors.New("Input may not be empty")
  }

	return user{
		firstName:    firstName,
		lastName:     lastName,
		birthdate:    birthdate,
		creationTime: time.Now(),
	}, nil
}

appUser, err := newUser(userFirstName, userLastName, userBirthdate)
if err != nil {
  fmt.Println(err)
  return

}
```

**Eine Methode einem Struct zuweisen (receiver)**

```go
func (u user) outputUserData() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}
```

**Eine Methode die Daten im Struct manipuliert erstellen (Pointer verwenden!)**

```go
func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

```

**Embedded Struct**

```go
type admin struct {
  email string,
  password string,
  user // erbt das user-struct
}
```

**Zuweisen über Constructor-Funktion inkl. Werte für Embedded Struct (user)**

```go
func newAdmin(email, password string) admin {
	return admin{
    email: email,
    password: password,
    user: user{
 		  firstName:    firstName,
		  lastName:     lastName,
		  birthdate:    birthdate,
		  creationTime: time.Now()
    }
  }
}
```

## Arrays / Slices

Neues Array deklarieren

```go
var myProducts [4]string
myProducts = [4]string{"Book","Pen","TV","Apple"}

// shorthand
myProducts := [4]string{"Book","Pen","TV","Apple"}
```

Über Index neue Werte setzen

```go
fmt.Println(myProducts[1]) // --> Pen

// zweites Element setzen (0-based index)
myProducts[1] = "Brush"

fmt.Println(myProducts[1]) // --> Brush
```

Dynamisches Array erstellen (no fixed size)

```go
myProducts := []string{"Book"}

// weitere Elemente hinzufügen
myProducts = append(myProducts,"Pen")
```

Slices verwenden

```go
// start-index ist inkludiert
// end-index ist exkludiert
fmt.Println(myProducts[1:3]) // --> Pen, TV
fmt.Println(myProducts[1:])  // --> Pen, TV, Apple
fmt.Println(myProducts[:1])  // --> Book
```

Unpacking Arrays `list...`

```go
productPrices := []float64{10.99,20.0,105.99}
newProductPrices := []float64{25.99, 200.0}

productPrices = append(productPrices,newProductPrices...) // ...-operator
```

## Maps

Maps sind die klassischen Hash-Tables und verwenden ein Key-Value Paar.

Map mit Werten erstellen

```go
// map[keyType]valueType

urls := map[string]string{
  "google":    "https://www.google.com/",
  "homepage":  "https://badricks-world.at/",
  "localhost": "http://127.0.0.1/",
}
```

Alternative kann man hier auch einen Type-Alias verwenden

```go
type urlsMap map[string]string

urls := urlsMap{
  "google":    "https://www.google.com/",
  "homepage":  "https://badricks-world.at/",
  "localhost": "http://127.0.0.1/",
}

```

Weitere/Neue Keys erstellen

```go
urls["deviantart"] = "https://www.deviantart.com/"
```

Einen Key löschen

```go
delete(urls, "google")
```

## Make

Wenn man eine Map oder ein Array mit `make` erstellt wird der Bereich im Speicher gleich reserviert und der Zugriff ist somit schneller und effizieneter.

Wenn mal z.B eine leere Map ohne `make` erstellen würde, könnte man ihr keine Werte zuweisen und es kommt ein Runtime-Error (panic)
`panic: assignment to entry in nil map`

### Slices mit make

```go
// make(slice, init-values, capacity )
prices := make([]float64,0,5)
fmt.Printf("%v",prices) // => []
fmt.Printf("len=%d cap=%d\n",len(prices),cap(prices)) // => 0, 5
```

### Maps mit make

```go
// make(map)
urls := make(map[string]string)
urls["pinterest"] = "https://www.pinterest.com"
```

## Pointer

Pointer sind im Gegensatz zu normalen Variablen nur ein "Zeiger" auf die Adresse in der die jew. Variable im Memory gespeichert ist.

Pointer machen Sinn wenn man z.B vermeiden möchte das gleich Wert öfters im RAM gespeichert wird.
Standardmäßig speichert nämlich GO ein Kopie eines Werts wenn man den einer Funktion übergibt. Mit Pointer wird dieser nur einmal im RAM abgespeichert.

Oder wenn man Werte direkt manipulieren möchte z.B innerhalb einer Funktion.
Somit muss man nicht den "neuen" Wert mittels return zurückgeben und abspeichern.

```go
var age int = 32
var agePointer *int = &age // Speichere nur die Adresse (im RAM) der Variable, nicht den Wert

fmt.Println(agePointer)  // zeigt nur die Adresse im Speicher an z.B 0xc000118008
fmt.Println(*agePointer) // dereferenziere den Pointer und erhalte den Wert

age = 22
fmt.Println(*agePointer) // würde nun 22 zurückgeben

```

## Stdout

Konsolenausgabe kann mit den jew. Print Functions im "fmt"-Package gemacht werden

```go
import "fmt"

fmt.Print("Hello World\n")
fmt.Println("Hello World") // with new-line

text := "World"
fmt.Printf("Hello %v", text) // with values
```

## Stdin

Werte in den STDIN einlesen mittels Scan-Function

```go
import "fmt"

var name string
fmt.Scan(&name) // verwende pointer zur variable (&-Symbol)
```

Für Input der mehr als einem Wert bzw. mehreren Wörtern besteht, kann ein Reader (IO-Buffer) verwendet werden

```go
reader := bufio.NewReader(os.Stdin)
text, err := reader.ReadString('\n') // rune with single-quotes ''

if err != nil {
  return ""
}

text = strings.TrimSuffix(text, "\n")
text = strings.TrimSuffix(text, "\r") // for windows line-endings
```

## Json Encoding

Structs als Json-Object encodieren.

Die Felder des Structs müssen "public" (mit UpperCase beginnen) sein, damit sie in ein JSON-Object konvertiert werden können.

```go
import "encoding/json"

// using tags for the keys in the json-object (`json:"key_name")
type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Save() error {
  json, err := json.Marshal(note) // encode
  if err != nil {
    return err
  }
  // code to save json-file ....
}

```

## Files lesen/schreiben

File schreiben

```go
import "os"

// Parameter
// 1. Filename/Filepfad (string)
// 2. Daten (konvertiert in bytes)
// 3. Berechtigungen (Linux-Style)
os.WriteFile("filename.txt",[]byte(data),0644)

```

File lesen

```go
import "os"

result, err := os.ReadFile("filename.txt")
```

## Error Handling

Wenn eine Funktion ein Error-Object zurückgibt kann dieses verwendet werden

```go
import "os"

results, err := os.ReadFile("results.txt")
if err != nil {
  panic("Can not read file - abort application")
}
```

Eigene Errors ausgeben

```go
import "errors" // errors-packe importieren um neuen Error (errors.New()) auszugeben
import "os"
import "fmt"

func readResultFile(filename string) (string,error) {
  result, err := os.ReadFile(filename)
  if err != nil {
    // gib leeren String und Error-Object zurück
    return "", errors.New("Can not read file " + filename)
  }
  // gib result-Sting und Nil-Object zurück, da kein Error aufgetreten ist
  return string(result),nil
}
```
