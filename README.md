# LA Tool

## Logic App tool for generating diffs between source code and/or deployed instance configerations and validation of workflows.

_La Tool_ was born out of frustrations when working with Logic Apps (Standard) in a highly restricted client environment mainly stemming from conflicts between manual changes made via the Azure Portal and code generated by the VS Code extension and managed via source control.

Numerous issues would slow development, such as:

- individuals contributing conflicting 'code' from different source and not adhering to proper source control practices (arguably a skill/process issue)
- inherrent low-code vs source control incompatibilites.
- differences in formatting of the underlying json code between VS Code and Portal generated code.

For the project I was on at the time, I just knocked up a quick Python tool for the CI/CD pipeline to perform a basic comparison between the source code and the deployed instance, as well as a few basic validations of the Workflow files.

I included a rudimentary estension system for adding further validation checks. The code was messy and very clearly a first draft. I wouldn't recommend anyone else use it...

The purpose of this project is to productionise my first draft and provide a more robust and extensible tool. This is by no means any sort of warranty or promise to maintain the project or provide any level of support. Use this tool at your own risk and only with proper understanding. My main reasons to creat this are:

- I want to get better at Go
- I want to have a better tool ready for any future client work relating to Logic Apps in a similar setting.
- I may do further work with the existing client and providing this tool to them would be a quick win for improving the existing pipeline, which I have no doubt nobody has touched since me.
