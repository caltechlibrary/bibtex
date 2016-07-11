
# Installation

## Compiled version

*bibfilter* is a command line program run from a shell like Bash. If you download the 
repository a compiled version is in the dist directory. The compiled binary matching
your computer type and operating system can be copied to a bin directory in your PATH.

Compiled versions are available for Mac OS X (amd64 processor), Linux (amd64), Windows
(amd64) and Rapsberry Pi (both ARM6 and ARM7)

### Mac OS X

1. Go to [github.com/caltechlibrary/bibtex/releases/latest](https://github.com/caltechlibrary/bibtex/releases/latest)
2. Find the link for [bibtex-binary-release.zip](https://github.com/caltechlibrary/bibtex/releases/latest/bibtex-binary-release.zip) and click on it.
3. Save bibtex-binary-release.zip in your "Downloads" folder
4. Open a finder window and unzip bibtex-binary-release.zip
5. Look in the unziped folder and find **dist/maxosx-amd64/bibfilter**
6. Drag (or copy) the *bibfilter* to a "bin" directory in your path
7. Open and "Terminal" and run `bibfilter -h` to confirm you were successful

### Windows

1. Go to [github.com/caltechlibrary/bibtex/releases/latest](https://github.com/caltechlibrary/bibtex/releases/latest)
2. Find the link for [bibtex-binary-release.zip](https://github.com/caltechlibrary/bibtex/releases/latest/bibtex-binary-release.zip) and click on it.
3. Save bibtex-binary-release.zip in your "Downloads" folder
4. Open the file manager find the downloaded file and unzip it (e.g. bibtex-binary-release.zip)
5. Look in the unziped folder and find **dist/windows-amd64/bibfilter.exe**
6. Drag (or copy) the *bibfilter.exe* to a "bin" directory in your path
7. Open Bash and and run `bibfilter -h` to confirm you were successful

### Linux

1. Go to [github.com/caltechlibrary/bibtex/releases/latest](https://github.com/caltechlibrary/bibtex/releases/latest)
2. Find the link for [bibtex-binary-release.zip](https://github.com/caltechlibrary/bibtex/releases/latest/bibtex-binary-release.zip) and click on it.
3. Save bibtex-binary-release.zip in your "Downloads" folder
4. Find the downloaded zip file and unzip it (e.g. unzip ~/Downloads/bibtex-binary-release.zip)
5. In the unziped directory and find for **dist/linux-amd64/bibfilter**
6. Copy the *bibfilter* to a "bin" directory (e.g. cp ~/Downloads/bibtex-master/dist/linux-amd64/bibfilter ~/bin/)
7. From the shell prompt run `bibfilter -h` to confirm you were successful

### Raspberry Pi

If you are using a Raspberry Pi 2 or later use the ARM7 binary, ARM6 is only for the first generaiton Raspberry Pi.

1. Go to [github.com/caltechlibrary/bibtex/releases/latest](https://github.com/caltechlibrary/bibtex/releases/latest)
2. Find the link for [bibtex-binary-release.zip](https://github.com/caltechlibrary/bibtex/releases/latest/bibtex-binary-release.zip) and click on it.
3. Save bibtex-binary-release.zip in your "Downloads" folder
4. Find the downloaded zip file and unzip it (e.g. unzip ~/Downloads/bibtex-binary-release.zip)
5. In the unziped directory and find for **dist/raspberrypi-arm7/bibfilter**
6. copy the *bibfilter* to a "bin" directory (e.g. cp ~/Downloads/bibtex-master/dist/raspberrypi-arm7/bibfilter ~/bin/)
    + if you are using an original Raspberry Pi you should copy the ARM6 version instead
7. From the shell prompt run `bibfilter -h` to confirm you were successful


## Compiling from source

If you have go v1.6.2 or better installed then should be able to "go get" to install all the **bibtex** utilities and
package. You will need the GOBIN environment variable set. In this example I've set it to $HOME/bin.

```
    GOBIN=$HOME/bin
    go get github.com/caltechlibrary/bibtex/...
```


