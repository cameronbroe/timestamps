# Timestamps

Super simple utility to print timestamps at a constant rate to stdout. Useful for analyzing screen sharing applications.

# Building

Just build with the standard Go toolchain

`go build -o ./timestamps main.go`

# Usage

You can run `./timestamps` without any arguments and get timestamps in the format `12:56:19.080` at a rate of 1ms.

There are two arguments you can use to control the behavior of the timestamps output:

* `-format`: a format string to display the timestamp, refer to https://github.com/cameronbroe/timeutil#strftime as a reference
* `-rate`: a number in milliseconds on how often to print a timestamp

# License

Check `LICENSE` for license information.
