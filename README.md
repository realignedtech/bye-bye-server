# bye-bye-server

Community Repository for the [Bye-Bye-Server.com](https://bye-bye-server.com/) Website

## PRs welcome

If you want to contribute a product - edit the products.json in the public directory

Make sure your product supports on premises deployments, as we want to provide alternatives for Atlassian Server customers.

There is a husky pre-commit hook that will automatically sort the product list in a alphabetical order. Committing with `--no-verify` and subsequently pushing will result in the github action validation failing if the products are not in alphabetical order. The pre-commit hook also runs prettier, maintaining a styling consistency.
