# GO Lang

## Basis Konzepte

### Main Module

Jedes Programm benötigt einen Einstiegspunkt: das **Main Package**.
Dies ist wie folgt aufgebaut und ruft beim Start sofoft die deklarierte main-Function auf:

```go
package main

func main() {
  // ... code here ...
}
```

### Module und Packages

Ein Module wird mittels `go mod init` erstellt.

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

### Externe Packages

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

### Kontrollstrukturen

IF / ELESE IF / ELSE

```go
if a < 10 {
  fmt.Println("kleiner als 10")
} else if a == 10 {
  fmt.Println("ist genau 10")
} else {
  fmt.Println("ist groesser als 10")
}
```

SWITCH

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

FOR

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
```

### Datentypen / Variablen

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

### Structs

Structs werden verwendet um Daten zu gruppieren (vergleichbar mit einem Dictonary oder JSON-Object).

Deklaration

```go
type user struct {
	firstName    string
	lastName     string
	birthdate    string
	creationTime time.Time
}
```

Zuweisen der Werte

```go
var newPerson user
newPerson = user{
  firstName: "Patrick",
  lastName: "Reiter",
  birthdate: "25.12.1988"
}
```

Zuweisen über Constructor-Funktion (inkl. Error-Handling)

```go
func newUser(firstName, lastName, birthdate string) (user, error) {
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

Eine Methode einem Struct zuweisen (receiver)

```go
func (u user) outputUserData() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}
```

Eine Methode die Daten im Struct manipuliert erstellen (Pointer verwenden!)

```go
func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

```

Embedded Struct

```go
type admin struct {
  email string,
  password string,
  user // erbt das user-struct
}
```

Zuweisen über Constructor-Funktion inkl. Werte für Embedded Struct (user)

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

### Pointer

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

### Stdout

Konsolenausgabe kann mit den jew. Print Functions im "fmt"-Package gemacht werden

```go
import "fmt"

fmt.Print("Hello World\n")
fmt.Println("Hello World") // with new-line

text := "World"
fmt.Printf("Hello %v", text) // with values
```

### Stdin

Werte in den STDIN einlesen mittels Scan-Function

```go
import "fmt"

var name string
fmt.Scan(&name) // verwende pointer zur variable (&-Symbol)
```

### Files lesen/schreiben

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

### Functions

```go
func name (param1 type, ...) returntype {
  // ... code ...//
  return
}

// Beispiel
func plus(a int, b int) int {
  return a + b
}
```

### Error Handling

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
