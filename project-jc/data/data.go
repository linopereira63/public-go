package data

import (
  "crypto/sha512"
  "encoding/base64"
  "time"
  "sync"
)

// Data contains the server data
type Data struct {
  // jobId tracks the assigned job IDs
  jobId int
  // totalHashNanoseconds tracks the total hash time from all hash requests
  totalHashNanoseconds int64
  // inProgressCount tracks the hash computations that have not completed yet
  inProgressCount int
  // hashMap contains the jobId-computedHash values
  hashMap map[int]string

}

var mutex = &sync.Mutex{}

// Create and initialize a Data structure
func NewData() *Data {
  d := &Data{}
  d.hashMap = make(map[int]string)
  return d
}

// Compute and return the nextJobId
func (d *Data)nextJobId() int {
  d.jobId++
  return d.jobId
}

// Get a new jobId, add it to the hashMap, and increment inProgressCount
func (d *Data)AddJob() int {
  id := d.nextJobId()
  mutex.Lock()
  d.hashMap[id] = ""
  d.inProgressCount++
  mutex.Unlock()
  return id
}

// Wait 5 seconds, compute the hash for s, store it in the hashMap, and
// decrement inProgressCount
func (d *Data)ComputeHash(k int, s string) {

  time.Sleep(5 * time.Second)

  h := sha512.New()
  h.Write([]byte(s))
  val := base64.StdEncoding.EncodeToString(h.Sum(nil))

  mutex.Lock()
  d.hashMap[k] = val
  d.inProgressCount--
  mutex.Unlock()
}

// Add hash request time t to totalHashNanoseconds
func (d *Data)AddToHashTime(t int64) {
  d.totalHashNanoseconds += t
}

// Return the computed hash value for jobId k, and its presence indicator
func (d *Data)Get(k int) (string, bool) {
  mutex.Lock()
  value, present := d.hashMap[k]
  mutex.Unlock()
  return value, present
}

// Return the total hash request count and the average time of a hash request
// in millisecs
func (d *Data)GetStats() (int, int64) {
  total := len(d.hashMap)
  var averageMilli int64
  if total > 0 {
    averageNano := d.totalHashNanoseconds / int64(total)
    averageMilli = averageNano / 1000
  }
  return total, averageMilli
}

// Return busy if there are hash computations in progress
func (d *Data)IsBusy() bool {
  return d.inProgressCount > 0
}
