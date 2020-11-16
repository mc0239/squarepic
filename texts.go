package main

const helpText = `
                             _
 ___ ___ _ _ ___ ___ ___ ___|_|___
|_ -| . | | | .'|  _| -_| . | |  _|
|___|_  |___|__,|_| |___|  _|_|___|
      |_|               |_|

https://github.com/mc0239/squarepic


+-----------+---------------------------------+-----------------------+
| Parameter |          Description            |     Example value     |
+-----------+---------------------------------+-----------------------+
| help      | Display this help               | (any non-empty value) |
| squares   | Number of squares per line      | 4                     |
| size      | Size of the picture in pixels   | 200                   |
| mirror    | Should the picture be symmetric | true                  |
+-----------+---------------------------------+-----------------------+


`

const defaultConfigText = `address=127.0.0.1:9001
images_folder=images/
default_squares_count=5
min_size=5
max_size=5000
default_size=250
mirror=false
`
