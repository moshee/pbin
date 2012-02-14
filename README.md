# pbin
## What?
A tool that uploads files or text from stdin to pastebin.com. Use this:

- If `pastebinit` is too long to type
- If you like Go
- For the novelty factor
- To support free software (no not really)

## How?

```
Usage: pbin [options] [filename]
  -f="text": Paste language
  -n="": Paste name
  -p=false: Private paste
  -s=false: Accept input from STDIN
  -x="never": Time before paste expires
Option ‘-x’ takes one of five arguments:
  never  Never expire (default)
minutes  Expire in 10 minutes
  hours  Expire in 1 hour
   days  Expire in 1 day
  month  Expire in 1 month
```

## Hatemail etc
moshee on Freenode or Rizon.
