@echo off
{{- if .Valid -}}
echo "go version go{{.GOVersion}} {{.GOOS}}/{{.GOArch}}"
{{-else-}}
echo "invalid go version output"
{{-endif-}}
