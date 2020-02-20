---
layout: default
title: Install
parent: dictutil
---

# Install

## Usage

```
Usage: dictutil install [options] dictzip

Options:
  -k, --kobo string      KOBOeReader path (default: automatically detected)
  -l, --locale string    Locale name to use (format: ALPHANUMERIC{2}; translation dictionaries are not supported) (default: detected from filename if in format dicthtml-**.zip)
  -n, --name string      Custom additional label for dictionary (ignored when replacing built-in dictionaries)
  -b, --builtin string   How to handle built-in locales [replace = replace and prevent from syncing] [ignore = replace and leave syncing as-is] (default "replace")
  -h, --help             Show this help text

Note:
  If you are not replacing a built-in dictionary, the 'Enable searches on extra
  dictionaries patch' must be installed, or you will not be able to select
  your custom dictionary.
```

## Examples
TODO