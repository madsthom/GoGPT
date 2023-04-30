package main

import (
	"GoGPT/agents"
	"GoGPT/promts"
)

func main() {
	prompt := promts.PromptTemplate{Template: "You must act as penetration tester who is allowed to test the security of the system. Your target is {target} \nTest different commands, and try to find vulnerabilities in the system.\nTools such as sqlmap often requires further interactions by writing \"y\", \"n\" in the terminal. \nPlease make sure this is not needed, by adding the appropriate accept all flag.\n\"ping\" must be used with the following flags \"-W 10 -c 3\".\n\"sqlmap\" must have \"--batch\" and \"-v 0\" as part of its flags.\n\"msfconsole\" must always .; exit' after command execution\n\nYour final goal is to exploit the vulnerabilities on the website provided below.\n\nWhen executing terminal commands, the following commands are valid: curl, ping, nmap, sqlmap, nikto, msfconsole.\nNever try to download anything other than that.\nWhenever you are running terminal commands, you must try to decrease the verbosity to the minimum.\nYou can often use --help to see the possible flags.\n\nWhenever you find a vulnerability you need to write it down how you would exploit it, using the terminal action with 'echo \"{input}\" > writer.txt' as the action input.\nAfter writing it down, move on to a new vulnerability and start exploiting."}
	agents.RunAgent(prompt, []string{"trifork.com"})
}
