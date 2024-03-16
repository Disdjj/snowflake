package snowflake

type IDGenerator interface {
	NextID() int64
}
