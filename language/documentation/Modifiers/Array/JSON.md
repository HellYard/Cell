JSON Modifier
==================
The JSON Modifier transforms an array value into a JSON-friendly string.

Syntax
--------------
```
Backend:
<?php
  $arr = ['test', 'test2' => ['testtest', 'testtest2']];
?>

Frontend:
{% arr json %}
{* Outputs: {"0":"test","test2":["testtest","testtest2"]} *}
```