# Contribution guidelines

## Contributing functions

Imagine you have a new function idea, let's call it Init. These are the steps you'd need to follow
to get a PR merged.

* Create Init.go under `/functions`
* Write the code for your function (the fun part!)
* Add your function in `functions/main.go` and specify for which types it applies
	* ForNumbers, ForStrings, ForStructs, ... 
* Try to generate your code for the built-in types by running the python build script
	* ./build 
* Write the respective unit tests in `types/*_test.go`

That's it! If everything went well you now have a PR ready for including your function in Hasgo. 

If you need some inspiration, check out the issues and pick one to start contributing. Feel free to
suggest a new function as well, but it's a good idea to first create the issue before you start
coding so we can help you get started in a good way!

If you get stuck, just create the PR of what you have and we'll help get you through the steps
:smiley:

Happy Coding! :raised_hands:
