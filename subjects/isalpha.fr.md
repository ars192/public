## isalpha

### Instructions

Écrire une fonction qui retourne `true` si la `string` passée en paramètre contient seulement des caractères alphanumériques, et qui retourne `false` autrement.

### Fonction attendue

```go
func IsAlpha(str string) bool {

}
```

### Utilisation

Voici un éventuel programme pour tester votre fonction :

```go
package main

import (
	"fmt"
	piscine ".."
)

func main() {
	fmt.Println(piscine.IsAlpha("Hello! How are you?"))
	fmt.Println(piscine.IsAlpha("HelloHowareyou"))
	fmt.Println(piscine.IsAlpha("What's this 4?"))
	fmt.Println(piscine.IsAlpha("Whatsthis4"))

}
```

Et son résultat :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
false
true
false
true
student@ubuntu:~/[[ROOT]]/test$
```
