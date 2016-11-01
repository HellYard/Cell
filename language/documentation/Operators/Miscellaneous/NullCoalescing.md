Null Coalescing Operator
==================
The Null Coalescing Operator returns the value if said value is not empty, otherwise it returns the value on the right
hand of the expression.

Syntax
--------------
```
Backend:
<?php
  $value = null;
  $default = 'Default';
?>

Frontend:
{% value?? default %}
{* Outputs: Default *}
```