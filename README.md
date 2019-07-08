# SUP

> Simple task tracker for daily standups

## Installation

1. Download the latest release from [releases](https://github.com/knicklabs/sup/releases).
2. Copy the binary to the desired directory on your computer.
3. Add the path to that directory to your path.

Here's an example of how to add a directory to your path in bash:

```
echo "PATH:\$PATH:/path/to/dir" >> ~/.bash_profile
```

In the above example, replace `/path/to/dir` with the path to the binary.

## Usage

In your terminal, run the `sup` command.

```
sup
```

By default it will output the help content:

```
NAME:
   sup - simple task tracker for daily standups

USAGE:
   main.exe [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
     add, a, new, n  Add a new task for Today
     copy, cp        Copy Yesterday's and Today's tasks
     list, ls        List Today's tasks
     edit, e         Edit Today's tasks
     print, p        Print Yesterday's and Today's tasks
     open, o         Opens the task directory
     which, w        Display location of tasks
     help, h         Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

You can append `--help` to a command to see usage instructions and additional options for a particular command.

For example, `sup print --help` will output:

```
NAME:
   main.exe print - Print Yesterday's and Today's tasks

USAGE:
   main.exe print [command options] [arguments...]

OPTIONS:
   --copy  Copy the output to the clipboard
```

To add a new task:

```
sup add "Write a blog post"
```

## Thanks

SUP is based on the original [SUP app for Node](https://github.com/ItsJonQ/sup). It includes code from the [open-golang](https://github.com/skratchdot/open-golang) library. Thanks to the developers and maintainers of these great projects.
