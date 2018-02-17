package matrix

var (
	f32Pool = newf32Pool()
)

type f32Bucket struct {
	vals []float32
}

type matf32Pool struct {
	pool chan *f32Bucket
}

func newf32Pool() *matf32Pool {
	return &matf32Pool{
		pool: make(chan *f32Bucket, 10),
	}
}

func (p *matf32Pool) get() *f32Bucket {
	var c *f32Bucket
	select {
	case c = <-p.pool:
	default:
		c = &f32Bucket{
			vals: make([]float32, 0),
		}
	}
	return c
}

func (p *matf32Pool) put(m *f32Bucket) {
	select {
	case p.pool <- m:
	default:
		return
	}
}
