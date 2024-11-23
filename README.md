
# sous-chef

Recipe organizer, ingredient tracker, meal planner and other cooking related stuff

## Local setup

Initialize Go modules:<br>
`go mod init github.com/TheNewThinkTank/sous-chef`

to add module requirements and sums:<br>
`go mod tidy`

Create dirs:<br>
`mkdir -p cmd/cli cmd/web pkg/recipes pkg/ingredients pkg/mealplanner internal/db api .config migrations docs test`

Install Cobra:<br>

```BASH
go install github.com/spf13/cobra-cli@latest
go get -u github.com/spf13/cobra
vim ~/.zshrc
# Add this:
export PATH=$PATH:~/go/bin
# then reload:
source ~/.zshrc
```

Generate a CLI skeleton:<br>
```BASH
cobra-cli init .

cobra-cli add recipes
cobra-cli add pantry
cobra-cli add mealplan
```

for loading YAML files:<br>
`go get gopkg.in/yaml.v3`


## Tech Stack

- CLI: `Cobra`
- Database: `SQLite` or `PostgreSQL` for storing data persistently.
- API Framework: `Gorilla Mux` or `Echo` for building REST APIs (if applicable).
- Frontend (Optional): Use Go templates or integrate with a frontend framework later.
- Task Scheduling: For notifications, use Go’s `time` package or a library like `robfig/cron`
