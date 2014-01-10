package context

import(
  "github.com/garyburd/redigo/redis"
  "github.com/boj/redistore"
  "launchpad.net/goamz/aws"
  "launchpad.net/goamz/s3"
)

/* 
 * Defaults 
 */
const VERSION = "0.0.1"
var Env string
var Port string
var Index string

/* 
 * S3 
 */
var S3Auth aws.Auth = aws.Auth{
  AccessKey: "XXXX",
  SecretKey: "XXXX",
}
var FileStorage *s3.Bucket
const BUCKET = "XXXXXXX"

/*
 * Redis DB Connection & Session Store
 */
var Conn redis.Conn
var Store *redistore.RediStore

func init() {
  Conn, _ = redis.Dial("tcp", ":6379")
  Store = redistore.NewRediStore(10, "tcp", ":6379", "", []byte("SECRET"))

  FileStorage = s3.New(S3Auth, aws.USEast).Bucket(BUCKET)
}
