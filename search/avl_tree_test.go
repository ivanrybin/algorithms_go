package search

import (
	"fmt"
	"math"
	"testing"

	"github.com/ivanrybin/algorithms_go/helpers"
)

var increasingIntSequence65536 = helpers.IncreasingInts(uint(65536))
var decreasingIntSequence65536 = helpers.DecreasingInts(uint(65536))

func TestAVLTree_Insert_Find_Overall_Sequential_Increasing_65536(t *testing.T) {
	avlInsertFindTestOverall(t, increasingIntSequence65536)
}

func TestAVLTree_Insert_Find_Overall_Sequential_Decreasing_65536(t *testing.T) {
	avlInsertFindTestOverall(t, decreasingIntSequence65536)
}

func TestAVLTree_Insert_Find_Overall_Random_16(t *testing.T) {
	avlInsertFindRandomTestOverall(t, 500, 32, 64)
}

func TestAVLTree_Insert_Find_Overall_Random_128(t *testing.T) {
	avlInsertFindRandomTestOverall(t, 100, 128, 128)
}

func TestAVLTree_Insert_Find_Overall_Random_512(t *testing.T) {
	avlInsertFindRandomTestOverall(t, 100, 512, 512)
}

func TestAVLTree_Insert_Find_Overall_Random_1024(t *testing.T) {
	avlInsertFindRandomTestOverall(t, 100, 1024, 1024)
}

func TestAVLTree_Insert_Find_Overall_Random_16384(t *testing.T) {
	avlInsertFindRandomTestOverall(t, 10, 16384, 16384)
}

func TestAVLTree_Insert_Find_Overall_Random_65536(t *testing.T) {
	avlInsertFindRandomTestOverall(t, 5, 65536, 65536)
}

func TestAVLTree_Insert_Find_Incremental_Sequential_Increasing_2048(t *testing.T) {
	avlInsertFindTestIncremental(t, increasingIntSequence65536[:2048])
}

func TestAVLTree_Insert_Find_Incremental_Sequential_Decreasing_2048(t *testing.T) {
	avlInsertFindTestIncremental(t, decreasingIntSequence65536[:2048])
}

func TestAVLTree_Insert_Find_Overall_Random_32(t *testing.T) {
	avlInsertFindTestRandomIncremental(t, 500, 32, 64)
}

func TestAVLTree_Insert_Find_Incremental_Random_128(t *testing.T) {
	avlInsertFindTestRandomIncremental(t, 10, 128, 64)
}

func TestAVLTree_Insert_Find_Incremental_Random_512(t *testing.T) {
	avlInsertFindTestRandomIncremental(t, 10, 512, 256)
}

func TestAVLTree_Insert_Find_Incremental_Random_1024(t *testing.T) {
	avlInsertFindTestRandomIncremental(t, 10, 1024, 512)
}

func TestAVLTree_Delete_Find_Overall_Sequential_Increasing_65536(t *testing.T) {
	avlDeleteFindTestOverall(t, increasingIntSequence65536)
}

func TestAVLTree_Delete_Find_Overall_Sequential_Decreasing_65536(t *testing.T) {
	avlDeleteFindTestOverall(t, decreasingIntSequence65536)
}

func TestAVLTree_Delete_Find_Incremental_Sequential_Increasing_2048(t *testing.T) {
	avlDeleteFindTestRepeatedValuesIncremental(t, increasingIntSequence65536[0:2048])
}

func TestAVLTree_Delete_Find_Incremental_Sequential_Decreasing_2048(t *testing.T) {
	avlDeleteFindTestRepeatedValuesIncremental(t, decreasingIntSequence65536[0:2048])
}

func TestAVLTree_Delete_Find_Random_Overall_128(t *testing.T) {
	avlDeleteFindRandomTestOverall(t, 10, 128, 64)
}

func TestAVLTree_Delete_Find_Overall_Random_512(t *testing.T) {
	avlDeleteFindRandomTestOverall(t, 100, 512, 256)
}

