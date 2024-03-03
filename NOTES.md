# Notes

> Some ideas and notes what I should implement.

## Store

- Store data locally in just a JSON? It sounds the least troublesome, and later I could upgrade it or whatever
- In the JSON file I could have the following structure as I have already thought of this before:

```json
{
  "globals": [],
  "locals": {}
}
```

- For the local ones we also need to store a path, hence the object instead of the array
- Then for a TODO I could have a structure like this:

```json
{
  "content": "Something",
  "resolved": false
}
```

And then could have a list something like this:

```json
{
  "globals": [
    {
      "content": "Something",
      "resolved": true
    },
    {
      "content": "Something else",
      "resolved": false
    }
  ],
  "locals": {
    "/home/foo/something": [
      {
        "content": "Something",
        "resolved": true
      },
      {
        "content": "Something else",
        "resolved": false
      }
    ],
    "/home/foo/something_else": [
      {
        "content": "Something else",
        "resolved": false
      }
    ],
    "/home/foo/idk": [
      {
        "content": "Something",
        "resolved": true
      }
    ]
  }
}
```

## Commands and subcommands

With commands we should be able to:

- Add
- Remove
- Resolve
- Edit? (Could be cool to have a `git commit --amend` kind of solution)
- Get status
  - Get only unresolved ones
  - Get only resolved ones
  - Get all

We should be able to do these for local, and global scopes, or even do it based on path.

## TODO

- [x] Should think through how to handle global TODOs and start to implement subcommand functionality for it.
- [x] Think about a TODO's ID, and also update the examples above.
- [x] Extend add & edit functionalities, to be able to execute the action without an editor
  - E.g.: Have something like this: `td -a -c "NEW TODO"` and `td -e -c "EDITED VALUE"`
- [ ] Add unit tests
- [ ] Make some CI/CD processes, e.g.: Run tests on push, and see a bit more about go-release-action

## Future ideas

- Export TODO status into a file or something like that
