## Tee Programming Language

Tee is a very simple language meant for learning the inner mechanisms of compilers (eg: lexers, parsers, etc..). 

Some rules:

1. All numbers are float64
2. if and for statements share scopes with the parent scope, meaning:
	```
	x = 2
	if ... {
		x = 3
	}
	```
	the parent x is now 3
3. Math operators have proper precedence, so * and / are computed properly. 

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