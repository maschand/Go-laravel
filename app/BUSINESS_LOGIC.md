# ./app

**Folder with business logic only**. This directory doesn't care about _what database driver you're using_ or _which caching solution your choose_ or any third-party things.

- `./app/controllers` folder for functional controllers (used in routes)
- `./app/models` folder for describe business models and methods of your project
- `./app/queries` folder for describe queries for models of your project
- `./app/helpers` folder for describe helpers of your project
- `./app/interfaces` folder for describe interfaces for repository, service, and controller of your project
- `./app/providers` folder for describe providers and binding interface of your project
- `./app/services` folder for describe services for controller of your project
- `./app/repositories` folder for describe repository for service of your project
- `./app/tests` folder for describe unit testing for your project
