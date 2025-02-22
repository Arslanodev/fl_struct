# Description
Simple CLI application for structuring files into folders  

<img src="https://github.com/Arslanodev/fl_struct/blob/main/imgs/showcase.png" alt="showcase" width="400"/>

# File Structurer
A CLI application that makes it easier to manage files through terminal interface

## Description


Here is an enhanced version of the text:

"Simplify your file management with this intuitive CLI application, featuring a seamless web interface for effortless storage analysis. Organize your files with ease by creating and moving them into structured folders, categorized by various criteria such as:

* File size (small, medium, large)
* Date (created, modified, accessed)
* Kind (documents, images, videos, audio)

Quickly locate files with CLI's powerful search functionality, which allows you to find files by name with just a few keystrokes. Application streamlines your file management workflow, saving you time and reducing clutter.

Key benefits include:

* **Effortless organization**: Create custom folders and categorize files based on your specific needs
* **Fast file search**: Quickly find files by name, eliminating tedious manual searches
* **Comprehensive storage analysis**: Get a clear picture of your storage usage through our intuitive web interface

Take control of your file storage and discover a more efficient way to manage your digital assets."

## Features

* List files with advanced sorting options (size, kind, date added, date modified)
* Search for files or folders
* Group files into their respective folders (group by: size(sm, md, lg), kind, date added, date modified)
* Index file system for later fast search
* Analyze Storage (returns web interface with analysis dashboard of your storage capacity)

## Tech Stack

* [Cobra](https://github.com/spf13/cobra) for CLI framework
* [PromptUI](https://github.com/manifoldco/promptui) for interactive CLI interface
* [Templ](https://github.com/a-h/templ) for web interface

## Installation
You can install the app directly with .dmg, or through this command:  
```bash
go install github.com/Arslanodev/fl_struct
```

### Prerequisites

* Go (version 1.14 or later)
* Git

### Steps

1. Clone the repository: `git clone https://github.com/Arslanodev/fl_struct.git`
2. Build the application: `go build`
3. Move the executable file to `/usr/local/bin` (optional, for macOS users)

## Usage

### Running the Application

1. Run the application with the following command: `fl_struct .` (or `fl_struct ./dirname` for a specific directory)
2. Select options from the interactive CLI interface

### Example Output

The application will list files with their respective sizes, kinds, and dates added.

## Contributing

Contributions are welcome! Please submit a pull request with your changes.

## License

This project is licensed under the MIT License.