func TestAVLTree_Delete_Find_Overall_Random_1024(t *testing.T) {
	avlDeleteFindRandomTestOverall(t, 100, 1024, 256)
}

func TestAVLTree_Delete_Find_Overall_Random_16384(t *testing.T) {
	avlDeleteFindRandomTestOverall(t, 10, 16384, 8192)
}

func TestAVLTree_Delete_Find_Overall_Random_65536(t *testing.T) {
	avlDeleteFindRandomTestOverall(t, 5, 65536, 32768)
}

func TestAVLTree_Delete_Find_Incremental_Random_32(t *testing.T) {
	avlDeleteFindRandomTestIncremental(t, 500, 32, 64)
}

func TestAVLTree_Delete_Find_Incremental_Random_128(t *testing.T) {
	avlDeleteFindRandomTestIncremental(t, 10, 128, 64)
}

func TestAVLTree_Delete_Find_Incremental_Random_512(t *testing.T) {
	avlDeleteFindRandomTestIncremental(t, 10, 512, 256)
}

func TestAVLTree_Delete_Find_Incremental_Random_1024(t *testing.T) {
	avlDeleteFindRandomTestIncremental(t, 10, 1024, 512)
}

func TestAVLTree_Delete_Find_Incremental_Random_Uniq_32(t *testing.T) {
	avlDeleteFindRandomUniqTestIncremental(t, 500, 32)
}

func TestAVLTree_Delete_Find_Incremental_Random_Uniq_128(t *testing.T) {
	avlDeleteFindRandomUniqTestIncremental(t, 10, 128)
}

func TestAVLTree_Delete_Find_Incremental_Random_Uniq_512(t *testing.T) {
	avlDeleteFindRandomUniqTestIncremental(t, 10, 512)
}

func avlInsertFindRandomTestOverall(t *testing.T, iters, size, maxValue int) {
	for i := 1; i <= iters; i++ {
		xs := helpers.RandomIntsExactSize(size, maxValue)
		t.Run(fmt.Sprintf("size=%v", size), func(t *testing.T) {
			avlInsertFindTestOverall(t, xs)
		})
	}
}

func avlInsertFindTestRandomIncremental(t *testing.T, iters, size, maxValue int) {
	for i := 1; i <= iters; i++ {
		xs := helpers.RandomIntsExactSize(size, maxValue)
		t.Run(fmt.Sprintf("size=%v", size), func(t *testing.T) {
			avlInsertFindTestIncremental(t, xs)
		})
	}
}

func avlDeleteFindRandomTestOverall(t *testing.T, iters, size, maxValue int) {
	for i := 1; i <= iters; i++ {
		xs := helpers.RandomIntsExactSize(size, maxValue)
		t.Run(fmt.Sprintf("size=%v", size), func(t *testing.T) {
			avlDeleteFindTestOverall(t, xs)
		})
	}
}

func avlDeleteFindRandomTestIncremental(t *testing.T, iters, size, maxValue int) {
	for i := 1; i <= iters; i++ {
		xs := helpers.RandomIntsExactSize(size, maxValue)
		t.Run(fmt.Sprintf("size=%v", len(xs)), func(t *testing.T) {
			avlDeleteFindTestRepeatedValuesIncremental(t, xs)
		})
	}
}

func avlDeleteFindRandomUniqTestIncremental(t *testing.T, iters, size int) {
	for i := 1; i <= iters; i++ {
		xs := make([]int, 4*size)
		copy(xs, increasingIntSequence65536[:helpers.MinInt(4*size, len(increasingIntSequence65536))])
		helpers.Shuffle(xs)
		xs = xs[:size]
		t.Run(fmt.Sprintf("size=%v", len(xs)), func(t *testing.T) {
			avlDeleteFindTestUniqValuesIncremental(t, xs)
		})
	}
}

func avlInsertFindTestOverall(t *testing.T, xs []int) {
	avl := BuildAVLTree(xs, helpers.LessInt[int])
	presenceTestOverall(t, avl, xs)
	invariantTest(t, avl, len(xs))
}

