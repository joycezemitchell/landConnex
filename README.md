# Name Sorter for LandConnex

This command-line application sorts a list of names from a file based on the last name and optionally by the first name. Users can specify the order of sorting (ascending or descending). This tool is part of the backend engineering assessment for LandConnex.

## Features

- **Sort by Last Name**: Sort names by last name by default.
- **Optional First Name Sorting**: Ability to sort by first name.
- **Order Sorting**: Supports both ascending and descending order.
- **File Input/Output**: Reads names from a file and outputs the sorted list to a new file.

## Getting Started

These instructions will guide you to get a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

Before you begin, ensure you have the following installed:
- Git
- Go (version 1.22 or higher)

### Installation

Clone the repository to your local machine:

``````
git clone https://github.com/joycezemitchell/landConnex.git
``````

Navigate to the project directory:

``````
cd landConnex 
go mod tidy
``````
No additional installations are required if Go is properly installed.

## Usage

Run the application from the command line by providing the path to the file containing the names:

``````
make build
``````

### Optional Arguments

- --sort-by first: Sort by first name instead of last name.
- --desc: Sort in descending order (default is ascending).

Example command to sort by first name in descending order:

``````
./sorter_app --file=file-name.txt
``````

``````
./sorter_app --file=file-name.txt --sort-by first --desc
``````

## Output

The sorted names will be saved to a new file named <original_file_name>-sorted.txt in the same directory as the input file.

## Running the Tests

To run the automated tests for this system, use the following command:

make tests

## Authors

- John Mitchell Villanueva

