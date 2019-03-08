# fksq  

**F**a**K**e **S**oftware **Q**uote generator using Markov Chains.  

## Installation

Download pre-compiled binaries for Mac, Windows and Linux from the [releases](https://github.com/vaithak/fksq/releases) page.  

### Mac OSX Installation of pre-compiled binaries
Please submit an issue if they don't work for you.  

1. Make sure you download and place the binary in a folder that's on your `$PATH`.  If you are unsure what this means, go to *step 2*. Otherwise, skip to *step 3*  
2. Create a `bin` directory under your home directory.  
```
$ mkdir ~/bin
$ cd ~/bin
```   
3. Add the following line at the end of your `~/.bash_profile` file.  [Link with instructions](https://natelandau.com/my-mac-osx-bash_profile/) on how to find this file  
```sh
export PATH=$PATH:$HOME/bin
```  

4. Download the `fksq` binary for OSX and rename it.  
```sh
$ wget https://github.com/vaithak/fksq/releases/download/v1.0/fksq-darwin-amd64  
$ mv fksq-darwin-amd64 fksq
```
5. Finally, make the binary an executable file and you are good to go!
```
$ chmod +x fksq
```  

### Installation of pre-compiled binaries on Ubuntu 
Please submit an issue if they don't work for you.  

1. Make sure you download and place the binary in a folder that's on your `$PATH`.  If you are unsure what this means, go to *step 2*. Otherwise, skip to *step 3*  
2. Create a `bin` directory under your home directory.  
```
$ mkdir ~/bin
$ cd ~/bin
```   
3. Add the following line at the end of your `~/.bashrc` file.  
```sh
export PATH=$PATH:$HOME/bin
```  

4. Download the `fksq` binary for 64 bit linux and rename it.  
```sh
$ wget https://github.com/vaithak/fksq/releases/download/v1.0/fksq-linux-amd64  
$ mv fksq-linux-amd64 fksq
```
5. Finally, make the binary an executable file and you are good to go!
```
$ chmod +x fksq
```  

### Downloading using go-get
```sh
go get github.com/vaithak/fksq  
```  

## Usage

```
  Usage:
      fksq
  
  Flags: 
    -in string
      	input file (default "data.txt")
    -n int
      	number of words to use as prefix (default 2)
    -words int
      	Minimum number of words. (default 8)
```  

## Sample Generated Quotes

```
Instead, the internals should be really excellent. Because this is not well written.  
For me, a comment is almost a note like "I should try to rewrite this code later."  
I think it is also easy to fall out of touch with the Universe to create bigger and  
better idiot-proof programs, while the Universe is trying to decide what to do,  
let us concentrate rather on explaining to human beings who can figure costs,  
and so they are forgiving of certain kinds of errors.
 -- FaKe Software Quote 
 
We all suck. The Internet? Is that thing still around?
 -- FaKe Software Quote 
 
All OO languages show some tendency to be too clever for the first time.
 -- FaKe Software Quote 
 
Turing and von Neumann understood from the beginning that a PL is a fundamentally different job.
 -- FaKe Software Quote  
 
Avoid OO whenever possible. Contrary to common belief, most problems are local and you don’t  
know quite where you are making one.
 -- FaKe Software Quote 
 
Necessity is the essence of things? The so-called "desktop metaphor" of today's workstation is  
instead an "airplane-seat" metaphor.
 -- FaKe Software Quote

```
