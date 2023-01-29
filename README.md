# Tasks
Program made in Go to handle [Best Mod's](https://bestmods.io/) tasks via sending HTTP requests to a REST API.

## Command Line
The following command line options are available from the help menu.

```
./tasks --cfg <cfgFile> --list --version --help
        --cfg => Path to config file. Default path is /etc/bestmods-tasks/tasks.conf.
        --list => Lists configuration file.
        --version => Prints the current version.
        --help => Prints the help menu.
```

## Configuration File
The configuration file is parsed using `JSON`. Here is a config example.

```JSON
{
    "debug": 2,
    "tasks": [
        {
            "cronstr": "* * * * *",
            "url": "mydomain.com",
            "auth": "Bearer TEST",
            "method": "GET"
        }
    ]
}
```

* **debug** - Debug level from `0` - `2`.
* **tasks** - An array of tasks to schedule via Cron.
    * **cronstr** - The cron string to schedule (read more [here](https://www.netiq.com/documentation/cloud-manager-2-5/ncm-reference/data/bexyssf.html)).
    * **url** - URL to send our HTTP/HTTPS request to.
    * **auth** - What to set the `Authorization` header to within the request.
    * **method** - HTTP method (e.g. `GET`, `POST`, etc.).

## Building
Building is simple via `git` and `go build`.

```bash
# Clone repository.
git clone https://github.com/bestmods/tasks

# Change directory.
cd tasks

# Build to `./tasks`.
go build -o tasks
```

## Credits
* [Christian Deacon](https://github.com/gamemann) - Creator