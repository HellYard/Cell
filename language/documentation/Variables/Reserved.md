Reserved Variables
==================
Reserved variables are variables that are contained within the core of Cell, which provide commonly used data
to front end developers, such as the date. Reserved variables require [Reserved Tags](../Tags/Reserved.md).

Syntax
--------------
```
{{ reserved_variable_name }}
```

Example
--------------
```
{* We use the date reserved variable to output the current year, for our copyright notice. *}
<p> Copyright &copy; {{ date.year }} Daniel Vidmar. All rights reserved.</p>
```