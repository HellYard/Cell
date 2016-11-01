Split Modifier
==================
The Split Modifier transforms a string variable into an array by means of splitting it into indexes at the delimiter.

Arguments
--------------
Character - random string - Defaults to ,

Syntax
--------------
```
Backend:
<?php
  $str = test1*test2;
?>

Frontend:
{% str split=* %}
{* Outputs: ['test1', 'test2'] *}
```