# Readme Topics to MarkDown

Simple program to take all the topics in the Readme DB and convert them to Markdown files.

### Flags

`--dir` : Specify what directory to export them too. Defaults to `.`

```bash
go run main.go --dir ~/Desktop
```

Will create a folder called `readme_topics` and place all the files in there as `.md` files.
