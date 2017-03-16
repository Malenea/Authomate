# Authomate

Authomate is a Golang written program meant to provide a simple way to fetch an author's full library on Goodreads.com providing the program with the author's id, the author's name or the XML response from Goodreads.com api from a book review from the author.

## Getting Started

### Building the binary

Building the Authomate's binary is done using the Golang compiler. Either by running

```
go install github.com/malenea/authomate
```

or by running the GoMake script.

```
/!\ It might be necessary to run the GoPath script in order to set the PATH and GOPATH env's variables, first
```

### Running Authomate

The Authomate's binary is in the bin folder.
To run it simply follow those examples' commands:

```
./bin/authomate "Douglas Adams"
./bin/authomate 4
./bin/authomate resources/Goodreads.xml
./bin/authomate https://www.goodreads.com/book/show/11.xml?key={your_dev_key}
```

You should then get that output:

```
```

### Overview

### Build with

### Authors

* **Olivier Conan** - *Initial work* - [Malenea](https://github.com/Malenea)