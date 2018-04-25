# sourceroostersvc



find ~/ -name "*.git" -type d | xargs -i{} git --git-dir={} ls-remote --get-url
find ~/ -name "*.git" -type d | xargs -i{} $(echo {}' => ' && git --git-dir={} ls-remote --get-url)