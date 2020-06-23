package stack

type Item interface {
}

// ItemStack：保存栈的item
type ItemStack struct {
    items []Item
}

// 新建一个ItemStack
func (s *ItemStack) New() *ItemStack {
    s.items = []Item{}
    return s
}

// 添加item到栈顶端
func (s *ItemStack) Push(t Item) {
    s.items = append(s.items, t)
}

// 从栈顶端移除一个item
func (s *ItemStack) Pop() *Item {
    item := s.items[len(s.items)-1] // 后进先出
    s.items = s.items[0:len(s.items)-1]
    return &item

}

