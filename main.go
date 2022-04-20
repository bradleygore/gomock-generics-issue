package main

import (
	"fmt"
	"gomock-generics-issue/workers"
)

func main() {
	fmt.Println("This is for demonstration purposes of gomock and generics, no need to run this.")
}

func getPrimitiveWorker() workers.PrimitiveWorker {
	return nil
}

func getCustomWorker() workers.CustomWorker {
	return nil
}

func getMultiWorker() workers.MultiWorker {
	return nil
}
