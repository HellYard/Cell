Sub Modifier
==================
The Sub Modifier returns an array with the specified start and end of an array variable.

Arguments
--------------
Start - numerical - 0
End - numerical - size of array - 1

Syntax
--------------
```
Backend:
<?php
  $arr = ['test', 'test2', 'test3', 'test4'];
?>

Frontend:
{% arr sub=1 %}
{* Outputs: ['test2', 'test3', 'test4'] *}
```