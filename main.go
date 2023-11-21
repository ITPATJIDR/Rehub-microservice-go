package main

import (
  "Rehub_Microservice/routers"
)

func main() {
  r := routers.SetupRouter()
  r.Run()
}
