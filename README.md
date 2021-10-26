## Tee Programming Language

Tee is a very simple language meant for learning the inner mechanisms of compilers (eg: lexers, parsers, etc..). 

Lexer creates tokens using a list of token parsers. Tokens are essentially non-white-space strings that get
matched to a certain type, like for, if, =. 

Parser takes these tokens and creates an Abstract Syntax Trees. Function calls, expressions, if statements, etc.. all have different looking trees.

Eval is an interpreter executes the each the code based on the AST. 

Some notes:

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

Currently, here's what can be parsed (see also dev/main.go for some test code):

```
x = 100.0 + 5.0 * 2.0 + 3.0 / 12.0
y = 1
func lol(a, b) {
	func x(h,j) {
		if h / 2 > 1 {
			print(h/2, j, "\n")
			y = h * j
		}	
	}
	x(a,b)
}

lol(3,4)

x = y * 2

i = 0
for 1 > 0 {
	i = i + 1
	if i > 100 {
		break
	}
	print("hello\n")
}
print(i)
```

To run:

```
make binary && ./dist/tee
```