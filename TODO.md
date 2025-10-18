# TODO: Convert Node.js Backend to Go

## 1. Create new Go files
- [x] Create `main.go`: Main server logic using Gin framework
- [x] Create `db.go`: MySQL connection using database/sql and mysql driver
- [x] Create `routes.go`: API endpoints (e.g., /test-db)
- [x] Create `auth.go`: OIDC client setup using go-oidc library

## 2. Create Go module
- [x] Create `go.mod`: Define Go modules and dependencies (gin, go-oidc, gorilla/sessions, etc.)

## 3. Update deployment files
- [x] Update `Dockerfile`: Change to Go base image and build/run Go binary
- [x] Update `docker-compose.yml`: Update app service for Go runtime

## 4. Update project files
- [x] Update `.gitignore`: Add Go-specific ignores (e.g., .exe, vendor/)
- [x] Update `README.md`: Change tech stack to Go backend

## 5. Remove Node.js files
- [ ] Remove `index.js`
- [ ] Remove `db.js`
- [ ] Remove `backend/routes.js`
- [ ] Remove `auth/faydaAuth.js`
- [ ] Remove `package.json`
- [ ] Remove `package-lock.json`

## 6. Followup steps
- [ ] Install Go dependencies via `go mod tidy`
- [ ] Test the Go server locally
- [ ] Verify frontend integration with new backend
