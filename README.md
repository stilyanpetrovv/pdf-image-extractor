# PDF Image Extractor

## Overview

PDF Image Extractor is a tool designed to extract images from PDF files and save them as separate image files. It's a simple, efficient solution for converting PDF content into image formats, making it easier to work with or share specific images from PDF documents.

## Features

- Extract images from each page of a PDF file.
- Save extracted images as separate IMAGE files.
- Easy-to-use command-line interface.

## Installation

### Using Go Install

To install PDF Image Extractor, use the `go install` command:

```bash
go install github.com/stilyanpetrovv/pdf-image-extractor@latest
```

### Download Windows Executable

If you prefer to use a pre-built Windows executable, you can download it from the [Releases page](https://github.com/stilyanpetrovv/pdf-image-extractor/releases) of this repository. Look for the `pdfImgExtract.exe` file in the latest release.

## Usage

Once installed, you can use the tool via the command line. Here’s a basic example:
```bash
./pdfImgExtract <path-to-your-pdf-file>
```
Replace <path-to-your-pdf-file> with the path to the PDF you want to process. The tool will extract images from the PDF and save them in the current directory.

### License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

