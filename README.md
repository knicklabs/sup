# SUP

> Simple task tracker for daily standups

## Installation

Download the latest release from [releases](https://github.com/knicklabs/sup/releases).
Choose the version for your operating system and architecture.

Unzip the desired version and rename it to `sup` or `sup.exe` (on Windows).

```
unzip sup_darwin_amd64.zip && mv sup_darwin_amd64 sup
```

On Linux and Mac, make the file executable. This is not required on Windows. 

```
chmod +x sup
```

Finally, move the executable binary into your path.

```
mv sup /usr/local/bin/sup
```

Alternatively, put `sup` in a custom directory and add it to your path.

```
echo "PATH:\$PATH:/path/to/dir" >> ~/.bash_profile
```

After you've installed `sup` on your system, just type `sup` in your terminal to use the application.

### Dependencies

There are no dependencies on Mac and Windows.

The Linux version requires xclip. Without xclip, copy functionality will not work. Install xclip on Debian-based distributions with `sudo apt install xclip`.

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
   sup [global options] command [command options] [arguments...]

VERSION:
   0.3.0

COMMANDS:
     add, a, new, n  Add a new task for Today
     config          configures SUP
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
   sup print - Print Yesterday's and Today's tasks

USAGE:
   sup print [command options] [arguments...]

OPTIONS:
   --copy  Copy the output to the clipboard
```

To add a new task:

```
sup add "Write a blog post"
```

You can use emojis when adding tasks:

```
sup add ":laughing: Tell a funny joke"
```

See this [Emoji Cheat Sheet](https://www.webfx.com/tools/emoji-cheat-sheet/) for a list of supported Emojis.

## Thanks

SUP is based on the original [SUP app for Node](https://github.com/ItsJonQ/sup). It includes code from the [open-golang](https://github.com/skratchdot/open-golang) library. Thanks to the developers and maintainers of these great projects.
