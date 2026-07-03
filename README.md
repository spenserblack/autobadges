# autobadges

## Description

This is a highly opinionated tool that generates badge text which you can copy and
paste into your README. All the cool coders use `| xclip -selection clipboard`.

See [`BADGES.md`](./BADGES.md) for a list of badges that are generated.

As an opinionated tool, this makes the following assumptions:

- The README is always in the same directory as manifest files (`Cargo.toml`, `package.json`, etc.).
- The README's title is a *level 1* header.

## Badge guidelines

These are not strict rules, but most badges should follow these guidelines.

The badge should be useful. Useful badges include these categories:

- Releases
- Documentation links
- Dynamic statistics that are useful for developers (e.g. the CI status)

Generally, badges should show *dynamic* information, and not be a hard-coded value.
Links to documentation are an exception to this guideline, because it is helpful
to provide an easily visible link to your project's documentation.

This project will typically avoid hard-coded badges, like
`https://img.shields.io/badge/<label>-<value>-<color>.svg`.
