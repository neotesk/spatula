<h1>Spatula<img align="left" width="42" height="42" alt="logo" src="https://github.com/user-attachments/assets/2c63d386-3582-4d22-b63c-7aaadf88ac6c" /></h1>

Spatula is an SPA server for debugging/testing purposes. It allows you to serve your Single Page Applications like React, Svelte and more. It also has live-server capabilities so you can update the web view as soon as you perform changes.

> [!CAUTION]
> This project is currently WIP, You can download the source code and build it using `truct do` however the application is merely tested and probably too unstable.

### Installation (Manual)
You can install Spatula manually through the [Releases](https://github.com/neotesk/spatula/releases)
section. Currently there are builds only for *Nix operating systems (Linux, OpenBSD, macOS etc.)
and Windows.

### Usage
You can use Spatula like so:
```
spatula serve build/index.html
```
and by default it will serve at port 8080

### Why does this exist?
I wasn't able to find any live-server for SPAs online (I am a quite lazy person don't ask.) and so I made my own. If it works it works and I don't question it.