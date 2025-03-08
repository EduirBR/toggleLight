# scroll-lock-led

A simple script to toggle the scroll lock LED/RGB on Ubuntu.

It finds the path to the scroll lock LED and toggles its state. It requires
sudo access to write to the brightness file.

usefull for ubuntu users who has rgb keyboards that turn on/off with the scroll lock key.

there its the build, and code to build by yourself

Go.

## make the build

`go build -o key_light main.go`

## move the file to bin

`sudo mv key_light /usr/local/bin/`
