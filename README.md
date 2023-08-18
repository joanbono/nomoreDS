# nomoreDS

Small utility to remove `.DS_Store` files recursively.

## Usage

```sh
$ nomoreDS -version

     📄  nomoreDS 1.0.0

# By default, deletes the .DS_Store file recursively from ${PWD}. No output
$ nomoreDS

# Add verbose output to list the files deleted
$ nomoreDS -V
[+] Removing /tmp/.DS_Store...
[+] Removing /tmp/A/B/.DS_Store...
[+] Removing /tmp/A/B/C/D/E/.DS_Store...
[i] Files deleted: 3
```