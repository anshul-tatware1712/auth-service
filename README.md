# Auth Service

Go service built with Bazel. Run all commands from this project directory.

## Prerequisites

- Go 1.25.5 or a compatible newer version for Go module commands.
- Bazel 9.x. The project declares rules_go and Gazelle in MODULE.bazel.

## Local Setup

1. Create a local `.env` file containing the server listen address:

   ```dotenv
   PORT=:8080
   ```

2. Download Go dependencies:

   ```sh
   go mod download
   ```

3. Build and start the service:

   ```sh
   bazel run //:app
   ```

The service listens at http://localhost:8080. Stop it with Ctrl+C.
`bazel run` builds the target before starting it; a separate build is optional.

## Everyday Commands

Build without starting the service:

```sh
bazel build //:app
```

After ordinary Go code edits, stop the running service and restart it:

```sh
bazel run //:app
```

After adding, moving, or deleting Go files, creating local packages, or changing
imports, regenerate build rules before running:

```sh
bazel run //:gazelle &&
bazel run //:app
```

## Add or Update a Dependency

1. Install the package with `go get`, not `go fetch`. For example:

   ```sh
   go get github.com/go-chi/chi/v5
   ```

   Use `go get package/path@version` to select a specific version.

2. Import and use the package in your Go code, then save the files.
   Do this before tidying: `go mod tidy` removes unused dependencies.

3. Synchronize dependencies, regenerate build rules, and run:

   ```sh
   go mod tidy &&
   bazel mod tidy &&
   bazel run //:gazelle &&
   bazel run //:app
   ```

`go mod tidy` updates go.mod and go.sum. `bazel mod tidy` synchronizes the
go_deps repository imports in MODULE.bazel. Gazelle updates BUILD.bazel files.
You normally do not need to add dependency labels manually.

Use the project-local `//:gazelle` target, not `@gazelle//:gazelle`.
Use `&&` between commands so a failed step stops subsequent steps.

## Environment and Version Control

- The app requires `.env` and a `PORT` value such as `:8080`.
- The app target declares `.env` in `data`, making it available to `bazel run`.
- This also makes `.env` a Bazel build input that may enter build caches.
  Do not use this arrangement for production secrets or shared remote caches.
- Keep `.env` and generated `bazel-*` paths out of Git.
- Commit Go source, BUILD.bazel files, MODULE.bazel, MODULE.bazel.lock,
  go.mod, and go.sum.

## Troubleshooting

- **Unknown external repository:** Ensure the package is imported in Go code,
  then run the dependency synchronization commands above.
- **Duplicate app target:** Keep only one `go_binary` named `app` in the root
  BUILD.bazel. Edit the existing target rather than appending another one.
- **Missing .env:** Create it in the project root and retain `data = [".env"]`
  in the app target.
- **Old executable runs after a failed build:** Chain commands with `&&`, or
  use `bazel run //:app` to build and run in one command.