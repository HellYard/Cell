Date Variable
==================
The date variable is a reserve variable for Cell. It contains information about the current date, which can be used
for copyrights, countdowns, and much more.

Syntax
--------------
```
{{ date }}
```

Example
--------------
```
{* We use the date reserved variable to output the current year, for our copyright notice. *}
<p> Copyright &copy; {{ date.year }} Daniel Vidmar. All rights reserved.</p>
{* We can also use the date reserved variable to output today's date. *}
<p>Today is the {{ date.day }} of {{ date.month }}, {{ date.year }}.</p>
```