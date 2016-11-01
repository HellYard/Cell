AND Operator
==================
The AND Operator determines if both sides of an expression are true.

Syntax
--------------
```
Backend:
<?php
  $value = false;
  $value1 = true;
?>

Frontend:
{% value AND value1 %}
{* Outputs: false *}

{% value1 AND 1 is 1 %}
{* Outputs: true *}
```