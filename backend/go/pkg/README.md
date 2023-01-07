# **/pkg**

In `/internal` we store code that we were unable to import into other applications, whereas in `/pkg` we store libraries used in third-party applications. This makes it possible to import them into a different project, and avoid the need to duplicate code from project to project. Generally speaking, it is our custom or shared libraries that are stored here.

You don’t have to use this directory if the project is very small and it makes no practical sense to add a new nesting level.
