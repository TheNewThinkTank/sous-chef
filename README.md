
# sous-chef

Recipe organizer, ingredient tracker, meal planner and other cooking related stuff

## Local setup

Initialize Go modules:<br>
`go mod init github.com/yourusername/sous-chef`

Create dirs:<br>
`mkdir -p cmd/cli cmd/web pkg/recipes pkg/ingredients pkg/mealplanner internal/db api .config migrations docs test`


## Tech Stack

- Database: `SQLite` or `PostgreSQL` for storing data persistently.
- API Framework: `Gorilla Mux` or `Echo` for building REST APIs (if applicable).
- Frontend (Optional): Use Go templates or integrate with a frontend framework later.
- Task Scheduling: For notifications, use Go’s `time` package or a library like `robfig/cron`
