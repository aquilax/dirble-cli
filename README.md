# dirble-cli
Command line tool for accessing https://dirble.com API

## Installation

```
go get -u github.com/aquilax/go-dirble
```

## Usage

```
NAME:
   dirble-cli - Fetches information from dirble.com

USAGE:
   dirble-cli [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
   stations, st         Get List of stations
   station        Get information about single station (requires station id as argument)
   song-history, sh     Get song history for station (requires station id as argument)
   similar-stations, ss    Get list of similar stations (requires station id as argument)
   categories, cat      Get list of categories
   primary-categories, pcat   Get list of primary categories
   categories-tree, tcat   Get the full category tree
   countries         Get list of countries
   country-stations     Get List of stations for country
   continents        Get list of continents
   countries         Get countries for continent
   search, s         Search for station
   help, h        Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --token, -t       API Token [$DIRBLE_API_TOKEN]
   --help, -h     show help
   --version, -v  print the version
```
