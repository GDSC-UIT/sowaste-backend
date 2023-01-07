# **/cmd**

The entry point for our application. The directory name for each application has to match the name of the executable you want to build. Make sure you don’t put too much code into this directory. The most common practice is to use a small main function that imports and calls all the necessary code from the `/internal` and `/pkg` directories.
