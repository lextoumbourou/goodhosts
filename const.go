// +build linux darwin

package goodhosts

const hostsFilePath = "/etc/hosts"
const eol = "\n"

/*
Whether or not to append multiple hosts pointing to the same IP on the same line
*/
const appendToLine = true
