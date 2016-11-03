Iterate Statement
==================
The iterate statement is a shorthand version of the for statement.
This shorthand method does not require you to specify a name for
each value iterated over, nor does it require an in keyword. Instead
it forces calls to the item that is being iterated over to be
the iteration item.

Syntax
--------------
```
Backend:
<?php
   $arr = [
      0 => [
        "one",
        "two"
      ]
      1 => [
        "three",
        "four"
      ]
   ];
?>

Frontend:
{% iterate arr %}
{% arr.0 %}
{% end %}

{* Output: onethree *}
```