## rotatevowels

### Instructions

Write a **program** that checks the arguments for vowels.

- If the argument contains vowels (for this exercise `y` is not considered a vowel) the program has to **"mirror"** the position of the vowels in the argument (see the examples).
- If the number of arguments is less than 1, the program display a new line ("`\n`").
- If the arguments does not have any vowels, the program just prints the arguments.

Example of output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./rotatevowels "Hello World" | cat -e
Hollo Werld$
student@ubuntu:~/[[ROOT]]/test$ ./rotatevowels "HEllO World" "problem solved"
Hello Werld problom sOlvEd
student@ubuntu:~/[[ROOT]]/test$ ./rotatevowels "str" "shh" "psst"
str shh psst
student@ubuntu:~/[[ROOT]]/test$ ./rotatevowels "happy thoughts" "good luck"
huppy thooghts guod lack
student@ubuntu:~/[[ROOT]]/test$ ./rotatevowels "aEi" "Ou"
uOi Ea
student@ubuntu:~/[[ROOT]]/test$ ./rotatevowels

student@ubuntu:~/[[ROOT]]/test$
```
