package goodhosts

const hostsFilePath = "${SystemRoot}/System32/drivers/etc/hosts"
const eol = "\r\n"

/*
Whether or not to append multiple hosts pointing to the same IP on the same line
*/
const appendToLine = false // On windows, there is a 9-alias limit per line, so we disable appending behavior
