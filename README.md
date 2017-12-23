# Track

Track is a simple command line time and interruption tracker.

[![Build Status](https://travis-ci.org/zlypher/track.svg?branch=master)](https://travis-ci.org/zlypher/track)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**Table of Contents**
* [Installation](#installation)
* [Usage](#usage)
* [Contributing](#contributing)
* [License](#license)

## Installation

1. Install via `go`

```bash
go get github.com/zlypher/track
go install github.com/zlypher/track
```

2. Download packaged version from [release page](https://github.com/zlypher/track/releases)

## Usage

**Commands**
* [Track Task](#track-task)
* [Stop Tracking](#stop-tracking)
* [Track Interrupt](#track-interrupt)
* [List Interrupts](#list-interrupts)
* [Location](#location)
* [Version](#version)

### Track Task

Starts tracking of a new task

```
track start "Refactor Application"
```

### Stop Tracking

Stops the tracking

```
track stop
```

### Track Interrupt

Creates a new interrupt entry

```
track int "John Doe"
```

...

### List Interrupts

Lists all tracked interrupts

```
track list
```

### Location

Returns the folder where track stores its data.

```bash
track location
// Output: C:/Users/<USERNAME>/.track
```

### Version

Outputs the current app version.

```bash
track version
// Output: 0.0.6
```

## Contributing

If you find any problems or have suggestions, please open a ticket or a pull request. (see CONTRIBUTING.md)

## LICENSE

MIT (see LICENSE.md)
