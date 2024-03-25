package services

import (
	"container/heap"
	"github.com/samber/lo"
	"github.com/sumitsj/url-shortener/models"
	"github.com/sumitsj/url-shortener/repositories"
	"github.com/sumitsj/url-shortener/utils"
	"log"
)

type HeapService interface {
	Get() *KVHeap
}

type heapService struct {
	h          *KVHeap
	repository repositories.UrlMappingRepository
}

func (receiver *heapService) Get() *KVHeap {

	if receiver.h == nil {
		receiver.h = &KVHeap{}
		urlMappings, err := receiver.repository.GetAll()
		if err != nil {
			panic(err.Error())
		}

		domains := extractDomains(urlMappings)

		countBy := lo.CountValues[string](domains)

		log.Println("countBy", countBy)

		receiver.createHeap(countBy)
	}

	return receiver.h
}

func (receiver *heapService) createHeap(m map[string]int) {
	heap.Init(receiver.h)
	for k, v := range m {
		heap.Push(receiver.h, KV{k, v})
	}
}

func InitialiseHeapService(repository repositories.UrlMappingRepository) HeapService {
	return &heapService{
		repository: repository,
	}
}

type KV struct {
	Key   string
	Value int
}

func extractDomains(list []models.URLMapping) []string {
	domains := []string{}

	for _, urlMapping := range list {
		domains = append(domains, utils.GetDomainName(urlMapping.OriginalUrl))
	}

	return domains
}

type KVHeap []KV

func (h KVHeap) Less(i, j int) bool { return h[i].Value > h[j].Value }
func (h KVHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h KVHeap) Len() int           { return len(h) }

func (h *KVHeap) Push(x interface{}) {
	*h = append(*h, x.(KV))
}

func (h *KVHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
