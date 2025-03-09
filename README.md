*(This project is still under development, but feel free to explore project internals)*

## Description
Simplify your file management with this intuitive CLI application. Organize your files with ease by creating and moving them into structured folders, categorized by various criteria such as:

* File size (small, medium, large)
* Date (created, modified, accessed)
* Kind (documents, images, videos, audio)

Quickly locate files with CLI's powerful search functionality, which allows you to find files by name with just a few keystrokes. Application streamlines your file management workflow, saving you time and reducing clutter. *(Under development)*

Key benefits include:

* **Effortless organization**: Create custom folders and categorize files based on your specific needs
* **Fast file search**: Quickly find files by name, eliminating tedious manual searches
* **Comprehensive storage analysis**: Get a clear picture of your storage usage through applicatin's intuitive web interface

Take control of your file storage and discover a more efficient way to manage your digital assets."

## Features

* List files with advanced sorting options (size, kind, date added, date modified)

*Features under development*:
* Search for files or folders
* Group files into their respective folders (group by: size(sm, md, lg), kind, date added, date modified)
* Index files for later fast search
* Analyze Storage (returns web interface with analysis dashboard of your storage capacity)

## Tech Stack

* [Cobra](https://github.com/spf13/cobra) for CLI framework
* [PromptUI](https://github.com/manifoldco/promptui) for interactive CLI interface

### Prerequisites

* Go (version 1.14 or later)
* Git

### Steps for usage

1. Clone the repository: `git clone https://github.com/Arslanodev/fl_struct.git`
2. Build the application: `make build`
3. Move the executable file to `/usr/local/bin` (optional, for macOS users)

## Usage

### Running the Application

1. After building the application, run this command inside bin dir: `fl_struct`
2. Select options from the interactive CLI interface

### Example


## License

This project is licensed under the MIT License.