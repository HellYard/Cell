Insert Statement
==================
The Insert Statement lets your include external template files, or Snippets in your template file.

Syntax
--------------
```
Index.cell
<p>test</p>

Home.cell
<p>Hello</p>
{% insert Index.cell %}
{* Output:
 * <p>Hello</p>
 * <p>test</p>
 *}
```