package main

import (
	"URL_CUTTER/service"
	"URL_CUTTER/storage"
	"URL_CUTTER/test"
	"fmt"
)

func main() {
	//////////////////////////////////////////////////////////
	///                      ТЕСТЫ                         ///
	//////////////////////////////////////////////////////////
	go test.StartTesting()
	//////////////////////////////////////////////////////////
	///                    ХРАНИЛИЩЕ                       ///
	//////////////////////////////////////////////////////////
	st, err := storage.NewDataStorage()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//////////////////////////////////////////////////////////
	///                      СЕРВЕР                        ///
	//////////////////////////////////////////////////////////
	service.CreateWebApp("localhost", "8080", st)
}
