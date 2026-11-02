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

## How the Application Works

`main.go` loads `.env`, reads `PORT` into the application configuration, creates
the application, and starts its HTTP server. The server uses `chi` and registers
the user route:

```text
POST /user
   -> UserRouter.Register
   -> UserController.RegisterUser
   -> UserService.CreateUser
   -> UserRepository.CreateUser
```

At the moment, creating a user does not write to a database. The repository and
service print messages, and the controller returns `user created`. The handler
does not yet decode or validate a request body.

## Dependency Injection

Dependencies are assembled manually in `Application.Run`; there is no DI
framework. Constructors receive the dependency they need, which keeps the
controller independent of the concrete repository implementation:

```text
UserRepoImpl -> UserServiceImpl -> UserController -> UserRouter -> HTTP server
```

`NewUserService` accepts the `UserRepository` interface. `NewUserController`
accepts the `UserService` interface. This is constructor injection: the app
creates the concrete objects at startup and passes them to the layer that uses
them. The application setup is the composition root, the place where concrete
implementations are selected and wired together.

There is also a `Storage` type whose `NewStorage` constructor creates a
`UserRepository`. However, `NewApplication` stores this value and `Run` does not
use it; `Run` separately calls `NewUserRepo`. As a result, `Storage` is not
currently part of the active dependency-injection chain. A consistent next step
is to create storage once in the composition root, then inject
`Storage.UserRepository` into `NewUserService`. When adding a real database,
create and configure the database connection there too, and pass it to the
repository implementation instead of opening connections inside handlers.

## Code Quality and Next Steps

The overall direction is good for a small service: HTTP handling, business
logic, and data access have separate packages; the service depends on a
repository interface; and dependencies are passed explicitly. This makes the
code easier to replace and test than code that constructs database clients
inside request handlers.

The implementation is still a starting point, not production-ready:

- `CreateUser` errors are discarded by both the service and controller. The
   service always returns `nil`, and the handler responds as though the request
   succeeded. Propagate errors and return an appropriate HTTP status.
- `NewStorage` is currently unused by the request path, so there are two
   different ways to create a user repository. Wire one shared storage instance
   through the application.
- `UserRepoImpl.CreateUser` does not persist data yet. Add request/domain data
   and a real database operation before treating this as user registration.
- `config.Load` and `GetString` panic on errors or missing values. Validate
   configuration during startup and return errors so startup failures are clear.
- Add tests around service behavior and controller HTTP responses. The service
   interface makes it straightforward to inject a small fake repository in
   tests.

So, the architecture is a reasonable foundation, and the explicit DI approach
fits this project. The main improvements are to make the existing wiring
consistent and complete the error, persistence, and test paths before relying
on it as an authentication service.

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