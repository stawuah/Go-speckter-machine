package main

import (
    "fmt"
    "log"
    "sync"
    "time"
)

type URLLocker struct {
    mu       sync.Mutex
    locks    map[string]*sync.Mutex
    refCount map[string]int
}

// NewURLLocker creates a new URLLocker instance.
func NewURLLocker() *URLLocker {
    return &URLLocker{
        locks:    make(map[string]*sync.Mutex),
        refCount: make(map[string]int),
    }
}

// Lock locks the specified URL.
func (ul *URLLocker) Lock(url string) {
    ul.mu.Lock()
    if _, exists := ul.locks[url]; !exists {
        ul.locks[url] = &sync.Mutex{}
        ul.refCount[url] = 0
    }
    ul.refCount[url]++
    mu := ul.locks[url]
    ul.mu.Unlock()

    mu.Lock()
}

// Unlock unlocks the specified URL.
func (ul *URLLocker) Unlock(url string) {
    ul.mu.Lock()
    if mu, exists := ul.locks[url]; exists {
        mu.Unlock()
        ul.refCount[url]--
        if ul.refCount[url] == 0 {
            delete(ul.locks, url)
            delete(ul.refCount, url)
        }
    }
    ul.mu.Unlock()
}

// Cleanup removes unused URL locks to prevent memory leaks.
func (ul *URLLocker) Cleanup() {
    ul.mu.Lock()
    defer ul.mu.Unlock()
    for url, count := range ul.refCount {
        if count == 0 {
            delete(ul.locks, url)
            delete(ul.refCount, url)
        }
    }
}

func main() {
    urlLocker := NewURLLocker()

    var wg sync.WaitGroup
    urls := []string{"http://example.com", "http://example.com", "http://another.com"}

    for _, url := range urls {
        wg.Add(1)
        go func(url string) {
            defer wg.Done()
            urlLocker.Lock(url)
            log.Printf("Locked URL: %s\n", url)
            // Simulate some work with the URL
            time.Sleep(1 * time.Second)
            urlLocker.Unlock(url)
            log.Printf("Unlocked URL: %s\n", url)
        }(url)
    }

    wg.Wait()
    urlLocker.Cleanup() // Clean up any unused locks
}
