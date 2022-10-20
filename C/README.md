# C
--
    import "github.com/kokizzu/gotro/C"


## Usage

#### func  IsAlpha

```go
func IsAlpha(ch byte) bool
```
IsAlpha check whether the character is a letter or not

    C.IsDigit('a') // true
    C.IsDigit('Z') // true

#### func  IsDigit

```go
func IsDigit(ch byte) bool
```
IsDigit check whether the character is a digit or not

    C.IsDigit('9') // true

#### func  IsIdent

```go
func IsIdent(ch byte) bool
```
IsIdent check whether the character is a valid identifier suffix alphanumeric
(letter/underscore/numeral)

    C.IsIdent('9'))

#### func  IsIdentStart

```go
func IsIdentStart(ch byte) bool
```
IsIdentStart check whether the character is a valid identifier prefix
(letter/underscore)

    C.IsIdentStart('-') // false
    C.IsIdentStart('_') // true

#### func  IsValidFilename

```go
func IsValidFilename(ch byte) bool
```
IsValidFilename check whether the character is a safe file-name characters
(alphanumeric/comma/full-stop/dash)

    C.IsValidFilename(' ') // output bool(true)
