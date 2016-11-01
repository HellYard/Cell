Join Modifier
==================
The Join Modifier transforms an array value into a string concatenated by a string.

Arguments
--------------
Character - random string - Defaults to ,

Syntax
--------------
```
Backend:
<?php
  $arr = ['test', 'test2'];
?>

Frontend:
{% arr join=* %}
{* Outputs: test*test2 *}
```