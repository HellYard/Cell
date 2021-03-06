Modifiers
===============
Cell supports the use of something called "modifies" for variables and expressions. They provide extra formatting
option for the values of variables and expressions without the need for added backend code. This allows you to
format variables without the extra work.

Syntax
-------------
Modifiers are used by simply adding a space after the variable or expression in a statement. You may use multiple
modifiers by concatenating them with a ","(comma) character. Some modifiers allow for parameters, these are denoted
by adding an "="(equals) sign after the modifier name followed by the parameters(again concatenated with a ","(comma)).

```
Backend:
<?php
  $hello = ' world';
  $num = 1.5345114;
?>

Frontend:
{* A simple modifier *}
Hello, {% hello uppercase %}
{* Outputs: Hello,  WORLD  *}

{* Multiple modifiers. *}
Hello, {% hello uppercase,trim %}
{* Outputs: Hello, WORLD *}

{* A modifier with a parameter *}
{% num round=1 %}
{* Outputs: 1.5 *}
```

Navigation
---------------
- [Array](Array/README.md)
  - [JSON](Array/JSON.md)
  - [Join](Array/Join.md)
  - [Sub](Array/Sub.md)
- [Generic](Generic/README.md)
  - [Ignore](Generic/Ignore.md)
  - [Default](Generic/Default.md)
- [Numerical](Numerical/README.md)
  - [Ceil](Numerical/Ceil.md)
  - [Datify](Numerical/Datify.md)
  - [Floor](Numerical/Floor.md)
  - [Pad](Numerical/Pad.md)
  - [Round](Numerical/Round.md)
- [String](String/README.md)
  - [Escape](String/Escape.md)
  - [Split](String/Split.md)
  - [Lower](String/Lower.md)
  - [Trim](String/Trim.md)
  - [Upper](String/Upper.md)
  - [UpperFirst](String/UpperFirst.md)