AS Operator
==================
The AS Operator assigns a value to a variable

Syntax
--------------
```
Backend:
<?php
  $value = 'Hello';
?>

Frontend:
{% temp as value %}
{% value as 'World' %}
{% temp + " " + value %}
{* Outputs: Hello World *}
```