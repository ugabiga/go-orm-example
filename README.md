# go(golang) orm example

This repository will compare ORMs of go languages.

Check example folder for the code. 

## ORM List
- [x] sqlboiler
- [ ] ent
- [ ] gorm
- [ ] sqlc

## Instruction for sqlboiler
1. Postgresql docker up
   ```
   make up
   ```
2. Run migration
   ```
   go run main.go bo up
   ```
3. Run example
   ```
   go run main.go bo run
   ```