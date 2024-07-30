# Windows JDK Version Switch

## Brief Introduction

Abbreviated as `wjvs`, a simple Windows JDK version switching tool. 
This tool is suitable for installing multiple JDK versions and 
has the system variable `JAVA_HOME` configured and the system variable `Path` 
configured with`%JAVA_SOME%\bin`. It is recommended to configure`%JAVA_SOME%\bin` 
in the system variable `Path` to increase the priority as much as possible and 
avoid the impact of the Java runtime environment of other Java developed applications 
on the priority of `JAVA_HOME`.

## Build

```shell
go build
```

## Usage

For ease of use, the directory where `wjvs.exe` is located can be added to the system variable `Path`, 
and the instructions of `wjvs` can be used anywhere in the future.

### help

Output instruction help information, `help` , `-h` , `--help`, These three instructions are equivalent.

```shell
# use
>.\wjvs.exe help
>.\wjvs.exe -h
>.\wjvs.exe --help

# If you add the directory where wjvs.exe is located to the system variable Path, you can do this:
>wjvs help
>wjvs -h
>wjvs --help

# For example, you will see the following information output:
A simple command-line tool for switching Java Development Kit versions on Windows.
You can list the Java Development Kit versions already installed on the device and switch to the currently used Java Development Kit version.

Usage:
  wjvm [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        It will list the Java Development Kit versions installed on the current device.
  use         Used to switch the current Java Development Kit version, must be the installed version.

Flags:
  -h, --help   help for wjvm

Use "wjvm [command] --help" for more information about a command.
```

### list

List the currently installed JDK versions

```shell
# use
>.\wjvs.exe list

# If you add the directory where wjvs.exe is located to the system variable Path, you can do this:
>wjvs list

# For example, you will see the following information output:

    1.8.0_281
  * 11.0.19 (Currently using)
```

### use

To switch to the desired version, you need to enter the version number of the currently installed version, 
such as `1.8.0_281` in the version information output by executing `wjvs list` above. 
When you execute `wjvs use [version]`, an administrator authorization confirmation will pop up. 
Please click `Yes` agree to authorize.

```shell
# use
>.\wjvs.exe use 1.8.0_281

# If you add the directory where wjvs.exe is located to the system variable Path, you can do this:
>wjvs use 1.8.0_281

# no output
```
Then open another command-line window and execute `java - version`
to verify if the JDK version has been switched to the version you need.
