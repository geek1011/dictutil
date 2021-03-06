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
  -k, --kobo string         KOBOeReader path (default: automatically detected)
  -l, --locale string       Locale name to use (format: ALPHANUMERIC{2}[-ALPHANUMERIC{2}]) (default: detected from filename if in format dicthtml-**.zip)
  -n, --name string         Custom additional label for dictionary (ignored when replacing built-in dictionaries) (doesn't have any effect on 4.20.14601+)
  -b, --builtin string      How to handle built-in locales [replace = replace and prevent from syncing] [ignore = replace and leave syncing as-is] (doesn't have any effect on 4.24.15672+) (default "replace")
  -B, --no-custom           Whether to force installation to .kobo/dict instead of .kobo/custom-dict (4.24.15672+ only)
      --use-extra-locales   Whether to use ExtraLocales on 4.24.15672+ if not a built-in dictionary (this is not required anymore since 4.24.15672) (4.24.15672+ only)
  -h, --help                Show this help text

Note:
  If you are not replacing a built-in dictionary and are using a firmware
  version before 4.24.15672, the 'Enable searches on extra dictionaries patch'
  must be installed or you will not be able to select your custom dictionary.
```

## Examples

**Install a dictionary with the locale in the filename (dicthtml-\*\*.zip):**

```sh
dictutil install dicthtml-aa.zip
```

**Install a dictionary with a different locale:**

```sh
dictutil install --locale aa mydictionary.zip
```

**Install a dictionary on a specific Kobo:**

```sh
dictutil install --kobo /path/to/KOBOeReader dicthtml-aa.zip
```

**Install a dictionary with a custom label (4.19.14123 and older):**

```sh
dictutil install --name "My Dictionary" dicthtml-aa.zip
```

## Details
See [installing dictionaries](../dicthtml/install.html) for more details on how this works.
