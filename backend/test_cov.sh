set -e
echo "" > coverage.txt
export TESTMODE=true
export GIN_MODE=release

for d in $(go list ./... | grep -v vendor); do
    go test -v -race -coverprofile=profile.out -covermode=atomic $d
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done
