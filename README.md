# dotctl - Simple software to help you manage your dotfiles

**dotctl** is a go program that works using git to copy and snapshot your dotfiles to a predefined directory plus some repository administration.

## How to install
### The lazy way

- Download the binary from the releases

### The gopher way (recommended)

- ```go get github.com/lucasew/dotctl```

## How to use
With this program you can do some useful things, like:

- ```add``` ```rules``` for files you want to bring to the repository

- ```edit``` the ```rules``` using the $EDITOR you had choosen or notepad if windows or nano for other systems

- ```ls``` the ```rules``` you have already defined.

- ```bring``` the files you configured to bring plus inject some temporary ones - the command is like a variadic function

- Use ```git``` from anywhere to manage your dotfile repository

- Or just ```snapshot```, that basically ```bring```s and commits the changes in a branch that is your $hostname, or your hostname with the current timestamp

## Why this is awesome

- Run on basically any OS that people usually use, including Windows and even Android via Termux

- The only dependency is Git in $PATH

- It's made using Go

- You can crosscompile it just changing two environment variables and ```go build```

- Supports .env files. If its useful for you, congradulations xD

- Rules pointing to directories are copied recursively

Happy Ricing :rice: :penguin:
