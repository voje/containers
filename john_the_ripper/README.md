# John

You'll need a `wordlist.txt` file and a `shadow` file (or at least a line from the shadow file).  
Put your files in `./workdir` and start up the container: 
```bash
$ make run
```

To run john, you'll want to `cd` into `/git/JohnTheRipper/run`:

```
./john --help
```

Generally, either create a wordlist (use the installed `crunch` or your own tools) or define a mask and provide a `/etc/shadow` file with password hashes.  

Example:
```bash
$ cd /git/JohnTheRipper/run
$ ./john --mask=passwd200?d /workdir/shadow.txt
```

## Wordlist generator
I've added my own wordlist generator -- run it whth `make generate`.  
It will read a list of possible passwords: `./workdir/wordlist.txt` and out of each password generate 
a list of permutations.  

To get a list of permutations:
```bash
./wordlist_generator.py -h
```
Permutations will be applied in order of parameters.  
