# Go + React Template

This is an all-in-one minimal solution for React SPA embedded within a single Go binary. It uses the following stack:

- Go with chi router v5
- React with react router v7
- air hot reloading
- mantine ui for components library

## Installation

Use this repository as a template, clone it and run:

```
make install
```

## Why is this template cursed?

Because I opt for the most minimal solutions, some decisions were made that may be read as cursed:

- React frontend and Go backend are interconnected within single directory level
- The development script (`make watch` or just `make`) does not run node development server but instead builds the production SPA and embeds it in Go binary
- Frontend uses traditional SPA without any complexity of SSR for snappy client-side routing
- I would love to use HTMX instead but the extraordinary Mantine library only supports React