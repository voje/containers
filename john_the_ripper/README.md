# John
John-the-ripper - cool tool for cracking Linux/Unix/MacOS passwords.

We have a line from a linux system `shadow` file and a password: `verifikacija1` which should be similar to the actual password stored in the hash.  

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
