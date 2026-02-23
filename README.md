a cli tool for when u are me and dont want to write an approriate commit message and also dont want to think of one cuz thats a lot of work and ur lazy

uses git hooks to run the cli code before commit, user can choose from the assortment of message options (3)
if they want to write their own message, continue as usual with git commit -m [msg]
only gets triggered when git commit

use cumit init to initialise
cumit -h for help

go install github.com/crizah/cumit@v0.1.2

add this to ~/.bashrc
```
export PATH=$PATH:$(go env GOPATH)/bin
```

ready to use!
(will make this automated and much smoother EVENTUALLY)