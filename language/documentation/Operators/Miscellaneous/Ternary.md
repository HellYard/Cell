Ternary Operator
==================
The Ternary Operator returns the result of the expression on the left of the ":"(colon) character if value is true,
 otherwise the result of the expression on the right is returned.

Syntax
--------------
```
Backend:
<?php
  $value = true;
  $default = 'Hello';
  $default2 = 'World';
?>

Frontend:
{% value? default : default2 %}
{* Outputs: Hello *}
```