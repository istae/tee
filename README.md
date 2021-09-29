## Tee Programming Language

Tee is a very simple language meant for learning the inner mechanisms of compilers (eg: lexers, parsers, etc..). 

Currently, here's what can be parsed:

```
x = 100.0 + 5.0 * 2.0 + 3.0 / 12.0
y = 12 - 10 * 2
z = x * y
if x > 10.0 {
	x = 123 * 12
}

for x < 1500 {
	x = x + 1
}
```

To run:

```
make binary && ./dist/tee
```