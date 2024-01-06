# Go - Should I Deloy Today - Cowsay

<img src="./_extras/cow.png">



__Should I Deploy Today with Cowsay__ Smart messages from [Should I Deploy Today](https://shouldideploy.today/) website and use it on Cowsay linux command line



## Configuration

1. Install Cowsay
   1. Ubuntu: `sudo apt install cowsay `
   1. Fedora: `sudo dnf install cowsay`
   2. Arch: `sudo pacman -S cowsay` 

1. Install ShouldIDeploy command line
   1. Download the binary (or build on your local machine)
   2. Move the binary to the directory: `/usr/bin`

1. Setup your terminal
   1. Open your `.bashrc` / `.zshrc` and add the command: `shouldideploytoday | cowsay`


# Building

Go Version: `go version go1.21.5 linux/amd64`


## Generating the binary 



__For the same machine__

    # Inside `src` folder
    go build -o shouldideploytoday main.go



*Copyright (c) 2023 Marcio Mangar*
