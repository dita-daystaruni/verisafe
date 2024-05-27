#   `internal/`

Contains the core application code, separated by domain (auth, handlers, models).

    internal/auth/
        auth.go: Contains the logic for authentication (e.g., token generation, password hashing).
        middleware.go: Contains the authentication middleware.

    internal/handlers/
        handlers.go: Contains HTTP handler functions (e.g., RegisterHandler, LoginHandler).

    internal/models/
        user.go: Defines the User model and any related database operations.
