# Readme Generator for Helm (Go)

> **Go port of [`bitnami/readme-generator-for-helm`](https://github.com/bitnami/readme-generator-for-helm)**
>
> Drop‑in replacement that keeps the same CLI **and metadata syntax**, but ships as a single statically‑linked Go binary.

---

* Autogenerates a **Parameters** table in your chart’s `README.md` from the metadata found in `values.yaml`.
* Optionally emits an **OpenAPI v3** JSON schema that describes the structure of `values.yaml`.

Both features behave *exactly* like the original Bitnami tool, so you can switch by merely replacing the binary in your pipeline.

---

## How it works

The generator looks for Javadoc‑style comments inside `values.yaml`. It validates that every real key has corresponding metadata (and vice‑versa); if everything lines up it rewrites the `## Parameters` section of `README.md` and/or writes an OpenAPI schema file.  If mismatches are detected it prints a detailed error list and exits with a non‑zero status, making it CI‑friendly.

The table it injects has the familiar structure

```markdown
## Parameters

### Section 1 title

| Name      | Description             | Default        |
|:----------|:------------------------|:---------------|
| `value_1` | Description for value 1 | `defaultValue` |
| `value_2` | Description for value 2 | `defaultValue` |

### Section 2 title

| Name      | Description             | Default        |
|:----------|:------------------------|:---------------|
| `value_a` | Description for value a | `defaultValue` |
```

The top‑level heading (`## Parameters`, `### Parameters`, …) is detected dynamically; its text can be customised via the [configuration file](#configuration-file).

---

## Requirements

* Go **1.22+** (any platform supported by Go)

---

## Installation

### Using `go install` (recommended)

```console
go install github.com/cozystack/readme-generator-for-helm@latest
```

The binary lands in `$(go env GOPATH)/bin` (usually `~/go/bin`). Make sure that directory is on your `$PATH`.

### Download a pre‑built release binary

Head over to [https://github.com/cozystack/readme-generator-for-helm/releases](https://github.com/cozystack/readme-generator-for-helm/releases), grab the archive for your OS/arch, unpack it somewhere on your `$PATH`.

### Build from source

```console
git clone https://github.com/cozystack/readme-generator-for-helm
cd readme-generator-for-helm
go build -o readme-generator-for-helm
```

---

## Basic usage

```console
readme-generator-for-helm [options]

Options:
  -v, --values  <file>   Path to the values.yaml file (required)
  -r, --readme  <file>   Path to the README.md file to update
  -c, --config  <file>   Path to config.json (optional; built‑in defaults if omitted)
  -s, --schema  <file>   Path for the generated OpenAPI Schema
      --version          Print program version and exit
  -h, --help             Show help
```

*At least one of* `--readme` *or* `--schema` *must be provided.*

---

## `values.yaml` metadata

The comment syntax, tags and modifiers are preserved from the original project:

* **Parameter:**     `## @param full.key.path [modifier1,modifier2] Description`
* **Section:**       `## @section Section Title`
* **Skip subtree:**  `## @skip full.key.path`
* **Intermediate object description:** `## @extra full.key.path Description`

Supported modifiers (customisable via the config file):

| Modifier        | Effect                                       |
| --------------- | -------------------------------------------- |
| `array`         | Treat parameter as array, default `[]`       |
| `object`        | Treat parameter as object, default `{}`      |
| `string`        | Force empty string default `""`              |
| `nullable`      | Parameter may be `null`; default stays as‑is |
| `default:VALUE` | Override default with given literal `VALUE`  |

> **Important:** Ordering of tags in the YAML file does not matter, *except* for `@section`, which groups all subsequent `@param`s until the next `@section`.

---

## Configuration file

If you need to change comment delimiters, tag names, or add custom modifiers, provide a JSON file via `--config`:

```json
{
  "comments": { "format": "##" },
  "tags": {
    "param": "@param",
    "section": "@section",
    "descriptionStart": "@descriptionStart",
    "descriptionEnd": "@descriptionEnd",
    "skip": "@skip",
    "extra": "@extra"
  },
  "modifiers": {
    "array": "array",
    "object": "object",
    "string": "string",
    "nullable": "nullable",
    "default": "default"
  },
  "regexp": { "paramsSectionTitle": "Parameters" }
}
```

Omit the flag entirely to use the built‑in defaults (same as above).

---

## License

Apache License 2.0 © 2025 Cozystack.
Portions of the code are adapted from the original Bitnami implementation.
