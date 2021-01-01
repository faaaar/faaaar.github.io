#!/bin/sh

set -e

printf "\033[0;32mDeploying updates to GitHub...\033[0m\n"

hugo

echo "gccio.com" > docs/CNAME

git add .

git commit --amend --date "$(date)" --author="LIYUNFAN<leeyunfans@gmail.com>"

git push -f origin RELEASE
