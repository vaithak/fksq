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
$ mv fksq-darwin-amd64 fsq
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
$ mv fsq-darwin-amd64 fsq
```
5. Finally, make the binary an executable file and you are good to go!
```
$ chmod +x fksq
```  

### Downloading using go-get
```sh
go get github.com/vaithak/fksq  
```
