OR Operator
==================
The OR Operator determines if one side of an expression is true.

Syntax
--------------
```
Backend:
<?php
  $value = false;
  $value1 = true;
?>

Frontend:
{% value OR value1 %}
{* Outputs: true *}

{% value or false %}
{* Outputs: false *}
```