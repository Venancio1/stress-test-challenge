package adapter

import (
	"fmt"
	"stress-test/client"
	"sync"
	"time"
)

func PrepareRequest(url string, requests int, concurrent int) {
	tarefas := make(chan int, requests)
	resultados := make(chan string, requests) // canal para coletar os status
	var wg sync.WaitGroup

	// Workers
	for w := 1; w <= concurrent; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for range tarefas {
				status, _ := client.RequestDefault(url)
				// envia o resultado para o canal
				resultados <- status
			}
		}(w)
	}

	start := time.Now()

	// Alimenta a fila
	for i := 1; i <= requests; i++ {
		tarefas <- i
	}
	close(tarefas)

	// Goroutine para fechar o canal de resultados quando todos os workers terminarem
	go func() {
		wg.Wait()
		close(resultados)
	}()

	// Aqui você consome os resultados
	var list []string
	for r := range resultados {
		list = append(list, r)

		// exemplo de armazenamento em slice
	}

	end := time.Now()
	duration := end.Sub(start)
	fmt.Printf("Tempo total de execução: %v\n", duration)

	analisarStatus(list)
}

func analisarStatus(statusList []string) {
	// Map para armazenar listas de acordo com o status
	statusMap := make(map[string][]string)

	// Itera sobre a lista principal e separa por status
	for _, status := range statusList {
		statusMap[status] = append(statusMap[status], status)
	}

	// Printar quantidade de cada status
	for codigo, lista := range statusMap {
		fmt.Printf("Status %s: %d ocorrências\n", codigo, len(lista))
	}
	fmt.Printf("Total de requisições realizadas: %d\n", len(statusList))
}
