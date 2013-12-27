package context

import(
  "github.com/garyburd/redigo/redis"
  "github.com/boj/redistore"
)

/* Defaults */
const VERSION = "0.0.1"
var Env string
var Port string

/*
 * Redis DB Connection & Session Store
 */
var Conn redis.Conn
var Store *redistore.RediStore

func init() {
  Conn, _ = redis.Dial("tcp", ":6379")
  Store = redistore.NewRediStore(10, "tcp", ":6379", "", []byte("SECRET"))
}
