package cache

import (
	"log"
	"math/rand"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Int32 int32

func (i Int32) Less(than Item) bool {
	l, ok := than.(Int32)
	if !ok {
		return false
	}
	return i < l
}

func TestNoSplit(t *testing.T) {
	btree := NewBTree()

	btree.Insert(Int32(1))
	btree.Insert(Int32(2))

	found, _ := btree.Find(Int32(1))
	assert.True(t, found)

	found, _ = btree.Find(Int32(2))
	assert.True(t, found)

	found, _ = btree.Find(Int32(3))
	assert.False(t, found)
}

func TestSplitParent(t *testing.T) {
	btree := NewBTree()

	btree.Insert(Int32(1))
	btree.Insert(Int32(2))
	btree.Insert(Int32(3))

	found, _ := btree.Find(Int32(1))
	assert.True(t, found)

	found, _ = btree.Find(Int32(2))
	assert.True(t, found)

	found, _ = btree.Find(Int32(3))
	assert.True(t, found)

	// test balance
	assert.Equal(t, btree.Top.Items[0], Int32(2))
	assert.Equal(t, btree.Top.Children[0].Items[0], Int32(1))
	assert.Equal(t, btree.Top.Children[1].Items[0], Int32(3))
}

func TestSplitChild(t *testing.T) {
	btree := NewBTree()
	btree.Insert(Int32(1))
	btree.Insert(Int32(2))
	btree.Insert(Int32(3))
	btree.Insert(Int32(4))
	btree.Insert(Int32(5))

	found, _ := btree.Find(Int32(1))
	assert.True(t, found)

	found, _ = btree.Find(Int32(2))
	assert.True(t, found)

	found, _ = btree.Find(Int32(3))
	assert.True(t, found)

	found, _ = btree.Find(Int32(4))
	assert.True(t, found)

	found, _ = btree.Find(Int32(5))
	assert.True(t, found)

	// test balance
	assert.Equal(t, btree.Top.Items[0], Int32(2))
	assert.Equal(t, btree.Top.Items[1], Int32(4))
	assert.Equal(t, btree.Top.Children[0].Items[0], Int32(1))
	assert.Equal(t, btree.Top.Children[1].Items[0], Int32(3))
	assert.Equal(t, btree.Top.Children[2].Items[0], Int32(5))
}

func TestBlanced(t *testing.T) {
	btree := NewBTree()
	btree.Insert(Int32(1))
	btree.Insert(Int32(2))
	btree.Insert(Int32(3))
	btree.Insert(Int32(4))
	btree.Insert(Int32(5))
	btree.Insert(Int32(6))
	btree.Insert(Int32(7))

	found, _ := btree.Find(Int32(1))
	assert.True(t, found)

	found, _ = btree.Find(Int32(2))
	assert.True(t, found)

	found, _ = btree.Find(Int32(3))
	assert.True(t, found)

	found, _ = btree.Find(Int32(4))
	assert.True(t, found)

	found, _ = btree.Find(Int32(5))
	assert.True(t, found)

	found, _ = btree.Find(Int32(6))
	assert.True(t, found)

	found, _ = btree.Find(Int32(7))
	assert.True(t, found)

	// test balance
	assert.Equal(t, btree.Top.Items[0], Int32(4))
	assert.Equal(t, btree.Top.Children[0].Items[0], Int32(2))
	assert.Equal(t, btree.Top.Children[1].Items[0], Int32(6))
	assert.Equal(t, btree.Top.Children[0].Children[0].Items[0], Int32(1))
	assert.Equal(t, btree.Top.Children[0].Children[1].Items[0], Int32(3))
	assert.Equal(t, btree.Top.Children[1].Children[0].Items[0], Int32(5))
	assert.Equal(t, btree.Top.Children[1].Children[1].Items[0], Int32(7))
}

func TestBlancedReversed(t *testing.T) {
	btree := NewBTree()
	btree.Insert(Int32(7))
	btree.Insert(Int32(6))
	btree.Insert(Int32(5))
	btree.Insert(Int32(4))
	btree.Insert(Int32(3))
	btree.Insert(Int32(2))
	btree.Insert(Int32(1))

	found, _ := btree.Find(Int32(1))
	assert.True(t, found)

	found, _ = btree.Find(Int32(2))
	assert.True(t, found)

	found, _ = btree.Find(Int32(3))
	assert.True(t, found)

	found, _ = btree.Find(Int32(4))
	assert.True(t, found)

	found, _ = btree.Find(Int32(5))
	assert.True(t, found)

	found, _ = btree.Find(Int32(6))
	assert.True(t, found)

	found, _ = btree.Find(Int32(7))
	assert.True(t, found)

	// test balance
	assert.Equal(t, btree.Top.Items[0], Int32(4))
	assert.Equal(t, btree.Top.Children[0].Items[0], Int32(2))
	assert.Equal(t, btree.Top.Children[1].Items[0], Int32(6))
	assert.Equal(t, btree.Top.Children[0].Children[0].Items[0], Int32(1))
	assert.Equal(t, btree.Top.Children[0].Children[1].Items[0], Int32(3))
	assert.Equal(t, btree.Top.Children[1].Children[0].Items[0], Int32(5))
	assert.Equal(t, btree.Top.Children[1].Children[1].Items[0], Int32(7))
}

func TestGet(t *testing.T) {
	btree := NewBTree()

	got := btree.Get(Int32(1))
	assert.Equal(t, got, nil)

	btree.Insert(Int32(1))
	btree.Insert(Int32(2))
	btree.Insert(Int32(3))
	btree.Insert(Int32(4))
	btree.Insert(Int32(5))
	btree.Insert(Int32(6))
	btree.Insert(Int32(7))

	item := btree.Get(Int32(1))
	i := item.(Int32)
	assert.Equal(t, i, Int32(1))

	item = btree.Get(Int32(7))
	i = item.(Int32)
	assert.Equal(t, i, Int32(7))
}

func TestRandom(t *testing.T) {
	btree := NewBTree()

	for i := 0; i < 10000; i++ {
		v := rand.Intn(1000)
		btree.Insert(Int32(v))
	}

	assert.Equal(t, btree.Len(), 10000)
}

func TestEmpty(t *testing.T) {
	btree := NewBTree()
	found, _ := btree.Find(Int32(1))
	assert.False(t, found)
}

func TestSerialize(t *testing.T) {
	btree := NewBTree()
	btree.Insert(Int32Item(1))
	b, err := SerializeBTree(btree)
	if err != nil {
		log.Fatal(err)
	}

	newTree, err := DeserializeBTree(b)
	if err != nil {
		log.Fatal(err)
	}

	item := newTree.Top.Items[0].(Int32Item)
	assert.Equal(t, item, Int32Item(1))
}

func TestIntItem_Less(t *testing.T) {
	type args struct {
		than Item
	}
	tests := []struct {
		name string
		i    Item
		args args
		want bool
	}{
		{
			name: "If the argument is greater than the value of the object",
			i:    Int32Item(1),
			args: args{
				than: Int32Item(2),
			},
			want: true,
		},
		{
			name: "If the object value is greater than the argument value",
			i:    Int32Item(2),
			args: args{
				than: Int32Item(1),
			},
			want: false,
		},
		{
			name: "Invalid argument",
			i:    Int32Item(1),
			args: args{
				than: Int64Item(1),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Less(tt.args.than); got != tt.want {
				t.Errorf("IntItem.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64Item_Less(t *testing.T) {
	type args struct {
		than Item
	}
	tests := []struct {
		name string
		i    Int64Item
		args args
		want bool
	}{
		{
			name: "If the argument is greater than the value of the object",
			i:    Int64Item(1),
			args: args{
				than: Int64Item(2),
			},
			want: true,
		},
		{
			name: "If the object value is greater than the argument value",
			i:    Int64Item(2),
			args: args{
				than: Int64Item(1),
			},
			want: false,
		},
		{
			name: "Invalid argument",
			i:    Int64Item(1),
			args: args{
				than: Int32Item(1),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Less(tt.args.than); got != tt.want {
				t.Errorf("Int64Item.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_get(t *testing.T) {
	type fields struct {
		Items    items
		Children []*node
	}
	type args struct {
		key Item
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Item
	}{
		{
			name: "No node",
			fields: fields{
				Items:    items{},
				Children: nil,
			},
			args: args{
				key: Int32Item(1),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &node{
				Items:    tt.fields.Items,
				Children: tt.fields.Children,
			}
			if got := n.get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("node.get() = %v, want %v", got, tt.want)
			}
		})
	}
}
