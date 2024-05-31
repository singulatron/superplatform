# dapper

Dapper is a funky little configuration management tool that exists to install you runtime dependencies of Singulatron locally on your computer - mostly Docker and related dependencies like WSL on Windows.

There were some design decisions that shaped it:

- No dependencies (single binary)
- Be quick - never redo an already done job
- Stream things are they are happening for user feedback

You can see it in action when you click on the "Install Runtime" button on the Start screen of Singulatron.

## How it works

Dapper runs [`Applications`](fixture/app.json) which are basically a collection of `Features`.

`Features` are a bunch of shell/powershell `Script` for different operating systems.

These `Script`s have two main parts: a `Check` and an `Execute` script. An `Execute` script only gets executed if a `Check` fails. If an `Execute` returns with an exit status other than `0` Dapper exits. If the `Execute` succeeds, it gets re`Check`ed. If that succeeds, all is well. If not, Dapper exits.

An other crucial detail is that after the first `Check` fails for a `Feature`, all dependencies get checked recursively.

## Run

### Example

```sh
go run main.go run --var-username=yourlogin fixture/app.json
```

```sh
PS C:\Users\dobro\mono\dapper> go run .\main.go run --var-username=dobro --var-assetfolder=$env:USERPROFILE .\fixture\app.json
```
