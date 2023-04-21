# DGTK - Geodesic Toolkit in Golang

> DGTK is a geodesic toolkit written in Golang. In its current version, it can solve direct and geodetic problems.

### Dependencies
- golang 1.19+
- fyne and fyne [depts](https://developer.fyne.io/started/#prerequisites)
- ##### Mac OS
  - glew
  - glew
### Instalation

To install DGTK, follow these steps:

1. Clone the repository:

```bash
git clone https://github.com/patriarch11/dgtk-golang-gui-app.git
```

2. Install the dependencies:

```bash
go mod download
```

### Usage

To use DGTK, simply run the executable:

```bash
./dgtk-golang-gui-app
```

This will launch the GUI application, where you can enter the input values and get the desired results.

### Compiling for different operating systems

DGTK can be compiled for different operating systems using the following commands:

#### Windows

```bash
GOOS=windows GOARCH=amd64 go build -o dgtk-golang-gui-app.exe cmd/dgtk-gui-app/main.go
```

#### Linux

```bash
GOOS=linux GOARCH=amd64 go build -o dgtk-golang-gui-app cmd/dgtk-gui-app/main.go
```

#### Mac OS

```bash
GOOS=darwin GOARCH=amd64 go build -o dgtk-golang-gui-app cmd/dgtk-gui-app/main.go
```

Note that the resulting executable file will have a different extension depending on the operating system (.exe for Windows and no extension for Linux and Mac OS).