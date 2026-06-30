# increment.nu

A Nushell script that **increments numbers** (dec, hex, bin, oct) and **toggles boolean-like values**, designed for Helix's `:pipe` command.

Replaces `C-a`/`C-x` with a version that also handles `true`/`false`, `yes`/`no`, `on`/`off`, etc.

---

## How it works

Reads the selected word from stdin, checks what it is, and:

| Input | `+1` (odd) | `-1` (odd) | `+2` (even) |
|-------|-----------|------------|-------------|
| `42` | `43` | `41` | `44` |
| `-5` | `-4` | `-6` | `-3` |
| `0xFF` | `0x100` | `0xFE` | `0x101` |
| `0b1010` | `0b1011` | `0b1001` | `0b1100` |
| `0o77` | `0o100` | `0o76` | `0o101` |
| `true` | `false` | `false` | `true` |
| `yes` | `no` | `no` | `yes` |
| `enable` | `disable` | `disable` | `enable` |
| `hello` | `hello` | `hello` | `hello` |

The toggle logic: **odd amounts** (1, -1, 3, -5, ...) trigger toggling for non-numeric words; **even amounts** (2, -2, 0, 4, ...) leave them unchanged. Numbers are always incremented/decremented regardless of parity.

---

## Usage

### Directly

```nu
echo "42" | nu --stdin increment.nu 1
# → 43

echo "0xFF" | nu --stdin increment.nu -1
# → 0xFE

echo "true" | nu --stdin increment.nu 1
# → false
```

**Must use `--stdin`** — Nushell's `$in` does not read from OS-level stdin without it.

### Helix

Add to your `config.toml`:

```toml
[keys.normal]
C-a = "@miw:pipe use <PATH>/increment.nu; $in | increment run 1<ret>"
C-x = "@miw:pipe use <PATH>/increment.nu; $in | increment run -1<ret>"
```

Replace `<PATH>` with the absolute path to the repo (e.g. `C:/Users/giuli/git/increment`).

The macro `@miw` selects inner word, then `:` enters command mode, types the pipe command, and `<ret>` runs it.

Your `shell` config must be:
```toml
[editor]
shell = ["nu", "--stdin", "-c"]
```

This makes Helix pass the pipe command as Nushell code (`nu --stdin -c "..."`), so `use` and `$in` work correctly.

---

## Supported formats

### Numbers

| Format | Prefix | Example | After `+1` |
|--------|--------|---------|------------|
| Decimal | none | `42` | `43` |
| Hex | `0x` / `0X` | `0xFF` | `0x100` |
| Binary | `0b` / `0B` | `0b1010` | `0b1011` |
| Octal | `0o` / `0O` | `0o77` | `0o100` |

All formats support negative values (e.g., `-5`, `-0xFF`). Output preserves the same format.

### Toggle values

| Input | Output |
|-------|--------|
| `true` | `false` |
| `True` | `False` |
| `TRUE` | `FALSE` |
| `yes` | `no` |
| `on` | `off` |
| `enable` | `disable` |
| `enabled` | `disabled` |

Each pair works in both directions and preserves case. All case variants of each pair are supported. Toggle only fires on **odd** amounts (1, -1, 3, -5, ...). Even amounts (2, 0, -2) pass non-numeric words through unchanged.

---

## Tests

```nu
nu test.nu
```

Run from the repo directory (`cd ~/git/increment`).

All 50+ test cases cover:
- Decimal increment/decrement (including negatives and large numbers)
- Hex, binary, octal with carry-over
- All toggle pairs with both `+1` and `-1`
- Even amounts leaving values unchanged
- Unknown values passed through
- Empty input
