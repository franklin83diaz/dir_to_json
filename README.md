# DirToJson

An application that takes the structure of a directory and represents it in a JSON.


## Installation

Install 

```bash
 go bluid -o dirtojson src/main.go
```
    
## Deployment

To deploy this project run

```bash
  go run src/main.go
```


## Running Tests

To run tests, run the following command

```bash
  go test ./...
```


## Usage/Examples

```
dirtojson /tmp/test1 0
```
or if you want include hidden files use:
```
dirtojson /tmp/test1 0 show
```


## Authors

- [@Franklin Diaz](https://github.com/franklin83diaz)


## License

[MIT](https://choosealicense.com/licenses/mit/)


## Contributing

Contributions are always welcome!

