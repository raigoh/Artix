# Art Encoder/Decoder

This is a simple command-line tool for encoding and decoding ASCII art. It supports both single-line and multi-line encoding/decoding.

## Content

1. [Usage](#usage)
    - [Decode a Single Line](#decode-a-single-line)
    - [Encode a Single Line](#encode-a-single-line)
    - [Encode Multiple Lines from a File](#encode-multiple-lines-from-a-file)
    - [Decode Multiple Lines from a File](#decode-multiple-lines-from-a-file)
    - [Help](#help)
2. [Installation](#installation)
3. [Features](#features)

## Usage

### Decode a Single Line
To decode a single line of ASCII art, run the following command:
```
go run . "[5 #][5 -_]-[5 #]"
```
This will decode the provided ASCII art and print the #####-\_-\_-\_-\_-\_-\_-##### result.

### Encode a Single Line
To encode a single line of ASCII art, run the following command:
```
go run . -e "#####-_-_-_-_-_-#####"
```
This will encode the provided ASCII art and print the [5 #][5 -_]-[5 #] result.

### Encode Multiple Lines from a File
To encode multiple lines of ASCII art from a file, run the following command:
```
go run . -me <inputfile>
```
Replace `<inputfile>` with the path to the file containing the ASCII art. This will encode each line of the file and print the encoded result.

### Decode Multiple Lines from a File
To decode multiple lines of ASCII art from a file, run the following command:
```
go run . -m <inputfile>
```
Replace `<inputfile>` with the path to the file containing the encoded ASCII art. This will decode each line of the file and print the decoded result.

### Help
To see this usage information again, run the following command:
```
go run . -h
```

## Installation

1. Clone this repository to your local machine:

```bash
git clone https://gitea.koodsisu.fi/raigohoim/art.git
cd your_repository
```

2. Navigate to the project directory:

```bash
cd art
```

3. To run commands, go art-decoder folder:

```bash
cd art-decoder
```

## Features

- **Single-Line Encoding**: Encode a single line of ASCII art by specifying the input string.
- **Single-Line Decoding**: Decode a single line of encoded ASCII art.
- **Multi-Line Encoding**: Encode multiple lines of ASCII art from a file. Each line of the file will be encoded individually.
- **Multi-Line Decoding**: Decode multiple lines of encoded ASCII art from a file. Each line of the file will be decoded individually.
- **Error Handling**: Handle various error scenarios gracefully, such as unbalanced brackets, missing arguments, invalid input, etc. Errors are displayed with informative messages.
- **Command-Line Interface**: Utilize a simple command-line interface to interact with the tool, making it easy to use and integrate into scripts or workflows.
- **Help Information**: Provide clear usage instructions and help information for users to understand how to use the tool effectively.
- **Customizable Encoding**: Encode consecutive characters and pairs in ASCII art into a compact format, making the output more concise and efficient.
- **Go Package Structure**: Organize the tool's functionality into separate packages, promoting modularity and maintainability.