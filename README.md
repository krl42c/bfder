# bfder

Byte index finder tool.

Usage:

```
bfder [search_term] [directory]
```

```sh
foo@bar:~$ ./bfder autocall ../autocall
{../autocall/.coverage}: 20033
{../autocall/Dockerfile}: 26
{../autocall/Makefile}: 138
{../autocall/README.md}: 123
{../autocall/pyproject.toml}: 94
{../autocall/run.py}: 40
```