% title: Autour du GO
% subtitle: Présentation du langage Go
% author: Martin Piegay
% thankyou: Thank you!
% contact: <a href="https://twitter.com/Martouf42">@Martouf42</a>
% contact: <span>github</span> <a href="http://github.com">cocap10</a>
% favicon: figures/golang.icon.png

---
title: Au programme :

- 15 minutes de présentation du langage

- 15 minutes d'atelier

Suivez cette présentation sur :

[https://cocap10.github.io/GolangTz/](https://cocap10.github.io/GolangTz/)


---
title: GO ?

- créer par Google
- GO 1.0 en Mars 2012
- open source
- simple, efficace et lisible
- on dit souvant **golang**

Lui c'est Gopher : <img height=200 src=figures/golang.png />

---
title: Les +

- compilation 80% à 90% plus rapide qu'en C
- garbage collecté
- fortement typé
- binaire linké statiquement

---
title: Les -

- mauvaise gestion des dépendance nativement
- pas vraiment orienté objet
- pas encore d'IDE parfait

---
title: Les outils natifs

- `go build`: build les binaires
- `go test`: lancer les tests unitaires, benchmark, couverture 
- `go fmt`: formate le code
- `go run`: raccourci pour builder et lancer le programme 
- `godoc`: génère la doc

---
title: Hello, Playground!

[http://play.golang.org/](http://play.golang.org/) Un service Web qui :

   - compile, lies, et exécute le programme à l'intérieur d'un bac à sable
   - puis renvoie le résultat.

<pre class="prettyprint" data-lang="go">
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
}
</pre>

---
title: Un exemple trivial

<pre class="prettyprint" data-lang="go">
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
</pre>

---
title: Les variables et constantes

Définition et initialisation des variables :
<pre class="prettyprint" data-lang="go">
// Definition puis initialisation
var myVar string
myVar = "Hello you"
// Ou tout en un avec :=
anotherVar := "Hello TechnoZore"
</pre>

Définition des constantes :
<pre class="prettyprint" data-lang="go">
package main
import "fmt"
const (
  Pi = 3.14
  Zero = 0
)
</pre>

---
title: Les types

Les types de base :

- bool, 
- string, 
- int(8 ou 16, 32, 64), 
- uint(8 ou 16, 32, 64), 
- byte, rune (Unicode),
- float32(ou 64),
- complex64(ou 128)

---
title: Les conditions

<pre class="prettyprint" data-lang="go">
  if num := -11 num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	num := 9
	switch {
	case num < 0:
		fmt.Println(num, "is negative")
	case num < 10:
		fmt.Println(num, "has 1 digit")
	default:
		fmt.Println(num, "has multiple digits")

	}
</pre>
[Essayer](https://play.golang.org/p/9XOsQMh7lO)

---
title: Les fonctions

<pre class="prettyprint" data-lang="go">
func addAndEven(x, y int) (sum int, even bool) {
	sum = x + y
	if sum%2 != 1 {
		even = true
	}
	return
}

func main() {
	x := 11
	y := 3
	mySum, myEven := addAndEven(x, y)
	fmt.Printf("%d + %d = %d is even ? %t\n", x, y, mySum, myEven)
}
</pre>
[Essayer](https://play.golang.org/p/sNsXJ_3eTE)

---
title: Les structs

<pre class="prettyprint" data-lang="go">
type Person struct {
  Name string
  Age  int
  isAdult bool
}

p := &Person{
  Name: "Greg",
  Age:  25,
  isAdult : true,
}

fmt.Println(p.Name)
</pre>
[Essayer](https://play.golang.org/p/Qcfgg_Px1_)

---
title: Les méthodes des structs

<pre class="prettyprint" data-lang="go">
func (p *Person) CanEnterPub() bool {
	return p.isAdult
}

person.CanEnterPub()
</pre>
[Essayer](https://play.golang.org/p/0Geqb9Gc81)

---
title: Les pointeurs

- Comme en C/C++ : Un pointeur contient l'adresse mémoire d'une variable
- Mais Go est garbage collecté donc pas besoin de gérer les allocations :)
- On utilise les même opérateur qu'en C/C++ : `*` et `&`

<pre class="prettyprint" data-lang="go">
var p *int
i := 42
p = &i
</pre>
[Essayer](https://play.golang.org/p/XmxN_QIrqd)

---
title: Les collections

Il y a plusieur type de collection de Go :
<pre class="prettyprint" data-lang="go">
  // Les Arrays : tableaux de taille fixe
	var intArray [10]int
	intArray[5] = 5
	// Les Slices : tableaux de taille variable
	var intSlice []int
	intSlice = append(intSlice, 1, 2, 3)
	// Les Maps
	var stringIntMap map[string]int
	stringIntMap["toto"] = 4
</pre>

---
title: Les boucles

Go a une seule structure de boucle, la boucle `for`:
<pre class="prettyprint" data-lang="go">
stringArray := [5]string{"un", "deux", "trois", "quatre", "cinq"}
for i, str := range stringArray {
	fmt.Printf("%d: %s\n", i, str)
}
</pre>
[Essayer](https://play.golang.org/p/xZyaDf0WtT)

---
title: Les Goroutines

Les Goroutines sont des processus léger :

- elles s'exécutent 
- dans le même espace d'adressage
- on peut en lancer pres de 100 000 à la fois
- la charge est distribué automatiquement sur les CPU

---
title: Les Goroutines exemple

<pre class="prettyprint" data-lang="go">
say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func main() {
	go say("world")
	say("hello")
}
</pre>
[Essayer](https://play.golang.org/p/6PHXHha_Uv)

---
title: Les channels

Les Goroutines communiques avec des channels. 
Ce sont des conduits typés à travers lesquels vous pouvez envoyer et recevoir des valeurs.
On utilise l'opérateur de canal : `<-`
exemple :
<pre class="prettyprint" data-lang="go">
ch := make(chan int, 2)
ch <- 1
ch <- 2
fmt.Println(<-ch)
Fmt.Println(<-ch)
</pre>
[Essayer](https://play.golang.org/p/ch4u4pu2ji)

---
title: Le Web

<pre class="prettyprint" data-lang="go">
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
</pre>
---
title: Atelier un FizzBuzz jusqu'à 100

- 1) Connectez-vous sur [https://play.golang.org/](https://play.golang.org/)
- 2) Codez une boucle qui affiche les 100 entier 
- 3) Avec un switch ou des if testez à chaque itération :
si l'entier est multiple de 5, afficher à la suite du nombre le mot « fizz »
si l'entier est multiple de 7, afficher à la suite du nombre le mot « buzz »
si l'entier est à la fois un multiple de 5 et de 7, afficher à la suite du nombre le mot « fizzbuzz »
- 4) Mettez le test dans une fonction appeler à chaque itération
- 5) Essayer d'utiliser la fonction en Goroutines (les paramètres seront des channels)  


Let's Go ;)