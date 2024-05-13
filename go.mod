module github.com/kitanoyoru/kitad

go 1.22

toolchain go1.22.3

replace (
    github.com/kitanoyoru/kitad/api/v2 => ./api
    github.com/kitanoyoru/kitad/server => ./server
)
