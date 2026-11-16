app_name := "changelogr"

main_path := "cmd/"+app_name+"/main.go"

run *args:
    go run {{main_path}} {{args}}

build app_version="dev":
    go build -ldflags="-X main.version={{app_version}}" -o {{app_name}} {{main_path}}

clean:
    rm ./{{app_name}} ./{{app_name}}.yaml ./{{app_name}}.yml ./{{app_name}}.json ./{{app_name}}.toml || true
