package customintslice

type visitor interface {
	InsertInMiddle(c CustomIntSlice) (err error)
}

type CustomIntSlice interface {
	ReversSlice()
	Insert(pos, x int)
	Delete(pos int)
	LenSlice() (l int)
	Accept(v visitor) (err error)
}

type customIntSlice struct {
	intSlice []int
}

func (c *customIntSlice) Accept(v visitor) (err error) {
	return v.InsertInMiddle(c)
}

func (c *customIntSlice) LenSlice() (l int) {
	return len(c.intSlice)
}

func (c *customIntSlice) ReversSlice() {
	for i := c.LenSlice()/2 - 1; i >= 0; i-- {
		opp := c.LenSlice() - 1 - i
		c.intSlice[i], c.intSlice[opp] = c.intSlice[opp], c.intSlice[i]
	}
}

func (c *customIntSlice) Insert(pos, x int) {
	c.intSlice = append(c.intSlice, 0)
	copy(c.intSlice[pos+1:], c.intSlice[pos:])
	c.intSlice[pos] = x
}

func (c *customIntSlice) Delete(pos int) {
	copy(c.intSlice[pos:], c.intSlice[pos+1:])
	c.intSlice[c.LenSlice()-1] = 0
	c.intSlice = c.intSlice[:c.LenSlice()-1]

}

func NewCustomIntSlice(slice []int) CustomIntSlice {
	return &customIntSlice{
		intSlice: slice,
	}
}
