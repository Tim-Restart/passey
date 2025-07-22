## Passey
Passey is a command line tool used to parse data from downloaded telegram chat backups.

It will obtain usernames throughout the data and return a slice of unique users identified.

If usernames are only a mobile number this will be sent to another slice for just numbers.

Any other channels posted will be parsed the channels slice.

If links are shared in the text, these should be captured and returned to the links slice.

## Installation:

#### Linux:
1. Download and install go // <https://go.dev/dl/>
2. Clone the repo using prefer method (download: <https://github.com/Tim-Restart/passey.git>)
3. `go build .`

#### Windows:
1. Download and install go // <https://go.dev/dl/>
2. Not yet tested the rest...

### Usage:
# Usage of the prorgram is easier from the same directory as the data to be parsed, however not required.

1. ./passey [html file name with no extension] [optional number]
2. eg: ./passey messages 10
3. Currently results will be printed to console, these can be collected to a document with >
eg. ./passey messages 10 > messages.txt

# In progress:
1. Parsing of Links and channel parsing still being worked on.
2. Formatting of output to html or MD to be completed.
3. Input option for type of phone numbers.
4. Input option for custom search term.
