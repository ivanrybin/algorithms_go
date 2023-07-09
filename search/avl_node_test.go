package search

import (
	"reflect"
	"testing"
)

func TestAVLNode_Predicats(t *testing.T) {
	type predicats struct {
		IsLeaf  bool
		IsLeft  bool
		IsRight bool
		IsFull  bool
	}
	for _, tt := range []struct {
		name string
		node *AVLNode[int]
		want predicats
	}{
		{
			name: "full",
			node: &AVLNode[int]{
				l: &AVLNode[int]{},
				r: &AVLNode[int]{},
			},
			want: predicats{
				IsLeaf:  false,
				IsLeft:  true,
				IsRight: true,
				IsFull:  true,
			},
		},
		{
			name: "left",
			node: &AVLNode[int]{
				l: &AVLNode[int]{},
			},
			want: predicats{
				IsLeaf:  false,
				IsLeft:  true,
				IsRight: false,
				IsFull:  false,
			},
		},
		{
			name: "right",
			node: &AVLNode[int]{
				r: &AVLNode[int]{},
			},
			want: predicats{
				IsLeaf:  false,
				IsLeft:  false,
				IsRight: true,
				IsFull:  false,
			},
		},
		{
			name: "leaf",
			node: &AVLNode[int]{},
			want: predicats{
				IsLeaf:  true,
				IsLeft:  false,
				IsRight: false,
				IsFull:  false,
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := predicats{
				IsLeaf:  tt.node.IsLeaf(),
				IsLeft:  tt.node.IsLeft(),
				IsRight: tt.node.IsRight(),
				IsFull:  tt.node.IsFull(),
			}
			if got != tt.want {
				t.Errorf("got = %+v != want %+v", got, tt.want)
			}
		})
	}
}

