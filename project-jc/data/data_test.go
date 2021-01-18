package data

import (
  "testing"
  "time"
)

func TestData(t *testing.T) {

  d := NewData()

  // Test initial stats
  total, average := d.GetStats()
  if total != 0 && average != 0 {
    t.Errorf("Expected total=0 and average=0, got total=%d, average=%d", total, average)
  }

  // Test AddJob
  jobId := d.AddJob()
  if jobId != 1 {
    t.Errorf("Expected jobId=1, got %d", jobId)
  }

  // Test ComputeHash
  go d.ComputeHash(jobId, "angryMonkey")

  // Test AddToHashTime
  start := time.Now()
  time.Sleep(10 * time.Millisecond)
  elapsed := time.Since(start)
  d.AddToHashTime(elapsed.Nanoseconds())

  // Test isBusy while hash is getting computed
  if !d.IsBusy() {
    t.Errorf("Expected busy to be true, got %v", d.IsBusy())
  }

  // Test Get while hash is getting computed
  val, present := d.Get(jobId)
  if val != "" || !present {
    t.Errorf("Expected hash for jobId %d to be \"\" because its not ready, got %v, present=%v",
      jobId, val, present)
  }

  // Test stats after adding 1st hash
  total, average = d.GetStats()
  if total != 1 && average != 10 {
    t.Errorf("Expected total=1 and average=10, got total=%d, average=%d", total, average)
  }

  // Wait for hash to complete, it has a 5sec delay, so waiting longer
  time.Sleep(6 * time.Second)

  // Test isBusy after hash computation is complete
  if d.IsBusy() {
    t.Errorf("Expected busy to be false, got %v", d.IsBusy())
  }

  // Test Get after hash computation is complete
  expectedHash := "ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q=="
  val, present = d.Get(jobId)
  if val != expectedHash && present {
    t.Errorf("Expected hash for jobId %d to be %v, got %v, present=%v",
      jobId, expectedHash, val, present)
  }

  // Test adding 999 more hashes
  for i := 0; i<999; i++ {
    jobId = d.AddJob()
    go d.ComputeHash(jobId, "angryMonkey")
    d.Get(jobId)
  }

  // Wait for hash to complete, it has a 5sec delay, so waiting longer
  time.Sleep(6 * time.Second)

  // Test isBusy after hash computation is complete
  if d.IsBusy() {
    t.Errorf("Expected busy to be false, got %v", d.IsBusy())
  }

  // Test Get for jobId of 1000 after hash computation is complete
  expectedHash = "ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q=="
  val, present = d.Get(jobId)
  if val != expectedHash && present {
    t.Errorf("Expected hash for jobId %d to be %v, got %v, present=%v",
      jobId, expectedHash, val, present)
  }

  // Test Get for a non-existent jobId
  jobId = 999999
  val, present = d.Get(jobId)
  if present {
    t.Errorf("Expected hash for non-present jobId %d to not be present, present=%v",
      jobId, present)
  }

}
