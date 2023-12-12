# Line Counter

Line Counter is a command-line tool written in Go that allows you to count the number of non-comment lines in a file or directory containing Go source code files.

## Installation

To use Line Counter, you need to have Go installed on your system. If you don't have Go installed, you can download it from the [official Go website](https://golang.org/dl/) and follow the installation instructions.

Once you have Go installed, you can clone the repository or download the code manually.

```shell
git clone https://github.com/icoder-new/linecounter.git
```

## Usage

The Line Counter tool provides two flags that you can use to specify the input:

- `-file <file>`: Path to a single Go source code file.
- `-dir <directory>`: Path to a directory containing Go source code files.

You can use either the `-file` flag or the `-dir` flag, but not both at the same time. If no flag is provided or both flags are provided, the tool will display the usage instructions.

To run the Line Counter tool, execute the following command:

```shell
go run main.go -file <file>
```

or

```shell
go run main.go -dir <directory>
```

The tool will count the number of non-comment lines in the specified file or directory and display the results.

## Example

Count the number of non-comment lines in a single file:

```shell
go run main.go -file myfile.go
```

Count the number of non-comment lines in all Go source code files in a directory:

```shell
go run main.go -dir mydirectory
```

## License

This project is licensed under the [MIT License](LICENSE).

Feel free to modify and use this tool to count lines in Go source code files as per your requirements.

## Contributions

Contributions to improve Line Counter are welcome! If you find any issues or want to add new features, please create a pull request.

## Acknowledgments

Line Counter is inspired by the need to quickly count lines of code in Go projects. Special thanks to the Go community and the developers of the Go programming language for providing a powerful toolset.
