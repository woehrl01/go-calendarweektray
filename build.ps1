$now = Get-Date -UFormat "%Y-%m-%d_%T"
$sha1 = (git rev-parse HEAD).Trim().Substring(0, 8)
$version = (git describe --tags $(git rev-list --tags --max-count=1)).Trim()

go build -ldflags "-X main.sha1ver=$sha1 -X main.buildTime=$now  -X main.semVer=$version -H=windowsgui" -o Kalenderwoche.exe
