Cache Variable
==================
The cache variable is used to enable caching for a certain template. This will tell cell to cache a template, and save
the parsed template for a simple read later. The default age for caching is two hours, and the cached template will be
saved to a file using the .pcell file format.

Syntax
--------------
```
{{ cache [age(in minutes)] }}
```

Example
--------------
```
{* We want to cache the example index.cell for five hours so the Cell implementation doesn't have to reevaluate
   everything. For this, we simply add the following line to the top of the file. *}
{{ cache 300 }}
{* This will save the parsed content of index.cell to a file called index.pcell, which will simply be outputted to the
   screen without parsing the next time the page is loaded within five hours. *}
```