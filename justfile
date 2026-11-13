app_name := "changelogr"

main_path := "cmd/"+app_name+"/main.go"

run *args:
    go run {{main_path}} {{args}}

build app_version="dev":
    go build -ldflags="-X main.version={{app_version}}" -o {{app_name}} {{main_path}}
