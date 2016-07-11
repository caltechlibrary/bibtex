
# Installation

## Compiled version

*bibfilter* is a command line program run from a shell like Bash. If you download the 
repository a compiled version is in the dist directory. The compiled binary matching
your computer type and operating system can be copied to a bin directory in your PATH.

Compiled versions are available for Mac OS X (amd64 processor), Linux (amd64), Windows
(amd64) and Rapsberry Pi (both ARM6 and ARM7)

### Mac OS X

1. Download [https://github.com/caltechlibrary/bibtex/releases/latest/bibtex-binary-release.zip](https://github.com/caltechlibrary/bibtex/releases/latest/bibtex-binary-release.zip)
2. Open a finder window, find and unzip **bibtex-binary-release.zip**
3. Look in the unziped folder and find **dist/maxosx-amd64/bibfilter**
4. Drag (or copy) the *bibfilter* to a "bin" directory in your path
5. Open and "Terminal" and run `bibfilter -h` to confirm you were successful

### Windows

1. Download [https://github.com/caltechlibrary/bibtex/releases/latest/bibtex-binary-release.zip](https://github.com/caltechlibrary/bibtex/releases/latest/bibtex-binary-release.zip)
2. Open the file manager find and unzip **bibtex-binary-release.zip**
3. Look in the unziped folder and find **dist/windows-amd64/bibfilter.exe**
4. Drag (or copy) the *bibfilter.exe* to a "bin" directory in your path
5. Open Bash and and run `bibfilter -h` to confirm you were successful

### Linux

1. Download [https://github.com/caltechlibrary/bibtex/releases/latest/bibtex-binary-release.zip](https://github.com/caltechlibrary/bibtex/releases/latest/bibtex-binary-release.zip)
2. Find and unzip **bibtex-binary-release.zip**
3. In the unziped directory and find for **dist/linux-amd64/bibfilter**
4. Copy the *bibfilter* to a "bin" directory (e.g. cp ~/Downloads/bibtex-master/dist/linux-amd64/bibfilter ~/bin/)
5. From the shell prompt run `bibfilter -h` to confirm you were successful

### Raspberry Pi

If you are using a Raspberry Pi 2 or later use the ARM7 binary, ARM6 is only for the first generaiton Raspberry Pi.

1. Download [https://github.com/caltechlibrary/bibtex/releases/latest/bibtex-binary-release.zip](https://github.com/caltechlibrary/bibtex/releases/latest/bibtex-binary-release.zip)
2. Find and unzip **bibtex-binary-release.zip**
3. In the unziped directory and find for **dist/raspberrypi-arm7/bibfilter**
4. copy the *bibfilter* to a "bin" directory (e.g. cp ~/Downloads/bibtex-master/dist/raspberrypi-arm7/bibfilter ~/bin/)
    + if you are using an original Raspberry Pi you should copy the ARM6 version instead
5. From the shell prompt run `bibfilter -h` to confirm you were successful


## Compiling from source

If you have go v1.6.2 or better installed then should be able to "go get" to install all the **bibtex** utilities and
package. You will need the GOBIN environment variable set. In this example I've set it to $HOME/bin.

```
    GOBIN=$HOME/bin
    go get github.com/caltechlibrary/bibtex/...
```

