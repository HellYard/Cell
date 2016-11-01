Defined Variables
==================
Defined variables are simply variables that exist in the backend of the software for which the template file is for. Defined
variables require [Statement](../Tags/Statement.md) or [Escape](../Tags/Escape.md) Tags.

Syntax
--------------
```
{% variable_name %}
```

Example
--------------
```
Backend:
<?php
   $hello = 'world';
?>

Frontend:
{* We are able to output variables from the backend using the defined variables system. *}
<p>Hello, {% hello %}</p>
```