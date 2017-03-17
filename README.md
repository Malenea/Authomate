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
The Hitchhiker's Guide to the Galaxy
The Ultimate Hitchhiker's Guide to the Galaxy
The Restaurant at the End of the Universe
Life, the Universe and Everything
So Long, and Thanks for All the Fish
Dirk Gently's Holistic Detective Agency
Mostly Harmless
The Long Dark Tea-Time of the Soul
The Hitchhiker's Guide to the Galaxy: A Trilogy in Four Parts
The Salmon of Doubt
Young Zaphod Plays It Safe
Last Chance to See
The Hitch Hiker's Guide to the Galaxy: A Trilogy in Five Parts
The Hitchhiker's Guide to the Galaxy: The Primary Phase
The Dirk Gently Omnibus
The Hitchhikers Guide to the Galaxy Live
The Deeper Meaning of Liff
The More Than Complete Hitchhiker's Guide
The Meaning of Liff
The Original Hitchhiker Radio Scripts
The Hitchhiker's Trilogy
The Hitchhiker's Guide to the Galaxy: The Complete Radio Series
The Hitchhiker's Guide to the Galaxy: The Secondary Phase
Per Anhalter durch die Galaxis/Das Restaurant am Ende des Universums
The Hitchhiker's Guide to the Galaxy: The Tertiary Phase
The Hitchhiker's Guide to the Galaxy: Quandary Phase
The Hitchhiker's Guide to the Galaxy: The Quintessential Phase
The Illustrated Hitchhiker's Guide To The Galaxy
Doctor Who: Shada
The Private Life of Genghis Khan
```

Each book title is sorted as title without associated series. This can easily be modified though as the title with series is also parsed in the code.

### Overview

```
*---------------*
| Author's id   |
| Author's name | -> Parameters are provided as follow, author's name (i.e "Stephen King"),
| XML File      |    author's id (i.e "3014"), XML file (i.e "Silmarilion_-_review.xml"),
| XML url       |    or XML url (i.e "https://goodreads.com/book/show/[...].xml[...]").
*------*--------*
       |
       |
       v               *----> url ----> http.Get --*
*------*--------*      |                           |
| Sorting       | -----*----> file ----------------*----> XML parsing to get id ----* 
*---------------*      |                           |                                |
                       *----> name ---> http.Get --*                                |
                       |                                                            |
                       *----> id ---------------------------------------------------*
                                                                                    |
                                                                                    |
                                                                                    v
                                                                   *----------------*----*
                                                                   | http.Get on books'  |
                                                                   | list                |
                                                                   *----------------*----*
                               *----------------------*                             |
                               | XML parsing to fetch |                             |
              Return list <----* each book and create *<----------------------------*
                               | a list               |
                               *----------------------*
```

### Build with

Golang compiler

### Authors

* **Olivier Conan** - *Initial work* - [Malenea](https://github.com/Malenea)