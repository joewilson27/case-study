- Case 1 -
Download file, then run with your favorite terminal / command prompt: node demo.js

- Case 2 -
Download FE (Frontend), then run: npm install && npm run serve

Download BE (Backend)
then run: 
go get -u github.com/gofiber/fiber/v2
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go install github.com/cosmtrek/air@latest

finally to start be, just run command: air

*make sure to create your DB first, before you run this BE (Backend). Don't forget to adjust DB Config in env settings*
