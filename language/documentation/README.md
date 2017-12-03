Cell
===============
The contents of this directory are meant to help front end developers that wish to utilize Cell's full feature set.

About
---------------
Cell is a template engine written in Go. It was designed to be dynamic, fast, and powerful while remaining as lightweight
as possible. Cell aims to provide a powerful feature set with a minimalist-style syntax.

Navigation
---------------
- [Tags](Tags/README.md)
  - [Comment](Tags/Comment.md) {* *}
  - [Escape](Tags/Escape.md) {! !}
  - [Reserved](Tags/Reserved.md) {{ }}
  - [Statement](Tags/Statement.md) {% %}
  - [Syntax](Tags/Syntax.md) {` `}
- [Variables](Variables/README.md)
  - [Defined](Variables/Defined.md)
  - [Reserved](Variables/Reserved.md)
    - [Cache](Variables/Reserved/Cache.md)
    - [Date](Variables/Reserved/Date.md)
    - [IPCache](Variables/Reserved/IPCache.md)
- [Modifiers](Modifiers/README.md)
  - [Array](Array/README.md)
    - [JSON](Array/JSON.md)
    - [Split](Array/Split.md)
    - [Sub](Array/Sub.md)
  - [Numerical](Modifiers/Numerical/README.md)
    - [Ceil](Modifiers/Numerical/Ceil.md)
    - [Datify](Modifiers/Numerical/Datify.md)
    - [Floor](Modifiers/Numerical/Floor.md)
    - [Pad](Modifiers/Numerical/Pad.md)
    - [Round](Modifiers/Numerical/Round.md)
  - [String](Modifiers/String/README.md)
    - [Escape](Modifiers/String/Escape.md)
    - [Join](Modifiers/String/Join.md)
    - [Lower](Modifiers/String/Lower.md)
    - [Trim](Modifiers/String/Trim.md)
    - [Upper](Modifiers/String/Upper.md)
    - [UpperFirst](Modifiers/String/UpperFirst.md)
- [Operators](Operators/README.md)
  - [Arithmetic](Operators/Arithmetic/README.md)
    - [Addition](Operators/Arithmetic/Addition.md)
    - [Division](Operators/Arithmetic/Division.md)
    - [Exponential](Operators/Arithmetic/Exponential.md)
    - [Multiplication](Operators/Arithmetic/Multiplication.md)
    - [Remainder](Operators/Arithmetic/Remainder.md)
    - [Subtraction](Operators/Arithmetic/Subtraction.md)
  - [Assignment](Operators/Assignment/README.md)
    - [AS](Operators/Assignment/AS.md)
  - [Comparison](Operators/Comparison/README.md)
    - [IS](Operators/Comparison/IS.md)
    - [NOT](Operators/Comparison/NOT.md)
    - [SUB](Operators/Comparison/SUB.md)
  - [Logical](Operators/Logical/README.md)
    - [AND](Operators/Logical/AND.md)
    - [OR](Operators/Logical/OR.md)
  - [Miscellaneous](Operators/Miscellaneous/README.md)
    - [Null Coalescing](Operators/Miscellaneous/NullCoalescing.md)
    - [Ternary](Operators/Miscellaneous/Ternary.md)
- [Statements](Statements/README.md)
  - [Copy](Statements/Copy.md)
  - [Def](Statements/Def.md)
  - [Empty](Statements/Empty.md)
  - [For](Statements/For.md)
  - [If-Else-Then](Statements/IfElseThen.md)
  - [Insert](Statements/Insert.md)
  - [Iterate](Statements/Iterate.md)
  - [Override](Statements/Override.md)
  - [Sinsert](Statements/SInsert.md)
  - [Snip](Statements/Snip.md)
  - [Throw](Statements/Throw.md)
  - [Where](Statements/Where.md)