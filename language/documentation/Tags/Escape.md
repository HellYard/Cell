Escape Tags
==================
Escape tags are used to limit what Cell functions are allowed to be used when, for example, importing a template file
that a third-party could modify. Putting a statement in escape tags means your inserts, raws, and non-safe statements
are securely outputted as plain text, and thus not evaluated. It is recommended to put all insert statements into escape
tags as it prevents unwanted, and potentially malicious template authors the ability to any secure variables and/or data.
The escape tag also has the added ability of automatically escaping any variable of the type string.

Syntax
--------------
```
{! !}
```

Example
--------------
```
{* Include an external template file using safe mode. *}
{! insert thirdparty/example.cell !}
```