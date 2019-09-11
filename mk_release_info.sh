#!/bin/sh
GIT_SHA1=$( (git show-ref --head --hash=8 2>/dev/null || echo 00000000) | head -n1)
GIT_DIRTY=$(git diff --no-ext-diff 2>/dev/null | wc -l)
BUILD_ID=$(uname -n)"-"$(date +%s)
if ! test -f "release.info"; then
  touch release.info
fi
(cat <release.info | grep SHA1 | grep "$GIT_SHA1") && \
(cat <release.info | grep DIRTY | grep "$GIT_DIRTY") && exit 0 # Already up-to-date
echo "REDIS_GIT_SHA1=\"$GIT_SHA1\"," >"release.info"
echo "REDIS_GIT_DIRTY=\"$GIT_DIRTY\"" >>"release.info"
echo "REDIS_BUILD_ID=\"$BUILD_ID\"" >>"release.info"
