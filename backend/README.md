## CCC Live Server Backend

Backend server/api for CCC Live Server.

### Backend Development setup

App runs as a standard REST api. Everything is saved in memory, but a Firebase DB can be connected to backup and persist data between builds.

#### Using Go (Dev build)
1. **Dependencies**  
   Make sure you have [Golang](https://golang.org/doc/install) installed.
2. **Environment**  
   Set the `CF_USER`, `CF_PASS` environment variables to a valid Codeforces login. The `FIREBASE_SERVICE` variable may optionally be set to a Firebase service account.
   ```
   $ export CF_USER=username
   $ export CF_PASS=password
   ```
   It is recommended not to use a personal account.
3. **Configuration**  
   Inside [backend/configs/config.go](./configs/config.go), set config.ContestId to a Codeforces gym contest that you have access to. The contest id can be found in the url of the contest page. Set config.FirebaseURL to the Firebase database url (if applicable).
4. **Build**  
   Run `make build` inside the **backend** directory to install all dependencies and compile the project.
5. **Run**  
   Run `make run` inside the **backend** directory to run the project. The server will be running at http://localhost:8000. Repeat steps 4 and 5 whenever changes to the backend are made. Steps 4 and 5 can be replaced by a single `make all` command.

#### Using Docker (Prod build)
1. **Dependencies**  
  Make sure you have [Docker](https://www.docker.com/get-started) installed and running on your system.
2. **Environment**  
   Set the `CF_USER`, `CF_PASS` environment variables to a valid Codeforces login. The `FIREBASE_SERVICE` variable may optionally be set to a Firebase service account.
   ```
   $ export CF_USER=username
   $ export CF_PASS=password
   ```
   It is recommended not to use a personal account.
3. **Configuration**  
   Inside [backend/configs/config.go](./configs/config.go), set config.ContestId to a Codeforces gym contest that you have access to. The contest id can be found in the url of the contest page. Set config.FirebaseURL to the Firebase database url (if applicable).
4. **Build + Run**  
   Run `sh build.sh` in the **root** directory of the repository to build and run the Docker image. The server will be running at http://localhost:8000. Use `sh build.sh` again to re-run when changes are made.

### Backend Linting/Testing

#### Linting
1. Make sure you have [Golang](https://golang.org/doc/install) installed.
2. Install [Golangci-lint](https://golangci-lint.run/usage/install/#local-installation)
3. Run `$ sh scripts/lint.sh` inside the **backend** directory to run the linters.

#### Testing
1. Make sure you have [Golang](https://golang.org/doc/install) installed.
2. Make sure the `CF_USER` and `CF_PASS` environment variables are set up correctly.
3. Run `$ sh scripts/test.sh` inside the **backend** directory to run the tests. The [config](backend/configs/config.go) file, and any test data in test files or folders may need to be modified when running this locally.