func TestAVLNode_Height(t *testing.T) {
	for _, tt := range []struct {
		name string
		node *AVLNode[int]
		want int
	}{
		{
			name: "nil",
			node: nil,
			want: 0,
		},
		{
			name: "42",
			node: &AVLNode[int]{h: 42},
			want: 42,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.node.Height(); got != tt.want {
				t.Errorf("Height() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAVLNode_IsBroken(t *testing.T) {
	for _, tt := range []struct {
		name string
		node *AVLNode[int]
		want bool
	}{
		{
			name: "ok empty",
			node: nil,
			want: false,
		},
		{
			name: "ok leaf",
			node: &AVLNode[int]{},
			want: false,
		},
		{
			name: "ok 1 left",
			node: &AVLNode[int]{
				l: &AVLNode[int]{h: 1},
			},
			want: false,
		},
		{
			name: "ok 1 right",
			node: &AVLNode[int]{
				r: &AVLNode[int]{h: 1},
			},
			want: false,
		},
		{
			name: "ok 42 left and 42 right",
			node: &AVLNode[int]{
				l: &AVLNode[int]{h: 42},
				r: &AVLNode[int]{h: 42},
			},
			want: false,
		},
		{
			name: "ok 42 left and 43 right",
			node: &AVLNode[int]{
				l: &AVLNode[int]{h: 42},
				r: &AVLNode[int]{h: 43},
			},
			want: false,
		},
		{
			name: "ok 43 left and 42 right",
			node: &AVLNode[int]{
				l: &AVLNode[int]{h: 43},
				r: &AVLNode[int]{h: 42},
			},
			want: false,
		},
		{
			name: "broken 40 left and 42 right",
			node: &AVLNode[int]{
				l: &AVLNode[int]{h: 40},
				r: &AVLNode[int]{h: 42},
			},
			want: true,
		},
		{
			name: "broken 2 left",
			node: &AVLNode[int]{
				l: &AVLNode[int]{h: 2},
			},
			want: true,
		},
		{
			name: "broken 2 right",
			node: &AVLNode[int]{
				r: &AVLNode[int]{h: 2},
			},
			want: true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.node.IsBroken(); got != tt.want {
				t.Errorf("IsBroken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAVLNode_Insert(t *testing.T) {
	root := &AVLNode[int]{v: 42}
	leq := func(l, r int) bool { return l <= r }
	// right
	n43 := root.Insert(43, leq)
	if root.r != n43 {
		t.Errorf("root.r != n43")
	}
	if n43.p != root {
		t.Errorf("n43.p != root")
	}
	// left
	n41 := root.Insert(41, leq)
	if root.l != n41 {
		t.Errorf("root.r != n41")
	}
	if n41.p != root {
		t.Errorf("n41.p != root")
	}
}

func TestAVLNode_Insert_Panic_Left(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("recover() == nil")
		}
	}()
	root := &AVLNode[int]{
		v: 42,
		l: &AVLNode[int]{},
	}
	root.Insert(41, func(l, r int) bool { return l <= r })
}

func TestAVLNode_Insert_Panic_Right(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("recover() == nil")
		}
	}()
	root := &AVLNode[int]{
		v: 42,
		r: &AVLNode[int]{},
	}
	root.Insert(43, func(l, r int) bool { return l <= r })
}

// TestAVLNode_Find
//
//			   42
//	    40    		   45
//	10			  43	     70
//		15               63		 95
//					  55
func TestAVLNode_Find(t *testing.T) {
	less := func(l, r int) bool { return l < r }
	n15 := &AVLNode[int]{v: 15}
	n10 := &AVLNode[int]{v: 10, r: n15}
	n40 := &AVLNode[int]{v: 40, l: n10}
	n43 := &AVLNode[int]{v: 43, l: n10}
	n55 := &AVLNode[int]{v: 55}
	n63 := &AVLNode[int]{v: 63, l: n55}
	n95 := &AVLNode[int]{v: 95}
	n70 := &AVLNode[int]{v: 70, l: n63, r: n95}
	n45 := &AVLNode[int]{v: 45, l: n43, r: n70}
	root := &AVLNode[int]{
		v: 42,
		l: n40,
		r: n45,
	}
	for _, tt := range []struct {
		name  string
		root  *AVLNode[int]
		value int
		want  *AVLNode[int]
	}{
		{
			name:  "nil",
			root:  nil,
			value: 42,
			want:  nil,
		},
		{
			name:  "root",
			root:  root,
			value: root.v,
			want:  root,
		},
		{
			name:  "15",
			root:  root,
			value: n15.v,
			want:  n15,
		},
		{
			name:  "40",
			root:  root,
			value: n40.v,
			want:  n40,
		},
		{
			name:  "45",
			root:  root,
			value: n45.v,
			want:  n45,
		},
		{
			name:  "55",
			root:  root,
			value: n55.v,
			want:  n55,
		},
		{
			name:  "63",
			root:  root,
			value: n63.v,
			want:  n63,
		},
		{
			name:  "70",
			root:  root,
			value: n70.v,
			want:  n70,
		},
		{
			name:  "95",
			root:  root,
			value: n95.v,
			want:  n95,
		},
		{
			name:  "9 and 10 expected as a parent",
			root:  root,
			value: 9,
			want:  n10,
		},
		{
			name:  "14 and 15 expected as a parent",
			root:  root,
			value: 14,
			want:  n15,
		},
		{
			name:  "16 and 15 expected as a parent",
			root:  root,
			value: 16,
			want:  n15,
		},
		{
			name:  "41 and 40 expected as a parent",
			root:  root,
			value: 41,
			want:  n40,
		},
		{
			name:  "62 and 55 expected as a parent",
			root:  root,
			value: 55,
			want:  n55,
		},
		{
			name:  "64 and 63 expected as a parent",
			root:  root,
			value: 64,
			want:  n63,
		},
		{
			name:  "100 and 95 expected as a parent",
			root:  root,
			value: 100,
			want:  n95,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.root.Find(tt.value, less); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAVLNode_DetectRotationType(t *testing.T) {
	for _, tt := range []struct {
		name string
		node *AVLNode[int]
		want RotationType
	}{
		{
			name: "no rotation 0 left and 0 right",
			node: &AVLNode[int]{
				l: nil,
				r: nil,
			},
			want: NoRotation,
		},
		{
			name: "no rotation 1 left and 1 right",
			node: &AVLNode[int]{
				l: &AVLNode[int]{h: 1},
				r: &AVLNode[int]{h: 1},
			},
			want: NoRotation,
		},
		{
			name: "no rotation 0 left and 1 right",
			node: &AVLNode[int]{
				l: nil,
				r: &AVLNode[int]{h: 1},
			},
			want: NoRotation,
		},
		{
			name: "no rotation 1 left and 0 right",
			node: &AVLNode[int]{
				l: &AVLNode[int]{h: 1},
				r: nil,
			},
			want: NoRotation,
		},
		{
			name: "small left 2 left 1 1 and 0 right",
			node: &AVLNode[int]{
				l: &AVLNode[int]{
					h: 2,
					l: &AVLNode[int]{h: 1},
					r: &AVLNode[int]{h: 1},
				},
				r: nil,
			},
			want: SmallLeft,
		},
		{
			name: "small left 2 left 1 0 and 0 right",
			node: &AVLNode[int]{
				l: &AVLNode[int]{
					h: 2,
					l: &AVLNode[int]{h: 1},
					r: nil,
				},
				r: nil,
			},
			want: SmallLeft,
		},
		{
			name: "big left 2 left 0 1 and 0 right",
			node: &AVLNode[int]{
				l: &AVLNode[int]{
					h: 2,
					l: nil,
					r: &AVLNode[int]{h: 1},
				},
				r: nil,
			},
			want: BigLeft,
		},
		{
			name: "small right 0 left and 2 right 1 1",
			node: &AVLNode[int]{
				l: nil,
				r: &AVLNode[int]{
					h: 2,
					l: &AVLNode[int]{h: 1},
					r: &AVLNode[int]{h: 1},
				},
			},
			want: SmallRight,
		},
		{
			name: "small right 0 left and 2 right 0 1",
			node: &AVLNode[int]{
				l: nil,
				r: &AVLNode[int]{
					h: 2,
					l: nil,
					r: &AVLNode[int]{h: 1},
				},
			},
			want: SmallRight,
		},
		{
			name: "big right 0 left and 2 right 1 0",
			node: &AVLNode[int]{
				l: nil,
				r: &AVLNode[int]{
					h: 2,
					l: &AVLNode[int]{h: 1},
					r: nil,
				},
			},
			want: BigRight,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.node.DetectRotationType(); got != tt.want {
				t.Errorf("DetectRotationType() = %v, want %v", got, tt.want)
			}
		})
	}
}
