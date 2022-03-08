# spirits

`spirits` is a turn-based battle framework.

To run:
```sh
$ docker build -t spirits .
$ docker run -p 12345:12345 spirits
$ # open browser to localhost:12345
```

## Muses

* Pokémon
* DnD
* Final Fantasy
* Fire Emblem

## Tech Stack

* Infrastructure: [Heroku](https://oh-great-spirits.herokuapp.com/)
* Backend: [Golang](api)
* Frontend: [React](web)
* Work Tracking: [Pivotal Tracker](https://www.pivotaltracker.com/n/projects/2556075)
* E2E Testing: [Node/Pupeteer](test)

## Directory Structure

| Path  | Description |
| ------------- | ------------- |
| `api`  | backend code  |
| `test`  | e2e tests  |
| `web` | frontend code |
