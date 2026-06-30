# increment

Go binary for incrementing numbers (dec, hex, bin, oct) and toggling boolean-like values. Designed for Helix's `:pipe` command.

Replaces `C-a`/`C-x` with a version that handles `true`/`false`, `yes`/`no`, `on`/`off`, etc.

---

## Install

```
go install github.com/eevleevs/increment@latest
```

The binary will be installed to `$GOPATH/bin/increment.exe` (make sure that's on your `PATH`).

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

## Helix

```toml
[keys.normal]
C-a = '@miw:pipe ^increment 1<ret>'
C-x = '@miw:pipe ^increment -1<ret>'
```

The `^` prefix tells Nu to run `increment` as an external command, inheriting stdin from Helix. The macro `@miw` selects inner word, then `:` enters command mode, types the pipe command, and `<ret>` runs it.

## Direct

```
echo 42 | increment 1
echo 0xFF | increment -1
echo true | increment 1
```

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

```
go test -v ./...
```

Run from the repo directory.
