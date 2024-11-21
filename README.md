# ASCII Art Decoder/Encoder

This project provides a web application for decoding and encoding ASCII art.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [File Structure](#file-structure)
- [Contributing](#contributing)
- [Screenshots](#screenshots)

## Introduction

This web application allows users to decode and encode ASCII art. It provides a simple interface where users can input ASCII art and choose to either decode or encode it.

## Features

- Decode ASCII art: Convert ASCII art encoded in base64 format to plain text.
- Encode ASCII art: Convert plain text to ASCII art encoded in base64 format.

## Installation

To run this project locally, follow these steps:

1. Clone the repository:

```bash
git clone https://github.com/raigoh/Artix.git
```

2. Navigate to the project directory:

```bash
cd Artix/art-interface
```

## Usage

To start the server, run:

```
go run .
```

or

```
go run server.go
```

The server will start running on http://localhost:8080.

## File Structure

The project has the following structure:

```markdown
Artix/
├───art-decode/
│ ├── flags.go
│ ├── README.md
│ ├── art-encoder.go
│ ├── art-decoder.go
│ ├── error_handling.go
│ │
│ └───Resources/
│ ├── cats.art.txt
│ ├── cats.encoded.txt
│ ├── kood.art.txt
│ ├── kood.encoded.txt
│ ├── lion.art.txt
│ ├── lion.encoded.txt
│ ├── plane.art.txt
│ └── plane.encoded.txt
│
├───art-interface/
│ ├── server.go
│ ├── README.md
│ │
│ ├───templates/
│ │ ├── index.html
│ │ └── shutdown.html
│ │
│ ├───static/
│ │ ├── style.css
│ │ └── style1.css
│ │
│ └───functions/
│ ├── render.go
│ ├── pagedata.go
│ ├── index.go
│ ├── shutdown.go
│ └── process.go
│
└── go.mod
```

## Contributing

Contributions are welcome! If you find any issues or want to add new features, please open an issue or submit a pull request.

## Screenshots

Here are some screenshots of the application:

![Screenshot 1](/art-interface/static/screenshots/webpage.jpg)
![Screenshot 2](/art-interface/static/screenshots/kood.jpg)
![Screenshot 3](/art-interface/static/screenshots/shutdown.jpg)
