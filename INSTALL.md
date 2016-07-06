
# Installation

*bibfilter* is a command line program run from a shell like Bash. If you download the 
repository a compiled version is in the dist directory. The compiled binary matching
your computer type and operating system can be copied to a bin directory in your PATH.

Compiled versions are available for Mac OS X (amd64 processor), Linux (amd64), Windows
(amd64) and Rapsberry Pi (both ARM6 and ARM7)

## Mac OS X

1. Go to [github.com/caltechlibrary/bibtex](https://github.com/caltechlibrary/bibtex)
2. Click on the green "Clone or Download" button on the right side of the page
3. A panel will open, click on "Download Zip"
4. Open a finder window downloaded file and unzip it (e.g. bibtex-master.zip)
5. Look in the Unziped folder and find dist/maxosx-amd64/bibfilter
6. Drag (or copy) the *bibfilter* to a "bin" directory in your path
7. Open and "Terminal" and run `bibfilter -h`

## Windows

1. Go to [github.com/caltechlibrary/bibtex](https://github.com/caltechlibrary/bibtex)
2. Click on the green "Clone or Download" button on the right side of the page
3. A panel will open, click on "Download Zip"
4. Open the file manager find the downloaded file and unzip it (e.g. bibtex-master.zip)
5. Look in the Unziped folder and find dist/windows/bibfilter.exe
6. Drag (or copy) the *bibfilter.exe* to a "bin" directory in your path
7. Open Bash and and run `bibfilter -h`

## Linux

1. Go to [github.com/caltechlibrary/bibtex](https://github.com/caltechlibrary/bibtex)
2. Click on the green "Clone or Download" button on the right side of the page
3. A panel will open, click on "Download Zip"
4. find the downloaded zip file and unzip it (e.g. unzip ~/Downloads/bibtex-master.zip)
5. In the Unziped directory and find for dist/linux-amd64/bibfilter
6. copy the *bibfilter* to a "bin" directory (e.g. cp ~/Downloads/bibtex-master/dist/linux-amd64/bibfilter ~/bin/)
7. From the shell prompt run `bibfilter -h`

## Raspberry Pi

If you are using a Raspberry Pi 2 or later use the ARM7 binary, ARM6 is only for the first generaiton Raspberry Pi.

1. Go to [github.com/caltechlibrary/bibtex](https://github.com/caltechlibrary/bibtex)
2. Click on the green "Clone or Download" button on the right side of the page
3. A panel will open, click on "Download Zip"
4. find the downloaded zip file and unzip it (e.g. unzip ~/Downloads/bibtex-master.zip)
5. In the Unziped directory and find for dist/raspberrypi-arm7/bibfilter
6. copy the *bibfilter* to a "bin" directory (e.g. cp ~/Downloads/bibtex-master/dist/raspberrypi-arm7/bibfilter ~/bin/)
    + if you are using an original Raspberry Pi you should copy the ARM6 version instead
7. From the shell prompt run `bibfilter -h`