func avlInsertFindTestIncremental(t *testing.T, xs []int) {
	avl := NewAVLTree(helpers.LessInt[int])
	for i, x := range xs {
		n := avl.Insert(x)
		if n.Value() != x {
			t.Errorf("n.Value()=%v != %v", n.Value(), x)
		}
		_presenceTest(t, avl, x)
		invariantTest(t, avl, i+1)
	}
}

func avlDeleteFindTestOverall(t *testing.T, xs []int) {
	avl := BuildAVLTree(xs, helpers.LessInt[int])
	for _, x := range xs {
		if !avl.Delete(x) {
			t.Errorf("avl.Delete(%v)=false", x)
		}
	}
	absenceTest(t, avl, xs)
	invariantTest(t, avl, 0)
}

func avlDeleteFindTestRepeatedValuesIncremental(t *testing.T, xs []int) {
	avl := BuildAVLTree(xs, helpers.LessInt[int])
	for i, x := range xs {
		if !avl.Delete(x) {
			t.Errorf("avl.Delete(%v)=false", x)
		}
		invariantTest(t, avl, len(xs)-i-1)
	}
}

func avlDeleteFindTestUniqValuesIncremental(t *testing.T, xs []int) {
	avl := BuildAVLTree(xs, helpers.LessInt[int])
	for i, x := range xs {
		if !avl.Delete(x) {
			t.Errorf("avl.Delete(%v)=false", x)
		}
		_absenceTest(t, avl, x)
		invariantTest(t, avl, len(xs)-i-1)
	}
}

func presenceTestOverall(t *testing.T, avl *AVLTree[int], xs []int) {
	for _, x := range xs {
		_presenceTest(t, avl, x)
	}
}

func absenceTest(t *testing.T, avl *AVLTree[int], xs []int) {
	for _, x := range xs {
		_absenceTest(t, avl, x)
	}
}

func _presenceTest[T comparable](t *testing.T, avl *AVLTree[T], x T) {
	if n := avl.Find(x); n == nil {
		t.Errorf("avl[%v].Find(%v)=nil", avl.Size(), x)
	} else if n.v != x {
		t.Errorf("avl[%v].Find(%v)=%v != %v", avl.Size(), x, n.v, x)
	}
}

func _absenceTest[T comparable](t *testing.T, avl *AVLTree[T], x T) {
	if n := avl.Find(x); n != nil {
		t.Errorf("avl[%v].Find(%v)=%v != nil", avl.Size(), n.Value(), x)
	}
}

func invariantTest[T comparable](t *testing.T, avl *AVLTree[T], size int) {
	nodes := avl.InorderTraverse()
	// size check
	if len(nodes) != size {
		t.Errorf("len(nodes)=%v != %v", len(nodes), size)
	}
	if len(nodes) != avl.Size() {
		t.Errorf("len(nodes)=%v != %v", len(nodes), avl.Size())
	}
	// tree height check
	if size != 0 {
		// https://en.wikipedia.org/wiki/AVL_tree#Properties
		maxHeight := int(math.Ceil(math.Log2(float64(size+2))/math.Log2(1.168) - 0.3277))
		if avl.Height() > maxHeight {
			t.Errorf("avl.Height()=%v > expected max height=%v", avl.Height(), maxHeight)
		}
	}
	// nodes height check
	for i, n := range nodes {
		if n.HeightDiff() > 2 {
			t.Errorf("n[%v]: l.h=%v r.h=%v: diff=%v l=%v r=%v", i, n.Left().Height(), n.Right().Height(), n.HeightDiff(), n.Left(), n.Right())
		}
	}
	// ordering check
	for i := 1; i < len(nodes); i++ {
		if !avl.Less()(nodes[i-1].Value(), nodes[i].Value()) && nodes[i-1].Value() != nodes[i].Value() {
			t.Errorf("n[%v].v=%v > n[%v].v=%v", i-1, nodes[i-1].Value(), i, nodes[i].Value())
		}
	}
}
